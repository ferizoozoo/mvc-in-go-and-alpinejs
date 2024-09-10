package internal

import (
	"bufio"
	"os"
	"strings"
)

func LoadEnvironmentVariables() {
	file, err := os.Open(".env")
	if err != nil {
		panic("Could not open the .env file.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "=")
		key, value := strings.TrimSpace(line[0]), strings.TrimSpace(line[1])
		os.Setenv(key, value)
	}
}
