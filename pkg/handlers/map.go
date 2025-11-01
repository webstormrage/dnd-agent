package handlers

import (
	"dnd-agent/pkg/domain"
	worldzone "dnd-agent/pkg/world-zone"
	"fmt"
)

func Map(w *domain.World, command *domain.Command) {
	// Проверка — есть ли уже юнит под контролем игрока
	uid := *w.PlayerUnitId
	if _, exists := w.Units[uid]; !exists {
		domain.Resolve(command)
		return
	}
	character := w.Units[uid]
	// Проверка — заспавлен ли юнит
	if character.ZoneId == nil {
		domain.Resolve(command)
		return
	}
	fmt.Println(worldzone.RenderLevel(w.Zones[*character.ZoneId]))
	domain.Resolve(command)
	return
}
