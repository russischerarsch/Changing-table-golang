package httpserver

type Employee struct {
	ID       int
	FullName string `json:"name"`
	Position string `json:"position"`
}

var Employees []Employee

func (e *Employee) AddID() {
	e.ID++
}
