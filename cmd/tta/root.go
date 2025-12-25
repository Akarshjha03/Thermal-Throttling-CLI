package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tta",
	Short: "Thermal Throttling Analyzer",
	Long:  `A Windows-only CLI diagnostic tool in Go that detects, analyzes, and explains CPU thermal throttling.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func matchRun() {
    Execute()
}
