package domain

import worldzone "dnd-agent/pkg/world-zone"

type Unit struct {
	Attributes interface{}
	Inventory  interface{}
	Equipment  interface{}
	States     interface{}
	ID         int
	ZoneId     *string
	Owner      string
}

type World struct {
	Units map[int]*Unit
	Zones map[string]*worldzone.Level
}

type Command struct {
	Procedure string
	Args      map[string]interface{}
	State     map[string]interface{}
	Stack
}

type Choice struct {
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	Options []string `json:"options"`
}
