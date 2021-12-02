package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/knudsenTaunus/sqlc-example/datasource"
	"github.com/knudsenTaunus/sqlc-example/helpers"
	_ "github.com/lib/pq"
)

func main() {
	conn, err := sql.Open("postgres", "user=postgres password=mypass dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = helpers.ValidateSchema(conn)
	if err != nil {
		log.Fatal(err)
	}

	db := datasource.New(conn)

	user, err := db.CreateUser(context.Background(), datasource.CreateUserParams{
		Firstname: "Jan-Philipp", Lastname: "Heinrich"})
	if err != nil {
		log.Printf("failed to create user: %s", err.Error())
		return
	}

	log.Println(user)

	_, err = db.CreateTodo(context.Background(), datasource.CreateTodoParams{Userid: user.ID, Task: "duschen", Done: false})
	if err != nil {
		log.Printf("failed to create todo: %s", err.Error())
		return
	}
	todos, err := db.FindTodoByUser(context.Background(), user.ID)
	if err != nil {
		log.Printf("failed to find todos", err.Error())
	}

	for _, todo := range todos {
		fmt.Println(todo)
	}

}
