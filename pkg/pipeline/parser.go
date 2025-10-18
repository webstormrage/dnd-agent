package pipeline

import (
	"regexp"
	"strings"
)

func ParseCommandLine(line string) (string, []string) {
	line = strings.TrimSpace(line)
	if line == "" {
		return "", nil
	}

	// Регулярка: находит либо "текст в кавычках", либо обычные слова
	re := regexp.MustCompile(`"([^"]+)"|(\S+)`)
	matches := re.FindAllStringSubmatch(line, -1)

	if len(matches) == 0 {
		return "", nil
	}

	parts := make([]string, 0, len(matches))
	for _, m := range matches {
		if m[1] != "" {
			parts = append(parts, m[1]) // аргумент в кавычках
		} else {
			parts = append(parts, m[2]) // обычное слово
		}
	}

	cmd := parts[0]
	args := []string{}
	if len(parts) > 1 {
		args = parts[1:]
	}

	return cmd, args
}
