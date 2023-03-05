package main

import (
	"tryFiberGo/database/migration"
	"tryFiberGo/route"

	"tryFiberGo/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.DatabaseInit()
	migration.RunMigration()

    app := fiber.New()

	//INITIAL ROUTE
	route.RouteInit(app)

	app.Listen(":8080")
}


