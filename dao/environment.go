package dao

import (
	"context"
	"cortify/models"
	"encoding/json"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Environment struct{}

type environmentPolicy struct {
	Ingress struct {
		Isolation string `json:"isolation"`
	} `json:"ingress"`
}

//Construct Environment
func ConstructEnvironment(env models.Environment) (*corev1.Namespace, error) {
	environment := corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:   env.Name,
			Labels: env.Namespace.Labels,
		},
	}

	//Map to Network Policy
	if env.Isolation != "" {
		ep := environmentPolicy{}
		ep.Ingress.Isolation = env.Isolation
		annotation, _ := json.Marshal(ep)
		environment.ObjectMeta.Annotations = map[string]string{
			"networking.k8s.io/v1": string(annotation),
		}
	}
	return &environment, nil
}

func CreateEnvironment(environment *corev1.Namespace) error {
	config, err := NewKubernetesClient()
	if err != nil {
		return fmt.Errorf("failed to get kubernetes config: %v", err)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create kubernetes client: %v", err)
	}
	_, err = client.CoreV1().Namespaces().Create(context.TODO(), environment, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create Environment: %v", err)
	}
	return nil
}
