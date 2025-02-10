package chaosware

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (c ChaosWare) readSettingsFromEnv() {
	envs := os.Environ()
	for _, env := range envs {
		if strings.HasPrefix(env, "CHAOSW_") {
			parts := strings.SplitN(env, "=", 2)
			switch parts[0] {
			case "CHAOSW_PANIC_CHANCE":
				i, err := parsePanicChance(parts[1])
				if err != nil {
					failParse(parts[1], err)
					continue
				}
				c.settings.PanicChance = i

			case "CHAOSW_PANIC_ENABLED":
				b, err := parseBool(parts[1])
				if err != nil {
					failParse(parts[1], err)
					continue
				}
				c.settings.PanicEnabled = b
			}
		}
	}
}

func failParse(key string, err error) {
	fmt.Printf("Failed to parse env, skipping %s: %s\n", key, err.Error())
}

func parsePanicChance(v string) (int, error) {
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}
	if i < 1 || i > 100 {
		return 0, fmt.Errorf("value out of range (1-100): %d", i)
	}
	return i, nil
}

func parseBool(s string) (bool, error) {
	v := strings.ToLower(s)
	if v != "true" && v != "false" {
		return false, fmt.Errorf("invalid boolean: %s", s)
	}
	return v == "true", nil
}

func validateSettings(settings Settings) error {
	if settings.PanicChance <= 0 || settings.PanicChance > 100 {
		return fmt.Errorf("invalid panic chance: %d", settings.PanicChance)
	}

	return nil
}
