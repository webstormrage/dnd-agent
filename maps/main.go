package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"dnd-agent/pkg/world-zone"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("DnD Level Viewer запущен.")
	fmt.Println("Введите путь к файлу уровня (или /quit для выхода).")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Ошибка ввода: %v\n", err)
			continue
		}

		path := strings.TrimSpace(input)
		if path == "" {
			continue
		}

		if path == "/quit" {
			fmt.Println("Завершение работы.")
			break
		}

		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("Ошибка чтения файла %s: %v\n", path, err)
			continue
		}

		var level worldzone.Level
		if err := level.Load(string(data)); err != nil {
			fmt.Printf("Ошибка загрузки уровня: %v\n", err)
			continue
		}

		fmt.Println()
		fmt.Println(worldzone.RenderLevel(&level))
		fmt.Println()
	}
}
