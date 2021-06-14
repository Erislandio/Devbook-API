package database

import (
	"api/src/config"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Connect - inicia as configuracoes de conexao com o baco sqlite
func Connect() (*sql.DB, error) {

	database, erro := sql.Open("sqlite3", config.Path)

	if erro != nil {
		fmt.Println(erro)
		return nil, erro
	}

	if erro = database.Ping(); erro != nil {
		database.Close()
		fmt.Println("Error - ping database: sqlite")
		return nil, erro
	}

	return database, nil

}
