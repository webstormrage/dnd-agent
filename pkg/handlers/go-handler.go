package handlers

import (
	"dnd-agent/pkg/domain"
	worldzone "dnd-agent/pkg/world-zone"
	"strconv"
	"strings"
)

func Go(w *domain.World, command *domain.Command) {
	defer domain.Resolve(command)

	argv := command.Args["argv"].([]string)
	dir := argv[0]
	times := 1
	if len(argv) > 1 {
		times, _ = strconv.Atoi(argv[1])
	}

	// TODO: extract current unitID from command
	unit := w.Units[1]
	if unit.ZoneId == nil {
		return
	}
	level := w.Zones[*unit.ZoneId]

	dx := 0
	dy := 0
	if strings.Contains(dir, "w") {
		dy = -1
	} else if strings.Contains(dir, "s") {
		dy = 1
	}
	if strings.Contains(dir, "a") {
		dx = -1
	} else if strings.Contains(dir, "d") {
		dx = 1
	}
	_, path := worldzone.Transition(unit.ID, dx, dy, times, level)

	target := path[len(path)-1]
	tunnelTarget := worldzone.GetZoneTunnel(target)
	if len(tunnelTarget) < 2 {
		return
	}

	nextZoneId := tunnelTarget[0]
	nextPosition := tunnelTarget[1]

	nextZone := w.Zones[nextZoneId]
	if nextZone == nil {
		return
	}
	targetObject := nextZone.GetObjectByName(nextPosition)
	if targetObject == nil {
		return
	}

	result := worldzone.FindNearestBFS(nextZone, targetObject.X, targetObject.Y, worldzone.IsPlaceble)
	if result == nil {
		return
	}

	// TODO: Cell должен знать свои координаты, у юнита не должно быть X и Y, координат
	targetCell := result[len(result)-1]
	level.UnSpawn(unit.ID)
	nextZone.SpawnUnit(unit.ID, targetCell.X, targetCell.Y)
}
