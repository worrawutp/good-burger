package initializers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadDotEnv() {
	env := os.Getenv("GOOD_BURGER_ENV")
	if "" == env {
		env = "development"
	}
	godotenv.Load(".env." + env)
	godotenv.Load()
	fmt.Println(os.Getenv("DATABASE_URL"))
}
