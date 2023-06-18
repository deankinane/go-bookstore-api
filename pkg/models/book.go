package models

import (
	"github.com/deankinane/go-bookstore-api/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name      string `gorm:"" json:"name"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
}

func init() {
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) *Book {
	var getBook Book
	db.Where("ID = ?", id).Find(&getBook)
	return &getBook
}

func DeleteBook(id int64) *Book {
	var bookToDelete Book
	result := db.First(&bookToDelete, id)
	if result.Error != nil {
		return nil
	}
	db.Delete(&Book{}, id)
	return &bookToDelete
}

func UpdateBook(id int64, book *Book) *Book {
	var bookToUpdate Book
	result := db.First(&bookToUpdate, id)
	if result.Error != nil {
		return nil
	}

	db.First(&bookToUpdate, id).Updates(*book)
	return &bookToUpdate
}
