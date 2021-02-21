package dao

import (
	"flag"
	"fmt"

	"k8s.io/client-go/tools/clientcmd"
	clientservingv1 "knative.dev/client/pkg/serving/v1"
	servingv1 "knative.dev/serving/pkg/client/clientset/versioned/typed/serving/v1"
)

func NewServingClient(namespace string) (clientservingv1.KnServingClient, error) {
	kubeconfig = flag.String("kubeconfig", "/root/DevOps-Cegedim/persona/Cortify/kubeconfig", "absolute path to the kubeconfig file")
	flag.Parse()
	fmt.Println("Kubeconfig :", kubeconfig)
	restConfig, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// restConfig, err := getClientConfig().ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to create new serving client: %v", err)
	}
	fmt.Errorf("Succeed 1 to create new serving client")
	servingClient, err := servingv1.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create new serving client: %v", err)
	}
	fmt.Errorf("Succeed 2 to create new serving client")
	client := clientservingv1.NewKnServingClient(servingClient, namespace)

	return client, nil
}
