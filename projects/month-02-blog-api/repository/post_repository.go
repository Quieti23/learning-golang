package repository

import (
	"errors"

	"month02blogapi/model"
)

var ErrPostNotFound = errors.New("post not found")

type PostRepository interface {
	List() ([]model.Post, error)
	Create(post model.Post) (model.Post, error)
	GetByID(id int) (model.Post, error)
	Update(post model.Post) (model.Post, error)
	Delete(id int) error
}
