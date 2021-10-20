package repository

import (
	model "gorestapi/pkg/model/author"

	"github.com/jinzhu/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthor(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{
		db: db,
	}
}

func (a *AuthorRepository) All() ([]model.Author, error) {
	authors := []model.Author{}
	err := a.db.Find(&authors).Error
	return authors, err
}

func (a *AuthorRepository) FindById(id uint) (*model.Author, error) {
	author := new(model.Author)
	err := a.db.Where(`id = ?`, id).First(&author).Error
	return author, err
}

func (a *AuthorRepository) Save(author *model.Author) (*model.Author, error) {
	err := a.db.Save(&author).Error
	return author, err
}

func (a *AuthorRepository) Delete(id uint) error {
	err := a.db.Delete(&model.Author{ID: id}).Error
	return err
}

func (a *AuthorRepository) Migrate() error {
	return a.db.AutoMigrate(&model.Author{}).Error
}
