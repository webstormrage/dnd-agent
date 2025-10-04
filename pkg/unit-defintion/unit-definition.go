package unitDefintion

import (
	"encoding/json"
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

// RunLuaFunctionWithAttributes — выполняет Lua-код из строки,
// вызывает указанную функцию, передаёт ей пустую таблицу attributes.
func RunLuaStringWithAttributes(luaCode, funcName string) (map[string]interface{}, error) {
	L := lua.NewState()

	// Выполняем Lua-код из строки
	if err := L.DoString(luaCode); err != nil {
		return nil, fmt.Errorf("ошибка при выполнении Lua-кода: %v", err)
	}

	fn := L.GetGlobal(funcName)
	if fn.Type() != lua.LTFunction {
		return nil, fmt.Errorf("в Lua не найдена функция '%s'", funcName)
	}

	attributesTable := L.NewTable()

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

func luaTableToMap(tbl *lua.LTable) map[string]interface{} {
	result := make(map[string]interface{})
	tbl.ForEach(func(k, v lua.LValue) {
		key := k.String()
		switch value := v.(type) {
		case lua.LNumber:
			result[key] = float64(value)
		case lua.LString:
			result[key] = string(value)
		case lua.LBool:
			result[key] = bool(value)
		case *lua.LTable:
			result[key] = luaTableToMap(value)
		default:
			result[key] = fmt.Sprintf("<unsupported:%s>", value.Type().String())
		}
	})
	return result
}

func PrettyPrintJSON(data map[string]interface{}) string {
	b, _ := json.MarshalIndent(data, "", "  ")
	return string(b)
}
