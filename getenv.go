package linkycore

import (
	"os"
	"strconv"
)

// GetEnv - Get env variable as string
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// GetBoolEnv - Get an boolean env var. This returns false to invalid values
func GetBoolEnv(key string, fallback bool) bool {
	if value, ok := os.LookupEnv(key); ok {
		boolV, err := strconv.ParseBool(value)
		if err == nil {
			return boolV
		}

		return false
	}

	return fallback
}

//GetIntEnv - Get and int env var. 
func GetenvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		intEnv, err := strconv.Atoi(value)
		if err == nil {
			return intEnv
		}
	}
	return fallback
}
