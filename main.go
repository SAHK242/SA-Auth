package main

import (
	"os"

	"auth/cmd"
)

func main() {
	if err := os.Setenv("TZ", "Asia/Saigon"); err != nil {
		panic(err)
	}

	cmd.Start()
}
