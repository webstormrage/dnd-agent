package worldzone

import "container/list"

// FindNearestBFS выполняет поиск ближайшей ячейки от (x, y),
// удовлетворяющей переданному предикату.
// Возвращает путь до неё в виде []Cell.
func FindNearestBFS(level *Level, x, y int, predicate func(Cell) bool) []Cell {
	grid := createMap(level)
	w := len(grid)
	if w == 0 {
		return nil
	}

	if isOutOfBounds(x, y, grid) || !IsPathible(grid[x][y]) {
		return nil
	}

	dirs := [8][2]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	prev := make(map[*Cell]*Cell)

	start := &grid[x][y]
	prev[start] = start // ✅ закольцевали стартовую клетку
	queue := list.New()
	queue.PushBack(start)

	for queue.Len() > 0 {
		elem := queue.Front()
		queue.Remove(elem)
		cell := elem.Value.(*Cell)

		// 🟢 Проверяем предикат — если подходит, возвращаем путь
		if predicate(*cell) {
			return reconstructPathCells(prev, start, cell)
		}

		// Проверяем всех соседей (8 направлений)
		for _, d := range dirs {
			nx, ny := cell.X+d[0], cell.Y+d[1]
			if isOutOfBounds(nx, ny, grid) {
				continue
			}
			next := &grid[nx][ny]

			// если уже был посещён — пропускаем
			if _, seen := prev[next]; seen {
				continue
			}
			if !IsPathible(*next) {
				continue
			}

			prev[next] = cell
			queue.PushBack(next)
		}
	}

	// 🚫 Не нашли подходящую ячейку
	return nil
}

// reconstructPathCells восстанавливает путь из prev-таблицы.
func reconstructPathCells(prev map[*Cell]*Cell, start, goal *Cell) []Cell {
	path := []Cell{}
	cur := goal
	for {
		path = append(path, *cur)
		if cur == start {
			break
		}
		cur = prev[cur]
		if cur == nil {
			break // на случай некорректных prev
		}
	}

	// Переворачиваем путь (от старта к цели)
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
