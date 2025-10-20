package handlers

import (
	"dnd-agent/pkg/domain"
	worldzone "dnd-agent/pkg/world-zone"
	"fmt"
)

func Map(w *domain.World, command *domain.Command) {
	// Проверка — есть ли уже юнит под контролем игрока
	if _, exists := w.Units[1]; !exists {
		domain.Resolve(command)
		return
	}
	character := w.Units[1]
	// Проверка — заспавлен ли юнит
	if character.ZoneId == nil {
		domain.Resolve(command)
		return
	}
	fmt.Println(worldzone.RenderLevel(w.Zones[*character.ZoneId]))
	domain.Resolve(command)
	return
}
