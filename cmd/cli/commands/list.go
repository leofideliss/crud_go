package commands

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
    Use: "list",
    Short: "Lista as categorias",
    Run: func(cmd *cobra.Command , args []string){
        result,_ := CategoryService.ListAllCategories(context.Background())
        fmt.Println("Lista de categorias:")
        for _,v := range result{
            fmt.Printf("Id:%s\t Name:%s\t Tag:%s\t Limit:%d\t\n" , string(v.Id.Hex()) , v.Name , v.Tag, v.Limit)
        }
    },
}

func init(){
    listCmd.Flags().StringP("list" , "l" , "" , "Lista todas as categorias")
}
