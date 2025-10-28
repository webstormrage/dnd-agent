package handlers

import (
	"dnd-agent/pkg/domain"
)

func UnitSpawn(w *domain.World, command *domain.Command) {

	unitID := int(command.Args["unitId"].(float64))
	x := int(command.Args["x"].(float64))
	y := int(command.Args["y"].(float64))
	owner, _ := command.Args["owner"].(string)
	gameZoneId, _ := command.Args["gameZoneId"].(string)

	unit := w.Units[unitID]
	unit.X = x
	unit.Y = y
	unit.Owner = owner
	unit.ZoneId = &gameZoneId
	zone := w.Zones[*unit.ZoneId]
	zone.SpawnUnit(unit.ID, x, y)
	domain.Resolve(command)
}
