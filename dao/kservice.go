package dao

import (
	"cortify/models"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	servingv1 "knative.dev/serving/pkg/apis/serving/v1"
)

type Service struct{}

// Construct Knative Service
func ConstructService(svc models.Service) (*servingv1.Service, error) {
	service := servingv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      svc.Name,
			Namespace: svc.Namespace,
			// Used to construct Domain Name per tenant
			Annotations: svc.KService.Config.Annotations,
		},
	}
	service.Spec.Template = servingv1.RevisionTemplateSpec{
		Spec: servingv1.RevisionSpec{},
		ObjectMeta: metav1.ObjectMeta{
			//Annotations per revision
			Annotations: svc.KService.Config.Annotations,
		},
	}

	service.Spec.Template.Spec.Containers = svc.KService.Specs.Containers
	return &service, nil
}

func CreateService(service *servingv1.Service) error {
	client, err := NewServingClient(service.Namespace)
	if err != nil {
		panic(err.Error())
	}
	err = client.CreateService(service)
	if err != nil {
		return err
	}
	return err
}
