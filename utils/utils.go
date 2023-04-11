package utils

import (
	"os"
)

func UseEnvOrDefault(env, def string) string {
	env, ok := os.LookupEnv(env)
	if !ok {
		return def
	}

	return env
}
