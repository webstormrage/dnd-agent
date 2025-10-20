package pipeline

import (
	"dnd-agent/pkg/domain"
	"dnd-agent/pkg/handlers"
	"fmt"
)

func HandleCommand(w *domain.World, command *domain.Command) {

	switch command.Procedure {
	case "/load":
		handlers.Load(w, command)
	case "/save":
		handlers.Save(w, command)
	case "/start":
		handlers.Start(w, command)
	case "Character.create":
		handlers.CharacterCreate(w, command)
	case "Unit.spawn":
		handlers.UnitSpawn(w, command)
	case "/map":
		handlers.Map(w, command)
	default:
		fmt.Printf("Неизвестная команда: %v\n", command.Procedure)
	}
}
