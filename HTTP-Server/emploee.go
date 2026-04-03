package httpserver

type Employee struct {
	ID       int
	FullName string `json:"name"`
	Position string `json:"position"`
}

var Employees []Employee
var GlobalID int

func (e *Employee) AddID() {
	GlobalID++
	e.ID = GlobalID
}
