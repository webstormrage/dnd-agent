package luaUtils

import (
	"dnd-agent/pkg/domain"
	"dnd-agent/pkg/utils"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"os"
)

// LoadLuaWithCore читает lua/core.lua и указанный скрипт, возвращая объединённый код.
func LoadLuaWithCore(scriptPath string) (string, error) {
	// читаем основной скрипт
	scriptData, err := os.ReadFile(scriptPath)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения скрипта %s: %v", scriptPath, err)
	}

	// читаем общий core.lua (всегда из lua/core.lua)
	coreData, err := os.ReadFile("lua/core.lua")
	if err != nil {
		return "", fmt.Errorf("ошибка чтения core.lua: %v", err)
	}

	// объединяем тексты
	fullScript := string(coreData) + "\n\n" + string(scriptData)
	return fullScript, nil
}

// CallLuaHandler — вызывает Lua-хэндлер и синхронизирует изменения обратно в command
func CallLuaHandler(scriptPath, handlerName string, command *domain.Command) error {

	fullScript, err := LoadLuaWithCore(scriptPath)

	// создаём Lua VM
	L := lua.NewState()
	defer L.Close()

	// выполняем объединённый Lua-код
	if err := L.DoString(fullScript); err != nil {
		return fmt.Errorf("ошибка загрузки Lua (%s): %v", scriptPath, err)
	}

	// достаём таблицу генераторов
	handlers := L.GetGlobal("generators")
	handlerTbl, ok := handlers.(*lua.LTable)
	if !ok {
		return fmt.Errorf("'generators' не таблица")
	}

	// достаём функцию по имени
	fn := L.GetField(handlerTbl, handlerName)
	if fn.Type() != lua.LTFunction {
		return fmt.Errorf("обработчик '%s' не найден в generators", handlerName)
	}

	// создаём Lua-таблицы для args, state, stack
	argsTbl := utils.MapToLuaTable(L, command.Args)
	stateTbl := utils.MapToLuaTable(L, command.State)
	stackTbl := utils.StructToLuaTable(L, command.Stack)

	// вызываем обработчик
	if err := L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    0,
		Protect: true,
	}, argsTbl, stateTbl, stackTbl); err != nil {
		return fmt.Errorf("ошибка вызова хэндлера '%s': %v", handlerName, err)
	}

	// обратное преобразование таблиц в Go-структуры
	command.Args = utils.LuaTableToMap(argsTbl)
	command.State = utils.LuaTableToMap(stateTbl)

	newStack, err := utils.LuaTableToStack(stackTbl)
	if err != nil {
		return fmt.Errorf("ошибка конвертации stack: %v", err)
	}
	command.Stack = *newStack

	return nil
}
