package utils

import (
	"dnd-agent/pkg/domain"
	"encoding/json"
	"fmt"
	lua "github.com/yuin/gopher-lua"
)

func LuaTableToInterface(tbl *lua.LTable) interface{} {
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
			slice[i-1] = LuaValueToInterface(val)
		}
		return slice
	}

	// Иначе — обычная таблица (map)
	result := make(map[string]interface{})
	tbl.ForEach(func(k, v lua.LValue) {
		key := k.String()
		result[key] = LuaValueToInterface(v)
	})
	return result
}

func LuaTableToMap(tbl *lua.LTable) map[string]interface{} {
	result := make(map[string]interface{})
	tbl.ForEach(func(k, v lua.LValue) {
		key := k.String()
		result[key] = LuaValueToInterface(v)
	})
	return result
}

// Вспомогательная функция для конвертации значения
func LuaValueToInterface(v lua.LValue) interface{} {
	switch value := v.(type) {
	case lua.LNumber:
		return float64(value)
	case lua.LString:
		return string(value)
	case lua.LBool:
		return bool(value)
	case *lua.LTable:
		return LuaTableToInterface(value)
	default:
		return fmt.Sprintf("<unsupported:%s>", value.Type().String())
	}
}

// luaTableToChoices — конвертирует таблицу Lua в []Choice
func LuaTableToChoices(tbl *lua.LTable) []domain.Choice {
	var result []domain.Choice

	tbl.ForEach(func(_, v lua.LValue) {
		if entry, ok := v.(*lua.LTable); ok {
			c := domain.Choice{
				Options: nil, // дефолтное значение
			}

			entry.ForEach(func(k, val lua.LValue) {
				key := k.String()
				switch key {
				case "name":
					c.Name = val.String()
				case "type":
					c.Type = val.String()
				case "options":
					if arr, ok := val.(*lua.LTable); ok {
						c.Options = LuaTableToStringSlice(arr)
					}
				}
			})

			result = append(result, c)
		}
	})

	return result
}

// luaTableToStringSlice — конвертирует Lua-массив строк в []string
func LuaTableToStringSlice(tbl *lua.LTable) []string {
	var res []string
	tbl.ForEach(func(_, v lua.LValue) {
		res = append(res, v.String())
	})
	return res
}

// MapToLuaTable — рекурсивно преобразует map[string]interface{} в *lua.LTable
func MapToLuaTable(L *lua.LState, data map[string]interface{}) *lua.LTable {
	tbl := L.NewTable()
	for k, v := range data {
		switch value := v.(type) {
		case string:
			L.SetField(tbl, k, lua.LString(value))
		case int:
			L.SetField(tbl, k, lua.LNumber(value))
		case int64:
			L.SetField(tbl, k, lua.LNumber(value))
		case float64:
			L.SetField(tbl, k, lua.LNumber(value))
		case bool:
			L.SetField(tbl, k, lua.LBool(value))
		case map[string]interface{}:
			L.SetField(tbl, k, MapToLuaTable(L, value)) // рекурсия
		case []interface{}:
			L.SetField(tbl, k, SliceToLuaTable(L, value))
		case []string:
			arr := make([]interface{}, len(value))
			for i, s := range value {
				arr[i] = s
			}
			L.SetField(tbl, k, SliceToLuaTable(L, arr))
		default:
			// игнорируем неподдерживаемые типы
		}
	}
	return tbl
}

// SliceToLuaTable — преобразует []interface{} в *lua.LTable
func SliceToLuaTable(L *lua.LState, data []interface{}) *lua.LTable {
	tbl := L.NewTable()
	for _, v := range data {
		switch value := v.(type) {
		case string:
			tbl.Append(lua.LString(value))
		case int:
			tbl.Append(lua.LNumber(value))
		case float64:
			tbl.Append(lua.LNumber(value))
		case bool:
			tbl.Append(lua.LBool(value))
		case map[string]interface{}:
			tbl.Append(MapToLuaTable(L, value))
		default:
			// пропускаем неизвестные типы
		}
	}
	return tbl
}

// LuaTableToCommands — преобразует Lua-таблицу next в срез []Command.
func LuaTableToCommands(tbl *lua.LTable) []domain.Command {
	var cmds []domain.Command

	tbl.ForEach(func(_ lua.LValue, v lua.LValue) {
		t, ok := v.(*lua.LTable)
		if !ok {
			return
		}
		m := LuaTableToInterface(t)
		js, _ := json.Marshal(m)
		var c domain.Command
		if err := json.Unmarshal(js, &c); err == nil {
			cmds = append(cmds, c)
		}
	})

	return cmds
}
