package unitDefintion

import (
	"encoding/json"
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func RunDefinition(L *lua.LState, luaCode string, attributesTable *lua.LTable) (interface{}, error) {

	// Выполняем Lua-код из строки
	if err := L.DoString(luaCode); err != nil {
		return nil, fmt.Errorf("ошибка при выполнении Lua-кода: %v", err)
	}

	fn := L.GetGlobal("unitDefinition")
	if fn.Type() != lua.LTFunction {
		return nil, fmt.Errorf("в Lua не найдена функция '%s'", "unitDefinition")
	}

	// вызываем функцию init(attributes)
	if err := L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    0,
		Protect: true,
	}, attributesTable); err != nil {
		return nil, fmt.Errorf("ошибка при вызове Lua-функции: %v", err)
	}

	return luaTableToMap(attributesTable), nil
}

func GetChoices(L *lua.LState, luaCode string, attributes *lua.LTable, choices *lua.LTable) (interface{}, error) {

	// Выполняем Lua-код из строки
	if err := L.DoString(luaCode); err != nil {
		return nil, fmt.Errorf("ошибка при выполнении Lua-кода: %v", err)
	}

	fn := L.GetGlobal("optionsDefinition")
	if fn.Type() != lua.LTFunction {
		return luaTableToMap(choices), nil
	}

	if err := L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    0,
		Protect: true,
	}, attributes, choices); err != nil {
		return nil, fmt.Errorf("ошибка при вызове Lua-функции: %v", err)
	}

	return luaTableToMap(choices), nil
}

func PrettyPrintJSON(data interface{}) string {
	b, _ := json.MarshalIndent(data, "", "  ")
	return string(b)
}
