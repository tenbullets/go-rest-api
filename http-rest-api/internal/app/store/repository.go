package store

import "github.com/tenbullets/go-rest-api.git/internal/app/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
