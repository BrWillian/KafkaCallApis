package config

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

func DbConn() bool {
	fmt.Println("... Iniciando conexão com banco!")

	db, err := sqlx.Open("godror", `user="xxxx" password="xxxx" connectString="(DESCRIPTION=(ADDRESS=(PROTOCOL=TCP)(HOST=DTF-LBDEXP-DEV.datatraffic.com.br)(PORT=1521))(CONNECT_DATA=(Service_name=xe)))"`)
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Printf("Error connecting to the database: %s\n", err)
		return false
	}
	fmt.Println("... Conectado com sucesso!")
	return true

}
