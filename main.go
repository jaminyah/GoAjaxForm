/* Main driver function. 
 * Run: go run main.go fetch.go submit.go
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"

	_"github.com/mattn/go-sqlite3"
)

var database, db_err = sql.Open("sqlite3", "./wxalert.db")


type UserComment struct {
	Id int				`json:"id"`
	Name string			`json:"name"`
	Message string		`json:"comment"`
	Timestamp string	`json:"date"`
}

func main() {


	if db_err != nil {
		log.Fatal(db_err)
	}
	defer database.Close()

	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS comments (id INTEGER PRIMARY KEY, username TEXT, comment TEXT, date TEXT)")
	statement.Exec()	

	http.Handle("/", http.FileServer(http.Dir("static")))

    http.HandleFunc("/submit", submitAjax)

	http.HandleFunc("/comments", fetchComments)

	fmt.Printf("Starting HTTP server...\n")
    if server_err := http.ListenAndServe(":8088", nil); server_err != nil {
        log.Fatal(server_err)
    }
}