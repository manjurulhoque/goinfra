/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "goinfra",
	Short: "A lightweight Infrastructure-as-Code CLI for AWS",
	Long: `GoInfra is a lightweight Infrastructure-as-Code CLI tool that helps you manage AWS resources.
It supports provisioning and managing various AWS services including EC2, S3, and RDS.

Features:
- AWS SDK v2 integration
- YAML/JSON configuration support
- Resource management for EC2, S3, and RDS
- CloudFormation template support

Example usage:
  goinfra init                  # Initialize a new infrastructure project
  goinfra plan                  # Show execution plan
  goinfra apply                 # Apply infrastructure changes
  goinfra destroy               # Remove all managed resources`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.goinfra.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
