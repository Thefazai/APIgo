package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "apigo",
	Short: "A lightweight CLI tool for API testing",
	Long:  `Apigo is a fast, minimal API testing tool built with Go.`,
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
