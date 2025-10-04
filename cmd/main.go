package main

import (
	"bufio"
	"dnd-agent/pkg/unit-defintion"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"os"
	"strconv"
	"strings"
)

// CollectInputFromChoices — запрашивает ввод пользователя для всех Choice
func CollectInputFromChoices(choices []unitDefintion.Choice) map[string]interface{} {
	reader := bufio.NewReader(os.Stdin)
	results := make(map[string]interface{})

	for _, ch := range choices {
		for {
			fmt.Printf("Введите значение для '%s' (%s): ", ch.Name, ch.Type)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if ch.Type == "string" {
				results[ch.Name] = input
				break
			}

			if ch.Type == "int" {
				val, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("❌ Ошибка: нужно ввести целое число.")
					continue
				}
				results[ch.Name] = val
				break
			}

			// если тип неизвестен
			fmt.Printf("⚠️  Тип '%s' не поддерживается, пропускаем.\n", ch.Type)
			break
		}
	}

	return results
}

func main() {
	// Читаем содержимое Lua-файла в строку
	data, err := os.ReadFile("unit-definition/character/character.lua")
	if err != nil {
		panic(err)
	}
	luaCode := string(data)

	L := lua.NewState()

	attrTable := L.NewTable()
	choicesTable := L.NewTable()
	// Выполняем функцию init(attributes) из Lua-кода
	choices, err := unitDefintion.GetChoices(L, luaCode, attrTable, choicesTable)
	if err != nil {
		panic(err)
	}
	equipmentTable := L.NewTable()

	// Просим пользователя ввести данные
	results := CollectInputFromChoices(choices)

	optionsTable := unitDefintion.MapToLuaTable(L, results)

	attributes, err := unitDefintion.RunDefinition(L, luaCode, attrTable, equipmentTable, optionsTable)

	fmt.Println(unitDefintion.PrettyPrintJSON(attributes))
}
