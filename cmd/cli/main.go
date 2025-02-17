package main

import (
	"fmt"
	"os"
    "log"
	"crud_go/cmd/cli/commands"
    "crud_go/internal/service"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "context"
)

func main() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("crud_go")

	// Criar o serviço de categorias
	categoryService := service.NewCategoryService(db)
    
	if err := commands.Execute(); err != nil {
		fmt.Println("Erro:", err)
		os.Exit(1)
	}
}
