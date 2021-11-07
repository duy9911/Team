package models

type Team struct {
	ID    string
	Name  string
	Staff []StaffRecieve
}

type StaffRecieve struct {
	ID     string
	Name   string
	Gender string
	Dob    string
	Salary float64
}
