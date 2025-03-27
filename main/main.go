package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Struct pour les données que nous allons renvoyer dans notre API
type Message struct {
	Text string `json:"text"`
}

// Structure pour représenter un utilisateur
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Handler pour la route "/"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Création d'un message de réponse
	response := Message{
		Text: "Bonjour, bienvenu dans l'API Go!",
	}

	// Convertir la réponse en JSON et l'envoyer
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println("error : ", err)
	}
}

// Handler pour obtenir un utilisateur par ID
func userHandler(w http.ResponseWriter, r *http.Request) {
	// Exemple d'utilisateur
	user := User{
		ID:   1,
		Name: "Samia Semlali",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Convertir l'utilisateur en JSON et l'envoyer en réponse
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Println("error : ", err)
	}
}

func main() {

	// Définir les routes
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/user", userHandler)

	// Démarrer le serveur sur le port 8080
	fmt.Println("Serveur en écoute sur http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}

}
