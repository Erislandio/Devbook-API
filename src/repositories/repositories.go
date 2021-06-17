package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
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

	statement, erro := userRepo.db.Prepare(SqlCreateUser)

	if erro != nil {
		return 0, erro
	}

	result, erro := statement.Exec(user.Name, user.Nick, user.Email, user.Password, user.CreatedAt)

	if erro != nil {
		return 0, erro
	}

	lastInsertId, erro := result.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(lastInsertId), nil
}

func (repo *User) GetAll() (*sql.Rows, error) {

	rows, erro := repo.db.Query(SelectAllUsers)

	if erro != nil {
		return nil, erro
	}

	return rows, nil

}

func (repo *User) GetByNickName(nick string) (*sql.Rows, error) {

	statement := fmt.Sprintf(`%s "%%%s%%" ;`, SelectByNickName, nick)

	rows, erro := repo.db.Query(statement)

	if erro != nil {
		return nil, erro
	}

	return rows, nil
}

func (repo *User) GetByID(id string) (*sql.Rows, error) {

	statement := fmt.Sprintf(`%s %s;`, SelectByID, id)

	rows, erro := repo.db.Query(statement)

	if erro != nil {
		return nil, erro
	}

	return rows, nil
}
