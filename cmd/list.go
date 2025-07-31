package cmd

import (
	"fmt"
	"strings"

	"github.com/SidraDarmawan/todo-cli/data"
	"github.com/SidraDarmawan/todo-cli/prompt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var showEverything bool

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all todo list",
	Long:  `Show all todo list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		todos := data.ReadAllTodos(showEverything)

		templates := promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "🟢 {{ if .Status }} {{ .Title | green }} {{ else }} {{ .Title | red }} {{ end }}",
			Inactive: "  {{ if .Status }} {{ .Title | green }} {{ else }} {{ .Title | red }} {{ end }}",
			Selected: "🟢 {{ if .Status }} {{ .Title | green }} {{ else }} {{ .Title | red }} {{ end }}",
			Details: `
----------- Details -----------
{{ "Title:" | faint }}  {{ .Title }}
{{ "Description:" | faint }}  {{ .Description }}
{{ "Status:" | faint }}  {{ if .Status }} {{ "✅ Done" }} {{ else }} {{ "❌ Work In Progress" }} {{ end }}
`,
		}

		todoSc := prompt.SelectContent{
			Label:     "Your Todo List",
			Items:     []interface{}{todos},
			Templates: &templates,
		}
		_, selectedTodo := prompt.PrompSelectContent(&todoSc)

		todoID := fetchTodoID(selectedTodo)

		todo := data.FindOneTodo(todoID)

		var items []string

		if todo.Status {
			items = []string{"Mark As Not Done", "Delete"}
		} else {
			items = []string{"Mark As Done", "Delete"}
		}
		actionSc := prompt.SelectContent{
			Label:     fmt.Sprintf("Choose An Action for %v", todo.Title),
			Items:     []interface{}{items},
			Templates: nil,
		}
		_, selectedAction := prompt.PrompSelectContent(&actionSc)

		switch selectedAction {
		case "Mark As Done":
			data.MarkTodoAsDone(&todo)
			fmt.Fprintln(cmd.OutOrStdout(), "======================")
			fmt.Fprintln(cmd.OutOrStdout(), todo.Title)
			fmt.Fprintln(cmd.OutOrStdout(), "======================")
			fmt.Fprintln(cmd.OutOrStdout(), "Marked as done")
		case "Mark As Not Done":
			data.MarkTodoAsNotDone(&todo)
			fmt.Fprintln(cmd.OutOrStdout(), "======================")
			fmt.Fprintln(cmd.OutOrStdout(), todo.Title)
			fmt.Fprintln(cmd.OutOrStdout(), "======================")
			fmt.Fprintln(cmd.OutOrStdout(), "Marked as not done")
		case "Delete":
			data.DeleteTodo(&todo)
			fmt.Fprintln(cmd.OutOrStdout(), "======================")
			fmt.Fprintln(cmd.OutOrStdout(), todo.Title)
			fmt.Fprintln(cmd.OutOrStdout(), "======================")
			fmt.Fprintln(cmd.OutOrStdout(), "Deleted")
		default:
			fmt.Fprintln(cmd.OutOrStdout(), "How the hell does it went here?")
		}
		return nil
	},
}

func fetchTodoID(s string) string {
	dirtyID := strings.Split(s, " ")[0]

	return dirtyID[2:]
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	listCmd.Flags().BoolVarP(&showEverything, "everything", "e", false, "Show everything including finished TODO.")
}
