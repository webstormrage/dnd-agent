package handlers

import (
	"dnd-agent/pkg/domain"
	luaUtils "dnd-agent/pkg/lua-utils"
	"fmt"
)

func Start(_ *domain.World, command *domain.Command) {
	err := luaUtils.CallLuaHandler(
		"lua/scenario/index.lua",
		"/start",
		command,
	)
	if err != nil {
		fmt.Println(err)
		domain.Resolve(command)
	}
}
