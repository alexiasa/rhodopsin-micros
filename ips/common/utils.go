package common

// todo: put these vars into heroku config vars: mongohost, mongouser, mongopass, dbname

import (
	"gopkg.in/mgo.v2"
	"log"
	"os"
	"time"
)
// database connection functions from microservices demo - altered to use Getenv since we are going to use heroku config vars

// Session holds the mongodb session for database access
var session *mgo.Session

// Get database session
func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.DialWithInfo(&mgo.DialInfo{
			Addrs:    []string{os.Getenv("mongohost")},
			Username: os.Getenv("mongouser"),
			Password: os.Getenv("mongopass"),
			Timeout:  60 * time.Second,
		})
		if err != nil {
			log.Fatalf("[GetSession]: %s\n", err)
		}
	}
	return session
}

// Create database session
func createDbSession() {
	var err error
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{os.Getenv("mongohost")},
		Username: os.Getenv("mongouser"),
		Password: os.Getenv("mongopass"),
		Timeout:  60 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	}
}
