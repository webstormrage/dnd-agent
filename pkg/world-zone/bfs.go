package worldzone

import "container/list"

// FindNearestBFS выполняет поиск ближайшей ячейки от (x, y),
// удовлетворяющей переданному предикату.
// Возвращает путь до неё в виде []Cell.
func FindNearestBFS(level *Level, x, y int, predicate func(Cell) bool) []Cell {
	grid := createMap(level)
	h := len(grid)
	if h == 0 {
		return nil
	}
	w := len(grid[0])

	if isOutOfBounds(x, y, grid) || !IsPathible(grid[x][y]) {
		return nil
	}

	dirs := [8][2]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	visited := make([][]bool, w)
	prev := make(map[*Cell]*Cell)

	for i := range visited {
		visited[i] = make([]bool, h)
	}

	start := &grid[x][y]
	queue := list.New()
	queue.PushBack(start)
	visited[x][y] = true

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
			if visited[nx][ny] {
				continue
			}
			next := &grid[nx][ny]
			if !IsPathible(*next) {
				continue
			}

			visited[nx][ny] = true
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
	for cur != nil {
		path = append(path, *cur)
		if cur == start {
			break
		}
		cur = prev[cur]
	}

	// Переворачиваем путь (от старта к цели)
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
