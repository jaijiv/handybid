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
	log.Print("Inserted user into mongo")
}

func (MongoRepo) FindUserById(id string) domain.User {
	return domain.User{}
}
