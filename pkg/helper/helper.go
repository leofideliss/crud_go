package helper

import (
    "net/http"
    "encoding/json"
	"io/ioutil"
)

type Headers map[string]string
type Response map[string]interface{}

func MakeResponse( w http.ResponseWriter , response Response , status int , header Headers) {
    for k , v := range header{
        w.Header().Set(k,v)
    }
    w.WriteHeader(status)
    json , _ := json.Marshal(response)
    w.Write(json)
}

func CheckMethodHttp(w http.ResponseWriter ,r *http.Request , method string) bool{
    if r.Method != method {
        MakeResponse(w,Response{"success" : false,"message":"Método inválido"},http.StatusMethodNotAllowed,Headers{"Content-type":"application/json"})
        return false
    }
    return true
}

func ReadBodyToStruct[T any](w http.ResponseWriter , r *http.Request , domain *T) *T{
    b, _ := ioutil.ReadAll(r.Body)
    errJson := json.Unmarshal(b, domain)
    if errJson != nil {
        MakeResponse(w,Response{"success" : false,"message":"Json invalido"},http.StatusInternalServerError,Headers{"Content-type":"application/json"})
        return nil
    }
    return domain
}

func ResponseAPI(response Response ,status int , w http.ResponseWriter){
    MakeResponse(w,response,status,Headers{"Content-type":"application/json"})
}
