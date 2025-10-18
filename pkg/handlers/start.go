package handlers

import (
	"dnd-agent/pkg/domain"
	luaUtils "dnd-agent/pkg/lua-utils"
	"dnd-agent/pkg/store"
	worldzone "dnd-agent/pkg/world-zone"
	"fmt"
)

func initializeWorld() *domain.World {
	w := &domain.World{
		Units: make(map[int]*domain.Unit),
		Zones: make(map[string]*worldzone.Level),
	}

	err := store.LoadAllMaps(w)
	if err != nil {
		panic(err)
	}
	return w
}

func Start(w *domain.World, command domain.Command, rest []domain.Command) []domain.Command {
	world := initializeWorld()
	*w = *world

	next, err := luaUtils.CallLuaHandler(
		"lua/scenario/init.lua",
		"/start",
		map[string]interface{}{
			"next": []domain.Command{},
		},
	)
	if err != nil {
		fmt.Println(err)
		return rest
	}

	rest = append(rest, next...)
	return rest
}
