package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-rest-api/handler"
	"go-rest-api/repo"
	"net/http"
)

func main() {
	// Se connecter à la base de données
	db, err := repo.Connect()
	if err != nil {
		fmt.Println("Erreur de connexion à la base de données:", err)
		return
	}
	defer db.Close()

	// Créer un repository pour les utilisateurs
	userRepo := repo.NewUserRepository(db)

	// Créer un handler pour les utilisateurs
	userHandler := handler.NewUserHandler(userRepo)

	// Créer un routeur et enregistrer les routes
	r := mux.NewRouter()
	userHandler.RegisterRoutes(r)

	// Démarrer le serveur sur le port 8080
	fmt.Println("Serveur en écoute sur http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Erreur lors du démarrage du serveur:", err)
	}
}
