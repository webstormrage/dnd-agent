package domain

type Unit struct {
	Attributes interface{}
	Inventory  interface{}
	Equipment  interface{}
	States     interface{}
	ID         int
	X          int
	Y          int
	ZoneId     *string
	Owner      string
}

type World struct {
	Units map[int]*Unit
}

type Command = map[string]interface{}

type Choice struct {
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	Options []string `json:"options"`
}
