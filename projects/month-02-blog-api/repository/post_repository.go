package repository

import "month02blogapi/model"

type PostRepository interface {
	List() []model.Post
	Create(post model.Post) model.Post
}
