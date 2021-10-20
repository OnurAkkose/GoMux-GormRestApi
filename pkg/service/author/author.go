package service

import (
	model "gorestapi/pkg/model/author"
	author "gorestapi/pkg/repository/author"
)

type AuthorService struct {
	AuthorRepository *author.AuthorRepository
}

func NewAuthorService(a *author.AuthorRepository) AuthorService {
	return AuthorService{AuthorRepository: a}
}

func (a *AuthorService) All() ([]model.Author, error) {
	return a.AuthorRepository.All()
}

func (a *AuthorService) FindById(id uint) (*model.Author, error) {
	return a.AuthorRepository.FindById(id)
}

func (a *AuthorService) Save(author *model.Author) (*model.Author, error) {
	return a.AuthorRepository.Save(author)
}

func (a *AuthorService) Delete(id uint) error {
	return a.AuthorRepository.Delete(id)
}

func (a *AuthorService) Migrate() error {
	return a.AuthorRepository.Migrate()
}
