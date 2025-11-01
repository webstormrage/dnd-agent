package handlers

import (
	"dnd-agent/pkg/domain"
)

func SetPlayerCurrentUnit(w *domain.World, command *domain.Command) {
	defer domain.Resolve(command)
	raw, ok := command.Args["unitId"]
	if !ok {
		return
	}
	unitId, ok := raw.(int)
	if !ok {
		return
	}
	w.PlayerUnitId = &unitId
}
