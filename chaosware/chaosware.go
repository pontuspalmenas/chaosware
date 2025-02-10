package chaosware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

type ChaosWare struct {
	settings Settings
}

type Settings struct {
	PanicChance int
}

func NewDefaultChaosMiddleware() ChaosWare {
	c := ChaosWare{}
	c.settings = Settings{}
	c.readSettingsFromEnv()

	return ChaosWare{}
}

func NewChaosMiddleware(settings Settings) ChaosWare {
	return ChaosWare{settings: settings}
}

func (c ChaosWare) readSettingsFromEnv() {
	envs := os.Environ()
	for _, env := range envs {
		if strings.HasPrefix(env, "CHAOSW_") {
			parts := strings.SplitN(env, "=", 2)
			fmt.Printf("%s=%s\n", parts[0], parts[1])
		}
	}
}

func (c ChaosWare) ChaosHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c.chaos(w, r)
		next.ServeHTTP(w, r)
	})
}

func (c ChaosWare) chaos(w http.ResponseWriter, r *http.Request) {

}
