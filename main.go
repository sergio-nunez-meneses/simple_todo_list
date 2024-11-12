/*
Gestion de liste de tâches (To-Do List)

Objectif : Créer un programme en Go qui permet de gérer une liste de tâches avec les opérations suivantes :
	1.	Ajouter une tâche
	2.	Afficher toutes les tâches
	3.	Marquer une tâche comme terminée
	4.	Supprimer une tâche

Étapes :
    1.	Créer une structure Task
    •	La structure doit contenir :
        - un ID (entier unique pour identifier la tâche)
        - un Titre (texte pour décrire la tâche)
        - un booléen Fait pour indiquer si la tâche est terminée
    2.	Définir les fonctions principales :
        - ajouterTache(titre string): Ajoute une nouvelle tâche à la liste avec le titre fourni
        - afficherTaches(): Affiche toutes les tâches, en précisant celles qui sont terminées et celles qui sont en
            attente
        - marquerCommeTerminee(id int): Marque une tâche comme terminée en utilisant son ID
        - supprimerTache(id int): Supprime une tâche de la liste en utilisant son ID
    3.	Interaction avec l’utilisateur :
    •	Dans main, proposer un menu textuel simple pour permettre à l’utilisateur de choisir une action :
        1. Ajouter une tâche
        2. Afficher les tâches
        3. Marquer une tâche comme terminée
        4. Supprimer une tâche
        5. Quitter
*/

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

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			viewTasks()
		case "2":
			addTask(scanner)
		case "3":
			updateTask(scanner)
		default:
			return
		}
	}
}

func viewTasks() {
	fmt.Println("\nList:")
	if len(tasks) == 0 {
		fmt.Println("No new tasks.")
		return
	}

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
	text := scanner.Text()

	task := Task{
		Id:    nextId,
		Title: text,
		Done:  false,
	}
	tasks = append(tasks, task)
	nextId++
	fmt.Printf("Task \"%s\" added successfully!\n", text)
}

func updateTask(scanner *bufio.Scanner) {
	fmt.Print("\nEnter the task id to update: ")
	scanner.Scan()
	input := scanner.Text()

	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid Id.")
		return
	}

	if len(tasks) == 0 {
		fmt.Println("You need to add a new task.")
		return
	}

	for _, task := range tasks {
		if task.Id == id {
			if !task.Done {
				task.Done = true
				fmt.Println(task.Id)
				fmt.Println(task.Title)
				fmt.Println(task.Done)
				fmt.Println("Task updated successfully!")
				break
			}
		}
	}
}
