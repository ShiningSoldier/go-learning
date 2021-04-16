package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"strconv"
)

var database, err = sql.Open("sqlite3", "./tasksDb.db")

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Until       string `json:"until"`
	Status      int    `json:"status"`
}

var statuses = map[int]string{
	1: "Ready",
	2: "In progress",
	3: "Cancelled",
	4: "Done",
	5: "Expired",
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/task/{id}", getTask).Methods("GET")
	router.HandleFunc("/tasks", addTask).Methods("POST")
	router.HandleFunc("/task/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8765", router))
}

func getTasks(res http.ResponseWriter, req *http.Request) {
	checkErr(err)

	rows, err := database.Query("SELECT id, name, status FROM tasks")
	defer rows.Close()
	checkErr(err)
	var id int
	var name string
	var until string
	var status int

	for rows.Next() {
		err = rows.Scan(&id, &name, &status, &until)
		checkErr(err)
		json.NewEncoder(res).Encode(fmt.Sprintf("Id: %d, name: %s, valid until: %s, status: %s", id, name, until, statuses[status]))
	}
}

func getTask(res http.ResponseWriter, req *http.Request) {
	var params = mux.Vars(req)
	var requestId, err = strconv.Atoi(params["id"])
	var (
		id          int
		name        string
		description string
		until       string
		status      int
	)
	checkErr(err)

	row, err := database.Query("SELECT * FROM tasks WHERE id = ?", requestId)
	checkErr(err)
	defer row.Close()

	for row.Next() {
		err := row.Scan(&id, &name, &description, &until, &status)
		checkErr(err)

		json.NewEncoder(res).Encode(fmt.Sprintf("%d %s", id, name))
	}
}

func addTask(res http.ResponseWriter, req *http.Request) {
	var task Task

	json.NewDecoder(req.Body).Decode(&task)

	insertQuery := `INSERT INTO tasks(name, description, until, status) VALUES (?, ?, ?, ?)`

	statement, err := database.Prepare(insertQuery)
	checkErr(err)

	_, errExec := statement.Exec(task.Name, task.Description, task.Until, task.Status)
	checkErr(errExec)

	json.NewEncoder(res).Encode("Success")
}

func updateTask(res http.ResponseWriter, req *http.Request) {
	var params = mux.Vars(req)
	var id, err = strconv.Atoi(params["id"])
	checkErr(err)
	var task Task

	json.NewDecoder(req.Body).Decode(&task)

	updateQuery := `UPDATE tasks SET name = ?, description = ?, until = ?, status = ? WHERE id = ?`

	statement, errPrepare := database.Prepare(updateQuery)
	checkErr(errPrepare)

	_, errExec := statement.Exec(task.Name, task.Description, task.Until, task.Status, id)
	checkErr(errExec)

	json.NewEncoder(res).Encode("Success")
}

func deleteTask(res http.ResponseWriter, req *http.Request) {
	var params = mux.Vars(req)
	var id, err = strconv.Atoi(params["id"])
	checkErr(err)

	deleteQuery := `DELETE FROM tasks WHERE id = ?`

	statement, errPrepare := database.Prepare(deleteQuery)
	checkErr(errPrepare)

	_, errExec := statement.Exec(id)
	checkErr(errExec)

	json.NewEncoder(res).Encode("Success")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
