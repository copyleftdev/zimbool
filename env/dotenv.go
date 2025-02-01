package env

import "github.com/joho/godotenv"

// LoadDotEnv reads a .env file and returns its key–value pairs.
func LoadDotEnv(filePath string) (map[string]string, error) {
	return godotenv.Read(filePath)
}
