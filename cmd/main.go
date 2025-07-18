package main

import "booking-system/cmd/app"

func main() {
	// Initialize the application
	app := app.App{}

	app.DBInit()
	app.RouterInit()
	app.Run()
}
