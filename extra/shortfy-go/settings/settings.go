package settings

import "github.com/joho/godotenv"

func Config() {

	// Load Environment Variables
	if err := godotenv.Load(); err != nil {
		panic("Could not load .env file")
	}

}
