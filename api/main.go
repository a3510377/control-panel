package main

import (
	"fmt"

	"github.com/a3510377/control-panel/service"
)

func main() {
	fmt.Println(service.NewSummonID().Generate().Base2())
	// se := node.Instance{ProcessArgs: []string{"java", "-jar", "server.jar"}, Root: "test"}
	// se.Run()
	// if err := cli.NewCommand().Execute(); err != nil {
	// 	os.Exit(1)
	// }
}
