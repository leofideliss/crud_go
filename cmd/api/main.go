package main

import (
    "flag"
    "fmt"
    "net/http"
    "log"

    "crud_go/internal"
)

var (
    porta *int
    baseUrl string
)

func init(){
    porta = flag.Int("p" , 8888 , "porta")
    flag.Parse()

    baseUrl = fmt.Sprintf("http://localhost:%d",porta)
}

func main (){

    category.NewRepository
    
    // Categorias
    http.HandleFunc("/api/addCategoria"  , category.Add)

    
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *porta), nil))
}

