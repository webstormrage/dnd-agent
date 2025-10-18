package handlers

import (
	characterCreation "dnd-agent/pkg/character-creation"
	"dnd-agent/pkg/domain"
)

func CharacterCreate(w *domain.World, command domain.Command, rest []domain.Command) []domain.Command {
	character := characterCreation.ScanCharacter()
	character.ID = 1
	w.Units[character.ID] = character

	// Добавляем команду Character.On.create
	next := domain.Command{
		"command": "Character.On.create",
		"characterOnCreate": map[string]interface{}{
			"unitId": float64(character.ID),
		},
	}
	rest = append(rest, next)
	return rest
}
