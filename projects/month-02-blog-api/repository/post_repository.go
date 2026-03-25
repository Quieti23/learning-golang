package repository

import (
	"context"
	"errors"

	"month02blogapi/model"
)

var ErrPostNotFound = errors.New("post not found")

type PostRepository interface {
	List(ctx context.Context) ([]model.Post, error)
	Create(ctx context.Context, post model.Post) (model.Post, error)
	GetByID(ctx context.Context, id int) (model.Post, error)
	Update(ctx context.Context, post model.Post) (model.Post, error)
	Delete(ctx context.Context, id int) error
}
