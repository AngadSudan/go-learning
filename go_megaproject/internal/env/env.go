package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	value, isFound := os.LookupEnv(key)

	if !isFound {
		return fallback
	}

	return value
}

func GetInt(key string, fallback int) int {
	value, isFound := os.LookupEnv(key)

	if !isFound {
		return fallback
	}

	parsedValue, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsedValue
}
