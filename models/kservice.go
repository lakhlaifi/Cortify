package models

import (
	"time"
)

type Service struct {
	ID        string         `json:"_id,omitempty" bson:"_id,omitempty"`
	KService  KnativeService `json:"service,omitempty" bson:"service,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}
type KnativeService struct {
	Base     KnativeServiceBase    `json:"base,omitempty" bson:"base,omitempty"`
	Metadata KnativeServiceMeta    `json:"metadata,omitempty" bson:"metadata,omitempty"`
	Specs    KnativeServiceSpecs   `json:"specs,omitempty" bson:"specs,omitempty"`
	Traffic  KnativeServiceTraffic `json:"traffic,omitempty" bson:"traffic,omitempty"`
	Config   ConfigMap             `json:"config,omitempty" bson:"config,omitempty"`
	// KnativeServiceContainers
	// KnativeServiceVolumes
}

type KnativeServiceBase struct {
	ApiVersion string `default:"serving.knative.dev/v1" json:"apiVersion,omitempty" bson:"apiVersion,omitempty"`
	Kind       string `default:"Service" json:"kind,omitempty" bson:"kind,omitempty"`
}

type KnativeServiceMeta struct {
	Name        string   `json:"name,omitempty" bson:"name,omitempty"`
	Namespace   string   `json:"namespace,omitempty" bson:"namespace,omitempty"`
	Labels      []string `json:"labels,omitempty" bson:"labels,omitempty"`
	Annotations []string `json:"annotations,omitempty" bson:"annotations,omitempty"`
}

type KnativeServiceSpecs struct {
	Replicas int `default:"1" json:"replicas,omitempty" bson:"replicas,omitempty"`
	// Selectors      []string                       `json:"selectors,omitempty" bson:"selectors,omitempty"`
	InitContainers []KnativeServiceinitContainers `json:"initcontainers" bson:"initcontainers,omitempty"`
	Containers     []KnativeServiceContainers     `json:"containers" bson:"containers,omitempty"`
}

type KnativeServiceTraffic struct {
	//TODO
}

type KnativeServiceContainers struct {
	Name  string   `json:"name,omitempty" bson:"name,omitempty"`
	Image string   `json:"image,omitempty" bson:"image,omitempty"`
	Envs  []string `json:"envs,omitempty" bson:"envs,omitempty"`
	// containerPorts  TODO
}

type KnativeServiceinitContainers struct {
	Name  string   `json:"name,omitempty" bson:"name,omitempty"`
	Image string   `json:"image,omitempty" bson:"image,omitempty"`
	Envs  []string `json:"envs,omitempty" bson:"envs,omitempty"`
	// containerPorts  TODO
}

type KnativeServiceVolumes struct {
	//TODO
}
