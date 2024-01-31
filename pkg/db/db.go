package db

import (
	"database/sql"
	"fmt"
	"log"

	"biletes/pkg/colors"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB = nil

func init() {
	Db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/biletes")
	if err != nil {
		panic(err.Error())
	}
	defer Db.Close()

	// check the connection
	err = Db.Ping()
	if err != nil {
		fmt.Print(colors.ColorRed, "Not Connected to DB!", colors.ColorEnd, "\n")
		log.Fatal(colors.ColorRed, err.Error(), colors.ColorEnd, "\n")
	}
	fmt.Print(colors.ColorGreen, "Connected to DB!", colors.ColorEnd, "\n")
	// return Db
}
