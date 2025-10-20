package handlers

import (
	"dnd-agent/pkg/domain"
	worldzone "dnd-agent/pkg/world-zone"
	"encoding/xml"
	"strconv"
)

func UnitSpawn(w *domain.World, command *domain.Command) {

	unitID := int(command.Args["unitId"].(float64))
	x := int(command.Args["x"].(float64))
	y := int(command.Args["y"].(float64))
	owner, _ := command.Args["x"].(string)
	gameZoneId, _ := command.Args["x"].(string)

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
	domain.Resolve(command)
}
