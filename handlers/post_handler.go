package handlers

import (
	"backend-article-portal/models"
	"backend-article-portal/services"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type PostHandlers struct {
	service *services.PostService
}

func NewPostHandler(service *services.PostService) *PostHandlers {
	return &PostHandlers{service: service}
}

func (h *PostHandlers) HandlePosts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.Create(w, r)
	case http.MethodGet:
		h.GetAll(w, r)
	}
}

func (h *PostHandlers) Create(w http.ResponseWriter, r *http.Request) {
	var post models.Posts
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid Request body", http.StatusBadRequest)
		return
	}

	err = h.service.Create(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandlers) GetAll(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	offset := r.URL.Query().Get("offset")

	Post, err := h.service.GetAll(limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Post)

}

func (h *PostHandlers) HandlePostByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetByID(w, r)
	case http.MethodPut:
		h.Update(w, r)
	case http.MethodDelete:
		h.Delete(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *PostHandlers) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/article/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Post Id", http.StatusBadRequest)
		return
	}

	post, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)

}

func (h *PostHandlers) Update(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/article/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Post Id", http.StatusBadRequest)
		return
	}

	var post models.Posts
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	post.ID = id

	err = h.service.Update(&post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  http.StatusOK,
		"message": "post updated successfully",
	})
}


func (h *PostHandlers) Delete(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/article/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Post Id", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "post deleted successfully",
	})
}