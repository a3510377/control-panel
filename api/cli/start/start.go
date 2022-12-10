package start

import (
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "start",
		Short: "Start the control plane",
		Long:  "This command starts the control plane.",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			// TODO: Start the control plane
		},
	}
}
