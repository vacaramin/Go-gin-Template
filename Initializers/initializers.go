package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func envInit() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// to access environment variables use os.Getenv("VAR_NAME")
}
func Init() {
	envInit()
	ConnectToDb()

}
