package book

import "time"

// untuk time ini akan otomatis diisi oleh Gorm
// table yang ada di database nanti akan bernama Books
// lalu kolomnya akan diisi sesuai dengan isi struct ini
// kita akan menggunakan Auto Migration, dimana akan otomatis dimapping ke kolom yang ada di database
type Book struct {
	ID          int
	Title       string
	Description string
	Price       int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
