package unitDefintion

import (
	"encoding/json"
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

type Choice struct {
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	Options []string `json:"options"`
}

// HandleChoicesFunc — callback, который обрабатывает []Choice и возвращает map[string]interface{}
type HandleChoicesFunc func(choices []Choice) map[string]interface{}

// ProcessUnitDefinition — обобщённая функция для полного цикла:
// 1. Получение choices из Lua,
// 2. Передача их в callback для заполнения,
// 3. Преобразование обратно в Lua-таблицу,
// 4. Вызов RunDefinition.
func ProcessUnitDefinition(
	L *lua.LState,
	luaCode string,
	attrTable *lua.LTable,
	inventoryTable *lua.LTable,
	handleChoices HandleChoicesFunc,
) (interface{}, interface{}, error) {

	// Создаём таблицу выбора
	choicesTable := L.NewTable()

	// Получаем список choices из Lua
	choices, err := GetChoices(L, luaCode, attrTable, choicesTable)
	if err != nil {
		return nil, nil, fmt.Errorf("ошибка получения choices: %v", err)
	}

	// Передаём их в callback для получения пользовательских значений
	results := handleChoices(choices)

	// Преобразуем результат обратно в Lua-таблицу
	optionsTable := MapToLuaTable(L, results)

	// Выполняем основное определение
	attributes, inventory, err := RunDefinition(L, luaCode, attrTable, inventoryTable, optionsTable)
	if err != nil {
		return nil, nil, fmt.Errorf("ошибка при выполнении RunDefinition: %v", err)
	}

	return attributes, inventory, nil
}

func RunDefinition(L *lua.LState, luaCode string, attributesTable, inventoryTable, optionsTable *lua.LTable) (interface{}, interface{}, error) {

	// Выполняем Lua-код из строки
	if err := L.DoString(luaCode); err != nil {
		return nil, nil, fmt.Errorf("ошибка при выполнении Lua-кода: %v", err)
	}

	fn := L.GetGlobal("unitDefinition")
	if fn.Type() != lua.LTFunction {
		return nil, nil, fmt.Errorf("в Lua не найдена функция '%s'", "unitDefinition")
	}

	// вызываем функцию init(attributes)
	if err := L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    0,
		Protect: true,
	}, attributesTable, inventoryTable, optionsTable); err != nil {
		return nil, nil, fmt.Errorf("ошибка при вызове Lua-функции: %v", err)
	}

	return luaTableToMap(attributesTable), luaTableToMap(inventoryTable), nil
}

func GetChoices(L *lua.LState, luaCode string, attributes *lua.LTable, choices *lua.LTable) ([]Choice, error) {

	// Выполняем Lua-код из строки
	if err := L.DoString(luaCode); err != nil {
		return nil, fmt.Errorf("ошибка при выполнении Lua-кода: %v", err)
	}

	fn := L.GetGlobal("optionsDefinition")
	if fn.Type() != lua.LTFunction {
		return luaTableToChoices(choices), nil
	}

	if err := L.CallByParam(lua.P{
		Fn:      fn,
		NRet:    0,
		Protect: true,
	}, attributes, choices); err != nil {
		return nil, fmt.Errorf("ошибка при вызове Lua-функции: %v", err)
	}

	return luaTableToChoices(choices), nil
}

func PrettyPrintJSON(data interface{}) string {
	b, _ := json.MarshalIndent(data, "", "  ")
	return string(b)
}
