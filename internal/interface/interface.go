package interface17

import "conn/internal/models"

type Books interface {
	GetBook(id int) (*models.Books, error)
	GetBooks() ([]models.Books, error)
	BookUpdate(book models.Books) (*models.Books, error)
	BookCreate(book models.Books) (*models.Books, error)
	DeleteBook(id int) error
}
