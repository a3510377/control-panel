package main

import (
	"embed"
	"fmt"

	"github.com/a3510377/control-panel/utils/system"
	_ "github.com/joho/godotenv/autoload" // auto load env
)

//go:embed all:dist
var webBuild embed.FS

func main() {
	fmt.Printf("%+v\n", system.GetNowMemInfo())
	fmt.Printf("%+v\n", system.GetNowCPUInfo())
	fmt.Printf("%+v\n", system.GetNowCPUUsage())
	// db, _ := database.NewDB("test.db")

	// for _, instance := range db.GetAutoStartInstances() {
	// 	go instance.Run()
	// }

	// ser := server.New(db)

	// /* init frontend -- start */
	// dir, err := fs.Sub(webBuild, "dist")
	// if err != nil {
	// 	panic(err)
	// }

	// ser.AddFileHandler(dir)
	// /* init frontend -- end */

	// ser.Start()
}
