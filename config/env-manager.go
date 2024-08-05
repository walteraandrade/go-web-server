package env_manager

import (
	"fmt"
	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:",
			err)
		return err
	}
	return nil
}
