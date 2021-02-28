package dao

import (
	"fmt"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	clientservingv1 "knative.dev/client/pkg/serving/v1"
	servingv1 "knative.dev/serving/pkg/client/clientset/versioned/typed/serving/v1"
)

//Kubeconfig file test
var kubeconfig = "../kubeconfig"

func NewServingClient(namespace string) (clientservingv1.KnServingClient, error) {
	// TODO remove it when U configured in cluster client
	// kubeconfig = flag.String("kubeconfig", "/root/DevOps-Cegedim/persona/Cortify/kubeconfig", "absolute path to the kubeconfig file")
	// flag.Parse()
	// restConfig, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	// if err != nil {
	// 	panic(err.Error())
	// }
	// // restConfig, err := getClientConfig().ClientConfig()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to create new serving client: %v", err)
	// }
	restConfig, err := NewKubernetesClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create new serving client: %v", err)
	}
	servingClient, err := servingv1.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create new serving client: %v", err)
	}
	client := clientservingv1.NewKnServingClient(servingClient, namespace)

	return client, nil
}

func NewKubernetesClient() (*rest.Config, error) {
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Config Done")
	return restConfig, err
}
