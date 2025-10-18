package luaUtils

import (
	"dnd-agent/pkg/domain"
	"dnd-agent/pkg/utils"
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
