package models

import (
	"time"

	corev1 "k8s.io/api/core/v1"
)

type Service struct {
	ID        string             `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Namespace string             `json:"namespace,omitempty" bson:"namespace,omitempty"`
	KService  KnativeService     `json:"kservice,omitempty" bson:"kservice,omitempty"`
	ConfigMap []corev1.ConfigMap `json:"configmap,omitempty" bson:"configmap,omitempty"`
	Secret    []corev1.Secret    `json:"secret,omitempty" bson:"secret,omitempty"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
type KnativeService struct {
	Config KnativeServiceConfig `json:"config,omitempty" bson:"config,omitempty"`
	Specs  KnativeServiceSpecs  `json:"specs,omitempty" bson:"specs,omitempty"`
	// Traffic  KnativeServiceTraffic `json:"traffic,omitempty" bson:"traffic,omitempty"`
	// KnativeServiceContainers
	// KnativeServiceVolumes
}

type KnativeServiceConfig struct {
	// Name        string            `json:"name,omitempty" bson:"name,omitempty"`
	// Namespace   string            `json:"namespace,omitempty" bson:"namespace,omitempty"`
	Labels      map[string]string `json:"labels,omitempty" bson:"labels,omitempty"`
	Annotations map[string]string `json:"annotations,omitempty" bson:"annotations,omitempty"`
	// Scale                  string            `json:"scale,omitempty" bson:"scale,omitempty"`
	// MinScale               int               `json:"minscale,omitempty" bson:"maxscale,omitempty"`
	// MaxScale               int
	// ConcurrencyTarget      int
	// ConcurrencyLimit       int
	// ConcurrencyUtilization int
	// AutoscaleWindow        string
	// LabelsService          []string
	// LabelsRevision         []string
	// RevisionName           string
	// AnnotationsService     []string
	// AnnotationsRevision    []string
	// ClusterLocal           bool
	// ScaleInit              int
}

type KnativeServiceSpecs struct {
	Replicas       int                `default:"1" json:"replicas,omitempty" bson:"replicas,omitempty"`
	InitContainers []corev1.Container `json:"initcontainers" bson:"initcontainers,omitempty"`
	Containers     []corev1.Container `json:"containers" bson:"containers,omitempty"`
	Volumes        []corev1.Volume    `json:"volumes" bson:"volumes,omitempty"`
	// Selectors      []string                       `json:"selectors,omitempty" bson:"selectors,omitempty"`
}

type VolumeSources struct {
	ConfigMap             corev1.ConfigMapVolumeSource             `json:"volconfigmap" bson:"volconfigmap,omitempty"`
	Secret                corev1.SecretVolumeSource                `json:"volsecret" bson:"volsecret,omitempty"`
	EmptyDir              corev1.EmptyDirVolumeSource              `json:"volemptydir" bson:"volemptydir,omitempty"`
	PersistentVolumeClaim corev1.PersistentVolumeClaimVolumeSource `json:"volpvc" bson:"volpvc,omitempty"`
	//EphemeralVolumeSource
	//GitRepoVolumeSource
	//NFSVolumeSource
}

// type KnativeServiceTraffic struct {
// 	//TODO
// }
