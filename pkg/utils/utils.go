package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func ParseBody(r *http.Request, v interface{}) error {
	if r == nil {
		return fmt.Errorf("nil request")
	}
	if r.Body == nil {
		return fmt.Errorf("request body is empty")
	}
	defer r.Body.Close()
	if v == nil {
		return fmt.Errorf("nil interface")
	}
	return json.NewDecoder(r.Body).Decode(v)
}

// LoadEnv loads environment variables from the .env file.
func LoadEnv() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	//use ../.env because main.go inside /cmd
	err = godotenv.Load(filepath.Join(pwd, "../../.env"))
	if err != nil {
		fmt.Println("Error loading .env file. Using system environment variables instead.")
	}
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	fmt.Fprintln(os.Stderr, "Warning: environment variable", key, "is not set. Using", fallback)
	return fallback
}
