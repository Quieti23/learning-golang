package repository

import (
	"context"
	"sync"
	"time"

	"month02blogapi/model"
)

type InMemoryPostRepository struct {
	mu     sync.RWMutex
	nextID int
	posts  []model.Post
}

func NewInMemoryPostRepository() *InMemoryPostRepository {
	return &InMemoryPostRepository{
		nextID: 1,
		posts:  make([]model.Post, 0),
	}
}

func (r *InMemoryPostRepository) List(ctx context.Context) ([]model.Post, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]model.Post, len(r.posts))
	copy(result, r.posts)
	return result, nil
}

func (r *InMemoryPostRepository) Create(ctx context.Context, post model.Post) (model.Post, error) {
	if err := ctx.Err(); err != nil {
		return model.Post{}, err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	post.ID = r.nextID
	post.CreatedAt = now
	post.UpdatedAt = now

	r.nextID++
	r.posts = append(r.posts, post)
	return post, nil
}

func (r *InMemoryPostRepository) GetByID(ctx context.Context, id int) (model.Post, error) {
	if err := ctx.Err(); err != nil {
		return model.Post{}, err
	}

	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, post := range r.posts {
		if post.ID == id {
			return post, nil
		}
	}

	return model.Post{}, ErrPostNotFound
}

func (r *InMemoryPostRepository) Update(ctx context.Context, post model.Post) (model.Post, error) {
	if err := ctx.Err(); err != nil {
		return model.Post{}, err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	for i, existing := range r.posts {
		if existing.ID == post.ID {
			post.CreatedAt = existing.CreatedAt
			post.UpdatedAt = time.Now()
			r.posts[i] = post
			return post, nil
		}
	}

	return model.Post{}, ErrPostNotFound
}

func (r *InMemoryPostRepository) Delete(ctx context.Context, id int) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	for i, post := range r.posts {
		if post.ID == id {
			r.posts = append(r.posts[:i], r.posts[i+1:]...)
			return nil
		}
	}

	return ErrPostNotFound
}
