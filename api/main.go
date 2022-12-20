package main

import (
	"github.com/a3510377/control-panel/database"
	"github.com/a3510377/control-panel/server"
	_ "github.com/joho/godotenv/autoload" // auto load env
)

func main() {
	db, _ := database.NewDB("test.db")

	for _, instance := range db.GetAutoStartInstances() {
		go instance.Run()
	}
	server.New().Start(db)
}
