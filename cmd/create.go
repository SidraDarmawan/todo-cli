package cmd

import (
	"fmt"

	"github.com/SidraDarmawan/todo-cli/data"
	"github.com/SidraDarmawan/todo-cli/prompt"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new todo list",
	Long:  `Create a new todo list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		todoTitlePc := prompt.PrompContent{
			ErrorMessage: "title is invalide",
			Label:        "Title",
			MaxChar:      125,
		}
		todoDescPc := prompt.PrompContent{
			ErrorMessage: "description is invalid",
			Label:        "Description",
			MaxChar:      150,
		}

		title := prompt.PromptGetInput(todoTitlePc, false)
		description := prompt.PromptGetInput(todoDescPc, true)
		data.InsertTodo(title, description, false)

		fmt.Fprintf(cmd.OutOrStdout(), "A Todo has been added with title: %v \n", title)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
