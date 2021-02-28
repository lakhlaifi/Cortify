package dao

import (
	"context"
	"cortify/models"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

//Handle CM per containerScope
func ConstructConfigMap(svc models.Service) (*corev1.ConfigMap, error) {
	// var configmaps []corev1.ConfigMap
	// for i, k := range svc.KService.Specs.Containers {
	configmaps := corev1.ConfigMap{
		TypeMeta:   metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{Name: svc.Name, Namespace: svc.Namespace},
		Data:       svc.ConfigMap.Data,
	}

	return &configmaps, nil
}

// func CreateConfigMap(namespace string, configmap *corev1.ConfigMap) error {
func CreateConfigMap(configmap *corev1.ConfigMap) error {
	config, err := NewKubernetesClient()
	if err != nil {
		return fmt.Errorf("failed to get kubernetes config: %v", err)
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return fmt.Errorf("failed to create kubernetes client: %v", err)
	}
	_, err = client.CoreV1().ConfigMaps(configmap.Namespace).Create(context.TODO(), configmap, metav1.CreateOptions{})
	if err != nil {
		return fmt.Errorf("failed to create ConfigMap: %v", err)
	}
	return nil
}
