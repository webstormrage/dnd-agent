package handlers

import (
	"bufio"
	"dnd-agent/pkg/domain"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func OptionScanf(world *domain.World, command *domain.Command) {
	reader := bufio.NewReader(os.Stdin)

	name, _ := command.Args["name"].(string)
	typ, _ := command.Args["type"].(string)

	var options []string
	if raw, ok := command.Args["options"].([]interface{}); ok {
		for _, v := range raw {
			if s, ok := v.(string); ok {
				options = append(options, s)
			}
		}
	} else if raw, ok := command.Args["options"].([]string); ok {
		options = raw
	}

	var result interface{}

	switch typ {

	case "string":
		fmt.Printf("Введите значение для '%s' (string): ", name)
		input, _ := reader.ReadString('\n')
		result = strings.TrimSpace(input)

	case "int":
		for {
			fmt.Printf("Введите значение для '%s' (int): ", name)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			val, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("❌ Ошибка: нужно ввести целое число.")
				continue
			}
			result = val
			break
		}

	case "select":
		if len(options) == 0 {
			fmt.Printf("⚠️  '%s' имеет тип 'select', но без options — пропуск.\n", name)
			domain.Resolve(command)
			return
		}

		fmt.Printf("\nВыберите значение для '%s':\n", name)
		for i, opt := range options {
			fmt.Printf("  %d) %s\n", i+1, opt)
		}

		for {
			fmt.Printf("Введите номер (1-%d): ", len(options))
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			num, err := strconv.Atoi(input)
			if err != nil || num < 1 || num > len(options) {
				fmt.Println("❌ Ошибка: введите корректный номер варианта.")
				continue
			}
			result = options[num-1]
			break
		}

	default:
		fmt.Printf("⚠️  Тип '%s' не поддерживается, пропускаем '%s'.\n", typ, name)
		domain.Resolve(command)
		return
	}

	// ✅ Записываем результат в command.Stack.Pop
	command.Stack.Pop = &result
}
