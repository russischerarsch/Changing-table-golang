package main

import (
	databaseconnection "DockerSTudy/DataBase_Connection"
	"fmt"

	"context"

	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()
	conn, err := databaseconnection.CheckConnection(ctx)
	if err != nil {
		fmt.Println("Не удалось подключиться к базе данных")
	}
	mux := mux.NewRouter()
	mux.HandleFunc("/employees", NewEmployeeHandler).Methods("POST")
	mux.HandleFunc("/employees", ListEmployeeHandler).Methods("GET")
	mux.HandleFunc("/employees", DeleteEmployeeHandler).Methods("DELETE")
}
