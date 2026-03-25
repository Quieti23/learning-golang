package service

import (
	"context"
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

type UpdatePostInput struct {
	Title   string
	Content string
	Author  string
}

type PostService interface {
	List(ctx context.Context) ([]model.Post, error)
	Create(ctx context.Context, input CreatePostInput) (model.Post, error)
	GetByID(ctx context.Context, id int) (model.Post, error)
	Update(ctx context.Context, id int, input UpdatePostInput) (model.Post, error)
	Delete(ctx context.Context, id int) error
}

type DefaultPostService struct {
	repository repository.PostRepository
}

func NewPostService(postRepository repository.PostRepository) PostService {
	return &DefaultPostService{repository: postRepository}
}

func (s *DefaultPostService) List(ctx context.Context) ([]model.Post, error) {
	return s.repository.List(ctx)
}

func (s *DefaultPostService) Create(ctx context.Context, input CreatePostInput) (model.Post, error) {
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

	return s.repository.Create(ctx, post)
}

func (s *DefaultPostService) GetByID(ctx context.Context, id int) (model.Post, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *DefaultPostService) Update(ctx context.Context, id int, input UpdatePostInput) (model.Post, error) {
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

	post, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return model.Post{}, err
	}

	post.Title = input.Title
	post.Content = input.Content
	post.Author = input.Author

	return s.repository.Update(ctx, post)
}

func (s *DefaultPostService) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
