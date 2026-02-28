package handlers

import (
	"backend-article-portal/models"
	"backend-article-portal/services"
	"encoding/json"
	"net/http"
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

	Post, err := h.service.GetAll(limit,offset)
	if err != nil{
		http.Error(w,err.Error(), http.StatusInternalServerError)
		return
	}


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Post)

}