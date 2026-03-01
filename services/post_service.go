package services

import (
	"backend-article-portal/models"
	"backend-article-portal/repositories"
)

type PostService struct {
	repo *repositories.PostRepository
}

func NewPostService(repo *repositories.PostRepository) *PostService {
	return &PostService{repo: repo}
}

func (s *PostService) Create(data *models.Posts) error {

	err := ValidatePost(data)
	if err != nil {
		return err
	}

	return s.repo.Create(data)
}

func (s *PostService) GetAll(limit, offset string) ([]models.Posts, error) {

	posts, err := s.repo.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}

	var response []models.Posts

	for _, p := range posts {
		response = append(response, models.Posts{
			ID:    p.ID,
			Title:    p.Title,
			Content:  p.Content,
			Category: p.Category,
			Status:   p.Status,
		})
	}

	return response, nil
}

func (s *PostService) GetByID(id int) (*models.ResponsePosts, error) {
	post, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	response := &models.ResponsePosts{

		Title:    post.Title,
		Content:  post.Content,
		Category: post.Category,
		Status:   post.Status,
	}

	return response, nil
}


func (s *PostService) Update(data *models.Posts) error {


	err := ValidatePost(data)
	if err != nil {
		return err
	}

	return s.repo.Update(data)
}


func (s *PostService) Delete(id int) error {
	return s.repo.Delete(id)
}