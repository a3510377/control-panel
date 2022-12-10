package cli

import (
	"fmt"
	"os"

	"github.com/a3510377/control-panel/cli/start"
	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	root := &cobra.Command{
		Use:     "control-plane",
		Short:   "control-plane is a control plane for the service mesh",
		Example: fmt.Sprintf("  %s <%s> [%s...]", os.Args[0], "command", "flags"),
	}

	createCliCommandTree(root)

	return root
}

func createCliCommandTree(cmd *cobra.Command) {
	cmd.AddCommand(start.NewCommand())
}
