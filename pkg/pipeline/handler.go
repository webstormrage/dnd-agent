package pipeline

import (
	"dnd-agent/pkg/domain"
	"dnd-agent/pkg/handlers"
	"fmt"
)

func HandleCommand(w *domain.World, cmds []domain.Command) []domain.Command {
	if len(cmds) == 0 {
		return cmds
	}

	command := cmds[0]
	rest := cmds[1:]

	cmdName, _ := command["command"].(string)

	switch cmdName {
	case "/load":
		return handlers.Load(w, command, rest)
	case "/save":
		return handlers.Save(w, command, rest)
	case "/start":
		return handlers.Start(w, command, rest)
	case "Character.create":
		return handlers.CharacterCreate(w, command, rest)
	case "Character.On.create":
		return handlers.CharacterOnCreate(w, command, rest)
	case "Unit.spawn":
		return handlers.UnitSpawn(w, command, rest)
	case "/map":
		return handlers.Map(w, command, rest)
	default:
		fmt.Printf("Неизвестная команда: %v\n", cmdName)
		return rest
	}
}
