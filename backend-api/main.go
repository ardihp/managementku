package main

import (
	"flag"
	"kanban-board/backend-api/config"
	"kanban-board/backend-api/database"
	seeds "kanban-board/backend-api/database/seeds"
	"kanban-board/backend-api/routes"

	socketController "kanban-board/backend-api/controllers/socket"
)

func main() {
	config.LoadEnv()

	database.InitDB()

	seedFlag := flag.Bool("seed", false, "Run seeders to database")
	migrationFlag := flag.Bool("migrate", false, "Run migration to database")
	flag.Parse()

	seed := *seedFlag
	migration := *migrationFlag

	if migration {
		database.MigrateDB()
		return
	}

	if seed {
		seeds.UserSeeds()
		return
	}

	go socketController.Manager.Start()

	r := routes.SetupRouter()

	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
