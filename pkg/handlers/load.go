package handlers

import (
	"dnd-agent/pkg/domain"
	"dnd-agent/pkg/store"
	"fmt"
)

func Load(w *domain.World, command domain.Command, rest []domain.Command) []domain.Command {
	args, _ := command["args"].([]string)
	file := args[0]

	save, err := store.LoadWorldFromJSON("temp/saves/" + file + ".json")
	if err != nil {
		fmt.Println(err)
		return rest
	}
	*w = *save
	return []domain.Command{}
}
