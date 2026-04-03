package main

import (
	databaseconnection "DockerSTudy/DataBase_Connection"
	"fmt"
	"net/http"

	httpserver "DockerSTudy/HTTP-Server"

	"github.com/gorilla/mux"
)

func main() {
	var err error
	databaseconnection.Conn, err = databaseconnection.CheckConnection(databaseconnection.Ctx)
	if err != nil {
		fmt.Println("Не удалось подключиться к базе данных. Error:", err)
		return
	}
	if databaseconnection.Conn == nil {
		fmt.Println("conn is nil — подключение не установлено")
		return
	}
	fmt.Println("Подключение к базе данных произведено успешно!")
	databaseconnection.CreateTable(databaseconnection.Ctx, databaseconnection.Conn)
	mux := mux.NewRouter()
	mux.HandleFunc("/employees", httpserver.NewEmployeeHandler).Methods("POST")
	mux.HandleFunc("/employees", httpserver.ListEmployeeHandler).Methods("GET")
	mux.HandleFunc("/employees", httpserver.DeleteEmployeeHandler).Methods("DELETE")
	http.ListenAndServe(":9091", mux)
}
