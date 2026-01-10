/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"tasks-cli/services"

	"github.com/spf13/cobra"
)

var TaskId string
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete one task by id",
	Long:  `Remove one task by id`,
	Run: func(cmd *cobra.Command, args []string) {
		err := services.DeletTasks(TaskId)
		fmt.Println(err)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&TaskId, "id", "i", "", "Delete one task using is id")
}
