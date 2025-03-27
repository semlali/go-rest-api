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

// Handler pour la route "/"
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Création d'un message de réponse
	response := Message{
		Text: "Bonjour, bienvenu dans l'API Go!",
	}

	// Convertir la réponse en JSON et l'envoyer
	json.NewEncoder(w).Encode(response)
}

func main() {

	// Définir les routes
	http.HandleFunc("/", helloHandler)

	// Démarrer le serveur sur le port 8080
	fmt.Println("Serveur en écoute sur http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}

}
