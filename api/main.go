package main

import "github.com/a3510377/control-panel/node"

func main() {
	se := node.Instance{ProcessArgs: []string{"java", "-jar", "server.jar"}}
	se.Run()
	// if err := cli.NewCommand().Execute(); err != nil {
	// 	os.Exit(1)
	// }
}
