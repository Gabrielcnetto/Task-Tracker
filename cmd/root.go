/*
Copyright © 2026 NAME HERE desenvolvedor.netto@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Task",
	Short: "Build a CLI app to track your tasks and manage your to-do list.",
	Long: `Task tracker is a project used 
	to track and manage your tasks. 
	In this task, you will build a simple command line interface (CLI) to track what you need to do, 
	what you have done, and what you are currently working on. 
	This project will help you practice your programming skills, 
	including working with the filesystem, handling user inputs, 
	and building a simple CLI application.
`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
