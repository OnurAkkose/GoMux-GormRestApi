package service

import (
	model "gorestapi/pkg/model/book"
	book "gorestapi/pkg/repository/book"
)

type BookService struct {
	BookRepository *book.BookRepository
}

func NewBookService(b *book.BookRepository) BookService {
	return BookService{BookRepository: b}
}

func (b *BookService) All() ([]model.Book, error) {
	return b.BookRepository.All()
}

func (b *BookService) FindById(id uint) (*model.Book, error) {
	return b.BookRepository.FindById(id)
}

func (b *BookService) Save(book *model.Book) (*model.Book, error) {
	return b.BookRepository.Save(book)
}

func (b *BookService) Delete(id uint) error {
	return b.BookRepository.Delete(id)
}

func (b *BookService) Migrate() error {
	return b.BookRepository.Migrate()
}
