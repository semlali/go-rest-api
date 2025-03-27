package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Structure pour représenter un utilisateur
type User struct {
	ID   int
	Name string
}

// Simuler une base de données d'utilisateurs
var users = []User{
	{ID: 1, Name: "user1 aa"},
	{ID: 2, Name: "user2 bb"},
	{ID: 3, Name: "user3 cc"},
}

// Handler pour la route "/"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Définir le type de contenu comme "texte" et renvoyer un message simple
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Bonjour, bienvenue dans l'API Go en texte brut!"))
}

// Handler pour la route "/user/{id}"
func userHandler(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de l'URL
	params := mux.Vars(r)
	idStr := params["id"]

	// Convertir l'ID en entier
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalide", http.StatusBadRequest)
		return
	}

	// Chercher l'utilisateur dans notre "base de données"
	var user *User
	for _, u := range users {
		if u.ID == id {
			user = &u
			break
		}
	}

	// Si l'utilisateur n'est pas trouvé
	if user == nil {
		http.Error(w, "Utilisateur non trouvé", http.StatusNotFound)
		return
	}

	// Définir le type de contenu comme "texte" et renvoyer un message avec les informations de l'utilisateur
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	message := fmt.Sprintf("Utilisateur : %s, ID : %d", user.Name, user.ID)
	_, err = w.Write([]byte(message))
	if err != nil {
		fmt.Println("Erreur lors de l'appel à l'utilisateur:", err)
	}

}

func main() {

	// Créer un nouveau routeur
	r := mux.NewRouter()

	// Définir les routes
	http.HandleFunc("/", helloHandler)
	r.HandleFunc("/user/{id}", userHandler).Methods("GET")

	// Démarrer le serveur sur le port 8080
	fmt.Println("Serveur en écoute sur http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}

}
