package store

import (
	"dnd-agent/pkg/domain"
	worldzone "dnd-agent/pkg/world-zone"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func LoadAllMaps(w *domain.World) error {
	dir := "./maps"

	entries, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("ошибка чтения директории %s: %v", dir, err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		// фильтруем только .xml файлы
		if !strings.HasSuffix(entry.Name(), ".xml") {
			continue
		}

		path := filepath.Join(dir, entry.Name())
		file, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("ошибка чтения файла %s: %v", path, err)
		}

		level, err := worldzone.LoadLevelFromXML(string(file))
		if err != nil {
			return fmt.Errorf("ошибка загрузки уровня из %s: %v", path, err)
		}

		// убираем расширение для ключа карты
		name := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
		w.Zones[name] = level
	}

	return nil
}
