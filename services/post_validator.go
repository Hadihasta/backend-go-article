package services

import (
	"backend-article-portal/models"
	"errors"
	"strings"
)

func ValidatePost(data *models.Posts) error {
	if strings.TrimSpace(data.Title) == "" {
		return errors.New("title is required")
	}
	if len(data.Title) < 20 {
		return errors.New("title minimal 20 characters")
	}

	if strings.TrimSpace(data.Content) == "" {
		return errors.New("content is required")
	}
	if len(data.Content) < 200 {
		return errors.New("content minimal 200 characters")
	}

	if strings.TrimSpace(data.Category) == "" {
		return errors.New("category is required")
	}
	if len(data.Category) < 3 {
		return errors.New("category minimal 3 characters")
	}

	if strings.TrimSpace(data.Status) == "" {
		return errors.New("status is required")
	}

	validStatus := map[string]bool{
		"publish": true,
		"draft":   true,
		"thrash":  true,
	}

	if !validStatus[strings.ToLower(data.Status)] {
		return errors.New("status must be publish, draft, or thrash")
	}

	return nil
}
