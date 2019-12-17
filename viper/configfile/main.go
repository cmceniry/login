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

	viper.SetEnvPrefix("VF")
	viper.BindEnv("Flag")

	// tag::configfile[]
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		// Only error on errors other than file not found
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			panic(err)
		}
	}
	// end::configfile[]

	rootCmd.Execute()
}