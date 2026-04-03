package httpserver

import (
	databaseconnection "DockerSTudy/DataBase_Connection"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

var employee Employee

func NewEmployeeHandler(w http.ResponseWriter, r *http.Request) {

	BodyByte, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Ошибка при конвертации текста"))
		w.WriteHeader(http.StatusBadRequest)
	}
	if err := json.Unmarshal(BodyByte, &employee); err != nil {
		panic(err)
	}
	employee.AddID()
	Employees = append(Employees, employee)
	err = databaseconnection.AddEmployeeDB(databaseconnection.Conn, databaseconnection.Ctx, employee.ID, employee.FullName, employee.Position, time.Now())
	if err != nil {
		panic(err)
	}
	w.Write([]byte("Добавление сотрудника прошло успешно!"))
}

func ListEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(Employees); err != nil {
		panic(err)
	}
}

func DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	EmployeeID, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Ошибка при конвертации текста"))
		w.WriteHeader(http.StatusBadRequest)
	}
	employeeID, err := strconv.Atoi(string(EmployeeID))
	if err != nil {
		panic(err)
	}
	for i, employee := range Employees {
		if employee.ID == employeeID {
			Employees = append(Employees[:i], Employees[i+1:]...)
		}
	}
	w.Write([]byte("Сотрудник был успешно удален из списка"))
	if err := databaseconnection.DeleteEmployeeDB(databaseconnection.Conn, databaseconnection.Ctx, employeeID); err != nil {
		panic(err)
	}
}
