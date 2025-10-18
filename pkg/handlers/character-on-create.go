package handlers

import (
	"dnd-agent/pkg/domain"
	luaUtils "dnd-agent/pkg/lua-utils"
	"fmt"
)

func CharacterOnCreate(w *domain.World, command domain.Command, rest []domain.Command) []domain.Command {
	// Извлекаем данные из команды
	characterOnCreate, _ := command["characterOnCreate"].(map[string]interface{})
	unitID, _ := characterOnCreate["unitId"].(float64) // JSON числа → float64

	next, err := luaUtils.CallLuaHandler(
		"lua/scenario/init.lua",
		"Character.On.create",
		map[string]interface{}{
			"next": []domain.Command{},
			"characterOnCreate": map[string]interface{}{
				"unitId": unitID,
			},
		},
	)
	if err != nil {
		fmt.Println(err)
		return rest
	}

	rest = append(rest, next...)
	return rest
}
