package main

import (
	characterCreation "dnd-agent/pkg/character-creation"
	"dnd-agent/pkg/domain"
	"dnd-agent/pkg/utils"
	"encoding/json"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"os"
)

func CallLuaCallback(script string, table string, index string, ctx map[string]interface{}) (map[string]interface{}, error) {
	L := lua.NewState()
	defer L.Close()

	// Выполняем скрипт
	if err := L.DoString(script); err != nil {
		return nil, fmt.Errorf("ошибка загрузки скрипта: %v", err)
	}

	// Получаем таблицу handlers
	handlers := L.GetGlobal(table)
	handlerTbl, ok := handlers.(*lua.LTable)
	if !ok {
		return nil, fmt.Errorf("'handlers' не таблица или отсутствует")
	}

	// Ищем функцию по ключу — например handlers["Scenario.start"]
	fn := L.GetField(handlerTbl, index)
	if fn.Type() != lua.LTFunction {
		return nil, fmt.Errorf("обработчик %s не найден в %s", index, table)
	}

	// Конвертируем ctx (Go → Lua)
	ctxTbl := utils.MapToLuaTable(L, ctx)

	// Вызываем обработчик с контекстом
	if err := L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    0, // Lua-хэндлер ничего не возвращает, он изменяет ctx
		Protect: true,
	}, ctxTbl); err != nil {
		return nil, fmt.Errorf("ошибка вызова обработчика %s %s: %v", table, index, err)
	}

	result := utils.LuaTableToMap(ctxTbl)

	return result, nil
}

func CallLuaHandler(scriptPath, handlerName string, ctxIn map[string]interface{}) ([]domain.Command, error) {
	data, err := os.ReadFile(scriptPath)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения скрипта: %v", err)
	}
	script := string(data)

	ctxOut, err := CallLuaCallback(script, "handlers", handlerName, ctxIn)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения Lua: %v", err)
	}

	rawNext, ok := ctxOut["next"]
	if !ok {
		return nil, nil // нет поля next
	}

	// Преобразуем []interface{} → []domain.Command
	nextIface, ok := rawNext.([]interface{})
	if !ok {
		return nil, fmt.Errorf("ожидался []interface{}, а получен %T", rawNext)
	}

	cmds := make([]domain.Command, 0, len(nextIface))
	for _, v := range nextIface {
		if m, ok := v.(map[string]interface{}); ok {
			cmds = append(cmds, domain.Command(m))
		} else {
			return nil, fmt.Errorf("элемент в next имеет неверный тип: %T", v)
		}
	}

	if len(cmds) == 0 {
		return nil, nil
	}
	return cmds, nil
}

func HandleCommand(w *domain.World, cmds []domain.Command) []domain.Command {
	if len(cmds) == 0 {
		return cmds
	}

	command := cmds[0]
	rest := cmds[1:]

	cmdName, _ := command["command"].(string)

	switch cmdName {

	case "Scenario.start":
		// Проверка — есть ли уже юнит под контролем игрока
		if _, exists := w.Units[1]; exists {
			return rest
		}

		next, err := CallLuaHandler(
			"lua/scenario/init.lua",
			"Scenario.start",
			map[string]interface{}{
				"next": []domain.Command{},
			},
		)
		if err != nil {
			fmt.Println(err)
			return rest
		}

		rest = append(rest, next...)
		return rest

	case "Character.create":
		// Создаём персонажа вручную (пока что)
		character := characterCreation.ScanCharacter()
		character.ID = 1
		w.Units[character.ID] = character
		fmt.Println("Character.create: unitId", character.ID)

		// Добавляем команду Character.On.create
		next := domain.Command{
			"command": "Character.On.create",
			"characterOnCreate": map[string]interface{}{
				"unitId": float64(character.ID),
			},
		}
		rest = append(rest, next)
		return rest

	case "Character.On.create":
		// Извлекаем данные из команды
		characterOnCreate, _ := command["characterOnCreate"].(map[string]interface{})
		unitID, _ := characterOnCreate["unitId"].(float64) // JSON числа → float64
		fmt.Println("Character.On.create unitId", unitID)

		next, err := CallLuaHandler(
			"lua/scenario/init.lua",
			"Character.On.create",
			map[string]interface{}{
				"next": []domain.Command{},
				"characterOnCreate": map[string]interface{}{
					"unitId": unitID,
				},
			},
		)
		if err != nil {
			fmt.Println(err)
			return rest
		}

		rest = append(rest, next...)
		return rest

	case "Unit.spawn":
		// Получаем spawn-данные
		spawnData, _ := command["spawn"].(map[string]interface{})

		unitID := int(spawnData["unitId"].(float64))
		x := int(spawnData["x"].(float64))
		y := int(spawnData["y"].(float64))
		owner, _ := spawnData["owner"].(string)
		gameZoneId, _ := spawnData["gameZoneId"].(string)

		fmt.Println("Unit.spawn unitId", unitID)
		unit := w.Units[unitID]
		unit.X = x
		unit.Y = y
		unit.Owner = owner
		unit.ZoneId = &gameZoneId

		data, err := json.MarshalIndent(unit, "", "  ")
		if err != nil {
			fmt.Println("Ошибка сериализации:", err)
			return rest
		}
		fmt.Println(string(data))
		return rest

	default:
		fmt.Printf("Неизвестная команда: %v\n", cmdName)
		return rest
	}
}

func main() {
	w := &domain.World{
		Units: make(map[int]*domain.Unit),
	}

	queue := []domain.Command{}

	fmt.Println("DnDAI запущен")

	for {
		// Пока есть команды в очереди — обрабатываем их
		for len(queue) > 0 {
			queue = HandleCommand(w, queue)
		}

		// Очередь пуста — ждём ввод пользователя
		fmt.Print("> ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			// Если пустая строка — просто пропускаем
			if err.Error() == "unexpected newline" {
				continue
			}
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		if input == "exit" || input == "quit" {
			fmt.Println("Завершение работы.")
			break
		}

		// Добавляем введённую команду в очередь
		queue = append(queue, domain.Command{"command": input})
	}
}
