package repository

import (
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

func (r *InMemoryPostRepository) List() []model.Post {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]model.Post, len(r.posts))
	copy(result, r.posts)
	return result
}

func (r *InMemoryPostRepository) Create(post model.Post) model.Post {
	r.mu.Lock()
	defer r.mu.Unlock()

	now := time.Now()
	post.ID = r.nextID
	post.CreatedAt = now
	post.UpdatedAt = now

	r.nextID++
	r.posts = append(r.posts, post)
	return post
}
