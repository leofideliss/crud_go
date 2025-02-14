package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

    "crud_go/internal/service"
    "crud_go/internal/handler"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "context"
)

var (
    porta *int
    baseUrl string
)

func init() {
    porta = flag.Int("p", 8888, "porta")
    flag.Parse()

    baseUrl = fmt.Sprintf("http://localhost:%d", porta)
}

func main() {

	// Conectar ao MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("crud_go")

	// Criar o serviço de categorias
	categoryService := service.NewCategoryService(db)

	// Criar o handler de categorias
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// Registrar a rota
	http.HandleFunc("/api/addCategory", categoryHandler.CreateCategoryHandler)

	// Iniciar o servidor HTTP
	log.Println("Servidor rodando na porta 8888...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *porta), nil))
}
