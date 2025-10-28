package worldzone

import (
	"strings"
)

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

// RenderLevel — строит карту и возвращает строковое представление
func RenderLevel(lvl *Level) string {
	lvlMap := createMap(lvl)
	return renderMap(lvlMap)
}
