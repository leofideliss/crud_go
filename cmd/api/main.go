package main

import (
	"crud_go/internal/app"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var (
	porta   *int
	baseUrl string
	repo    app.Repository
)

func init() {
	porta = flag.Int("p", 8888, "porta")
	flag.Parse()

	baseUrl = fmt.Sprintf("http://localhost:%d", porta)
}

func main() {

	repo = app.CreateRepository()

	http.HandleFunc("/api/addCategoria", AddCategory)
	http.HandleFunc("/api/readCategoria", readCategory)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *porta), nil))
}

func AddCategory(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
	}
	var category app.Category
	errJson := json.Unmarshal(b, &category)
	if errJson != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
	}
	repo.Create(&category)
}

func readCategory(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	pos, _ := strconv.Atoi(id)
	categoryPos , _  := repo.Read(pos)

    if categoryPos != nil{
        fmt.Println(*categoryPos)
        w.WriteHeader(http.StatusOK)
        return
    }
    http.NotFound(w,r)

}
