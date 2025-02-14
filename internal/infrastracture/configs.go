package infrastracture

import (
	"log"
	"os"
	"pasour/internal/infrastracture/utils"

	"github.com/joho/godotenv"
)

type config struct {
	RootDir string
	Debug   bool
}

var Configs = LoadConfigs()

func LoadConfigs() *config {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env does not exists in root directory.")
	}

	rootDir, _ := os.Getwd()
	debug, err := utils.GetEnv("DEBUG", false)
	if err != nil {
		log.Fatal(err)
	}
	return &config{
		RootDir: rootDir,
		Debug:   debug,
	}
}
