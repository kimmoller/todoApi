package model

type Todo struct {
	ID         int
	Task       string
	Status     string
	IdentityId int
}

type Identity struct {
	ID       int
	Username string
	Password string
}
