package handlers

import (
	"dnd-agent/pkg/domain"
	"dnd-agent/pkg/store"
	"fmt"
)

func Save(w *domain.World, command domain.Command, rest []domain.Command) []domain.Command {
	args, _ := command["args"].([]string)
	file := args[0]

	err := store.SaveWorldToJSON(w, "temp/saves/"+file+".json")
	if err != nil {
		fmt.Println(err)
		return rest
	}
	return rest
}
