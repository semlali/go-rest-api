package main

import (
	"fmt"
	"net/http"
)

// Structure pour représenter un utilisateur
type User struct {
	ID   int
	Name string
}

// Handler pour la route "/"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Définir le type de contenu comme "texte" et renvoyer un message simple
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Bonjour, bienvenue dans l'API Go en texte brut!"))
}

// Handler pour obtenir un utilisateur par ID
func userHandler(w http.ResponseWriter, r *http.Request) {
	// Exemple d'utilisateur
	user := User{
		ID:   1,
		Name: "Samia Semlali",
	}

	// Définir le type de contenu comme "texte" pour une réponse en texte brut
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Créer un message avec les informations de l'utilisateur
	message := fmt.Sprintf("Utilisateur : %s, ID : %d", user.Name, user.ID)

	// Envoyer le message
	w.Write([]byte(message))

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
