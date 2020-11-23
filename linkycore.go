package linkycore

import (
	"net/http"
	"strconv"
)

// LinkyCoreOptions - Linky core configuration options
type LinkyCoreOptions struct {
	LogMode         bool
	DbURI           string
	CacheDBAddr     string
	CacheDBPassword string
	CacheDBIndex    int
}

// Init -
func Init() {
	opts := LinkyCoreOptions{
		LogMode:         false,
		DbURI:           GetEnv("DB_URI", "project:project@(database)/project"),
		CacheDBAddr:     GetEnv("CACHE_DB_ADDR", "localhost:6379"),
		CacheDBPassword: GetEnv("CACHE_DB_PASSWORD", ""),
		CacheDBIndex:    0,
	}

	opts.LogMode = GetBoolEnv("LOG_MODE", true)

	cacheDBIndex, _ := strconv.Atoi(GetEnv("CACHE_DB_DB", "0"))
	if cacheDBIndex != 0 {
		opts.CacheDBIndex = cacheDBIndex
	}

	HttpClient = &http.Client{}

	InitSanitizer(&opts)
	InitDB(&opts)
	InitCache(&opts)
}
