package dao

import (
	"cortify/db"
	"cortify/models"
	"errors"

	"gopkg.in/mgo.v2/bson"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"
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

// Construct Service
func ConstructService(name string, namespace string, svc models.Service) (*servingv1.Service, error) {
	if name == "" || namespace == "" {
		return nil, errors.New("internal: no name or namespace provided when constructing a service")
	}
	service := servingv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      svc.KService.Metadata.Name,
			Namespace: namespace,
		},
	}

	service.Spec.Template = servingv1.RevisionTemplateSpec{
		Spec: servingv1.RevisionSpec{},
		ObjectMeta: metav1.ObjectMeta{
			Annotations: map[string]string{
				"foo": "bar",
			},
		},
	}
	service.Spec.Template.Spec.Containers = []corev1.Container{
		{
			Name:  svc.KService.Specs.Containers[0].Name,
			Image: svc.KService.Specs.Containers[0].Image,
		},
	}
	return &service, nil
}
