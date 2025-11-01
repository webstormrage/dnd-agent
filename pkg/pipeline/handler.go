package pipeline

import (
	"dnd-agent/pkg/domain"
	"dnd-agent/pkg/handlers"
)

func HandleCommand(w *domain.World, command *domain.Command) {

	switch command.Procedure {
	case "/load":
		handlers.Load(w, command)
	case "/save":
		handlers.Save(w, command)
	case "/start":
		handlers.LuaHandle(w, command, "lua/scenario/index.lua", "/start")
	case "Unit.createCharacter":
		handlers.LuaHandle(w, command, "lua/unit-definition/create-character.lua", "Unit.createCharacter")
	case "Unit.addBase":
		handlers.LuaHandle(w, command, "lua/unit-definition/base/index.lua", "Unit.addBase")
	case "Unit.addCharacterName":
		handlers.LuaHandle(w, command, "lua/unit-definition/character/index.lua", "Unit.addCharacterName")
	case "Unit.addFighter_1":
		handlers.LuaHandle(w, command, "lua/unit-definition/classes/index.lua", "Unit.addFighter_1")
	case "Unit.addAbilities":
		handlers.LuaHandle(w, command, "lua/unit-definition/abilities/index.lua", "Unit.addAbilities")
	case "Unit.addBackground":
		handlers.LuaHandle(w, command, "lua/unit-definition/background/index.lua", "Unit.addBackground")
	case "Unit.addRace":
		handlers.LuaHandle(w, command, "lua/unit-definition/races/index.lua", "Unit.addRace")
	case "Unit.spawn":
		handlers.UnitSpawn(w, command)
	case "World.addUnit":
		handlers.WorldAddUnit(w, command)
	case "World.setPlayerCurrentUnit":
		handlers.SetPlayerCurrentUnit(w, command)
	case "option.scanf":
		handlers.OptionScanf(w, command)
	case "/map":
		handlers.Map(w, command)
	case "/go":
		handlers.Go(w, command)
	default:
		handlers.Default(w, command)
	}
}
