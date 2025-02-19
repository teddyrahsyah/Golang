# Task Tracker CLI

- **Description**: Task tracker is a project used to track and manage your tasks. A simple command line interface (CLI) to track what you need to do, what you have done, and what you are currently working on.
- **Features**:
    - **Add Tasks**: Create new tasks with a description.
    - **Update Tasks**: Modify existing task descriptions.
    - **Delete Tasks**: Remove tasks by their ID.
    - **Mark Tasks**: Set tasks as 'in-progress' or 'done'.
    - **List Tasks**:
        - All tasks
        - Tasks by status: 'done', 'to-do', or 'in-progress'
- **Tech Stack**: Go.

## Requirements

- **Programming Language**: Golang
- **Storage**: Tasks are stored in a JSON file in the current directory.

## 🛠️ Installation & Setup

To run this project, follow these steps:

```go
go build -o task-cli // Build task-tracker app binary

mv task-cli /usr/local/bin/ // Move it to a system-wide location

Now you can just run:
	task-cli {action} {parameter}
        ex: task-cli add "Programming with Go"
```

## ✒️ Example
```bash
The list of commands and their usage is given below:

# Adding a new task
task-cli add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
task-cli update 1 "Buy groceries and cook dinner"
task-cli delete 1

# Marking a task as in progress or done
task-cli mark-in-progress 1
task-cli mark-done 1

# Listing all tasks
task-cli list

# Listing tasks by status
task-cli list done
task-cli list todo
task-cli list in-progress
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📧 Contact

For any questions or collaborations, feel free to reach out:
- Email: teddyrahsyah@gmail.com
- LinkedIn: [Muhammad Teddy Rahmansyah](https://www.linkedin.com/in/teddy-rahsyah/)
- GitHub: [teddyrahsyah](https://github.com/teddyrahsyah)
---

⭐ If you like this repository, don't forget to give it a star! ⭐

