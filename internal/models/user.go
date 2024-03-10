package models

type User struct {
	Id        string `db:"id"`
	Email     string `db:"email"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Picture   string `db:"picture"`
}
