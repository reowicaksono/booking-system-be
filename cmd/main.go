package main

import (
	"booking-system/cmd/app"
	"context"
)

func main() {
	// Initialize the application
	app := app.App{}

	app.DBInit()
	app.RouterInit(context.Background())
	app.Run()
}
