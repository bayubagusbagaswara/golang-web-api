package book

import "gorm.io/gorm"

// kita buar Repository dengan tipe interface
// dan didalamnya kita buat function-function yang berkaitan dengan data book
// misalnya mencari buku, mengambil data buku, menambah buku dll
type Repository interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book Book) (Book, error)
}

// kita bikin implementasi dari interface Repository
type repository struct {
	// untuk berkomunikasi database, kita butuh object db dari gorm nya
	db *gorm.DB
}

// kita buat function untuk instance dari struct diatas
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// kita buat function untuk method dari repository
func (r *repository) FindAll() ([]Book, error) {
	var books []Book

	err := r.db.Find(&books).Error

	return books, err
}

func (r *repository) FindById(ID int) (Book, error) {
	var book Book
	err := r.db.Find(&book, ID).Error
	return book, err
}

func (r *repository) Create(book Book) (Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}
