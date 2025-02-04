package category

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Category struct {
    Id int
    Name string
    Tag string
}

func Add(w http.ResponseWriter , r *http.Request){
    body  , err := ioutil.ReadAll(r.Body)
    if err != nil{
        http.Error(w,"Failed to read body" , http.StatusInternalServerError)
        return
    }

    var category Category
   
    if err:= json.Unmarshal(body,&category); err != nil{
        http.Error(w,err.Error(),http.StatusBadRequest)
        return
    }

    fmt.Printf("%v" , category)
}
