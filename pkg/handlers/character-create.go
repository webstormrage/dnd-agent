package handlers

import (
	"dnd-agent/pkg/domain"
	luaUtils "dnd-agent/pkg/lua-utils"
	"fmt"
)

/*func CharacterCreate(w *domain.World, command *domain.Command) {
	character := characterCreation.ScanCharacter()
	// TODO: fix this
	character.ID = 1
	w.Units[character.ID] = character

	ret := interface{}(character.ID)
	command.Pop = &ret
}*/

func CharacterCreate(w *domain.World, command *domain.Command) {
	err := luaUtils.CallLuaHandler(
		"lua/unit-definition/create-character.lua",
		"Unit.createCharacter",
		command,
	)
	if err != nil {
		fmt.Println(err)
		domain.Resolve(command)
	}
}
