package initializer

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

var Db *sql.DB

func ConnectToDb() {
	var err error
	Db, err = sql.Open("pgx", "host=localhost user=cna password=1388 dbname=weblog sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to database", err)
	}
	fmt.Println("connected to database")
	err = Db.Ping()
	if err != nil {
		log.Fatal("cannot ping to database", err)
	}
	fmt.Println("ping database successfully")
}
