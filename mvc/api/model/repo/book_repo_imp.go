package repo

import (
	"database/sql"

	"github.com/Podre-Henrique/arquitetura-api/mvc/api/model"
)

func NewMysqlBookRepo(db *sql.DB) BookRepo {
	return &MysqlBookRepo{DB: db}
}

type MysqlBookRepo struct {
	DB *sql.DB
}

func (b *MysqlBookRepo) CreateBook(bs *model.Book) {
	query := "INSERT INTO book(isbn,name,description,pages,quantity,category) values(?,?,?,?,?,?)"
	b.DB.Exec(query, bs.ISBN, bs.Name, bs.Description, bs.Pages, bs.Quantity, bs.Category)
}

func (b *MysqlBookRepo) AddBooks(isbn uint64, quantity uint16) {
	query := "UPDATE book SET quantity=quantity+? WHERE isbn=?"
	b.DB.Exec(query, quantity, isbn)
}

func (b *MysqlBookRepo) ReadDisponibleBooks() *[]model.Book {
	query := "SELECT isbn,name,description,pages,category FROM book WHERE disponible=TRUE"
	rows, err := b.DB.Query(query)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		rows.Scan(&book.ISBN, &book.Name, &book.Description, &book.Pages, &book.Category)
		books = append(books, book)
	}
	return &books
}

func (b *MysqlBookRepo) ReadLoanUserBooks(userId uint64) *[]model.Book {

	query := `SELECT b.isbn,b.name,b.description,b.pages,b.category FROM book b
			  INNER JOIN loan l on b.isbn=l.bookISBN
			  INNER JOIN users u on b.userId=u.id
			  WHERE u.id=? AND l.WasReturned = FALSE
			 `
	rows, err := b.DB.Query(query, userId)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var books []model.Book
	for rows.Next() {
		var book model.Book
		rows.Scan(&book.ISBN, &book.Name, &book.Description, &book.Pages, &book.Category)
		books = append(books, book)
	}
	return &books
}
