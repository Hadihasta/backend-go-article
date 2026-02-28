package handlers

import (
	"backend-article-portal/services"
	"log"
	"net/http"
)

type PostHandlers struct {
	service *services.PostService
}                                                                                                                              

func NewPostHandler(service *services.PostService) *PostHandlers{
	return &PostHandlers{service: service}
}

func (h *PostHandlers) HandlePosts(w http.ResponseWriter, r *http.Request){
	log.Println(w,r)
}