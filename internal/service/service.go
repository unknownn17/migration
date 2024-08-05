package service

import (
	interface17 "conn/internal/interface"
	"conn/internal/models"
)

type Service struct {
	Psql interface17.Books
}

func NewService(repo interface17.Books) *Service {
	return &Service{Psql: repo}
}
func (u *Service) GetBook(id int) (*models.Books, error) {
	return u.Psql.GetBook(id)
}

func (u *Service) GetBooks() ([]models.Books, error) {
	return u.Psql.GetBooks()
}

func (u *Service) BookUpdate(book models.Books) (*models.Books, error) {
	return u.Psql.BookUpdate(book)
}

func (u *Service) BookCreate(book models.Books) (*models.Books, error) {
	return u.Psql.BookCreate(book)
}

func (u *Service) DeleteBook(id int) error {
	return u.Psql.DeleteBook(id)
}
