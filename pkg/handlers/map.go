package handlers

import (
	"dnd-agent/pkg/domain"
	worldzone "dnd-agent/pkg/world-zone"
	"fmt"
)

func Map(w *domain.World, command domain.Command, rest []domain.Command) []domain.Command {
	// Проверка — есть ли уже юнит под контролем игрока
	if _, exists := w.Units[1]; !exists {
		return rest
	}
	character := w.Units[1]
	// Проверка — заспавлен ли юнит
	if character.ZoneId == nil {
		return rest
	}
	fmt.Println(worldzone.RenderLevel(w.Zones[*character.ZoneId]))
	return rest
}
