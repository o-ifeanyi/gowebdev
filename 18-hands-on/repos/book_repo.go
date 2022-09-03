package repos

import (
	"context"
	"errors"
	"fmt"
	"gowebdev/18-hands-on/models"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type BookRepo struct {
	db *mongo.Database
}

func NewBookRepo(db *mongo.Database) BookRepo {
	return BookRepo{db}
}

func (br BookRepo) AllBooks() ([]models.Book, error) {
	bks := []models.Book{}
	cur, err := br.db.Collection("books").Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	err = cur.All(context.Background(), &bks)
	if err != nil {
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
	res := br.db.Collection("books").FindOne(context.Background(), bson.M{"isbn": isbn})

	err := res.Decode(&bk)
	if err != nil {
		fmt.Println(err)
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

	_, err = br.db.Collection("books").InsertOne(context.Background(), bk)
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
		fmt.Println(err)
		return bk, errors.New("406. Not Acceptable. Enter number for price")
	}
	bk.Price = float32(f64)

	_, err = br.db.Collection("books").UpdateOne(context.Background(), bson.M{"isbn": bk.Isbn}, bson.M{"$set": &bk})
	if err != nil {
		fmt.Println(err)
		return bk, err
	}
	return bk, nil
}

func (br BookRepo) DeleteBook(r *http.Request) error {
	isbn := r.FormValue("isbn")
	if isbn == "" {
		return errors.New("400. Bad Request")
	}

	_, err := br.db.Collection("books").DeleteOne(context.Background(), bson.M{"isbn": isbn})
	if err != nil {
		return errors.New("500. Internal Server Error")
	}
	return nil
}
