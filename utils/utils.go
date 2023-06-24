package utils

import (
	"net/http"
	"net/url"
	"os"
)

func UseEnvOrDefault(env, def string) string {
	env, ok := os.LookupEnv(env)
	if !ok {
		return def
	}

	return env
}

func GetDebugClient() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				return url.Parse("http://192.168.100.167:9090") //this sshould be dynamic based on the proxyman url
			},
		},
	}
}
