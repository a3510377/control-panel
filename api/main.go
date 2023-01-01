package main

import (
	"embed"
	"math/rand"
	"time"

	"github.com/a3510377/control-panel/database"
	"github.com/a3510377/control-panel/server"
	_ "github.com/joho/godotenv/autoload" // auto load env
)

//go:embed all:dist
var webBuild embed.FS

func main() {
	rand.Seed(time.Now().UnixNano())

	db, _ := database.NewDB("test.db")

	for _, instance := range db.GetAutoStartInstances() {
		go instance.Run()
	}

	ser := server.New(db)

	ser.AddFileHandler(webBuild)

	ser.Start()
}
