package model

import (
	model "gorestapi/pkg/model/author"
	"time"
)

type Book struct {
	ID        uint         `gorm:"primary_key" json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	Title     string       `json:"title"`
	AuthorId  int          `json:"author_id"`
	Author    model.Author `gorm:"foreignkey:id"`
}

type BookDto struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Title    string `json:"title"`
	AuthorId int    `json:"author_id"`
}

func ToBook(bookDto BookDto) *Book {
	return &Book{
		ID:       bookDto.ID,
		Title:    bookDto.Title,
		AuthorId: bookDto.AuthorId,
	}
}
func ToBookDto(book *Book) *BookDto {
	return &BookDto{
		ID:       book.ID,
		Title:    book.Title,
		AuthorId: book.AuthorId,
	}
}
