package repositories

import (
	"github.com/jaijiv/handybid/domain"
)

type MongoRepo struct{}

func (MongoRepo) StoreUser(user domain.User) {
}

func (MongoRepo) FindUserById(id string) domain.User {
	return domain.User{}
}
