package usecases

import (
	"github.com/jaijiv/handybid/domain"
	"time"
)

// These data types are exposed to interfaces layer.
// These are data types that satisfies usecases for this application.
type User struct {
	Id         string `json:"id"`
	DispName   string `json:"disp_name"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	EmailId    string `json:"email_id"`
	Password   string `json:"password"`
	Address    string `json:"address"`
	IsProvider bool   `json:"is_provider"`
}

type UserInteractor struct {
	UserRepository domain.UserRepository
}

func (interactor *UserInteractor) RegisterUser(user User) error {
	// map user to dUser
	interactor.UserRepository.StoreUser(mapToDomainUser(user))
	return nil
}

func mapToDomainUser(user User) (dUser domain.User) {
	dUser.DispName = user.DispName
	dUser.Name = user.Name
	dUser.Phone = user.Phone
	dUser.EmailId = user.EmailId
	dUser.Password = user.Password
	dUser.Address = user.Address
	dUser.IsProvider = user.IsProvider
	dUser.CreateTs = time.Now().UTC()
	return
}
