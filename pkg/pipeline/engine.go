package pipeline

import (
	"bufio"
	"dnd-agent/pkg/store"
	worldzone "dnd-agent/pkg/world-zone"
	"fmt"
	"os"
	"strings"

	"dnd-agent/pkg/domain"
)

func initializeWorld() *domain.World {
	w := &domain.World{
		Units: make(map[int]*domain.Unit),
		Zones: make(map[string]*worldzone.Level),
	}

	err := store.LoadAllMaps(w)
	if err != nil {
		panic(err)
	}
	return w
}

// Engine — главный управляющий цикл со стеком команд
type Engine struct {
	world *domain.World
	stack []*domain.Command // LIFO
}

// NewEngine — конструктор
func NewEngine() *Engine {
	world := initializeWorld()
	return &Engine{
		world: world,
		stack: make([]*domain.Command, 0),
	}
}

// Run — основной цикл обработки
func (e *Engine) Run() {
	fmt.Println("DnDAI (stack mode) запущен")
	reader := bufio.NewReader(os.Stdin)

	for {
		// Пока есть команды в стеке — обрабатываем вершину
		for len(e.stack) > 0 {
			top := e.stack[len(e.stack)-1]
			HandleCommand(e.world, top)
			e.handleStackResult(top)
		}

		// Стэк пуст — ждём пользовательский ввод
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		cmd, argv := ParseCommandLine(line)
		if cmd == "/quit" {
			fmt.Println("Завершение работы.")
			break
		}

		// создаём новую команду
		e.PushCommand(&domain.Command{
			Procedure: cmd,
			Args:      map[string]interface{}{"argv": argv},
			State:     make(map[string]interface{}),
			Stack:     domain.Stack{},
		})
	}
}

// PushCommand — добавляет команду в стек
func (e *Engine) PushCommand(cmd *domain.Command) {
	e.stack = append(e.stack, cmd)
}

// PopCommand — снимает верхнюю команду
func (e *Engine) PopCommand() *domain.Command {
	if len(e.stack) == 0 {
		return nil
	}
	top := e.stack[len(e.stack)-1]
	e.stack = e.stack[:len(e.stack)-1]
	return top
}

// handleStackResult — логика стэка
func (e *Engine) handleStackResult(cmd *domain.Command) {
	if cmd == nil {
		_ = e.PopCommand()
		return
	}

	// 🔹 Если есть следующая команда — пушим в стек
	if cmd.Stack.Push != nil {
		e.PushCommand(cmd.Stack.Push)
		return
	}

	// 🔹 Если есть pop — снимаем текущую
	if cmd.Stack.Pop != nil {
		popVal := *cmd.Stack.Pop
		e.PopCommand()

		// если стек не пуст — передаём значение наверх
		if len(e.stack) > 0 {
			parent := e.stack[len(e.stack)-1]
			target := parent.Stack.Target
			parent.State[target] = popVal
			parent.Stack.Push = nil
		}

		cmd.Stack.Pop = nil
	}
}
