package main

import (
	"dnd-agent/pkg/unit-defintion"
	"fmt"
	"os"
)

func main() {
	// Читаем содержимое Lua-файла в строку
	data, err := os.ReadFile("unit-definition/base/base.lua")
	if err != nil {
		panic(err)
	}
	luaCode := string(data)

	// Выполняем функцию init(attributes) из Lua-кода
	attrs, err := unitDefintion.RunLuaStringWithAttributes(luaCode, "unitDefinition")
	if err != nil {
		panic(err)
	}

	// Выводим результат в виде JSON
	fmt.Println(unitDefintion.PrettyPrintJSON(attrs))
}
