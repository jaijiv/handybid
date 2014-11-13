package main

import (
	"flag"
	"fmt"
	"github.com/adampresley/sigint"
	"github.com/gorilla/mux"
	"github.com/jaijiv/handybid/infrastructure"
	"github.com/jaijiv/handybid/interfaces/repositories/mongo"
	"github.com/jaijiv/handybid/interfaces/webcontrollers"
	"github.com/jaijiv/handybid/interfaces/webcontrollers/middleware"
	"github.com/jaijiv/handybid/usecases"
	"github.com/justinas/alice"
	"log"
	"net/http"
	"os"
	"runtime"
)

func init() {
	var err error
	err = infrastructure.ConnectMongo()
	if err != nil {
		log.Fatal(err)
	}
}

var (
	port              = flag.Int("port", 8080, "Port for web server")
	host              = flag.String("host", "localhost", "Address for web server")
	webserviceHandler = new(webcontrollers.WebserviceHandler)
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	/*
	 * Handle SIGINT (CTRL+C)
	 */
	sigint.ListenForSIGINT(func() {
		log.Println("Shutting down handybid server...")
		os.Exit(0)
	})

	// Instantiate UserInteractor from usecases
	userInteractor := new(usecases.UserInteractor)
	// Inject mongo repository that implements domain methods into userInteractor
	userInteractor.UserRepository = mongo.MongoRepo{}
	// Inject userInteractor from usecases into web service handler
	webserviceHandler.UserInteractor = userInteractor

	router := mux.NewRouter()
	router.HandleFunc("/users", RegisterUser).Methods("POST")
	router.HandleFunc("/users", ListUsers).Methods("GET")
	middleware := alice.New(middleware.Logger).Then(router)

	log.Printf("handybid services started on %s:%d\n\n", *host, *port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *port), middleware)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	webserviceHandler.RegisterUser(w, r)
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	webserviceHandler.ListUsers(w, r)
}
