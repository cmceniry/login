package gofs

import (
	"fmt"

	"github.com/spf13/cobra"
)

// tag::cmd[]
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the gofs server side",
	Long: `This will run the gofs server.
The gofs server provides a remote file system management interface.`,
	// end::cmd[]
	// tag::args[]
	Args: cobra.NoArgs,
	// end::args[]
	// tag::run[]
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Starting server on '%s'\n", serverAddress)
	},
	// end::run[]
}
