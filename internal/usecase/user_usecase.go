package usecase

import (
	"boilerplate/internal/domain"
	"boilerplate/internal/repository"
)

type UserUsecase interface {
	GetUser(id string) *domain.User
}

// Hapus deklarasi ganda
type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) GetUser(id string) *domain.User {
	return u.repo.FindByID(id)
}
