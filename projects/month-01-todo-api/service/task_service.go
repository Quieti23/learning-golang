package service

import (
	"errors"
	"strings"

	"month01todoapi/model"
	"month01todoapi/store"
)

var ErrTitleRequired = errors.New("title is required")

type TaskService struct {
	store *store.TaskStore
}

type UpdateTaskInput struct {
	Title string
	Done  bool
}

func NewTaskService(taskStore *store.TaskStore) *TaskService {
	return &TaskService{store: taskStore}
}

func (s *TaskService) Create(title string) (model.Task, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return model.Task{}, ErrTitleRequired
	}

	return s.store.Create(title), nil
}

func (s *TaskService) List() []model.Task {
	return s.store.List()
}

func (s *TaskService) GetByID(id int) (model.Task, error) {
	return s.store.GetByID(id)
}

func (s *TaskService) Update(id int, input UpdateTaskInput) (model.Task, error) {
	input.Title = strings.TrimSpace(input.Title)
	if input.Title == "" {
		return model.Task{}, ErrTitleRequired
	}

	return s.store.Update(id, input.Title, input.Done)
}

func (s *TaskService) Delete(id int) error {
	return s.store.Delete(id)
}
