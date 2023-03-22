package utils

import (
	"net/http"
	"os"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func UseEnvOrDefault(env, def string) string {
	env, ok := os.LookupEnv(env)
	if !ok {
		return def
	}

	return env
}
