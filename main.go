package main

import (
	"github.com/ashkanamani/dummygame/cmd"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	cmd.Execute()
}
