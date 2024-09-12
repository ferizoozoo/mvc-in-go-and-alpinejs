package internal

import (
	"fmt"
	"os/exec"
)

// NOTE: This is not complete and should be tested and developed later!
const command = "swag init"

func GenerateSwaggerFiles() {
	_, err := exec.Command(command).Output()
	if err != nil {
		fmt.Println("Error generating swagger files.\n", err)
	}
}
