package infrastructure

import (
	"github.com/jaijiv/handybid/config"
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

type MgoSession struct {
	Session *mgo.Session
}

var mongoSession *mgo.Session

func ConnectMongo() error {
	var err error
	/*
	 * Create the connection
	 */
	log.Println("Connecting to mongo database...")

	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{config.Conf.Db.Host},
		Timeout:  60 * time.Second,
		Database: config.Conf.Db.Database,
		Username: config.Conf.Db.User,
		Password: config.Conf.Db.Password,
	}

	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	mongoSession, err = mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}

	mongoSession.SetMode(mgo.Monotonic, true)

	return nil
}

func MongoSession() *MgoSession {
	return &MgoSession{mongoSession.Copy()}
}

func (ms *MgoSession) UserCol() *mgo.Collection {
	return ms.Session.DB(config.Conf.Db.Database).C("user")
}
