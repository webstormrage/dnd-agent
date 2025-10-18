package pipeline

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"dnd-agent/pkg/domain"
)

// Engine — основной игровой цикл ввода / обработки команд.
type Engine struct {
	world *domain.World
	queue []domain.Command
}

// NewEngine — конструктор.
func NewEngine() *Engine {
	return &Engine{
		world: &domain.World{},
		queue: make([]domain.Command, 0),
	}
}

// Run — главный цикл: обрабатывает очередь и ждёт команды пользователя.
func (e *Engine) Run() {
	fmt.Println("DnDAI запущен")
	reader := bufio.NewReader(os.Stdin)

	for {
		// Пока есть команды в очереди — обрабатываем их
		for len(e.queue) > 0 {
			e.queue = HandleCommand(e.world, e.queue)
		}

		// Очередь пуста — ждём ввод пользователя
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

		cmd, args := ParseCommandLine(line)

		if cmd == "/quit" {
			fmt.Println("Завершение работы.")
			break
		}

		e.AddCommand(cmd, args)
	}
}

// AddCommand — добавляет новую команду в очередь.
func (e *Engine) AddCommand(cmd string, args []string) {
	e.queue = append(e.queue, domain.Command{
		"command": cmd,
		"args":    args,
	})
}
