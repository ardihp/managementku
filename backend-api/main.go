package main

import (
	"flag"
	"fmt"
	"kanban-board/backend-api/config"
	"kanban-board/backend-api/database"
	seeds "kanban-board/backend-api/database/seeds"
	"kanban-board/backend-api/routes"

	socketController "kanban-board/backend-api/controllers/socket"
)

func main() {
	config.LoadEnv()

	database.InitDB()

	tableFlag := flag.String("table", "", "specify the table")
	flag.Parse()
	table := *tableFlag

	if table != "" {
		switch {
		case table == "users":
			seeds.UserSeeds()
		default:
			fmt.Println("your argument not found")
		}
		return
	}

	go socketController.Manager.Start()

	r := routes.SetupRouter()

	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
