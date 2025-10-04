package main

import (
	"dnd-agent/pkg/world-zone"
	"fmt"
	"os"
)

func main() {
	// Чтение XML из файла в строку
	data, err := os.ReadFile("maps/frey-tavern.xml")
	if err != nil {
		panic(err)
	}

	// Создание структуры уровня из XML через пакет worldzone
	level, err := worldzone.LoadLevelFromXML(string(data))
	if err != nil {
		panic(err)
	}

	// Рендеринг уровня в строку
	result := worldzone.RenderLevel(level)

	// Вывод карты в стандартный вывод
	fmt.Println(result)
}
