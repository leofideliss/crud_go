package commands

import (
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "crudcli",
	Short: "API CRUD",
	Long:  "Gerenciamento da api CRUD",
}

func Execute() error {
	return RootCommand.Execute()
}

// init adiciona os subcomandos
func init() {
	RootCommand.AddCommand(incrementCmd)
}
