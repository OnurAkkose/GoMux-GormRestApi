package repository

import (
	model "gorestapi/pkg/model/book"

	"github.com/jinzhu/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBook(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (b *BookRepository) All() ([]model.Book, error) {
	books := []model.Book{}
	err := b.db.Find(&books).Error
	return books, err
}

func (b *BookRepository) FindById(id uint) (*model.Book, error) {
	book := new(model.Book)
	err := b.db.Where(`id = ?`, id).First(&book).Error
	return book, err
}

func (b *BookRepository) Save(book *model.Book) (*model.Book, error) {
	err := b.db.Save(&book).Error
	return book, err
}

func (b *BookRepository) Delete(id uint) error {
	err := b.db.Delete(&model.Book{ID: id}).Error
	return err
}

func (b *BookRepository) Migrate() error {
	return b.db.AutoMigrate(&model.Book{}).Error
}
