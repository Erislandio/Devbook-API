package repositories

import (
	"api/src/models"
	"database/sql"
)

// User struct -repositorio de usuarios
type User struct {
	db *sql.DB
}

// NewUserRepo -cria uma nova instancia de repositorio de usuarios
func NewUserRepo(db *sql.DB) *User {
	return &User{
		db: db,
	}
}

func (userRepo *User) Create(user models.User) (uint64, error) {
	return 0, nil
}
