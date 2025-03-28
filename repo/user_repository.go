package repo

import (
	"database/sql"
	"fmt"
)

// Structure pour représenter un utilisateur
type User struct {
	ID   int
	Name string
}

// Interface pour un repository d'utilisateurs
type UserRepository interface {
	FindByID(id int) (*User, error)
}

// Implémentation du repository pour les utilisateurs
type userRepository struct {
	db *sql.DB
}

// Crée une nouvelle instance de UserRepository
func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

// Trouver un utilisateur par son ID
func (r *userRepository) FindByID(id int) (*User, error) {
	var user User
	query := "SELECT id, name FROM users WHERE id = ?"
	row := r.db.QueryRow(query, id)

	if err := row.Scan(&user.ID, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Utilisateur non trouvé")
		}
		return nil, fmt.Errorf("Erreur lors de la récupération des données : %w", err)
	}

	return &user, nil
}
