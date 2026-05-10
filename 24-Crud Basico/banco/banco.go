package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Conectar banco
func Conectar(*sql.DB, error) {
	stringConexao := "root:password@tcp(127.0.0.1:3306)/golang"

	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
