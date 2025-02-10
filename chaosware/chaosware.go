package chaosware

import (
	"fmt"
	"math/rand"
	"net/http"
)

type ChaosWare struct {
	settings *Settings
}

type Settings struct {
	PanicChance  int // A chance between 1-100 representing probability in percent of how likely a panic will be.
	PanicEnabled bool
}

func NewDefaultChaosMiddleware() *ChaosWare {
	c := &ChaosWare{}
	c.settings = &Settings{}
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
	if c.settings.PanicEnabled {
		fmt.Println("[trace] chaosware: enabled")
		if rand.Intn(100) < c.settings.PanicChance {
			fmt.Println("[trace] chaosware: panic")
			panic("chaosware: controlled panic")
		}
		fmt.Println("[trace] chaosware: no panic")
	}
}
