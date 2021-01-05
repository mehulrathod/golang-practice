package main

import (
	"HirebasePractice/mux/app"
	"HirebasePractice/mux/app/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
