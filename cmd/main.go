package main

import (
	characterCreation "dnd-agent/pkg/character-creation"
	"dnd-agent/pkg/domain"
	"dnd-agent/pkg/utils"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"os"
)

func CallLuaProcedure(script string, procedure string, ctx map[string]interface{}) (map[string]interface{}, error) {
	L := lua.NewState()
	defer L.Close()

	// Загружаем скрипт
	if err := L.DoString(script); err != nil {
		return nil, fmt.Errorf("ошибка загрузки скрипта: %v", err)
	}

	// Получаем ссылку на функцию
	fn := L.GetGlobal(procedure)
	if fn.Type() != lua.LTFunction {
		return nil, fmt.Errorf("функция %q не найдена в скрипте", procedure)
	}

	// Преобразуем ctx (Go) → Lua-таблицу
	ctxTbl := utils.MapToLuaTable(L, ctx)

	// Вызываем Lua-функцию (без ожидания возвращаемого значения)
	if err := L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    0,
		Protect: true,
	}, ctxTbl); err != nil {
		return nil, fmt.Errorf("ошибка вызова функции %s: %v", procedure, err)
	}

	// После выполнения функция могла изменить ctxTbl внутри Lua.
	// Преобразуем обратно Lua-таблицу → Go map
	result := utils.LuaTableToMap(ctxTbl)

	return result, nil
}

func CallLuaHandler(scriptPath, handlerName string, ctxIn map[string]interface{}) ([]domain.Command, error) {
	data, err := os.ReadFile(scriptPath)
	if err != nil {
		panic(err)
	}
	script := string(data)

	ctxOut, err := CallLuaProcedure(script, handlerName, ctxIn)
	if err != nil {
		panic(err)
	}

	cmds, ok := ctxOut["next"].([]domain.Command)
	if !ok {
		return nil, fmt.Errorf("ожидался []domain.Command, а получен %T", ctxOut["next"])
	}

	if len(cmds) == 0 {
		return nil, nil
	}
	return cmds, nil
}

var character *domain.Unit

func HandleCommand(w *domain.World, cmd []domain.Command) []domain.Command {
	if len(cmd) == 0 {
		return cmd
	}
	command := cmd[0]
	rest := cmd[1:]
	switch command.Command {
	case "Scenario.start":
		// TODO: заменить на проверку наличия игроков под контролем игрока
		_, exists := w.Units[1]
		if exists {
			return rest
		}
		next, err := CallLuaHandler(
			"lua/scenario/init.lua",
			"Scenario.start",
			map[string]interface{}{
				"next": []domain.Command{},
			})
		if err != nil {
			fmt.Println(err)
			return rest
		}
		rest = append(rest, next...)
		return rest
	case "Character.create":
		// TODO: здесь будет ожидание запроса с клиента
		character = characterCreation.ScanCharacter()
		// TODO: здесь сохранение в базу
		character.ID = 1
		w.Units[character.ID] = character
		rest = append(rest, domain.Command{Command: "Character.On.create", CharacterOnCreate: &domain.CharacterOnCreateCommand{UnitId: character.ID}})
		return rest
	case "Character.On.create":
		next, err := CallLuaHandler(
			"lua/scenario/init.lua",
			"Scenario.start",
			map[string]interface{}{
				"next": []domain.Command{},
				"characterOnCreate": map[string]interface{}{
					"unitId": command.CharacterOnCreate.UnitId,
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
		id := command.UnitSpawn.UnitId
		w.Units[id].X = command.UnitSpawn.X
		w.Units[id].X = command.UnitSpawn.X
		w.Units[id].Owner = command.UnitSpawn.Owner
		w.Units[id].ZoneId = &command.UnitSpawn.GameZoneId
		return rest
	default:
		fmt.Println("Неизвестная команда: %s\n", command.Command)
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
		queue = append(queue, domain.Command{Command: input})
	}
}
