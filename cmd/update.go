/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"tasks-cli/services"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var NewStatus int8

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Task status",
	Long:  `Update task status using is id. status be like 1,2,3 (1 pending, 2 todo, 3 done)`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TaskId:", TaskId)
		fmt.Println("status:", NewStatus)
		err := services.Update(TaskId, NewStatus)
		fmt.Println(err)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&TaskId, "id", "i", "", "Use to identify task")
	updateCmd.Flags().Int8VarP(&NewStatus, "status", "s", 0, "Status do add")
}
