package main

import (
	"bufio"
	characterCreation "dnd-agent/pkg/character-creation"
	"dnd-agent/pkg/domain"
	luaUtils "dnd-agent/pkg/lua-utils"
	worldzone "dnd-agent/pkg/world-zone"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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

	case "/load":
		args, _ := command["args"].([]string)
		file := args[0]

		save, err := LoadWorldFromJSON("temp/saves/" + file + ".json")
		if err != nil {
			fmt.Println(err)
			return rest
		}
		*w = *save
		return []domain.Command{}

	case "/save":
		args, _ := command["args"].([]string)
		file := args[0]

		err := SaveWorldToJSON(w, "temp/saves/"+file+".json")
		if err != nil {
			fmt.Println(err)
			return rest
		}
		return rest
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
			XMLName: xml.Name{
				Local: "objects",
				Space: "",
			},
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

func initializeWorld() *domain.World {
	w := &domain.World{
		Units: make(map[int]*domain.Unit),
		Zones: make(map[string]*worldzone.Level),
	}

	err := loadAllMaps(w)
	if err != nil {
		panic(err)
	}
	return w
}

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

func main() {
	world := initializeWorld()

	queue := []domain.Command{}

	fmt.Println("DnDAI запущен")
	reader := bufio.NewReader(os.Stdin)

	for {
		// Пока есть команды в очереди — обрабатываем их
		for len(queue) > 0 {
			queue = HandleCommand(world, queue)
		}

		// Очередь пуста — ждём ввод пользователя
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		line = strings.TrimSpace(line)
		if line == "" {
			fmt.Println("Ошибка ввода: пустая строка", err)
			continue
		}
		command, args := ParseCommandLine(line)

		if command == "/quit" {
			fmt.Println("Завершение работы.")
			break
		}

		// TODO: сделать универсальный парсинг в команды
		queue = append(queue, domain.Command{"command": command, "args": args})
	}
}
