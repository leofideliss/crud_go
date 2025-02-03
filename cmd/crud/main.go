package main

import (
    "fmt"
    "flag"
    "net/http"

    "github.com/leofideliss/crud_go/internal/controller"
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
    http.HandleFunc("/api/add"  , Add)

    http.ListenAndServe(fmt.Sprintf(":%d", *porta), nil)
}

