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
    porta *int
    baseUrl string
    repo app.Repository
)

func init() {
    porta = flag.Int("p", 8888, "porta")
    flag.Parse()

    baseUrl = fmt.Sprintf("http://localhost:%d", porta)
}

func main() {
    // na memória
    repo = app.CreateRepository()
    
    app.Connect()
    
    http.HandleFunc("/api/addCategory", addCategory)
    http.HandleFunc("/api/deleteCategory", deleteCategory)
    http.HandleFunc("/api/readCategory", readCategory)
    http.HandleFunc("/api/updateCategory",updateCategory)

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *porta), nil))
}

func addCategory(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        makeResponse(w,Response{"success" : false,"message":"Método inválido"},http.StatusMethodNotAllowed,Headers{"Content-type":"application/json"})
        return
    }
    
    b, _ := ioutil.ReadAll(r.Body)
    var category app.Category
    errJson := json.Unmarshal(b, &category)
    if errJson != nil {
        makeResponse(w,Response{"success" : false,"message":"Json invalido"},http.StatusInternalServerError,Headers{"Content-type":"application/json"})
        return
    }

    repo.Create(&category)
    makeResponse(w,Response{"success" : true,"message":"Categoria adicionada"},http.StatusOK,Headers{"Content-type":"application/json"})
}

func deleteCategory(w http.ResponseWriter , r *http.Request){
    if r.Method != http.MethodDelete {
        makeResponse(w,Response{"success" : false,"message":"Método inválido"},http.StatusMethodNotAllowed,Headers{"Content-type":"application/json"})
        return
    }

    id := r.URL.Query().Get("id")
    pos, _ := strconv.Atoi(id)
    hasDeleted := repo.Delete(pos)
    if hasDeleted {
        makeResponse(w,Response{"success" : true,"message":"Categoria deletada"},http.StatusOK,Headers{"Content-type":"application/json"})
        return
    }

    makeResponse(w,Response{"success" : false,"message":"Erro ao deletar"},http.StatusNotFound,Headers{"Content-type":"application/json"})
    return
}

func readCategory(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        makeResponse(w,Response{"success" : false,"message":"Método inválido"},http.StatusMethodNotAllowed,Headers{"Content-type":"application/json"})
        return
    }
    
    id := r.URL.Query().Get("id")
    pos, _ := strconv.Atoi(id)
    categoryPos, _ := repo.Read(pos)

    if categoryPos != nil {
        makeResponse(w,Response{"success" : true,"message":"Categoria consultada" ,"name":string(categoryPos.Name) ,"tag":string(categoryPos.Tag)},http.StatusOK,Headers{"Content-type":"application/json"})
        return
    }

       makeResponse(w,Response{"success" : false,"message":"Categoria não encontrada"},http.StatusNotFound,Headers{"Content-type":"application/json"})
        return
}

func updateCategory(w http.ResponseWriter , r *http.Request){
    if r.Method != http.MethodPatch {
        makeResponse(w,Response{"success" : false,"message":"Método inválido"},http.StatusMethodNotAllowed,Headers{"Content-type":"application/json"})
        return
    }

    var category app.Category
    
    id := r.URL.Query().Get("id")
    pos, _ := strconv.Atoi(id)
    b, _ := ioutil.ReadAll(r.Body)
    errJson := json.Unmarshal(b, &category)
    if errJson != nil {
        makeResponse(w,Response{"success" : false,"message":"Json invalido"},http.StatusInternalServerError,Headers{"Content-type":"application/json"})
        return
    }

    err := repo.Update(&category,pos)
    if err {
        makeResponse(w,Response{"success" : true,"message":"Categoria atualizada com sucesso!"},http.StatusOK,Headers{"Content-type":"application/json"})
        return
    }

    makeResponse(w,Response{"success" : false,"message":"Erro ao atualizar"},http.StatusNotFound,Headers{"Content-type":"application/json"})
    return
}


type Headers map[string]string
type Response map[string]interface{}

func makeResponse( w http.ResponseWriter , response Response ,  status int , header Headers) {
    for k , v := range header{
        w.Header().Set(k,v)
    }
    w.WriteHeader(status)
    json , _ := json.Marshal(response)
    w.Write(json)
}
