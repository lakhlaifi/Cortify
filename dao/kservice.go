package dao

import (
	"cortify/db"
	"cortify/models"
	"errors"

	"gopkg.in/mgo.v2/bson"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"
)

//  struct
type Service struct{}

// COLLECTION of the database
const SERVICE_COLLECTION = "services"

//Kubeconfig file test
var kubeconfig *string

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

// Construct go Service
func ConstructService(name string, namespace string, ksvc models.Service) (*servingv1.Service, error) {
	if name == "" || namespace == "" {
		return nil, errors.New("internal: no name or namespace provided when constructing a service")
	}
	service := servingv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ksvc.KService.Metadata.Name,
			Namespace: namespace,
		},
	}
	service.Spec.Template = servingv1.RevisionTemplateSpec{
		Spec: servingv1.RevisionSpec{},
		ObjectMeta: metav1.ObjectMeta{
			Annotations: ksvc.KService.Metadata.Annotations,
		},
	}
	service.Spec.Template.Spec.Containers = ksvc.KService.Specs.Containers
	return &service, nil
}
