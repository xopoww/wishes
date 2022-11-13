package models

type User struct {
	ID   int64
	Name string

	Fname string
	Lname string

	PassHash []byte
}
