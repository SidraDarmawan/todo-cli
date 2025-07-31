package prompt

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

type SelectContent struct {
	Label     string
	Items     []interface{}
	Templates *promptui.SelectTemplates
}

func PrompSelectContent(sc *SelectContent) (int, string) {
	prompt := promptui.Select{
		Label:        sc.Label,
		Items:        sc.Items,
		Templates:    sc.Templates,
		HideSelected: true,
	}

	index, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failder %v\n", err)
		return 0, ""
	}

	return index, result
}
