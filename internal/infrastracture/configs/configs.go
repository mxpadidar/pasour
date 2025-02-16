package configs

import (
	"os"
	"time"
)

type config struct {
	RootDir     string
	Debug       bool
	Secret      string
	JwtDuration time.Duration
}

var Configs = NewConfig()

func NewConfig() *config {
	rootDit, _ := os.Getwd()
	return &config{
		RootDir:     rootDit,
		Debug:       getEnv("DEBUG", "flase") == "true",
		Secret:      getEnv("SECRET_KEY", "CHANGE-ME"),
		JwtDuration: getJwtDuration("1h"),
	}
}
