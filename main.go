/*
Gestion de liste de tâches (To-Do List)

Objectif : Créer un programme en Go qui permet de gérer une liste de tâches avec les opérations suivantes :
	1.	Ajouter une tâche.
	2.	Afficher toutes les tâches.
	3.	Marquer une tâche comme terminée.
	4.	Supprimer une tâche.

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
        - marquerCommeTerminee(id int): Marque une tâche comme terminée en utilisant son ID.
        - supprimerTache(id int): Supprime une tâche de la liste en utilisant son ID.
    3.	Interaction avec l’utilisateur :
    •	Dans main, proposer un menu textuel simple pour permettre à l’utilisateur de choisir une action :
        1. Ajouter une tâche
        2. Afficher les tâches
        3. Marquer une tâche comme terminée
        4. Supprimer une tâche
        5. Quitter
*/
