package commands

import (
    //	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
    Use: "list",
    Short: "Lista as categorias",
    Run: func(cmd *cobra.Command , args []string){
        
    },
}

func init(){
    listCmd.Flags().StringP("list" , "l" , "" , "Lista todas as categorias")
}
