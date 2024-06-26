package usecase

import "clean-arch-with-go/api/domain"

type UserRepository interface {
	FindById(id int) (domain.User, error)
	FindAll() (domain.Users, error)
	Store(domain.User) (domain.User, error)
	Update(domain.User) (domain.User, error)
	DeleteById(domain.User) error
}
