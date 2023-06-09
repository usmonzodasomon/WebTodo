package repository

type Authorization interface {
}

type Todo interface {
}

type Repository struct {
	Authorization
	Todo
}

func NewRepository() *Repository {
	return &Repository{}
}
