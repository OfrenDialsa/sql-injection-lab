/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/OfrenDialsa/lab/sql-injection/api"

	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run vulnerable API server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting API...")
		api.Run()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
