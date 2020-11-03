package linkycore

import (
	"net/http"
	"strconv"
)

type LinkyCoreOptions struct {
	LogMode         bool
	DbURI           string
	CacheDBAddr     string
	CacheDBPassword string
	CacheDBIndex    int
}

func Init() {
	opts := LinkyCoreOptions{
		LogMode:         false,
		DbURI:           GetEnv("DB_URI", "project:project@(database)/project"),
		CacheDBAddr:     GetEnv("CACHE_DB_ADDR", "localhost:6379"),
		CacheDBPassword: GetEnv("CACHE_DB_PASSWORD", ""),
		CacheDBIndex:    0,
	}

	envLogMode := GetEnv("LOG_MODE", "")
	if envLogMode != "" {
		opts.LogMode = true
	}

	cacheDBIndex, _ := strconv.Atoi(GetEnv("CACHE_DB_DB", "0"))
	if cacheDBIndex != 0 {
		opts.CacheDBIndex = cacheDBIndex
	}

	HttpClient = &http.Client{}

	InitSanitizer(&opts)
	InitDB(&opts)
	InitCache(&opts)
}
