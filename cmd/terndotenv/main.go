package main

import (
	"fmt"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/db/migrations/",
		"--config",
		"./internal/db/migrations/tern.conf",
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Command execution faleid:", err)
		fmt.Println("Output:", string(output))
		panic(err)
	}

	fmt.Println("Command executed successfully: ", string(output))
}
