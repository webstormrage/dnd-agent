package handlers

import (
	"dnd-agent/pkg/domain"
	worldzone "dnd-agent/pkg/world-zone"
	"encoding/xml"
	"strconv"
)

func UnitSpawn(w *domain.World, command domain.Command, rest []domain.Command) []domain.Command {
	// Получаем spawn-данные
	spawnData, _ := command["spawn"].(map[string]interface{})

	unitID := int(spawnData["unitId"].(float64))
	x := int(spawnData["x"].(float64))
	y := int(spawnData["y"].(float64))
	owner, _ := spawnData["owner"].(string)
	gameZoneId, _ := spawnData["gameZoneId"].(string)

	unit := w.Units[unitID]
	unit.X = x
	unit.Y = y
	unit.Owner = owner
	unit.ZoneId = &gameZoneId
	zone := w.Zones[*unit.ZoneId]
	zone.Objects = append(zone.Objects, worldzone.Object{
		XMLName: xml.Name{
			Local: "objects",
			Space: "",
		},
		Type:   "unit",
		X:      unit.X,
		Y:      unit.Y,
		Width:  1,
		Height: 1,
		Name:   "unit#" + strconv.Itoa(unit.ID),
	})
	return rest
}
