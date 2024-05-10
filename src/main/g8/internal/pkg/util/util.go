package util

import "os"

func GetENV() string {
	env := os.Getenv("ENV")
	if env == "" {
		return "DEVELOPMENT"
	}

	return env
}
