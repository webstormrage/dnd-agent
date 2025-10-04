package worldzone

import (
	"encoding/xml"
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

type Cell struct {
	Type   string
	Room   *Room
	Object *Object
}

// ====== Основные публичные методы ======

// LoadLevelFromXML — читает XML-строку и создаёт структуру Level
func LoadLevelFromXML(data string) (*Level, error) {
	var lvl Level
	if err := xml.Unmarshal([]byte(data), &lvl); err != nil {
		return nil, err
	}
	return &lvl, nil
}

// RenderLevel — строит карту и возвращает строковое представление
func RenderLevel(lvl *Level) string {
	lvlMap := createMap(lvl)
	return renderMap(lvlMap)
}

// ====== Создание карты ======

func createMap(lvl *Level) [][]Cell {
	lvlMap := createLevel(lvl)
	createWalls(lvl.Rooms, lvlMap)
	for _, t := range []string{"door", "furnace", "kip", "table", "barrel", "boundary", "path-blocker", "unit"} {
		createObjects(lvl.Objects, t, lvlMap)
	}
	return lvlMap
}

func createLevel(lvl *Level) [][]Cell {
	lvlMap := make([][]Cell, lvl.Width)
	for i := 0; i < lvl.Width; i++ {
		lvlMap[i] = make([]Cell, lvl.Height)
		for j := 0; j < lvl.Height; j++ {
			lvlMap[i][j] = Cell{Type: "empty"}
		}
	}
	return lvlMap
}

func isOutOfBounds(x, y int, lvlMap [][]Cell) bool {
	return x < 0 || x >= len(lvlMap) || y < 0 || y >= len(lvlMap[x])
}

func createRoomWalls(room *Room, lvlMap [][]Cell) {
	for i := room.X; i < room.X+room.Width; i++ {
		if !isOutOfBounds(i, room.Y, lvlMap) {
			lvlMap[i][room.Y] = Cell{Type: "wall", Room: room}
		}
		if !isOutOfBounds(i, room.Y+room.Height-1, lvlMap) {
			lvlMap[i][room.Y+room.Height-1] = Cell{Type: "wall", Room: room}
		}
	}
	for j := room.Y; j < room.Y+room.Height; j++ {
		if !isOutOfBounds(room.X, j, lvlMap) {
			lvlMap[room.X][j] = Cell{Type: "wall", Room: room}
		}
		if !isOutOfBounds(room.X+room.Width-1, j, lvlMap) {
			lvlMap[room.X+room.Width-1][j] = Cell{Type: "wall", Room: room}
		}
	}
}

func createWalls(rooms []Room, lvlMap [][]Cell) {
	for i := range rooms {
		createRoomWalls(&rooms[i], lvlMap)
	}
}

func createObject(obj *Object, lvlMap [][]Cell) {
	for i := obj.X; i < obj.X+obj.Width; i++ {
		for j := obj.Y; j < obj.Y+obj.Height; j++ {
			if !isOutOfBounds(i, j, lvlMap) {
				lvlMap[i][j] = Cell{Type: obj.Type, Object: obj}
			}
		}
	}
}

func createObjects(objects []Object, objectType string, lvlMap [][]Cell) {
	for i := range objects {
		if objects[i].Type == objectType {
			createObject(&objects[i], lvlMap)
		}
	}
}

// ====== Рендер ======

func renderMap(lvlMap [][]Cell) string {
	var sb strings.Builder
	for y := 0; y < len(lvlMap[0]); y++ {
		for x := 0; x < len(lvlMap); x++ {
			sb.WriteString(renderCell(&lvlMap[x][y]))
			if x < len(lvlMap)-1 {
				sb.WriteString("  ")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func renderCell(cell *Cell) string {
	switch cell.Type {
	case "empty":
		return "."
	case "wall":
		return "w"
	case "door":
		return "d"
	case "kip":
		return "k"
	case "table":
		return "t"
	case "barrel":
		return "b"
	case "furnace":
		return "f"
	case "boundary":
		return "~"
	case "path-blocker":
		return " "
	case "unit":
		return "u"
	default:
		return "?"
	}
}
