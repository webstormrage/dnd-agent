package worldzone

import "container/list"

type point struct{ X, Y int }

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

	if !IsPathible(grid[y][x]) {
		return nil
	}

	dirs := [8][2]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
		{1, 1}, {1, -1}, {-1, 1}, {-1, -1},
	}

	visited := make([][]bool, h)
	prev := make(map[point]*point)

	for i := range visited {
		visited[i] = make([]bool, w)
	}

	start := point{X: x, Y: y}
	queue := list.New()
	queue.PushBack(start)
	visited[y][x] = true

	for queue.Len() > 0 {
		elem := queue.Front()
		queue.Remove(elem)
		cell := elem.Value.(point)

		curCell := grid[cell.Y][cell.X]

		// 🟢 Проверяем предикат — если подходит, возвращаем путь
		if predicate(curCell) {
			return reconstructPath(grid, prev, start, cell)
		}

		// иначе — добавляем соседей
		for _, d := range dirs {
			nx, ny := cell.X+d[0], cell.Y+d[1]
			if nx < 0 || ny < 0 || nx >= w || ny >= h {
				continue
			}
			if visited[ny][nx] {
				continue
			}
			if !IsPathible(grid[ny][nx]) {
				continue
			}

			visited[ny][nx] = true
			prev[point{X: nx, Y: ny}] = &cell
			queue.PushBack(point{X: nx, Y: ny})
		}
	}

	// 🚫 Если ничего не нашли
	return nil
}

// reconstructPath восстанавливает путь от старта до найденной клетки
func reconstructPath(grid [][]Cell, prev map[point]*point, start, goal point) []Cell {
	path := []Cell{}
	cur := &goal
	for cur != nil {
		path = append(path, grid[cur.Y][cur.X])
		if *cur == start {
			break
		}
		cur = prev[*cur]
	}

	// переворачиваем путь
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}
