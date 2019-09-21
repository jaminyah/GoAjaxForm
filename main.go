package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"database/sql"
	"encoding/json"

	_"github.com/mattn/go-sqlite3"
)

type UserComment struct {
	Id int				`json:"id"`
	Name string			`json:"name"`
	Message string		`json:"comment"`
	Timestamp string	`json:"date"`
}

func receiveAjax(w http.ResponseWriter, r *http.Request) {

	// Setup database
	database, _ := sql.Open("sqlite3", "./wxalert.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS comments (id INTEGER PRIMARY KEY, username TEXT, comment TEXT, date TEXT)")
	statement.Exec()	
	
	fmt.Println("Received ajax data.")

	// Invoke ParseForm before reading form values
	r.ParseForm()
	for key, value := range r.Form {
		fmt.Printf("%s: %s\n", key, value)
	}

	// Acknowledge receiving the data
	fmt.Printf("User name: %s\n", r.FormValue("username"))
	fmt.Printf("Message body: %s\n", r.FormValue("message"))

	// Format current time
	now := time.Now()
	//fmt.Println(now.Format("Jan 2, 2006 - 3:04 pm MST"))
	var now_time string = now.Format("Jan 2, 2006 - 3:04 pm MST")

	// Client data received
	user_name := r.FormValue("username")
	user_message := r.FormValue("message")
	
	// Insert into database
	statement, _ = database.Prepare("INSERT INTO comments (username, comment, date) VALUES (?, ?, ?)")
	statement.Exec(user_name, user_message, now_time)

	// Read database
	rows, _ := database.Query("SELECT id, username, comment, date FROM comments ORDER BY id DESC LIMIT 0, 10")

	var db_id int
	var db_name string
	var db_message string
	var db_timestamp string
	var user_comment UserComment
	var user_comments []UserComment

	for rows.Next() {
		rows.Scan(&db_id, &db_name, &db_message, &db_timestamp)
		//fmt.Println(strconv.Itoa(id) + ": " + first + " " + last)
	
		user_comment.Id = db_id
		user_comment.Name = db_name
		user_comment.Message = db_message
		user_comment.Timestamp = db_timestamp
		user_comments = append(user_comments, user_comment)
	}

	json_str, _ := json.Marshal(user_comments)
	fmt.Printf("%s\n", json_str)
	
	// Send back data to client
	//fmt.Fprintf(w, "%s: <span> %s</span><br/>", r.FormValue("username"), r.FormValue("message"))
	fmt.Fprint(w, json_str)
	
 }

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))

    http.HandleFunc("/receive", receiveAjax)

	fmt.Printf("Starting HTTP server...\n")
    if err := http.ListenAndServe(":8088", nil); err != nil {
        log.Fatal(err)
    }
}