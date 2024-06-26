package tables

import (
	"database/sql"
	"new/genproto/generator"
)

type BookRepo struct {
	Db *sql.DB
}

func ConnectBook(db *sql.DB) *BookRepo {
	return &BookRepo{Db: db}
}

func (u *BookRepo) CreateBook(book *generator.AddBookRequest, id *generator.AddBookResponse) error {
	row := u.Db.QueryRow("insert into books(name, author, year) values($1,$2,$3) returning id", book.Title, book.Author, book.YearPublished)
	err := row.Scan(&id.BookId)
	return err
}

func (u *BookRepo) AddBook(book *generator.BorrowBookRequest) error {
	_, err := u.Db.Exec(`update books set user_id=$1 where id=$2`, book.UserId, book.BookId)
	return err
}

func (u *BookRepo) SearchBook(book *generator.SearchBookRequest) (*generator.SearchBookResponse, error) {
	rows, err := u.Db.Query(book.Query)
	if err != nil {
		return nil, err
	}
	books := generator.SearchBookResponse{}
	for rows.Next() {
		b := generator.Book{}
		err = rows.Scan(&b.BookId, &b.Title, &b.Author, &b.YearPublished)
		if err != nil {
			return nil, err
		}
		books.Books = append(books.Books, &b)
	}
	return &books, nil
}
