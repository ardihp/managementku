package main

import (
	"kanban-board/backend-api/config"
	"kanban-board/backend-api/database"
	"kanban-board/backend-api/routes"
)

func main() {
	config.LoadEnv()

	database.InitDB()

	r := routes.SetupRouter()

	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
