package worldzone

import "strings"

func GetZoneTunnel(cell Cell) []string {
	if cell.Object == nil {
		return []string{}
	}
	return strings.Split(cell.Object.To, ".")
}
