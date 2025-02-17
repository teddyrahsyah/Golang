package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"task-tracker/model"
	"time"
)

type TaskRepository interface {
	GetAll() ([]model.Task, error)
	GetById(id int) (*model.Task, error)
	GetByStatus(status string) ([]model.Task, error)
	Add(task model.Task) error
	Update(id int, updatedTask model.Task) error
	Delete(id int) error
}

type taskRepository struct {
	filePath string
	mu       sync.Mutex
}

func NewTaskRepository() TaskRepository {
	dataDir := "data"
	filePath := filepath.Join(dataDir, "tasks.json")

	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		os.Mkdir(dataDir, os.ModePerm)
	}

	// Check if file exists, if not create it with valid JSON []
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		os.WriteFile(filePath, []byte("[]"), 0644)
	} else {
		// If file exists but is empty, write valid JSON []
		fixEmptyFile(filePath)
	}
	return &taskRepository{filePath: filePath}
}

func fixEmptyFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return // Can't open file, assume it's fine
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return
	}

	// If file is empty, write valid JSON []
	if stat.Size() == 0 {
		file.Write([]byte("[]")) // Write valid JSON
	}
}

func (repo *taskRepository) readTasks() ([]model.Task, error) {
	file, err := os.Open(repo.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []model.Task
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func (repo *taskRepository) writeTasks(task []model.Task) error {
	file, err := os.Create(repo.filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(task); err != nil {
		return err
	}
	return nil
}

/*
METHOD IMPLEMENTATION
*/
func (repo *taskRepository) GetAll() ([]model.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	tasks, err := repo.readTasks()
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (repo *taskRepository) GetById(id int) (*model.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	tasks, err := repo.readTasks()
	if err != nil {
		return nil, err
	}

	for _, task := range tasks {
		if task.Id == id {
			return &task, nil
		}
	}
	return nil, errors.New("Task with ID : " + strconv.Itoa(id) + " not found")
}

func (repo *taskRepository) GetByStatus(status string) ([]model.Task, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	tasks, err := repo.readTasks()
	if err != nil {
		return nil, err
	}

	var filteredTask []model.Task
	for _, task := range tasks {
		if task.Status == status {
			filteredTask = append(filteredTask, task)
		}
	}

	if len(filteredTask) == 0 {
		return nil, errors.New("Task with status '" + status + "' not found")
	}

	return filteredTask, nil
}

func (repo *taskRepository) Add(task model.Task) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	tasks, err := repo.readTasks()
	if err != nil {
		return err
	}

	task.Id = len(tasks) + 1
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	tasks = append(tasks, task)
	err = repo.writeTasks(tasks)
	if err != nil {
		return err
	}
	fmt.Println("Task added successfully (ID : " + strconv.Itoa(task.Id) + ")")
	return nil
}

func (repo *taskRepository) Update(id int, updatedTask model.Task) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	tasks, err := repo.readTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.Id == id {
			tasks[i] = updatedTask
			err := repo.writeTasks(tasks)
			if err != nil {
				return err
			}
			fmt.Println("Task updated successfully (ID : " + strconv.Itoa(id) + ")")
			return err
		}
	}

	return errors.New("Task with ID : " + strconv.Itoa(id) + " not found")
}

func (r *taskRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	tasks, err := r.readTasks()
	if err != nil {
		return err
	}

	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted successfully (ID : " + strconv.Itoa(id) + ")")
			return r.writeTasks(tasks)
		}
	}
	return errors.New("Task with ID : " + strconv.Itoa(id) + " not found")
}
