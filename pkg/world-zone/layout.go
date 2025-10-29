package worldzone

// TODO: делать классом layout
type Cell struct {
	X, Y   int
	Type   string
	Room   *Room
	Object *Object
}

// ====== Хелпер ======
func setCell(lvlMap [][]Cell, x, y int, cellType string, room *Room, obj *Object) {
	if isOutOfBounds(x, y, lvlMap) {
		return
	}
	lvlMap[x][y] = Cell{
		X:      x,
		Y:      y,
		Type:   cellType,
		Room:   room,
		Object: obj,
	}
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
	for x := 0; x < lvl.Width; x++ {
		lvlMap[x] = make([]Cell, lvl.Height)
		for y := 0; y < lvl.Height; y++ {
			setCell(lvlMap, x, y, "empty", nil, nil)
		}
	}
	return lvlMap
}

func isOutOfBounds(x, y int, lvlMap [][]Cell) bool {
	return x < 0 || x >= len(lvlMap) || y < 0 || y >= len(lvlMap[x])
}

func createRoomWalls(room *Room, lvlMap [][]Cell) {
	for x := room.X; x < room.X+room.Width; x++ {
		setCell(lvlMap, x, room.Y, "wall", room, nil)
		setCell(lvlMap, x, room.Y+room.Height-1, "wall", room, nil)
	}
	for y := room.Y; y < room.Y+room.Height; y++ {
		setCell(lvlMap, room.X, y, "wall", room, nil)
		setCell(lvlMap, room.X+room.Width-1, y, "wall", room, nil)
	}
}

func createWalls(rooms []Room, lvlMap [][]Cell) {
	for i := range rooms {
		createRoomWalls(&rooms[i], lvlMap)
	}
}

func createObject(obj *Object, lvlMap [][]Cell) {
	for x := obj.X; x < obj.X+obj.Width; x++ {
		for y := obj.Y; y < obj.Y+obj.Height; y++ {
			setCell(lvlMap, x, y, obj.Type, nil, obj)
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
