package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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
	if handler.EmptyHandler() {
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

	if handler.Lists[listName] != nil {
		fmt.Println("\nList name already in use")
		return
	}

	handler.Lists[listName] = NewTodoList(listName)
	fmt.Printf("\nNew list \"%s\" created successfully!\n", listName)
}

func (handler *TodoListHandler) UpdateList(scanner *bufio.Scanner) {
	if handler.EmptyHandler() {
		fmt.Println("\nYou need to add a new list.")
		return
	}

	fmt.Println("\nEnter the list name to update: ")
	scanner.Scan()
	listName := scanner.Text()

	list, exists := handler.Lists[listName]
	if !exists {
		fmt.Println("\nList does not exist.")
		return
	}

	for {
		fmt.Printf("\nUpdating list \"%s\":\n", listName)
		fmt.Println("1. Show tasks")
		fmt.Println("2. Create task")
		fmt.Println("3. Mark task as \"done\"")
		fmt.Println("4. Delete task")
		fmt.Println("5. Back to main menu")

		fmt.Print("Choose an option: ")
		scanner.Scan()

		switch scanner.Text() {
		case "1":
			list.ShowTasks()
		case "2":
			list.AddTask(scanner, list)
		case "3":
			list.UpdateTask(scanner, list)
		case "4":
			list.DeleteTask(scanner, list)
		case "5":
			return
		default:
			fmt.Println("Invalid option, choose an one from 1 to 5.")
		}
	}
}

func (handler *TodoListHandler) DeleteList(scanner *bufio.Scanner) {
	if handler.EmptyHandler() {
		fmt.Println("\nYou need to add a new list.")
		return
	}

	fmt.Println("\nEnter the list name to delete: ")
	scanner.Scan()
	listName := scanner.Text()

	if handler.Lists[listName] == nil {
		fmt.Println("\nList does not exist.")
		return
	}

	delete(handler.Lists, listName)
	fmt.Printf("\nList \"%s\" deleted successfully!\n", listName)
}

func (handler *TodoListHandler) EmptyHandler() bool {
	return len(handler.Lists) == 0
}

// Class TodoList
type TodoList struct {
	Id    int
	Name  string
	Tasks []*Task
}

// Constructor
func NewTodoList(name string) *TodoList {
	return &TodoList{
		Id:    1,
		Name:  name,
		Tasks: []*Task{},
	}
}

// Methods
func (todoList *TodoList) ShowTasks() {
	if todoList.EmptyTasks() {
		fmt.Println("\nNo new tasks.")
		return
	}

	fmt.Printf("\nTasks in list \"%s\":\n", todoList.Name)
	for _, task := range todoList.Tasks {
		taskDone := "(task in progress)"
		if task.Done {
			taskDone = "(task done)"
		}
		fmt.Printf("%d. %s %s\n", task.Id, task.Title, taskDone)
	}
}

func (todoList *TodoList) AddTask(scanner *bufio.Scanner, list *TodoList) {
	fmt.Println("\nEnter task name:")
	scanner.Scan()
	taskTitle := scanner.Text()

	list.Tasks = append(list.Tasks, NewTask(list.Id, taskTitle))
	list.Id++
	fmt.Printf("\nTask \"%s\" added successfully to list \"%s\"!\n", taskTitle, list.Name)
}

func (todoList *TodoList) UpdateTask(scanner *bufio.Scanner, list *TodoList) {
	if list.EmptyTasks() {
		fmt.Println("\nYou need to add a new task.")
		return
	}

	fmt.Println("\nEnter the task id to update: ")
	scanner.Scan()
	input := scanner.Text()

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("\nInvalid Id.")
		return
	}

	for _, task := range list.Tasks {
		if task.Id == id {
			if !task.Done {
				task.MarkAsDone()
				fmt.Printf("\nTask \"%s\" updated successfully from list \"%s\"!\n", task.Title, list.Name)
				break
			}
		}
	}
}

func (todoList *TodoList) DeleteTask(scanner *bufio.Scanner, list *TodoList) {
	if list.EmptyTasks() {
		fmt.Println("\nYou need to add a new task.")
		return
	}

	fmt.Print("\nEnter the task id to delete: ")
	scanner.Scan()
	input := scanner.Text()

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("\nInvalid id.")
		return
	}

	if len(list.Tasks) < id {
		fmt.Printf("\nTask with id %d not found.\n", id)
		return
	}

	id = id - 1
	task := list.Tasks[id]
	taskTitle := task.Title
	list.Tasks = append(list.Tasks[:id], list.Tasks[id+1:]...)
	fmt.Printf("\nTask \"%s\" deleted successfully from list \"%s\"!\n", taskTitle, list.Name)
}

func (todoList *TodoList) EmptyTasks() bool {
	return len(todoList.Tasks) == 0
}

// Class Task
type Task struct {
	Id    int
	Title string
	Done  bool
}

// Constructor
func NewTask(id int, name string) *Task {
	return &Task{
		Id:    id,
		Title: name,
		Done:  false,
	}
}

// Method
func (task *Task) MarkAsDone() {
	task.Done = true
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
