package model

import "time"

type Author struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	FullName  string     `json:"fullname"`
	Email     string     `json:"email"`
}

type AuthorDto struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

func ToAuthor(authorDto AuthorDto) *Author { // converts authordto to author
	return &Author{
		FullName: authorDto.FullName,
		Email:    authorDto.Email,
	}
}

func ToAuthorDto(author *Author) *AuthorDto { // converts author to authorDto
	return &AuthorDto{
		ID:       author.ID,
		FullName: author.FullName,
		Email:    author.Email,
	}
}
