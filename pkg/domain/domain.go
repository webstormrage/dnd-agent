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

type UnitSpawnCommand struct {
	UnitId     int    `json:"unitId"`
	GameZoneId string `json:"gameZoneId"`
	X          int    `json:"x"`
	Y          int    `json:"y"`
	Owner      string `json:"owner"`
}

type CharacterOnCreateCommand struct {
	UnitId int `json:"unitId"`
}

type Command struct {
	Command           string                    `json:"command"`
	UnitSpawn         *UnitSpawnCommand         `json:"unitSpawn,omitempty"`
	CharacterOnCreate *CharacterOnCreateCommand `json:"characterOnCreate,omitempty"`
}

type Choice struct {
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	Options []string `json:"options"`
}
