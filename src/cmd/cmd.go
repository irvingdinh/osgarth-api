package cmd

import (
	"github.com/irvingdinh/osgarth-api/src/component/config"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Lorem ipsum dolor sit amet",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
		os.Exit(0)
	},
}

func Execute() {
	config.Load()

	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(configCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err.Error())
	}
}
