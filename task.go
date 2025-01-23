package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

const filename = "todo.json"

func about() {
	fmt.Println("\nAvailable commands:")
	fmt.Println("\"add\"     adds a task         Usage: task Add <description>")
	fmt.Println("\"update\"  updates the task    Usage: task Update <task id> <task description>")
	fmt.Println("\"delete\"  Delete the task     Usage: task Delete <task id>")
	fmt.Println("\"list\"    List the task       Usage: task List --Flag")
	fmt.Println("\nFlags:")
	fmt.Println("--done    List all tasks which are done")
	fmt.Println("--todo    List all tasks that are not done")
}

func readTasks() ([]task, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []task{}, nil
		}
		return nil, err
	}

	if len(data) == 0 {
		return []task{}, nil
	}

	var tasks []task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func writeTasks(tasks []task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func Add(description string) error {
	tasks, err := readTasks()
	if err != nil {
		return err
	}

	newTask := task{
		ID:          generateID(tasks),
		Description: description,
		Status:      false,
		CreatedAt:   time.Now().Format("2006-1-2 15:4:5"),
		UpdatedAt:   time.Now().Format("2006-1-2 15:4:5"),
	}

	tasks = append(tasks, newTask)

	if err := writeTasks(tasks); err != nil {
		return err
	}

	fmt.Printf("Task added successfully with ID: %d\n", newTask.ID)
	return nil
}

func Update(id int, newDescription string) error {
	tasks, err := readTasks()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now().Format("2006-1-2 15:4:5")

			if err := writeTasks(tasks); err != nil {
				return err
			}

			fmt.Printf("Task %d updated successfully\n", id)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func Delete(id int) error {
	tasks, err := readTasks()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)

			if err := writeTasks(tasks); err != nil {
				return err
			}

			fmt.Printf("Task %d deleted successfully\n", id)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func List(flag string) error {
	tasks, err := readTasks()
	if err != nil {
		return err
	}

	fmt.Printf("%-5s %-30s %-10s %-20s %-20s\n", "ID", "Description", "Status", "Created", "Updated")
	fmt.Println(strings.Repeat("-", 85))

	for _, t := range tasks {
		switch flag {
		case "done":
			if t.Status {
				fmt.Printf("%-5d %-30s %-10v %-20s %-20s\n", t.ID, t.Description, t.Status, t.CreatedAt, t.UpdatedAt)
			}
		case "todo":
			if !t.Status {
				fmt.Printf("%-5d %-30s %-10v %-20s %-20s\n", t.ID, t.Description, t.Status, t.CreatedAt, t.UpdatedAt)
			}
		default:
			fmt.Printf("%-5d %-30s %-10v %-20s %-20s\n", t.ID, t.Description, t.Status, t.CreatedAt, t.UpdatedAt)
		}
	}

	return nil
}

func Done(id int) error {
	tasks, err := readTasks()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = true
			tasks[i].UpdatedAt = time.Now().Format("2006-1-2 15:4:5")

			if err := writeTasks(tasks); err != nil {
				return err
			}

			fmt.Printf("Task %d completed\n", id)
			return nil
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func generateID(tasks []task) int {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].ID + 1
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: tasks <command> <args>")
		fmt.Println("Use 'tasks about' to see available commands")
		return
	}

	switch os.Args[1] {
	case "about":
		about()
	case "add":
		if len(os.Args) != 3 {
			fmt.Println("Usage: tasks Add <description>")
			return
		}
		if err := Add(os.Args[2]); err != nil {
			fmt.Println("Error:", err)
		}
	case "update":
		if len(os.Args) != 4 {
			fmt.Println("Usage: tasks Update <id> <new description>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		if err := Update(id, os.Args[3]); err != nil {
			fmt.Println("Error:", err)
		}
	case "delete":
		if len(os.Args) != 3 {
			fmt.Println("Usage: tasks Delete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID")
			return
		}
		if err := Delete(id); err != nil {
			fmt.Println("Error:", err)
		}
	case "done":
		if len(os.Args) != 3 {
			fmt.Println("Usage: tasks done <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			println("%s is not a valid int id", os.Args[2])
		}
		if err := Done(id); err != nil {
			fmt.Println("Error:", err)
		}
	case "list":
		if len(os.Args) != 3 {
			fmt.Println("Usage: tasks List --done or tasks List --todo")
			return
		}
		flag := os.Args[2][2:] // Remove --
		if err := List(flag); err != nil {
			fmt.Println("Error:", err)
		}
	default:
		fmt.Println("Invalid command")
		fmt.Println("Use 'tasks about' to see available commands")
	}
}
