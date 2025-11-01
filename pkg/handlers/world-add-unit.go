package handlers

import (
	"dnd-agent/pkg/domain"
	"dnd-agent/pkg/utils"
)

func WorldAddUnit(w *domain.World, command *domain.Command) {
	unit, err := utils.ToUnit(command.Args["unit"])
	unit.ID = 1 // TODO: использовать sequence из базы данных
	if err != nil {
		domain.Resolve(command)
	}
	w.Units[unit.ID] = unit
	v := interface{}(unit.ID)
	command.Stack.Pop = &v
}
