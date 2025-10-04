package unitDefintion

import (
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

func luaTableToMap(tbl *lua.LTable) interface{} {
	// Проверяем, выглядит ли таблица как массив
	isArray := true
	maxIndex := 0
	length := 0

	tbl.ForEach(func(k, _ lua.LValue) {
		if keyNum, ok := k.(lua.LNumber); ok {
			i := int(keyNum)
			if i > maxIndex {
				maxIndex = i
			}
			length++
		} else {
			isArray = false
		}
	})

	// Если это действительно массив (ключи 1..N без пропусков)
	if isArray && maxIndex == length {
		slice := make([]interface{}, length)
		for i := 1; i <= length; i++ {
			val := tbl.RawGetInt(i)
			slice[i-1] = luaValueToInterface(val)
		}
		return slice
	}

	// Иначе — обычная таблица (map)
	result := make(map[string]interface{})
	tbl.ForEach(func(k, v lua.LValue) {
		key := k.String()
		result[key] = luaValueToInterface(v)
	})
	return result
}

// Вспомогательная функция для конвертации значения
func luaValueToInterface(v lua.LValue) interface{} {
	switch value := v.(type) {
	case lua.LNumber:
		return float64(value)
	case lua.LString:
		return string(value)
	case lua.LBool:
		return bool(value)
	case *lua.LTable:
		return luaTableToMap(value)
	default:
		return fmt.Sprintf("<unsupported:%s>", value.Type().String())
	}
}
