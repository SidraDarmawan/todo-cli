package main

import (
	"github.com/SidraDarmawan/todo-cli/cmd"
	"github.com/SidraDarmawan/todo-cli/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()
}