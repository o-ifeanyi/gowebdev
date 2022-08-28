package repos

import (
	"database/sql"
	"errors"
	"gowebdev/17-hands-on/models"
	"net/http"
	"strconv"
)

type BookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) BookRepo {
	return BookRepo{db}

}

func (br BookRepo) AllBooks() ([]models.Book, error) {
	rows, err := br.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bks := []models.Book{}
	for rows.Next() {
		bk := models.Book{}
		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price) // order matters
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return bks, nil
}

func (br BookRepo) OneBook(r *http.Request) (models.Book, error) {
	bk := models.Book{}
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return bk, errors.New("400. Bad Request")
	}

	row := br.db.QueryRow("SELECT * FROM books WHERE isbn = $1", isbn)

	err := row.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
	if err != nil {
		return bk, err
	}

	return bk, nil
}

func (br BookRepo) PutBook(r *http.Request) (models.Book, error) {
	bk := models.Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad request. All fields must be complete")
	}

	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Price must be a number")
	}
	bk.Price = float32(f64)

	_, err = br.db.Exec("INSERT INTO books (isbn, title, author, price) VALUES ($1, $2, $3, $4)", bk.Isbn, bk.Title, bk.Author, bk.Price)
	if err != nil {
		return bk, errors.New("500. Internal Server Error." + err.Error())
	}
	return bk, nil
}

func (br BookRepo) UpdateBook(r *http.Request) (models.Book, error) {
	bk := models.Book{}
	bk.Isbn = r.FormValue("isbn")
	bk.Title = r.FormValue("title")
	bk.Author = r.FormValue("author")
	p := r.FormValue("price")

	if bk.Isbn == "" || bk.Title == "" || bk.Author == "" || p == "" {
		return bk, errors.New("400. Bad Request. Fields can't be empty")
	}

	f64, err := strconv.ParseFloat(p, 32)
	if err != nil {
		return bk, errors.New("406. Not Acceptable. Enter number for price")
	}
	bk.Price = float32(f64)

	_, err = br.db.Exec("UPDATE books SET isbn = $1, title=$2, author=$3, price=$4 WHERE isbn=$1;", bk.Isbn, bk.Title, bk.Author, bk.Price)
	if err != nil {
		return bk, err
	}
	return bk, nil
}

func (br BookRepo) DeleteBook(r *http.Request) error {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("400. Bad Request")
	}

	_, err := br.db.Exec("DELETE FROM books WHERE isbn=$1;", isbn)
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}
