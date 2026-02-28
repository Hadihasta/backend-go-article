package repositories

import (
	"backend-article-portal/models"
	"database/sql"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (repo *PostRepository) Create(post *models.Posts) error {
	query := `INSERT INTO posts (title, content, category, status) VALUES (?, ?, ?, ?)
	`

	result, err := repo.db.Exec(
		query,
		post.Title,
		post.Content,
		post.Category,
		post.Status,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	post.ID = int(id)
	return nil
}

func (repo *PostRepository) GetAll(limit, offset string) ([]models.Posts, error) {

	query := `
	SELECT id, title, content, category, status
	FROM posts
	LIMIT ? OFFSET ?
	`

	rows, err := repo.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Posts

	for rows.Next() {
		var post models.Posts

		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.Category,
			&post.Status,
		)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}