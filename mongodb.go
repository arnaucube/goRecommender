package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//MongoConfig stores the configuration of mongodb to connect
type MongoConfig struct {
	Ip       string `json:"ip"`
	Database string `json:"database"`
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
func getCollection(session *mgo.Session, collection string) *mgo.Collection {

	c := session.DB(mongoConfig.Database).C(collection)
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

func saveItem(c *mgo.Collection, item ItemModel) {
	//first, check if the item already exists
	result := ItemModel{}
	err := c.Find(bson.M{"id": item.ID}).One(&result)
	if err != nil {
		//item not found, so let's add a new entry
		err = c.Insert(item)
		check(err)
	} else {
		/*result.Data = append(result.Data, dataEntry)
		err = c.Update(bson.M{"id": dataEntry.ContratoCOD}, result)
		if err != nil {
			log.Fatal(err)
		}*/
	}

}

func saveUser(c *mgo.Collection, user UserModel) {
	//first, check if the item already exists
	result := UserModel{}
	err := c.Find(bson.M{"id": user.ID}).One(&result)
	if err != nil {
		//item not found, so let's add a new entry
		err = c.Insert(user)
		check(err)
	} else {
		/*result.Data = append(result.Data, dataEntry)
		err = c.Update(bson.M{"id": dataEntry.ContratoCOD}, result)
		if err != nil {
			log.Fatal(err)
		}*/
	}

}

func getUserById(id string) (UserModel, error) {
	result := UserModel{}
	err := userCollection.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		//user not exist
		return result, err
	} else {
		//user exist
		return result, err
	}
}
func getItemById(id string) (ItemModel, error) {
	result := ItemModel{}
	err := itemCollection.Find(bson.M{"id": id}).One(&result)
	if err != nil {
		//user not exist
		return result, err
	} else {
		//user exist
		return result, err
	}
}
func updateItem(item ItemModel) (ItemModel, error) {
	err := itemCollection.Update(bson.M{"id": item.ID}, item)
	if err != nil {
		//log.Fatal(err)
		return item, err
	}
	return item, err
}
func updateUser(user UserModel) (UserModel, error) {
	err := userCollection.Update(bson.M{"id": user.ID}, user)
	if err != nil {
		return user, err
	}
	return user, err
}

func getAllItems() ([]ItemModel, error) {
	result := []ItemModel{}
	iter := itemCollection.Find(nil).Limit(100).Iter()
	err := iter.All(&result)
	return result, err
}

func getItemsNotActed(actedItems []string) ([]ItemModel, error) {
	result := []ItemModel{}
	iter := itemCollection.Find(bson.M{"id": bson.M{"$nin": actedItems}}).Limit(100).Iter()
	err := iter.All(&result)
	return result, err
}
