package gofs

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// tag::cmd[]
var mvCmd = &cobra.Command{
	Use:   "mv source ... target",
	Short: "Move a file to a destination file, or moves multiple files into a directory",
	Long: `If given two arguments, moves the first file to the second file. If that second file is a directory, the first is moved into it.
If more than two arguments are given, it moves all but the last file into the last one which it expects to be a directory`,
	Args: cobra.MinimumNArgs(2),
	// end::cmd[]
	//tag::run[]
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 2 {
			fmt.Printf("Moving file '%s' to '%s'", args[0], args[1])
		} else {
			fmt.Printf(
				"Moving files '%s' into directory '%s'\n",
				strings.Join(args[0:len(args)-1], ","),
				args[len(args)-1],
			)
		}
	},
}

// end::run[]
