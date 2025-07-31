package cmd

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/SidraDarmawan/todo-cli/data"
	"github.com/SidraDarmawan/todo-cli/prompt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var showEverything bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all todo list",
	Long:  `Show all todo list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		todos := data.ReadAllTodos(showEverything)

		if len(todos) == 0 {
			fmt.Fprintln(cmd.OutOrStdout(), "No todo items found. Create one using 'todo create'.")
			return nil
		}

		items := make([]interface{}, len(todos))
		for i, t := range todos {
			items[i] = t
		}

		templates := promptui.SelectTemplates{
			Label:    "{{ . }}?",
			Active:   "🟢 {{ .ID | printf \"%.0d\" }} {{ if .Status }} {{ .Title | green }} {{ else }} {{ .Title | red }} {{ end }}",
			Inactive: "  {{ .ID | printf \"%.0d\" }} {{ if .Status }} {{ .Title | green }} {{ else }} {{ .Title | red }} {{ end }}",
			Selected: "🟢 {{ .ID | printf \"%.0d\" }} {{ if .Status }} {{ .Title | green }} {{ else }} {{ .Title | red }} {{ end }}",
			Details: `
----------- Details -----------
{{ "ID:" | faint }}          {{ .ID }}
{{ "Title:" | faint }}       {{ .Title }}
{{ "Description:" | faint }} {{ .Description }}
{{ "Status:" | faint }}      {{ if .Status }} {{ "✅ Done" }} {{ else }} {{ "❌ Work In Progress" }} {{ end }}
{{ "Created At:" | faint }}  {{ .CreatedAt }}  
`,
		}

		todoSc := prompt.SelectContent{
			Label:     "Your Todo List",
			Items:     items,
			Templates: &templates,
		}

		_, selectedString := prompt.PrompSelectContent(&todoSc)

		re := regexp.MustCompile(`\d+`)
		idMatch := re.FindString(selectedString)

		if idMatch == "" {
			return fmt.Errorf("failed to extract ID from selected string: %q", selectedString)
		}

		parsedTodoID, err := strconv.ParseUint(idMatch, 10, 64)
		if err != nil {

			return fmt.Errorf("failed to parse extracted ID %q: %w", idMatch, err)
		}

		todo, err := data.FindOneTodo(uint(parsedTodoID))
		if err != nil {
			return fmt.Errorf("failed to find todo with ID %d: %w", parsedTodoID, err)
		}

		var actionItems []string
		if todo.Status {
			actionItems = []string{"Mark As Not Done", "Delete"}
		} else {
			actionItems = []string{"Mark As Done", "Delete"}
		}

		actionScItems := make([]interface{}, len(actionItems))
		for i, item := range actionItems {
			actionScItems[i] = item
		}

		actionSc := prompt.SelectContent{
			Label:     fmt.Sprintf("Choose An Action for %v", todo.Title),
			Items:     actionScItems,
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
			fmt.Fprintln(cmd.OutOrStdout(), "An unexpected action was selected.")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&showEverything, "everything", "e", false, "Show everything including finished TODO.")
}
