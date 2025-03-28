package repo

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Connexion à la base de données MariaDB
func Connect() (*sql.DB, error) {
	// Configurer la connexion à la base de données MariaDB
	dsn := "root@tcp(127.0.0.1:3306)/userdb"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Erreur de connexion à la base de données : %w", err)
	}

	// Vérifier la connexion
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Erreur lors de la vérification de la connexion à la base de données : %w", err)
	}

	return db, nil
}
