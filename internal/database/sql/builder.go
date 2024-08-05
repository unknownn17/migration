package sqlbuilder

import (
	"conn/internal/models"

	"github.com/Masterminds/squirrel"
)

func GetBook(id int) (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").From("books").Where(squirrel.Eq{"id": id}).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return "", nil, err
	}
	return query, args, nil
}

func GetBooks() (string, []interface{}, error) {
	query, args, err := squirrel.Select("*").From("books").ToSql()
	if err != nil {
		return "", nil, err
	}
	return query, args, nil
}

func CreateBook(req *models.Books) (string, []interface{}, error) {
	query, args, err := squirrel.
		Insert("books").
		Columns("title", "author", "published_at").
		Values(req.Title, req.Author, req.Published_at).
		PlaceholderFormat(squirrel.Dollar).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return "", nil, err
	}
	return query, args, nil
}

func UpdateBook(req *models.Books) (string, []interface{}, error) {
	updateData := map[string]interface{}{
		"title":        req.Title,
		"author":       req.Author,
		"published_at": req.Published_at,
	}
	query, args, err := squirrel.Update("books").SetMap(updateData).PlaceholderFormat(squirrel.Dollar).Where(squirrel.Eq{"id": req.ID}).Suffix("RETURNING *").ToSql()
	if err != nil {
		return "", nil, err
	}
	return query, args, nil
}

func DeleteBook(id int) (string, []interface{}, error) {
	query, args, err := squirrel.Delete("books").Where(squirrel.Eq{"id": id}).PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return "", nil, err
	}
	return query, args, nil
}
