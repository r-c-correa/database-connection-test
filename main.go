package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func main() {
	host := os.Getenv("HOST")
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	if host == "" {
		panic(errors.New("HOST environment variable cannot be empty"))
	}

	if port < 1 || port > 65535 {
		panic(errors.New("PORT environment variable must be between 1 and 65535"))
	}

	if user == "" {
		panic(errors.New("USER environment variable cannot be empty"))
	}

	if password == "" {
		panic(errors.New("PASSWORD environment variable cannot be empty"))
	}

	if dbname == "" {
		panic(errors.New("DBNAME environment variable cannot be empty"))
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	var version string
	query := "SELECT version();"

	err = db.QueryRow(query).Scan(&version)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully connected: %s\n", version)
}
