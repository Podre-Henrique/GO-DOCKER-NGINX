package repo

import (
	"database/sql"

	"github.com/Podre-Henrique/arquitetura-api/mvc/api/model"
)

type MysqlLoanRepo struct {
	DB *sql.DB
}

func (r *MysqlLoanRepo) CreateLoan(ls *model.Loan) {
	query := "INSERT INTO loan(userId,bookISBN) values(?,?)"
	r.DB.Exec(query, ls.UserId, ls.BookISBN)
}

func (r *MysqlLoanRepo) ReturnLoan(userId uint64, isbn uint64) {
	query := "UPDATE loan SET wasReturned=TRUE where isbn=? and userId=? and wasReturned=FALSE"
	r.DB.Exec(query, userId, isbn)
}

func (b *MysqlLoanRepo) ReadLoanDependencies() *[]model.Book {

	query := `SELECT b.isbn,b.name,b.category,u.email,u.name FROM book b
			  INNER JOIN loan l on b.isbn=l.bookISBN
			  INNER JOIN users u on b.userId=u.id
			  WHERE u.id=? AND l.WasReturned = FALSE
			`
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
