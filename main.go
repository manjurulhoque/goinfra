/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/manjurulhoque/goinfra/cmd"
	"github.com/manjurulhoque/goinfra/internal/commands"
)

func main() {
	// Add subcommands
	cmd.RootCmd.AddCommand(commands.GetInitCommand())
	cmd.RootCmd.AddCommand(commands.GetPlanCommand())

	// Execute the root command
	cmd.Execute()
}
