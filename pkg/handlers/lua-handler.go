package handlers

import (
	"dnd-agent/pkg/domain"
	luaUtils "dnd-agent/pkg/lua-utils"
	"fmt"
)

func LuaHandle(_ *domain.World, command *domain.Command, script string, handler string) {
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
