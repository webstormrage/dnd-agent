package main

import (
	characterCreation "dnd-agent/pkg/character-creation"
	"dnd-agent/pkg/domain"
	luaUtils "dnd-agent/pkg/lua-utils"
	worldzone "dnd-agent/pkg/world-zone"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func HandleCommand(w *domain.World, cmds []domain.Command) []domain.Command {
	if len(cmds) == 0 {
		return cmds
	}

	command := cmds[0]
	rest := cmds[1:]

	cmdName, _ := command["command"].(string)

	switch cmdName {

	case "/start":
		// Проверка — есть ли уже юнит под контролем игрока
		if _, exists := w.Units[1]; exists {
			return rest
		}

		next, err := luaUtils.CallLuaHandler(
			"lua/scenario/init.lua",
			"/start",
			map[string]interface{}{
				"next": []domain.Command{},
			},
		)
		if err != nil {
			fmt.Println(err)
			return rest
		}

		rest = append(rest, next...)
		return rest

	case "Character.create":
		// Создаём персонажа вручную (пока что)
		character := characterCreation.ScanCharacter()
		character.ID = 1
		w.Units[character.ID] = character

		// Добавляем команду Character.On.create
		next := domain.Command{
			"command": "Character.On.create",
			"characterOnCreate": map[string]interface{}{
				"unitId": float64(character.ID),
			},
		}
		rest = append(rest, next)
		return rest

	case "Character.On.create":
		// Извлекаем данные из команды
		characterOnCreate, _ := command["characterOnCreate"].(map[string]interface{})
		unitID, _ := characterOnCreate["unitId"].(float64) // JSON числа → float64

		next, err := luaUtils.CallLuaHandler(
			"lua/scenario/init.lua",
			"Character.On.create",
			map[string]interface{}{
				"next": []domain.Command{},
				"characterOnCreate": map[string]interface{}{
					"unitId": unitID,
				},
			},
		)
		if err != nil {
			fmt.Println(err)
			return rest
		}

		rest = append(rest, next...)
		return rest

	case "Unit.spawn":
		// Получаем spawn-данные
		spawnData, _ := command["spawn"].(map[string]interface{})

		unitID := int(spawnData["unitId"].(float64))
		x := int(spawnData["x"].(float64))
		y := int(spawnData["y"].(float64))
		owner, _ := spawnData["owner"].(string)
		gameZoneId, _ := spawnData["gameZoneId"].(string)

		unit := w.Units[unitID]
		unit.X = x
		unit.Y = y
		unit.Owner = owner
		unit.ZoneId = &gameZoneId
		zone := w.Zones[*unit.ZoneId]
		zone.Objects = append(zone.Objects, worldzone.Object{
			Type:   "unit",
			X:      unit.X,
			Y:      unit.Y,
			Width:  1,
			Height: 1,
			Name:   "unit#" + strconv.Itoa(unit.ID),
		})
		return rest

	case "/map":
		// Проверка — есть ли уже юнит под контролем игрока
		if _, exists := w.Units[1]; !exists {
			return rest
		}
		character := w.Units[1]
		// Проверка — заспавлен ли юнит
		if character.ZoneId == nil {
			return rest
		}
		fmt.Println(worldzone.RenderLevel(w.Zones[*character.ZoneId]))
		return rest
	default:
		fmt.Printf("Неизвестная команда: %v\n", cmdName)
		return rest
	}
}

func loadAllMaps(w *domain.World) error {
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
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("ошибка чтения файла %s: %v", path, err)
		}

		level, err := worldzone.LoadLevelFromXML(string(data))
		if err != nil {
			return fmt.Errorf("ошибка загрузки уровня из %s: %v", path, err)
		}

		// убираем расширение для ключа карты
		name := strings.TrimSuffix(entry.Name(), filepath.Ext(entry.Name()))
		w.Zones[name] = level
	}

	return nil
}

func main() {
	w := &domain.World{
		Units: make(map[int]*domain.Unit),
		Zones: make(map[string]*worldzone.Level),
	}

	err := loadAllMaps(w)
	if err != nil {
		panic(err)
	}

	queue := []domain.Command{}

	fmt.Println("DnDAI запущен")

	for {
		// Пока есть команды в очереди — обрабатываем их
		for len(queue) > 0 {
			queue = HandleCommand(w, queue)
		}

		// Очередь пуста — ждём ввод пользователя
		fmt.Print("> ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			// Если пустая строка — просто пропускаем
			if err.Error() == "unexpected newline" {
				continue
			}
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		if input == "/quit" {
			fmt.Println("Завершение работы.")
			break
		}

		// Добавляем введённую команду в очередь
		queue = append(queue, domain.Command{"command": input})
	}
}
