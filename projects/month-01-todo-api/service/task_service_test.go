package service

import (
	"errors"
	"testing"

	"month01todoapi/store"
)

func TestTaskServiceCreate(t *testing.T) {
	tests := []struct {
		name      string
		title     string
		wantTitle string
		wantErr   error
	}{
		{
			name:      "create task success",
			title:     "learn go testing",
			wantTitle: "learn go testing",
		},
		{
			name:    "empty title returns error",
			title:   "   ",
			wantErr: ErrTitleRequired,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			service := NewTaskService(store.NewTaskStore())

			task, err := service.Create(test.title)
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("got err %v, want %v", err, test.wantErr)
			}

			if test.wantErr != nil {
				return
			}

			if task.ID != 1 {
				t.Fatalf("got ID %d, want 1", task.ID)
			}

			if task.Title != test.wantTitle {
				t.Fatalf("got title %q, want %q", task.Title, test.wantTitle)
			}
		})
	}
}

func TestTaskServiceUpdate(t *testing.T) {
	service := NewTaskService(store.NewTaskStore())
	created, err := service.Create("learn go")
	if err != nil {
		t.Fatalf("create task failed: %v", err)
	}

	tests := []struct {
		name    string
		id      int
		input   UpdateTaskInput
		wantErr error
	}{
		{
			name: "update success",
			id:   created.ID,
			input: UpdateTaskInput{
				Title: "learn go updated",
				Done:  true,
			},
		},
		{
			name: "empty title returns error",
			id:   created.ID,
			input: UpdateTaskInput{
				Title: "   ",
				Done:  true,
			},
			wantErr: ErrTitleRequired,
		},
		{
			name: "missing task returns not found",
			id:   999,
			input: UpdateTaskInput{
				Title: "missing task",
				Done:  true,
			},
			wantErr: store.ErrTaskNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			task, err := service.Update(test.id, test.input)
			if !errors.Is(err, test.wantErr) {
				t.Fatalf("got err %v, want %v", err, test.wantErr)
			}

			if test.wantErr != nil {
				return
			}

			if task.Title != test.input.Title {
				t.Fatalf("got title %q, want %q", task.Title, test.input.Title)
			}

			if task.Done != test.input.Done {
				t.Fatalf("got done %v, want %v", task.Done, test.input.Done)
			}
		})
	}
}

func TestTaskServiceDelete(t *testing.T) {
	service := NewTaskService(store.NewTaskStore())
	created, err := service.Create("learn delete")
	if err != nil {
		t.Fatalf("create task failed: %v", err)
	}

	if err := service.Delete(created.ID); err != nil {
		t.Fatalf("delete existing task failed: %v", err)
	}

	if _, err := service.GetByID(created.ID); !errors.Is(err, store.ErrTaskNotFound) {
		t.Fatalf("got err %v, want %v", err, store.ErrTaskNotFound)
	}

	if err := service.Delete(999); !errors.Is(err, store.ErrTaskNotFound) {
		t.Fatalf("got err %v, want %v", err, store.ErrTaskNotFound)
	}
}
