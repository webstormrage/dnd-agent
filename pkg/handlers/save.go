package handlers

import (
	"dnd-agent/pkg/domain"
	"dnd-agent/pkg/store"
	"fmt"
)

func Save(w *domain.World, command *domain.Command) {
	argv, _ := command.Args["argv"].([]string)
	file := argv[0]

	err := store.SaveWorldToJSON(w, "temp/saves/"+file+".json")
	domain.Resolve(command)
	if err != nil {
		fmt.Println(err)
	}
}
