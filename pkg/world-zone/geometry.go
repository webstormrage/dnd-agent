package worldzone

func IsPathible(cell Cell) bool {
	switch cell.Type {
	case "door":
		return true
	case "unit":
		return true
	case "empty":
		return true
	default:
		return false
	}
}

func IsPlaceble(cell Cell) bool {
	return cell.Type == "empty"
}

// TODO: здесь должен быть учет размера
func Transition(unitId int, dx int, dy int, times int, level *Level) (int, []Cell) {
	layout := createMap(level)
	unit := level.FindUnit(unitId)
	x := unit.X
	y := unit.Y
	count := 0
	path := []Cell{}
	for i := 1; i <= times; i++ {
		cx := unit.X + dx*i
		cy := unit.Y + dy*i
		if cx < 0 || cx >= len(layout) || cy < 0 || cy >= len(layout[cx]) {
			break
		}
		// TODO: здесь должна быть проверка пересечения зон (например opportunity attack)
		cell := layout[cx][cy]
		path = append(path, cell)
		if IsPathible(cell) {
			break
		}
		if IsPlaceble(cell) {
			x = cx
			y = cy
		}
		count++
	}
	level.MoveUnit(unitId, x, y)
	return count, path
}

func GetChebyshevDistance(x1, y1, x2, y2 int) int {
	dx := x2 - x1
	if dx < 0 {
		dx = -dx
	}
	dy := y2 - y1
	if dy < 0 {
		dy = -dy
	}

	if dx > dy {
		return dx
	}
	return dy
}
