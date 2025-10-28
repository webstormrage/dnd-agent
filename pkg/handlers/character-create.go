package handlers

import (
	characterCreation "dnd-agent/pkg/character-creation"
	"dnd-agent/pkg/domain"
)

func CharacterCreate(w *domain.World, command *domain.Command) {
	character := characterCreation.ScanCharacter()
	// TODO: fix this
	character.ID = 1
	w.Units[character.ID] = character

	ret := interface{}(character.ID)
	command.Pop = &ret
}
