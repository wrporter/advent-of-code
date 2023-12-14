package runegrid

import (
	"strings"
)

func String(grid [][]rune) string {
	var sb strings.Builder
	for _, row := range grid {
		sb.WriteString(string(row))
		sb.WriteByte('\n')
	}
	return sb.String()
}
