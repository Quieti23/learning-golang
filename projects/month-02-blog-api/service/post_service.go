package service

import (
	"errors"
	"strings"

	"month02blogapi/model"
	"month02blogapi/repository"
)

var ErrTitleRequired = errors.New("title is required")
var ErrContentRequired = errors.New("content is required")
var ErrAuthorRequired = errors.New("author is required")

type CreatePostInput struct {
	Title   string
	Content string
	Author  string
}

type PostService interface {
	List() []model.Post
	Create(input CreatePostInput) (model.Post, error)
}

type DefaultPostService struct {
	repository repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) PostService {
	return &DefaultPostService{repository: postRepository}
}

func (s *DefaultPostService) List() []model.Post {
	return s.repository.List()
}

func (s *DefaultPostService) Create(input CreatePostInput) (model.Post, error) {
	input.Title = strings.TrimSpace(input.Title)
	input.Content = strings.TrimSpace(input.Content)
	input.Author = strings.TrimSpace(input.Author)

	if input.Title == "" {
		return model.Post{}, ErrTitleRequired
	}
	if input.Content == "" {
		return model.Post{}, ErrContentRequired
	}
	if input.Author == "" {
		return model.Post{}, ErrAuthorRequired
	}

	post := model.Post{
		Title:   input.Title,
		Content: input.Content,
		Author:  input.Author,
	}

	return s.repository.Create(post), nil
}
