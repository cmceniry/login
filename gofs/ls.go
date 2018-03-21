package gofs

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tag::cmd[]
var lsCmd = &cobra.Command{
	Use:   "ls [path to ls]",
	Short: "Shows the directory contents of a path",
	Long:  "Show the directory contents of a path. If given no path, uses the running path for the gofs server.",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("ls server='%s' path='%s'\n", serverAddress, args[0])
	},
}

// end::cmd[]
