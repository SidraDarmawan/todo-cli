package prompt

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

type PrompContent struct {
	ErrorMessage string
	Label        string
	MaxChar      int
}

func PromptGetInput(pc PrompContent, optional bool) string {
	validate := func(input string) error {
		if !optional && (input == "") {
			return errors.New(pc.ErrorMessage)
		}
		if len(input) > pc.MaxChar {
			return fmt.Errorf("text is too long. max: %v; current: %v", pc.MaxChar, len(input))
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:       pc.Label,
		Validate:    validate,
		HideEntered: true,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}
	return result
}
