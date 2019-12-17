package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	rootCmd := &cobra.Command{
		Run: func(c *cobra.Command, args []string) {
			fmt.Println(viper.GetString("Flag"))
		},
	}

	viper.SetDefault("Flag", "default")

	rootCmd.Flags().String("flag", "", "help for flag")
	viper.BindPFlag("Flag", rootCmd.Flags().Lookup("flag"))

	// tag::envvar[]
	viper.SetEnvPrefix("VF")
	viper.BindEnv("Flag")
	// end::envvar[]

	rootCmd.Execute()
}