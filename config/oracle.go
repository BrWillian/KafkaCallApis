package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

func DbConn() bool {
	fmt.Println("... Iniciando conexão com banco!")

	db, err := sqlx.Open("godror", os.Getenv("DATASOURCE_URL"))
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Printf("Error connecting to the database: %s\n", err)
		return false
	}
	fmt.Println("... Conectado com sucesso!")
	return true

}
