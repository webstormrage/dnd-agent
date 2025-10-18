package characterCreation

import (
	"bufio"
	"dnd-agent/pkg/domain"
	unitDefintion "dnd-agent/pkg/unit-defintion"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"os"
	"strconv"
	"strings"
)

func collectInputFromChoices(choices []domain.Choice) map[string]interface{} {
	reader := bufio.NewReader(os.Stdin)
	results := make(map[string]interface{})

	for _, ch := range choices {
		for {
			switch ch.Type {
			case "string":
				fmt.Printf("Введите значение для '%s' (string): ", ch.Name)
				input, _ := reader.ReadString('\n')
				results[ch.Name] = strings.TrimSpace(input)
				break

			case "int":
				fmt.Printf("Введите значение для '%s' (int): ", ch.Name)
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)
				val, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("❌ Ошибка: нужно ввести целое число.")
					continue
				}
				results[ch.Name] = val
				break

			case "select":
				if len(ch.Options) == 0 {
					fmt.Printf("⚠️  '%s' имеет тип 'select', но без options — пропуск.\n", ch.Name)
					break
				}

				fmt.Printf("\nВыберите значение для '%s':\n", ch.Name)
				for i, opt := range ch.Options {
					fmt.Printf("  %d) %s\n", i+1, opt)
				}

				var choiceIndex int
				for {
					fmt.Printf("Введите номер (1-%d): ", len(ch.Options))
					input, _ := reader.ReadString('\n')
					input = strings.TrimSpace(input)
					num, err := strconv.Atoi(input)
					if err != nil || num < 1 || num > len(ch.Options) {
						fmt.Println("❌ Ошибка: введите корректный номер варианта.")
						continue
					}
					choiceIndex = num - 1
					break
				}

				results[ch.Name] = ch.Options[choiceIndex]
				break

			default:
				fmt.Printf("⚠️  Тип '%s' не поддерживается, пропускаем '%s'.\n", ch.Type, ch.Name)
				break
			}

			break // выходим из внутреннего цикла, если всё успешно
		}
	}

	return results
}

func getTemplate(template string) string {
	data, err := os.ReadFile("lua/unit-definition/" + template + ".lua")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func ScanCharacter() *domain.Unit {
	L := lua.NewState()

	attrTable := L.NewTable()
	inventoryTable := L.NewTable()

	//TODO: перевести на command и handlers
	templates := []string{
		"base/base",
		"abilities/abilities",
		"races/human",
		"backgrounds/outlander",
		"classes/fighter-1",
		"character/character",
	}

	var attributes interface{}
	var inventory interface{}
	var err error

	for _, template := range templates {
		attributes, inventory, err = unitDefintion.ProcessUnitDefinition(
			L,
			getTemplate(template),
			attrTable,
			inventoryTable,
			collectInputFromChoices,
		)

		if err != nil {
			panic(err)
		}
	}

	// 🧱 Формируем итоговый объект Unit
	return &domain.Unit{
		Attributes: attributes,
		Inventory:  inventory,
		Equipment:  nil,
		States:     nil,
		ID:         0,
		X:          0,
		Y:          0,
		ZoneId:     nil,
		Owner:      "",
	}
}
