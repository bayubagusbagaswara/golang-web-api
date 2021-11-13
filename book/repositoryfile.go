package book

import (
	"errors"
	"fmt"
)

// disini kita bikin repository yang gunanya untuk menyimpan data kedalam sebuah file
// jadi selain kita menyimpan di database, kita juga akan menyimpan data kedalam sebuah file

// kita bikin sebuah struct dulu
type fileRepository struct {
}

func NewFileRepository() *fileRepository {
	return &fileRepository{}
}

func (r *fileRepository) FindAll() ([]Book, error) {
	var books []Book

	fmt.Println("FindAll")
	return books, errors.New("dummy error")
}

func (r *fileRepository) FindById(ID int) (Book, error) {
	var book Book

	fmt.Println("FindById")
	return book, nil
}

func (r *fileRepository) Create(book Book) (Book, error) {
	fmt.Println("Create")
	return book, nil
}
