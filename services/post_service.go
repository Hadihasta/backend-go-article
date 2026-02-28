package services

import (
	"backend-article-portal/models"
	"backend-article-portal/repositories"
)

type PostService struct {
	repo *repositories.PostRepository
}

func NewPostService(repo *repositories.PostRepository) *PostService{
	return &PostService{repo: repo}
}


func (s *PostService) Create(data *models.Posts) error  {

	err := ValidatePost(data)
	if err != nil{
		return err
	}

	return s.repo.Create(data)
}

func (s *PostService) GetAll(limit, offset string) ([]models.ResponsePosts, error) {

	posts, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	var response []models.ResponsePosts

	for _, p := range posts {
		response = append(response, models.ResponsePosts{
			Title:    p.Title,
			Content:  p.Content,
			Category: p.Category,
			Status:   p.Status,
		})
	}

	return response, nil
}
