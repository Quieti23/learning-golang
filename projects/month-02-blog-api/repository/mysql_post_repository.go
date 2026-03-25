package repository

import (
	"context"
	"database/sql"

	"month02blogapi/model"
)

type MySQLPostRepository struct {
	db *sql.DB
}

func NewMySQLPostRepository(db *sql.DB) *MySQLPostRepository {
	return &MySQLPostRepository{db: db}
}

func (r *MySQLPostRepository) List(ctx context.Context) ([]model.Post, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT id, title, content, author, created_at, updated_at
		FROM posts
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, normalizeContextError(ctx, err)
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
		return nil, normalizeContextError(ctx, err)
	}

	return posts, nil
}

func (r *MySQLPostRepository) Create(ctx context.Context, post model.Post) (model.Post, error) {
	result, err := r.db.ExecContext(ctx, `
		INSERT INTO posts (title, content, author, created_at, updated_at)
		VALUES (?, ?, ?, NOW(), NOW())
	`, post.Title, post.Content, post.Author)
	if err != nil {
		return model.Post{}, normalizeContextError(ctx, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return model.Post{}, err
	}

	row := r.db.QueryRowContext(ctx, `
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
		return model.Post{}, normalizeContextError(ctx, err)
	}

	return createdPost, nil
}

func (r *MySQLPostRepository) GetByID(ctx context.Context, id int) (model.Post, error) {
	row := r.db.QueryRowContext(ctx, `
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
		err = normalizeContextError(ctx, err)
		if err == sql.ErrNoRows {
			return model.Post{}, ErrPostNotFound
		}

		return model.Post{}, err
	}

	return post, nil
}

func (r *MySQLPostRepository) Update(ctx context.Context, post model.Post) (model.Post, error) {
	result, err := r.db.ExecContext(ctx, `
		UPDATE posts
		SET title = ?, content = ?, author = ?, updated_at = NOW()
		WHERE id = ?
	`, post.Title, post.Content, post.Author, post.ID)
	if err != nil {
		return model.Post{}, normalizeContextError(ctx, err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return model.Post{}, err
	}
	if affected == 0 {
		return model.Post{}, ErrPostNotFound
	}

	return r.GetByID(ctx, post.ID)
}

func (r *MySQLPostRepository) Delete(ctx context.Context, id int) error {
	result, err := r.db.ExecContext(ctx, `DELETE FROM posts WHERE id = ?`, id)
	if err != nil {
		return normalizeContextError(ctx, err)
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

func normalizeContextError(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}

	if ctxErr := ctx.Err(); ctxErr != nil {
		return ctxErr
	}

	return err
}
