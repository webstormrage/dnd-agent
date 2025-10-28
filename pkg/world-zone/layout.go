package worldzone

type Cell struct {
	Type   string
	Room   *Room
	Object *Object
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
