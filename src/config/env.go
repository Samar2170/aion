package config

import (
	"os"

	"github.com/joho/godotenv"
)

var DBFile string
var BaseDir string

var NasaAPIKey string
var FileupAPIKey string

var FileupUsername string

func init() {
	// currentFile, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }
	// BaseDir = filepath.Dir(filepath.Dir(currentFile))

	BaseDir = "/Users/samararora/Desktop/PROJECTS/aion"
	godotenv.Load(BaseDir + "/.env")

	DBFile = os.Getenv("DB_FILE")

	NasaAPIKey = os.Getenv("NASA_API_KEY")

	FileupAPIKey = os.Getenv("FILEUP_API_KEY")
	FileupUsername = os.Getenv("FILEUP_USERNAME")
}
