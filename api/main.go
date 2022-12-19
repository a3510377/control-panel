package main

import (
	"github.com/a3510377/control-panel/database"
	"github.com/a3510377/control-panel/server"
)

func main() {
	db, _ := database.NewDB("test.db")

	// fmt.Println(db.GetAutoStartInstances())

	for _, instance := range db.GetAutoStartInstances() {
		instance.Run()
	}
	server.New().Start(db)
}
