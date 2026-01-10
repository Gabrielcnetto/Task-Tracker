/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"tasks-cli/models"
	"tasks-cli/services"

	"github.com/spf13/cobra"
)

var description string
var NewCmd = &cobra.Command{
	Use:   "New",
	Short: "Add new Task to list",
	Long:  `Add new task to json database`,
	Run: func(cmd *cobra.Command, args []string) {
		newTask := models.Task{
			Description: description,
			Status:      1,
		}
		newTask.GenerateId()
		newTask.CreatedTime()
		newTask.UpdatedTime()
		err := services.AddTask(newTask)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("Task created with sucess!")
	},
}

func init() {
	rootCmd.AddCommand(NewCmd)
	NewCmd.Flags().StringVarP(&description, "description", "d", "", "Task that you will do")
	NewCmd.MarkFlagRequired("description")
}
