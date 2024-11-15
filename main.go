package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Class TodoList
type TodoList struct {
	Id    int
	Name  string
	Tasks []Task
}

// Class TodoListHandler
type TodoListHandler struct {
	Lists map[string]*TodoList
}

// Constructor
func ListHandler() *TodoListHandler {
	return &TodoListHandler{
		Lists: make(map[string]*TodoList),
	}
}

// Methods
func (handler *TodoListHandler) ShowLists() {
	if handler.IsEmpty() {
		fmt.Println("\nNo new lists.")
		return
	}

	fmt.Println("\nCurrent lists:")
	for _, list := range handler.Lists {
		fmt.Printf("- %s\n", list.Name)
	}
}

func (handler *TodoListHandler) CreateList(scanner *bufio.Scanner) {
	fmt.Println("\nEnter list name: ")
	scanner.Scan()
	listName := scanner.Text()

	if handler.GetList(listName) != nil {
		fmt.Println("\nList name already in use")
		return
	}

	handler.SetEmptyList(listName)
	fmt.Printf("\nNew list \"%s\" created successfully!", listName)
}

func (handler *TodoListHandler) UpdateList(scanner *bufio.Scanner) {
	if handler.IsEmpty() {
		fmt.Println("\nYou need to add a new list.")
		return
	}

	fmt.Println("\nEnter the list name to update:")
	scanner.Scan()
	listName := scanner.Text()
	list := handler.GetList(listName)

	if list == nil {
		fmt.Println("\nList does not exist.")
		return
	}

	for {
		fmt.Printf("\nUpdating list \"%s\":", listName)
		fmt.Println("1. Show tasks")
		fmt.Println("2. Create task")
		fmt.Println("3. Mark task as \"done\"")
		fmt.Println("4. Delete task")
		fmt.Println("5. Back to main menu")

		fmt.Print("Choose an option: ")
		scanner.Scan()

		switch scanner.Text() {
		// TODO: Create Task methods
		case "1":
			showTasks(list)
		case "2":
			addTask(scanner, list)
		case "3":
			updateTask(scanner, list)
		case "4":
			deleteTask(scanner, list)
		case "5":
			return
		default:
			fmt.Println("Invalid option, choose an one from 1 to 5.")
		}
	}
}

func (handler *TodoListHandler) DeleteList(scanner *bufio.Scanner) {
	if handler.IsEmpty() {
		fmt.Println("\nYou need to add a new list.")
		return
	}

	fmt.Println("\nEnter the list name to delete:")
	scanner.Scan()
	listName := scanner.Text()

	if handler.GetList(listName) == nil {
		fmt.Println("\nList does not exist.")
		return
	}

	delete(handler.Lists, listName)
	fmt.Printf("\nList \"%s\" deleted successfully!", listName)
}

func (handler *TodoListHandler) SetEmptyList(name string) {
	handler.Lists[name] = &TodoList{
		Id:    1,
		Name:  name,
		Tasks: []Task{},
	}
}

func (handler *TodoListHandler) GetList(name string) *TodoList {
	return handler.Lists[name]
}

func (handler *TodoListHandler) IsEmpty() bool {
	return len(handler.Lists) == 0
}

type Task struct {
	Id    int
	Title string
	Done  bool
}

func showTasks(list *TodoList) {
	if len(list.Tasks) == 0 {
		fmt.Println("No new tasks.")
		return
	}

	fmt.Printf("Tasks in list \"%s\":\n", list.Name)
	for _, task := range list.Tasks {
		taskDone := "(task in progress)"
		if task.Done {
			taskDone = "(task done)"
		}
		fmt.Printf("%d. %s %s\n", task.Id, task.Title, taskDone)
	}
}

func handleUserInput(scanner *bufio.Scanner) {
	handler := ListHandler()

	for {
		fmt.Println("\nTODO list handler:")
		fmt.Println("1. Show lists")
		fmt.Println("2. Create new list")
		fmt.Println("3. Add tasks to a list")
		fmt.Println("4. Delete a list")
		fmt.Println("5. Exit")

		fmt.Print("Choose an option: ")
		scanner.Scan()

		switch scanner.Text() {
		case "1":
			handler.ShowLists()
		case "2":
			handler.CreateList(scanner)
		case "3":
			handler.UpdateList(scanner)
		case "4":
			handler.DeleteList(scanner)
		case "5":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option, choose an one from 1 to 5.")
		}
	}
}

func main() {
	handleUserInput(bufio.NewScanner(os.Stdin))
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
	list.Tasks = append(list.Tasks, task)
	list.Id++
	fmt.Printf("Task \"%s\" added successfully to list \"%s\"!\n", taskTitle, list.Name)
}

func updateTask(scanner *bufio.Scanner, list *TodoList) {
	if len(list.Tasks) == 0 {
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

	for i := range list.Tasks {
		if list.Tasks[i].Id == id {
			if !list.Tasks[i].Done {
				list.Tasks[i].Done = true
				fmt.Printf("Task \"%s\" updated successfully from list \"%s\"!\n", list.Tasks[i].Title, list.Name)
				break
			}
		}
	}
}

func deleteTask(scanner *bufio.Scanner, list *TodoList) {
	if len(list.Tasks) == 0 {
		fmt.Println("You need to add a new task.")
		return
	}

	fmt.Print("\nEnter the task id to delete: ")
	scanner.Scan()
	input := scanner.Text()

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid id.")
		return
	}

	idExists := false
	for i := range list.Tasks {
		if list.Tasks[i].Id == id {
			taskTitle := list.Tasks[i].Title
			list.Tasks = append(list.Tasks[:i], list.Tasks[i+1:]...)
			fmt.Printf("Task \"%s\" deleted successfully from list \"%s\"!\n", taskTitle, list.Name)
			idExists = true
			break
		}
	}

	if !idExists {
		fmt.Printf("Task with id %d not found.\n", id)
	}
}
