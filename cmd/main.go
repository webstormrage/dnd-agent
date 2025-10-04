package main

import (
	"dnd-agent/pkg/unit-defintion"
	"fmt"
	lua "github.com/yuin/gopher-lua"
	"os"
)

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

	// Выводим результат в виде JSON
	fmt.Println(unitDefintion.PrettyPrintJSON(choices))
}
