package handler

import (
    "context"
    "encoding/json"
    "net/http"
    "crud_go/internal/domain"
    "crud_go/internal/service"
)

type CategoryHandler struct {
    service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
    return &CategoryHandler{service: service}
}

func (h *CategoryHandler) CreateCategoryHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

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

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Categoria criada com sucesso",
        "id": result.InsertedID,
    })
}
