package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Id    int
	Title string
	Done  bool
}

var tasks []Task
var nextId = 1

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	handleUserInput(scanner)
}

func handleUserInput(scanner *bufio.Scanner) {
	for {
		fmt.Println("\nMy First TODO List App in Go:")
		fmt.Println("1. View Tasks")
		fmt.Println("2. Add Task")
		fmt.Println("3. Mark task as done")
		fmt.Println("4. Delete task")
		fmt.Println("5. Exit")

		fmt.Print("Choose an option: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			viewTasks()
		case "2":
			addTask(scanner)
		case "3":
			updateTask(scanner)
		case "4":
			deleteTask(scanner)
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			return
		}
	}
}

func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("No new tasks.")
		return
	}

	fmt.Println("\nList:")
	for _, task := range tasks {
		taskDone := "(task in progress)"
		if task.Done {
			taskDone = "(task done)"
		}
		fmt.Printf("%d. %s %s\n", task.Id, task.Title, taskDone)
	}
}

func addTask(scanner *bufio.Scanner) {
	fmt.Println("Enter task description: ")
	scanner.Scan()
	taskTitle := scanner.Text()

	task := Task{
		Id:    nextId,
		Title: taskTitle,
		Done:  false,
	}
	tasks = append(tasks, task)
	nextId++
	fmt.Printf("Task \"%s\" added successfully!\n", taskTitle)
}

func updateTask(scanner *bufio.Scanner) {
	if len(tasks) == 0 {
		fmt.Println("You need to add a new task.")
		return
	}

	fmt.Print("\nEnter the task id to update: ")
	scanner.Scan()
	input := scanner.Text()

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid Id.")
		return
	}

	for i := range tasks {
		if tasks[i].Id == id {
			if !tasks[i].Done {
				tasks[i].Done = true
				fmt.Printf("Task \"%s\" updated successfully!\n", tasks[i].Title)
				break
			}
		}
	}
}

func deleteTask(scanner *bufio.Scanner) {
	if len(tasks) == 0 {
		fmt.Println("You need to add a new task.")
		return
	}

	fmt.Print("\nEnter the task id to delete: ")
	scanner.Scan()
	input := scanner.Text()

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid Id.")
		return
	}

	idExists := false
	for i := range tasks {
		if tasks[i].Id == id {
			taskTitle := tasks[i].Title
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Printf("Task \"%s\" deleted successfully!\n", taskTitle)
			idExists = true
			break
		}
	}

	if !idExists {
		fmt.Printf("Task with Id %d not found.\n", id)
	}
}
