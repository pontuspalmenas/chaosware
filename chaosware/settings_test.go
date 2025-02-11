package chaosware

import "testing"

func TestChaosWare_readSettingsFromEnv(t *testing.T) {
	t.Setenv("CHAOSW_PANIC_CHANCE", "37")
	t.Setenv("CHAOSW_FREEZE_CHANCE", "18")

	c := &ChaosWare{settings: &Settings{}}
	c.readSettingsFromEnv()

	if c.settings.PanicChance != 37 {
		t.Errorf("c.settings.PanicChance should be 37, was %d", c.settings.PanicChance)
	}

	if c.settings.FreezeChance != 18 {
		t.Errorf("c.settings.FreezeChance should be 18, was %d", c.settings.FreezeChance)
	}
}
