package main

import (
	"fmt"
	"os"
)

func about() {
	fmt.Println("")
	fmt.Println("Available commands: ")
	fmt.Println("\"Add\"     adds a task         Usage: task Add <description>")
	fmt.Println("\"Update\"  updates the task    Usage: task Update <task id> <task description>")
	fmt.Println("\"Delete\"  Delete the task     Usage: task Delete <task id>")
	fmt.Println("\"List\"    List the task    	 Usage: task List --Flag")
	fmt.Println()
	fmt.Println("Flags: ")
	fmt.Println("--done  			List all the task which are not done")
	fmt.Println("--todo   			List all the task that are done")
	fmt.Println("--in-progress   	List all the task that are done")
	fmt.Println()
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("	Usage: tasks <command> <description>")
		fmt.Println("	Use tasks about to see valid commands")
		return
	}

	switch os.Args[1] {
	case "Add":
		
	case "Update":
	case "Delete":
	case "List":
	case "about":
		about()
	default:
		fmt.Println("Invalid Command:  ", os.Args[1])
	}
}
