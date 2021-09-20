package book

import "gorm.io/gorm"

type Repository interface {
	Index() ([]Book, error)
	Store(book Book) (Book, error)
	Show(ID int) (Book, error)
	Update(book Book) (Book, error)
	Destroy(book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Index() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *repository) Store(book Book) (Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *repository) Show(ID int) (Book, error) {
	var book Book
	err := r.db.Find(&book, ID).Error
	return book, err
}

func (r *repository) Update(book Book) (Book, error) {
	err := r.db.Save(&book).Error
	return book, err
}

func (r *repository) Destroy(book Book) (Book, error) {
	err := r.db.Delete(&book).Error
	return book, err
}
