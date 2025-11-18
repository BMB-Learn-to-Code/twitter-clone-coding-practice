package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}

func GetInt(key string, fallback uint32) uint32 {
	val := GetString(key, "")

	if len(val) == 0 {
		return fallback
	}

	n, err := strconv.Atoi(val)

	if err != nil {
		return fallback
	}
	return uint32(n)
}

func GetBool(key string, fallback bool) bool {
	val := GetString(key, "")
	b, err := strconv.ParseBool(val)

	if err != nil {
		return fallback
	}

	return b
}
