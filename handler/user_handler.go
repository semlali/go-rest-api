package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-rest-api/repo"
	"net/http"
	"strconv"
)

// Structure pour gérer les utilisateurs via HTTP
type UserHandler struct {
	repo repo.UserRepository
}

// Crée un nouveau UserHandler
func NewUserHandler(repo repo.UserRepository) *UserHandler {
	return &UserHandler{repo}
}

// Handler pour la route "/user/{id}"
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	// Récupérer l'ID de l'URL
	params := mux.Vars(r)
	idStr := params["id"]

	// Convertir l'ID en entier
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID invalid", http.StatusBadRequest)
		return
	}

	// Chercher l'utilisateur dans la base de données via le repository
	user, err := h.repo.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Définir le type de contenu comme "texte" et renvoyer un message avec les informations de l'utilisateur
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	message := fmt.Sprintf("Utilisateur : %s, ID : %d", user.Name, user.ID)
	w.Write([]byte(message))
}

// Méthode pour enregistrer les routes : utilisée pour lier les routes HTTP au handler.
func (h *UserHandler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
}
