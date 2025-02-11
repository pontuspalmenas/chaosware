package chaosware

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Settings struct {
	PanicChance  int // A chance between 0-100 representing probability in percent of how likely a panic will be.
	FreezeChance int // A chance between 0-100 representing probability in percent of how likely an infinite freeze will be.
}

func (c *ChaosWare) readSettingsFromEnv() {
	envs := os.Environ()
	for _, env := range envs {
		if strings.HasPrefix(env, "CHAOSW_") {
			parts := strings.SplitN(env, "=", 2)
			switch parts[0] {
			case "CHAOSW_PANIC_CHANCE":
				i, err := parseChance(parts[1])
				if err != nil {
					failParse(parts[0], err)
					continue
				}
				c.settings.PanicChance = i
			case "CHAOSW_FREEZE_CHANCE":
				i, err := parseChance(parts[1])
				if err != nil {
					failParse(parts[0], err)
					continue
				}
				c.settings.FreezeChance = i
			}
		}
	}
}

func failParse(key string, err error) {
	fmt.Printf("chaosware: failed to parse env, skipping %s: %s\n", key, err.Error())
}

func parseChance(v string) (int, error) {
	i, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}
	if i < 0 || i > 100 {
		return 0, fmt.Errorf("value out of range (0-100): %d", i)
	}
	return i, nil
}

func validateSettings(settings *Settings) error {
	if settings.PanicChance < 0 || settings.PanicChance > 100 {
		return fmt.Errorf("invalid panic chance: %d", settings.PanicChance)
	}

	return nil
}
