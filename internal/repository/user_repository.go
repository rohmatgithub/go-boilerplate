package repository

import "boilerplate/internal/domain"

type UserRepository interface {
	FindByID(id string) *domain.User
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindByID(id string) *domain.User {
	// Implementasi dummy
	return &domain.User{ID: id, Name: "John Doe"}
}
