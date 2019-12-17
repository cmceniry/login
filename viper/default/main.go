// tag::intro[]
package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
// end::intro[]
	// tag::cobra[]
	rootCmd := &cobra.Command{
		Run: func(c *cobra.Command, args []string) {
			fmt.Println(viper.GetString("Flag"))
		},
	}
	// end::cobra[]

	// tag::viper[]
	viper.SetDefault("Flag", "default")
	// end::viper[]

	// tag::execute[]
	rootCmd.Execute()
	// end::execute[]
}