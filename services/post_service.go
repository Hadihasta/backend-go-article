package services

import "backend-article-portal/repositories"

type PostService struct {
	repo *repositories.PostRepository
}

func NewPostService(repo *repositories.PostRepository) *PostService{
	return &PostService{repo: repo}
}

