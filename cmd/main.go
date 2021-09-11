package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	/*	log := log.New()

		cfg, err := newConfig()
		if err != nil {
			log.Fatal("cant get config: ", err)
		}
	*/

	type user struct {
		Id        int
		FirstName string
		LastName  string
	}

	storage, err := pgxpool.Connect(context.Background(), "user=postgres password=qwe123 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connection to database: %v\n", err)
	}
	defer storage.Close()

	rows, err := storage.Query(context.Background(), "select id,first_name,last_name from employee_accounting.staff s ")
	if err != nil {
		fmt.Println("11 ", err)
	}

	defer rows.Close()
	var users []user
	var user1 user
	for rows.Next() {

		err = rows.Scan(&user1.Id, &user1.FirstName, &user1.LastName)
		users = append(users, user1)
		if err != nil {
			fmt.Println("222 ", err)
		}
	}
	for _, val := range users {
		fmt.Println(val)
	}
	//	server := APIServer.New(storage, log, APIServer.NewConfig(cfg.servPort))
	//	server.Start()

	fmt.Println("happy end!")
}
