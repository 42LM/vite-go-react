package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

type Data struct {
	Token string
}

func init() {
	var err error

	// Open connection to database
	Db, err = sql.Open("sqlite3", "./vite-go-react.db")
	if err != nil {
		log.Fatal(err)
	}
	// Ping database
	if err = Db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You connected to your database.")
}

func getPort() int {
	if port := os.Getenv("PORT"); port == "" {
		return 8080
	} else {
		portNum, _ := strconv.Atoi(port)
		return portNum
	}
}

func main() {
	port := getPort()
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./web"))
	mux.Handle("/", fs)

	st := http.StripPrefix("/static/", http.FileServer(http.Dir("./web/static")))
	mux.Handle("/static/", st)

	mux.HandleFunc("/api/timestamp", TestReactHandler)

	fmt.Println("Port:", port)

	http.ListenAndServe(":"+fmt.Sprint(port), mux)
}

func TestReactHandler(w http.ResponseWriter, r *http.Request) {
	d := Data{}

	if err := d.Show(); err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(500)
		w.Write([]byte(http.StatusText(500)))
		return
	}

	w.Write([]byte(d.Token))
}

func CreateTimeStamp() {
	_, err := Db.Exec("INSERT INTO data (token) VALUES ($1)", time.Now())
	if err != nil {
		log.Println("Could not insert timestamp")
		log.Println(err)
	}
}

func (d *Data) Show() (err error) {
	CreateTimeStamp()

	statement := `SELECT token FROM data ORDER BY token DESC LIMIT 1`

	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}

	err = stmt.QueryRow().Scan(&d.Token)
	return
}
