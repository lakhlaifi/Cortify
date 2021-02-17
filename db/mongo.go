package db

import (
	"cortify/common"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
)

// DataStruct Manage MongoDB Connection
type DataStruct struct {
	Session      *mgo.Session
	DatabaseName string
}

// TODO Manage Multitenancy in Database
// Init Initializes a mongo database
func (db *DataStruct) Init() error {
	db.DatabaseName = common.Config.DBName

	// DialInfo
	dialInfo := mgo.DialInfo{
		Addrs:    []string{common.Config.DBAddrs},
		Timeout:  60 * time.Second,
		Database: db.DatabaseName,
		Username: common.Config.DBUser,
		Password: common.Config.DBPwd,
	}
	var err error
	db.Session, err = mgo.DialWithInfo(&dialInfo)
	if err != nil {
		log.Debug("Couldn't connect to database : ", err)
	}
	return err
}

// Close the existing connection
func (db *DataStruct) Close() {
	if db.Session != nil {
		db.Session.Close()
	}
}
