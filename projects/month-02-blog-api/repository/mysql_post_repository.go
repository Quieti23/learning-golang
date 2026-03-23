package repository

import (
	"database/sql"

	"month02blogapi/model"
)

type MySQLPostRepository struct {
	db *sql.DB
}

func NewMySQLPostRepository(db *sql.DB) *MySQLPostRepository {
	return &MySQLPostRepository{db: db}
}

func (r *MySQLPostRepository) List() ([]model.Post, error) {
	rows, err := r.db.Query(`
		SELECT id, title, content, author, created_at, updated_at
		FROM posts
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := make([]model.Post, 0)
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.Author,
			&post.CreatedAt,
			&post.UpdatedAt,
		); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (r *MySQLPostRepository) Create(post model.Post) (model.Post, error) {
	result, err := r.db.Exec(`
		INSERT INTO posts (title, content, author, created_at, updated_at)
		VALUES (?, ?, ?, NOW(), NOW())
	`, post.Title, post.Content, post.Author)
	if err != nil {
		return model.Post{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return model.Post{}, err
	}

	row := r.db.QueryRow(`
		SELECT id, title, content, author, created_at, updated_at
		FROM posts
		WHERE id = ?
	`, id)

	var createdPost model.Post
	if err := row.Scan(
		&createdPost.ID,
		&createdPost.Title,
		&createdPost.Content,
		&createdPost.Author,
		&createdPost.CreatedAt,
		&createdPost.UpdatedAt,
	); err != nil {
		return model.Post{}, err
	}

	return createdPost, nil
}

func (r *MySQLPostRepository) GetByID(id int) (model.Post, error) {
	row := r.db.QueryRow(`
		SELECT id, title, content, author, created_at, updated_at
		FROM posts
		WHERE id = ?
	`, id)

	var post model.Post
	if err := row.Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.Author,
		&post.CreatedAt,
		&post.UpdatedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return model.Post{}, ErrPostNotFound
		}

		return model.Post{}, err
	}

	return post, nil
}

func (r *MySQLPostRepository) Update(post model.Post) (model.Post, error) {
	result, err := r.db.Exec(`
		UPDATE posts
		SET title = ?, content = ?, author = ?, updated_at = NOW()
		WHERE id = ?
	`, post.Title, post.Content, post.Author, post.ID)
	if err != nil {
		return model.Post{}, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return model.Post{}, err
	}
	if affected == 0 {
		return model.Post{}, ErrPostNotFound
	}

	return r.GetByID(post.ID)
}

func (r *MySQLPostRepository) Delete(id int) error {
	result, err := r.db.Exec(`DELETE FROM posts WHERE id = ?`, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected == 0 {
		return ErrPostNotFound
	}

	return nil
}