package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// planCmd represents the plan command
var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Display a preview of the changes that would be made by the apply command",
	Long: `This command SHOULD take a configuration file as input and display a preview of 
	the changes that would be made by the apply command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("plan called")
	},
}

func GetPlanCommand() *cobra.Command {
	return planCmd
}
