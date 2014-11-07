package domain

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Implementation in repositories package
type UserRepository interface {
	StoreUser(user User)
	FindUserById(id string) User
}

type User struct {
	ID           bson.ObjectId `bson:"_id,omitempty"`
	DispName     string
	Name         string
	Phone        string
	EmailId      string
	Password     string
	Address      string
	Lat          float64
	Long         float64
	IsProvider   bool
	ProviderInfo []Provider
	CreateTs     time.Time
	UpdateTs     time.Time
}

type Provider struct {
	Type        string  // "cleaner", "handyman"
	Rate        float64 // example $22.50 / hour
	DiscountPct float64 // discount percentage - example 10% - applied to Rate
	AutoReply   bool    // If on and available, then system auto replies to jobs
	Available   bool
}
