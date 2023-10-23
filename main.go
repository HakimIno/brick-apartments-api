package main

import (
	"apartments-api/routes"

	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
)

func main() {
	godotenv.Load()

	app := iris.Default()

	location := app.Party("/brick-api/location")
	{
		location.Get("/autocomplete", routes.Autocomplete)
		location.Get("/search")
	}

	app.Listen(":4000")
}
