package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

var incrementCmd = &cobra.Command{
	Use:   "increment",
	Short: "incrementa o limite de uma categoria",
	Run: func(cmd *cobra.Command, args []string) {
		
		id, _ := cmd.Flags().GetString("id")
        if id == ""{
            fmt.Println("Informar o Id da categoria")
        }
		
	},
 }

func init() {
	incrementCmd.Flags().StringP("id", "d", "", "Id da categoria que será incrementada")
}
