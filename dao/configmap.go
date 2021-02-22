package dao

import (
	"context"
	"cortify/models"
	"errors"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ConstructConfigMap(name string, namespace string, ksvc models.Service) (*corev1.ConfigMap, error) {
	if name == "" || namespace == "" {
		return nil, errors.New("internal: no name or namespace provided when constructing a service")
	}
	configmap := corev1.ConfigMap{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{Name: ksvc.KService.Metadata.Name, Namespace: namespace},
		Immutable:  new(bool),
		Data:       ksvc.ConfigMap.Data,
		BinaryData: map[string][]byte{},
	}
	return &configmap, nil
}

func CreateConfigMap(namespace string, configmap *corev1.ConfigMap) error {
	config, err := NewKubernetesClient()
	if err != nil {
		return fmt.Errorf("failed to get kubernetes config: %v", err)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create kubernetes client: %v", err)
	}
	_, err = client.CoreV1().ConfigMaps(namespace).Create(context.TODO(), configmap, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create ConfigMap: %v", err)
	}
	return nil
}
