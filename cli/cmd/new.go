/*
Copyright © 2025 Nima Yonten bugloper@hey.com
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "gho new",
	Short: "gho new will generate a new gho application",
	Long:  `gho new will generate a new gho application`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("new called")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
