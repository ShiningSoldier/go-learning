package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

var db, err = sqlx.Connect("mysql", "root:@tcp(127.0.0.1:3306)/go_orders")

func main() {
	handleDatabase()
	checkErr(err)

}

func handleDatabase() {
	ordersQuery := `CREATE TABLE IF NOT EXISTS orders (
        uuid INT PRIMARY KEY AUTO_INCREMENT,
        book_uuid INT NOT NULL,
        description TEXT,
		created_at datetime default CURRENT_TIMESTAMP,
        updated_at datetime default CURRENT_TIMESTAMP,
        deleted_at datetime
    );`

	db.MustExec(ordersQuery)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
