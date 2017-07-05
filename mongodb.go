package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	mgo "gopkg.in/mgo.v2"
)

//MongoConfig stores the configuration of mongodb to connect
type MongoConfig struct {
	Ip         string `json:"ip"`
	Database   string `json:"database"`
	Collection string `json:"collection"`
}

var mongoConfig MongoConfig

func readMongodbConfig(path string) {
	file, e := ioutil.ReadFile(path)
	if e != nil {
		fmt.Println("error:", e)
	}
	content := string(file)
	json.Unmarshal([]byte(content), &mongoConfig)
}

func getSession() (*mgo.Session, error) {
	session, err := mgo.Dial("mongodb://" + mongoConfig.Ip)
	if err != nil {
		panic(err)
	}
	//defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session, err
}
func getCollection(session *mgo.Session) *mgo.Collection {

	c := session.DB(mongoConfig.Database).C(mongoConfig.Collection)
	return c
}
func connectMongodb() *mgo.Collection {
	session, err := getSession()
	if err != nil {
		log.Fatal(err)
	}
	c := getCollection(session)
	return c
}

func saveDataEntryToMongo(c *mgo.Collection, user UserModel) {
	/*
		how to call this function
			var auxArr []string
			user := UserModel{"123", auxArr}
			saveDataEntryToMongo(c, user)
	*/

	err := c.Insert(user)
	if err != nil {
		log.Fatal(err)
	}

}
