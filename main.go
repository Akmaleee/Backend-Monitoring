package main

import (
	"it-backend/cmd"
	"it-backend/internal/helper"
)

func main() {
	helper.SetupLogger()

	cmd.ServeHTTP()
}
