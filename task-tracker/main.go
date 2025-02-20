package main

import (
	"fmt"
	_ "fmt"
	"os"
	"strconv"
	"task-tracker/enum"
	"task-tracker/repository"
	"task-tracker/service"
)

func main() {
	taskRepository := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepository)

	inputs := os.Args[1:]

	if len(inputs) == 0 {
		fmt.Println("Please select an action (add, update, delete, get, list, mark-in-progress, mark-done)")
		os.Exit(0)
	} else {
		switch inputs[0] {
		case "add":
			if len(inputs) != 2 {
				fmt.Println("Please provide parameter (Task Description)")
				os.Exit(0)
			}
			err := taskService.AddTask(inputs[1])
			if err != nil {
				fmt.Println(err.Error())
			}
		case "list":
			if len(inputs) > 1 {
				_, err := taskService.GetTaskByStatus(inputs[1])
				if err != nil {
					fmt.Println(err.Error())
				}
			} else {
				_, err := taskService.GetAllTasks()
				if err != nil {
					fmt.Println(err.Error())
				}
			}
		case "get":
			if len(inputs) != 2 {
				fmt.Println("Please provide parameter (Task ID)")
				os.Exit(0)
			}
			id, _ := strconv.Atoi(inputs[1])
			_, err := taskService.GetTaskById(id)
			if err != nil {
				fmt.Println(err.Error())
			}
		case "mark-in-progress":
			if len(inputs) != 2 {
				fmt.Println("Please provide parameter (Task ID)")
				os.Exit(0)
			}
			id, _ := strconv.Atoi(inputs[1])
			err := taskService.UpdateTaskStatus(id, enum.StatusInProgress)
			if err != nil {
				fmt.Println(err.Error())
			}
		case "mark-done":
			if len(inputs) != 2 {
				fmt.Println("Please provide parameter (Task ID)")
				os.Exit(0)
			}
			id, _ := strconv.Atoi(inputs[1])
			err := taskService.UpdateTaskStatus(id, enum.StatusDone)
			if err != nil {
				fmt.Println(err.Error())
			}
		case "update":
			if len(inputs) == 1 {
				fmt.Println("Please provide parameter (Task ID)")
				os.Exit(0)
			} else if len(inputs) == 2 {
				fmt.Println("Please provide parameter (Task Description)")
				os.Exit(0)
			}

			id, _ := strconv.Atoi(inputs[1])
			description := inputs[2]
			err := taskService.UpdateTaskDescription(id, description)
			if err != nil {
				fmt.Println(err.Error())
			}
		case "delete":
			if len(inputs) != 2 {
				fmt.Println("Please provide parameter (Task ID)")
				os.Exit(0)
			}
			id, _ := strconv.Atoi(inputs[1])
			err := taskService.DeleteTask(id)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}
