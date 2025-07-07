package main

import (
	"flag"
	"fmt"
	"kanban-board/backend-api/config"
	"kanban-board/backend-api/database"
	seeds "kanban-board/backend-api/database/seeds"
	"kanban-board/backend-api/routes"
)

func main() {
	config.LoadEnv()

	database.InitDB()

	tableFlag := flag.String("table", "all-table", "specify the table")
	flag.Parse()
	table := *tableFlag

	if table != "" {
		switch {
		case table == "statuses":
			seeds.WorkspaceStatusSeed()
		case table == "types":
			seeds.WorkspaceTypeSeed()
		case table == "all-table":
			seeds.WorkspaceTypeSeed()
			seeds.WorkspaceStatusSeed()
		default:
			fmt.Println("your argument not found")
		}
		return
	}

	r := routes.SetupRouter()

	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
