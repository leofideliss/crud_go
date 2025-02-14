package handler

import (
	"context"
	"crud_go/internal/domain"
	"crud_go/internal/service"
	"crud_go/pkg/helper"
	"encoding/json"
	"net/http"
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
