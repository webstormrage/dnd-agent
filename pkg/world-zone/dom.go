package worldzone

import (
	"encoding/xml"
	"strconv"
	"strings"
)

// ====== DOM-структуры ======

type Level struct {
	XMLName xml.Name `xml:"level"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
	Rooms   []Room   `xml:"room"`
	Objects []Object `xml:"object"`
	Areas   []Object `xml:"area"`
}

type Room struct {
	XMLName xml.Name `xml:"room"`
	X       int      `xml:"x,attr"`
	Y       int      `xml:"y,attr"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
}

type Object struct {
	XMLName xml.Name `xml:"object"`
	X       int      `xml:"x,attr"`
	Y       int      `xml:"y,attr"`
	Type    string   `xml:"type,attr"`
	Width   int      `xml:"width,attr"`
	Height  int      `xml:"height,attr"`
	Name    string   `xml:"name,attr"`
	To      string   `xml:"to,attr"`
}

func (lvl *Level) Load(data string) error {
	if err := xml.Unmarshal([]byte(data), lvl); err != nil {
		return err
	}
	return nil
}

func unitId2Name(id int) string {
	return "unit#" + strconv.Itoa(id)
}

func name2UnitId(name string) int {
	const prefix = "unit#"
	if !strings.HasPrefix(name, prefix) {
		return 0
	}
	idStr := strings.TrimPrefix(name, prefix)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0
	}
	return id
}

func (lvl *Level) SpawnUnit(unitId int, x int, y int) {
	if lvl.FindUnit(unitId) != nil {
		return
	}
	lvl.Objects = append(lvl.Objects, Object{
		XMLName: xml.Name{
			Local: "objects",
			Space: "",
		},
		Type:   "unit",
		X:      x,
		Y:      y,
		Width:  1,
		Height: 1,
		Name:   unitId2Name(unitId),
	})
}

func (lvl *Level) FindUnit(unitId int) *Object {
	unitName := unitId2Name(unitId)
	for i := range lvl.Objects {
		if lvl.Objects[i].Name == unitName {
			return &lvl.Objects[i]
		}
	}
	return nil
}

func (lvl *Level) MoveUnit(unitId int, x int, y int) {
	unit := lvl.FindUnit(unitId)
	if unit == nil {
		return
	}
	unit.X = x
	unit.Y = y
}

func (lvl *Level) GetObjectByName(name string) *Object {
	for i := range lvl.Objects {
		if lvl.Objects[i].Name == name {
			return &lvl.Objects[i]
		}
	}
	return nil
}

func (lvl *Level) UnSpawn(unitId int) bool {
	unitName := unitId2Name(unitId)

	for i := range lvl.Objects {
		if lvl.Objects[i].Name == unitName {
			// Удаляем элемент из среза безопасно
			lvl.Objects = append(lvl.Objects[:i], lvl.Objects[i+1:]...)
			return true // успешно удалили
		}
	}

	return false // юнит не найден
}
