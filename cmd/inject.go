package cmd

import (
	"fmt"

	"github.com/OfrenDialsa/lab/sql-injection/inject"
	"github.com/spf13/cobra"
)

var targetURL string
var mode string

var injectCmd = &cobra.Command{
	Use:   "inject",
	Short: "Run SQL injection tester",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running SQLi tester...")
		fmt.Println("Target:", targetURL)
		fmt.Println("Mode:", mode)

		inject.Run(targetURL, mode)
	},
}

func init() {
	rootCmd.AddCommand(injectCmd)

	injectCmd.Flags().StringVar(&targetURL, "url", "http://localhost:8080", "Target base URL")
	injectCmd.Flags().StringVar(&mode, "mode", "basic", "Mode: basic | boolean | time")
}
