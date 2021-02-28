package dao

import (
	"cortify/db"
	"cortify/models"

	"gopkg.in/mgo.v2/bson"
)

//  struct
type Service struct{}

// COLLECTION of the database
const SERVICE_COLLECTION = "services"

// GetAll Function retreive all objects
func (s *Service) GetAll() ([]models.Service, error) {
	sessionCopy := db.Database.Session.Copy()
	defer sessionCopy.Close()

	// Retrieve collection
	collection := sessionCopy.DB(db.Database.DatabaseName).C(SERVICE_COLLECTION)

	var service []models.Service
	err := collection.Find(bson.M{}).All(&service)
	return service, err
}

// GetByID Function retrive One Object
func (s *Service) GetByID(id string) (models.Service, error) {
	var service models.Service
	sessionCopy := db.Database.Session.Copy()
	defer sessionCopy.Close()

	// Retrieve collection
	collection := sessionCopy.DB(db.Database.DatabaseName).C(SERVICE_COLLECTION)

	err := collection.FindId(id).One(&service)
	return service, err
}

// Insert Service
func (s *Service) Insert(service models.Service) error {
	sessionCopy := db.Database.Session.Copy()
	defer sessionCopy.Close()
	// Retrieve collection
	collection := sessionCopy.DB(db.Database.DatabaseName).C(SERVICE_COLLECTION)
	err := collection.Insert(&service)
	return err
}
