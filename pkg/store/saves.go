package store

import (
	"dnd-agent/pkg/domain"
	worldzone "dnd-agent/pkg/world-zone"
	"encoding/json"
	"fmt"
	"os"
)

// SaveWorldToJSON сохраняет состояние мира в JSON-файл.
func SaveWorldToJSON(w *domain.World, path string) error {
	data, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка сериализации мира: %v", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("ошибка записи файла %s: %v", path, err)
	}

	return nil
}

// LoadWorldFromJSON считывает состояние мира из JSON-файла.
func LoadWorldFromJSON(path string) (*domain.World, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла %s: %v", path, err)
	}

	var w domain.World
	if err := json.Unmarshal(data, &w); err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %v", err)
	}

	// гарантируем, что карты инициализированы
	if w.Units == nil {
		w.Units = make(map[int]*domain.Unit)
	}
	if w.Zones == nil {
		w.Zones = make(map[string]*worldzone.Level)
	}

	return &w, nil
}
