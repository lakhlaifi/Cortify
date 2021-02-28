package dao

import (
	"context"
	"cortify/models"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ConstructSecret(svc models.Service) (*corev1.Secret, error) {
	secret := corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{Name: svc.Name, Namespace: svc.Namespace},
	}
	return &secret, nil
}

func CreateSecret(secret *corev1.Secret) error {
	config, err := NewKubernetesClient()
	if err != nil {
		return fmt.Errorf("failed to get kubernetes config: %v", err)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create kubernetes client: %v", err)
	}
	_, err = client.CoreV1().Secrets(secret.Namespace).Create(context.TODO(), secret, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create Secret: %v", err)
	}
	return nil
}
