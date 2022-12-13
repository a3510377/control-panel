package start

import (
	"github.com/a3510377/control-panel/database"
	"github.com/a3510377/control-panel/server"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start the control plane",
		Long:  "This command starts the control plane.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			db, _ := database.NewDB("test.db")
			server.New().Start(db)
		},
	}
}
