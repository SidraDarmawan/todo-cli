package cmd

import (
	"github.com/SidraDarmawan/todo-cli/data"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A logger description that spans multiple lines and likely contains examples
	and usage of using your command. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		data.MigrateDatabase()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
