package database

import (
	sqlbuilder "conn/internal/database/sql"
	interface17 "conn/internal/interface"
	"conn/internal/models"
	"database/sql"
)

type Postgres struct {
	Db *sql.DB
}

func NewPostgres(repo *sql.DB) interface17.Books {
	return &Postgres{Db: repo}
}

func (u *Postgres) GetBook(id int) (*models.Books, error) {
	query, args, err := sqlbuilder.GetBook(id)
	if err != nil {
		return nil, err
	}
	var res models.Books
	if err := u.Db.QueryRow(query, args...).Scan(&res.ID, &res.Title, &res.Author, &res.Published_at); err != nil {
		return nil, err
	}
	return &res, nil
}

func (u *Postgres) GetBooks() ([]models.Books, error) {
	query, args, err := sqlbuilder.GetBooks()
	if err != nil {
		return nil, err
	}
	var slc []models.Books

	rows, err := u.Db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var all models.Books

		if err := rows.Scan(&all.ID, &all.Title, &all.Author, &all.Published_at); err != nil {
			return nil, err
		}

		slc = append(slc, all)
	}
	return slc, nil
}

func (u *Postgres) BookCreate(book models.Books) (*models.Books, error) {
	query, args, err := sqlbuilder.CreateBook(&book)
	if err != nil {
		return nil, err
	}
	var res models.Books
	if err := u.Db.QueryRow(query, args...).Scan(&res.ID, &res.Title, &res.Author, &res.Published_at); err != nil {
		return nil, err
	}
	return &res, nil
}

func (u *Postgres) BookUpdate(book models.Books) (*models.Books, error) {
	query, args, err := sqlbuilder.UpdateBook(&book)
	if err != nil {
		return nil, err
	}
	var res models.Books
	if err := u.Db.QueryRow(query, args...).Scan(&res.ID, &res.Title, &res.Author, &res.Published_at); err != nil {
		return nil, err
	}
	return &res, nil
}

func (u *Postgres) DeleteBook(id int) error {
	query, args, err := sqlbuilder.DeleteBook(id)
	if err != nil {
		return err
	}
	_, err = u.Db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}
