package service

import (
	"encoding/json"
	"errors"
	"os"
	"task-tracker/enum"
	"task-tracker/model"
	"task-tracker/repository"
	"time"
)

type TaskService interface {
	GetAllTasks() ([]model.Task, error)
	GetTaskById(id int) (*model.Task, error)
	GetTaskByStatus(status string) ([]model.Task, error)
	AddTask(desc string) error
	UpdateTaskStatus(id int, status string) error
	UpdateTaskDescription(id int, desc string) error
	DeleteTask(id int) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) GetAllTasks() ([]model.Task, error) {
	tasks, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	err = json.NewEncoder(os.Stdout).Encode(tasks)
	if err != nil {
		return nil, err
	}
	return tasks, err
}

func (s *taskService) GetTaskById(id int) (*model.Task, error) {
	if id == 0 {
		return nil, errors.New("Invalid Task ID : Please enter a number")
	}

	tasks, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}
	err = json.NewEncoder(os.Stdout).Encode(tasks)
	if err != nil {
		return nil, err
	}
	return tasks, err
}

func (s *taskService) GetTaskByStatus(status string) ([]model.Task, error) {

	validStatuses := map[string]bool{
		enum.StatusTodo:       true,
		enum.StatusInProgress: true,
		enum.StatusDone:       true,
	}

	if !validStatuses[status] {
		return nil, errors.New("invalid status (" + status + ")")
	}

	tasks, err := s.repo.GetByStatus(status)
	if err != nil {
		return nil, err
	}
	err = json.NewEncoder(os.Stdout).Encode(tasks)
	if err != nil {
		return nil, err
	}
	return tasks, err
}

func (s *taskService) AddTask(desc string) error {
	task := model.Task{
		Description: desc,
		Status:      enum.StatusTodo,
	}
	return s.repo.Add(task)
}

func (s *taskService) UpdateTaskStatus(id int, status string) error {
	if id == 0 {
		return errors.New("Invalid Task ID : Please enter a number")
	}

	task, err := s.repo.GetById(id)
	if err != nil {
		return err
	}
	task.Status = status
	task.UpdatedAt = time.Now()
	return s.repo.Update(id, *task)
}

func (s *taskService) UpdateTaskDescription(id int, desc string) error {
	if id == 0 {
		return errors.New("Invalid Task ID : Please enter a number")
	}

	task, err := s.repo.GetById(id)
	if err != nil {
		return err
	}
	task.Description = desc
	task.UpdatedAt = time.Now()
	return s.repo.Update(id, *task)
}

func (s *taskService) DeleteTask(id int) error {
	if id == 0 {
		return errors.New("Invalid Task ID : Please enter a number")
	}

	return s.repo.Delete(id)
}
