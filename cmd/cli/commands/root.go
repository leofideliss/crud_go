package commands

import (
	"github.com/spf13/cobra"
    "crud_go/internal/service"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "context"
    "log"
)
var CategoryService *service.CategoryService

var RootCommand = &cobra.Command{
	Use:   "crudcli",
	Short: "API CRUD",
	Long:  "Gerenciamento da api CRUD",
}

func Execute() error {
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("crud_go")

	// Criar o serviço de categorias
	CategoryService = service.NewCategoryService(db)
    
	return RootCommand.Execute()
}

// init adiciona os subcomandos
func init() {
	RootCommand.AddCommand(updateCmd)
	RootCommand.AddCommand(listCmd)
}
