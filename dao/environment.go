package dao

import (
	"context"
	"cortify/models"
	"encoding/json"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
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
func ConstructEnvironment(env models.Environment) *corev1.Namespace {
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

	//Network Policy :
	return &environment
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

	//Create Network Policy
	np := CreateBaseNP(environment)
	_, err = client.NetworkingV1().NetworkPolicies(environment.Name).Create(context.TODO(), np, metav1.CreateOptions{})

	return nil
}

func CreateBaseNP(env *corev1.Namespace) *networkingv1.NetworkPolicy {
	np := networkingv1.NetworkPolicy{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: env.Name,
			Name:      env.Name + "-default-policy",
		},
		Spec: networkingv1.NetworkPolicySpec{
			// An empty PodSelector selects all pods in this Namespace.
			PodSelector: metav1.LabelSelector{},
			Ingress: []networkingv1.NetworkPolicyIngressRule{
				networkingv1.NetworkPolicyIngressRule{
					From: []networkingv1.NetworkPolicyPeer{
						networkingv1.NetworkPolicyPeer{
							// An empty PodSelector selects all pods in this Namespace.
							PodSelector: &metav1.LabelSelector{},
						},
					},
				},
			},
		},
	}

	return &np
}
