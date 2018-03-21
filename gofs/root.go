package gofs

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// tag::serveraddress[]
var (
	serverAddress string
)

// end::serveraddress[]

// tag::rootcmd[]
var rootCmd = &cobra.Command{
	// end::rootcmd[]
	// tag::rootusage[]
	Long: `A simple interface to a remote file system.
It provides a remote interface similar to the standard tools of
ls, cp, mv, etc.`,
}

// end::rootusage[]

// tag::flag[]
func init() {
	rootCmd.PersistentFlags().StringVarP(
		&serverAddress,
		"server",
		"s",
		"localhost:4270",
		"address:port of the server to connect to",
	)
	// end::flag[]
	// tag::wireserve[]
	rootCmd.AddCommand(serveCmd)
	// end::wireserve[]
	// tag::wirelsmv[]
	rootCmd.AddCommand(lsCmd)
	rootCmd.AddCommand(mvCmd)
	// end::wirelsmv[]
}

// Execute is the main entry from any main func
// tag::execute[]
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// end::execute[]
