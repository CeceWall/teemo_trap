package mushroom

import (
	"gopkg.in/mgo.v2"
)

type MongoConfig struct {
	url string
}

type Database struct {
	config *MongoConfig
}

var database *Database
var config *MongoConfig

func InitDatabase(c *MongoConfig) {
	config = c
}
func GetDatabase() *Database {
	if database == nil {
		database = &Database{
			config,
		}
	}
	return database
}

func (db *Database) GetSession(config MongoConfig) (*mgo.Session, error) {
	session, err := mgo.Dial(config.url)
	if err != nil {
		return nil, err
	}
	return session, nil
}
