package main

import (
	"github.com/jaijiv/handybid/infrastructure"
	"github.com/jaijiv/handybid/interfaces/repositories/mongo"
	"github.com/jaijiv/handybid/interfaces/webcontrollers"
	"github.com/jaijiv/handybid/usecases"
	"log"
	"net/http"
)

func init() {
	var err error
	err = infrastructure.ConnectMongo()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Instantiate UserInteractor from usecases
	userInteractor := new(usecases.UserInteractor)
	// Inject mongo repository that implements domain methods into userInteractor
	userInteractor.UserRepository = mongo.MongoRepo{}

	webserviceHandler := new(webcontrollers.WebserviceHandler)
	// Inject userInteractor from usecases into web service handler
	webserviceHandler.UserInteractor = userInteractor

	// Register user /list users
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		webserviceHandler.HandleUser(w, r)
	})
	http.ListenAndServe(":8080", nil)
}
