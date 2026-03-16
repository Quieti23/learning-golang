package store

import (
	"errors"
	"sync"

	"month01todoapi/model"
)

var ErrTaskNotFound = errors.New("task not found")

type TaskStore struct {
	mu     sync.Mutex
	nextID int
	tasks  []model.Task
}

func NewTaskStore() *TaskStore {
	return &TaskStore{
		nextID: 1,
		tasks:  make([]model.Task, 0),
	}
}

func (s *TaskStore) Create(title string) model.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := model.Task{
		ID:    s.nextID,
		Title: title,
		Done:  false,
	}
	s.nextID++
	s.tasks = append(s.tasks, task)
	return task
}

func (s *TaskStore) List() []model.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]model.Task, len(s.tasks))
	copy(result, s.tasks)
	return result
}

func (s *TaskStore) GetByID(id int) (model.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, task := range s.tasks {
		if task.ID == id {
			return task, nil
		}
	}

	return model.Task{}, ErrTaskNotFound
}

func (s *TaskStore) Update(id int, title string, done bool) (model.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, task := range s.tasks {
		if task.ID == id {
			task.Title = title
			task.Done = done
			s.tasks[i] = task
			return task, nil
		}
	}

	return model.Task{}, ErrTaskNotFound
}

func (s *TaskStore) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, task := range s.tasks {
		if task.ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}

	return ErrTaskNotFound
}