/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"tasks-cli/services"

	"github.com/spf13/cobra"
)

var status int

var listFilterCmd = &cobra.Command{
	Use:   "listFilter",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := services.GetAllTasks(status)
		if err != nil {
			fmt.Println("Error:", err.Error())
		}
		fmt.Println(tasks)
	},
}

func init() {
	rootCmd.AddCommand(listFilterCmd)
	listFilterCmd.Flags().IntVarP(&status, "status", "s", 0, "Add which status you wanna get")
}
