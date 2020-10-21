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
		CacheDBAddr:     GetEnv("CACHE_DB_ADDR"),
		CacheDBPassword: GetEnv("CACHE_DB_PASSWORD"),
		CacheDBIndex:    strconv.Atoi(GetEnv("CACHE_DB_DB", "1")),
	}

	envLogMode := GetEnv("LOG_MODE")
	if envLogMode {
		opts.LogMode = true
	}

	HttpClient = &http.Client{}

	InitSanitizer(opts)
	InitDB(opts)
}
