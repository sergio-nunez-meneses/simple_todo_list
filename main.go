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

type TodoList struct {
	Id   int
	Name string
	List []Task
}

var todoLists = make(map[string]*TodoList)
var tasks []Task
var nextId = 1

func main() {
	handleUserInput(bufio.NewScanner(os.Stdin))
}

func handleUserInput(scanner *bufio.Scanner) {
	for {
		fmt.Println("\nTODO Lists:")
		fmt.Println("1. Show lists")
		fmt.Println("2. Create new list")
		fmt.Println("3. Add tasks to a list")
		fmt.Println("4. Delete a list")
		fmt.Println("5. Exit")

		fmt.Print("Choose an option: ")
		scanner.Scan()

		switch scanner.Text() {
		case "1":
			showLists()
		case "2":
			createList(scanner)
		case "3":
			handleList(scanner)
		case "4":
			deleteList(scanner)
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, choose an one from 1 to 5.")
		}
	}
}

func showLists() {
	if len(todoLists) == 0 {
		fmt.Println("No new lists.")
		return
	}

	fmt.Println("\nLists:")
	for _, list := range todoLists {
		fmt.Printf("%d. %s\n", list.Id, list.Name)
	}
}

func createList(scanner *bufio.Scanner) {
	fmt.Println("Enter list name: ")
	scanner.Scan()
	listName := scanner.Text()

	if listExists(listName) {
		fmt.Println("List name already in use")
		return
	}

	todoLists[listName] = &TodoList{
		Id:   1,
		Name: listName,
		List: []Task{},
	}
	fmt.Printf("New list \"%s\" created successfully!", listName)
}

func handleList(scanner *bufio.Scanner) {
	if len(todoLists) == 0 {
		fmt.Println("You need to add a new list.")
		return
	}

	fmt.Println("Enter the list name to update: ")
	scanner.Scan()
	listName := scanner.Text()

	if !listExists(listName) {
		fmt.Println("List does not exist.")
		return
	}

	list := todoLists[listName]
	for {
		fmt.Printf("Updating list \"%s\":\n", listName)
		fmt.Println("1. Show tasks")
		fmt.Println("2. Add task")
		fmt.Println("3. Mark task as \"done\"")
		fmt.Println("4. Delete task")
		fmt.Println("5. Back to main menu")

		fmt.Print("Choose an option: ")
		scanner.Scan()

		switch scanner.Text() {
		case "1":
			showTasks(list)
		case "2":
			addTask(scanner, list)
		case "3":
			updateTask(scanner)
		case "4":
			deleteTask(scanner)
		case "5":
			return
		default:
			fmt.Println("Invalid option, choose an one from 1 to 5.")
		}
	}
}

func deleteList(scanner *bufio.Scanner) {
	if len(todoLists) == 0 {
		fmt.Println("You need to add a new list.")
		return
	}

	fmt.Println("Enter the list name to delete: ")
	scanner.Scan()
	listName := scanner.Text()

	if !listExists(listName) {
		fmt.Println("List does not exist.")
		return
	}

	delete(todoLists, listName)
	fmt.Printf("List \"%s\" deleted successfully!", listName)
}

func listExists(listName string) bool {
	_, exists := todoLists[listName]
	return exists
}

func showTasks(list *TodoList) {
	if len(list.List) == 0 {
		fmt.Println("No new tasks.")
		return
	}

	fmt.Printf("Tasks in list \"%s\":\n", list.Name)
	for _, task := range list.List {
		taskDone := "(task in progress)"
		if task.Done {
			taskDone = "(task done)"
		}
		fmt.Printf("%d. %s %s\n", task.Id, task.Title, taskDone)
	}
}

func addTask(scanner *bufio.Scanner, list *TodoList) {
	fmt.Println("Enter task: ")
	scanner.Scan()
	taskTitle := scanner.Text()

	task := Task{
		Id:    list.Id,
		Title: taskTitle,
		Done:  false,
	}
	list.List = append(list.List, task)
	list.Id++
	fmt.Printf("Task \"%s\" added successfully to list \"%s\"!\n", taskTitle, list.Name)
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
