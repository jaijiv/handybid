package webcontrollers

import (
	"encoding/json"
	"github.com/jaijiv/handybid/usecases"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// Implementation in use cases package
type UserInteractor interface {
	RegisterUser(user usecases.User) error
}

type WebserviceHandler struct {
	UserInteractor UserInteractor
}

func (handler WebserviceHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		serve404(w)
		return
	}
	log.Println("RegisterUser called...")
	// Parse json request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		serveError(w, err)
		return
	}
	log.Println(string(body))
	var user usecases.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		serveError(w, err)
		return
	}
	log.Println(user)
	handler.UserInteractor.RegisterUser(user)
}

func serve404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, "404 error")
}

func serveError(w http.ResponseWriter, err error) {
	log.Println(err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
}
