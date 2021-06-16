package database

import (
	"api/src/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// Connect - inicia as configuracoes de conexao com o baco sqlite
func Connect() (*sql.DB, error) {

	database, erro := sql.Open("mysql", config.SqlConnectionString)

	fmt.Println(config.SqlConnectionString)

	if erro != nil {
		fmt.Println(erro)
		return nil, erro
	}

	if erro = database.Ping(); erro != nil {
		database.Close()
		fmt.Println("Error - ping database: mysql")
		return nil, erro
	}

	return database, nil

}
