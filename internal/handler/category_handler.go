package handler

import (
	"context"
	"crud_go/internal/domain"
	"crud_go/internal/service"
	"crud_go/pkg/helper"
	"encoding/json"
	"fmt"
	"net/http"
    "strconv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryHandler struct {
    service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
    return &CategoryHandler{service: service}
}

func (h *CategoryHandler) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
    if helper.CheckMethodHttp(w,r,http.MethodPost) {   
        var category domain.Category
        if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
            http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
            return
        }

        result, err := h.service.CreateCategory(context.Background(), &category)
        if err != nil {
            http.Error(w, "Erro ao criar categoria", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "message": "Categoria criada com sucesso",
            "id": result.InsertedID,
        })
    }
    return
}

func (h *CategoryHandler) ListCategoryHandler(w http.ResponseWriter, r *http.Request) {
    if helper.CheckMethodHttp(w,r,http.MethodGet) {
        result, err := h.service.ListAllCategories(context.Background())
        if err != nil {
            http.Error(w, "Erro ao listar categoria", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "message": "Categoria listada com sucesso",
            "list" : result,
        })

    }
    return
}

func (h *CategoryHandler) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request) {
    if helper.CheckMethodHttp(w,r,http.MethodDelete) {
        id := r.URL.Query().Get("id")
        objId , _ := primitive.ObjectIDFromHex(id)
        result, err := h.service.DeleteCategory(context.Background(),objId)
        if err != nil {
            http.Error(w, "Erro ao deletar categoria", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "message": "Categoria deletada com sucesso",
            "deleted" : result.DeletedCount,
        })

    }
    return
}

func (h *CategoryHandler) ReadCategoryHandler(w http.ResponseWriter, r *http.Request) {
    if helper.CheckMethodHttp(w,r,http.MethodGet) {
        id := r.URL.Query().Get("id")
        objId , _ := primitive.ObjectIDFromHex(id)
        result, err := h.service.ReadCategories(context.Background(),objId)
        if err != nil {
            http.Error(w, "Erro ao buscar categoria", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "message": "Categoria consultada com sucesso",
            "categoria" : result,
        })

    }
    return
}

func (h *CategoryHandler) UpdateCategoryHandler(w http.ResponseWriter, r *http.Request) {
    if helper.CheckMethodHttp(w,r,http.MethodPatch) {
        id := r.URL.Query().Get("id")
        objId , _ := primitive.ObjectIDFromHex(id)

        var category domain.Category
        if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
            http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
            return
        }
        
        result, err := h.service.UpdateCategory(context.Background(), &category ,objId)
        if err != nil {
            fmt.Println(err)
            http.Error(w, "Erro ao atualizar categoria", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "message": "Categoria atualizada com sucesso",
            "categoria" : result.MatchedCount,
        })

    }
    return
}

func (h *CategoryHandler) IncrementCategory(w http.ResponseWriter , r *http.Request){
    if helper.CheckMethodHttp(w,r,http.MethodPost){
        operation := r.URL.Query().Get("operation")
        value , _ := strconv.Atoi(r.URL.Query().Get("value"))
        id := r.URL.Query().Get("id")
 
        objId , _ := primitive.ObjectIDFromHex(id)
        result, err := h.service.ReadCategories(context.Background(),objId)
        if err != nil {
            http.Error(w, "Erro ao buscar categoria", http.StatusInternalServerError)
            return
        }

        switch operation {
        case "add" :
            result.Limit += value
            break
        case "sub" :
            result.Limit -= value
            break
        }
        
        category :=  domain.Category{Name:result.Name , Tag:result.Tag , Limit:result.Limit,}
        resultUpdate, errUpdate := h.service.UpdateCategory(context.Background(), &category ,objId)
        if errUpdate != nil {
            http.Error(w, "Erro ao buscar categoria", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(map[string]interface{}{
            "message": "Categoria atualizada com sucesso",
            "categoria" : resultUpdate.MatchedCount,
        })

    }
}
