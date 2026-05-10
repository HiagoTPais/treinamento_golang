package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//stringConexao := "usuario:senha@tcp(127.0.0.1:3306)/meubanco"
	stringConexao := "root:password@tcp(127.0.0.1:3306)/golang"
	db, erro := sql.Open("mysql", stringConexao)

	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conexao esta aberta")
}
