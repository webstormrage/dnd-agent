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

func getTemplate(template string) string {
	data, err := os.ReadFile("unit-definition/" + template + "/" + template + ".lua")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func main() {
	L := lua.NewState()

	attrTable := L.NewTable()
	equipmentTable := L.NewTable()

	templates := []string{
		"base",
		"character",
	}

	var attributes interface{}
	var err error

	for _, template := range templates {
		attributes, err = unitDefintion.ProcessUnitDefinition(
			L,
			getTemplate(template),
			attrTable,
			equipmentTable,
			CollectInputFromChoices,
		)

		if err != nil {
			panic(err)
		}
	}

	fmt.Println(unitDefintion.PrettyPrintJSON(attributes))
}
