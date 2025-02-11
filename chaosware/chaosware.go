package chaosware

import (
	"fmt"
	"math/rand"
	"net/http"
)

type ChaosWare struct {
	settings *Settings
}

func NewDefaultChaosMiddleware() *ChaosWare {
	c := &ChaosWare{settings: &Settings{}}
	c.readSettingsFromEnv()

	return c
}

func NewChaosMiddleware(settings *Settings) (ChaosWare, error) {
	if err := validateSettings(settings); err != nil {
		return ChaosWare{}, fmt.Errorf("chaosware: failed to load settings: %v", err)
	}

	return ChaosWare{settings: settings}, nil
}

func (c *ChaosWare) ChaosHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.chaos(w, r)
		next.ServeHTTP(w, r)
	})
}

func (c *ChaosWare) chaos(w http.ResponseWriter, r *http.Request) {
	if c.settings.PanicChance > 0 {
		if rand.Intn(100) < c.settings.PanicChance {
			panic("chaosware: controlled panic")
		}
	}
	if c.settings.FreezeChance > 0 {
		if rand.Intn(100) < c.settings.FreezeChance {
			fmt.Println("chaosware: infinite freeze")
			select {} // Block forever without eating up cpu
		}
	}
}
