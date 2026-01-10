/*
Copyright © 2026 NAME HERE desenvolvedor.netto@gmail.com
*/
package cmd

import (
	"fmt"
	"tasks-cli/services"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Get all tasks",
	Long:  `Get all tasks added by you`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := services.GetAllTasks(0)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("\n%v\n", tasks)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
