package mongo

import (
	"github.com/jaijiv/handybid/domain"
	"github.com/jaijiv/handybid/infrastructure"
	"log"
)

type MongoRepo struct{}

func (MongoRepo) StoreUser(user domain.User) {
	ms := infrastructure.MongoSession()
	defer ms.Session.Close()

	ms.UserCol().Insert(&user)
	log.Println("Inserted user into db")
}

func (MongoRepo) ListUsers() ([]domain.User, error) {
	ms := infrastructure.MongoSession()
	defer ms.Session.Close()

	var dusers []domain.User
	err := ms.UserCol().Find(nil).All(&dusers)
	if err != nil {
		return nil, err
	}
	log.Println("Listed all users from db")
	return dusers, nil
}

func (MongoRepo) FindUserById(id string) domain.User {
	return domain.User{}
}
