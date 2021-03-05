package models

import (
	"time"

	corev1 "k8s.io/api/core/v1"
)

type Environment struct {
	ID        string           `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string           `json:"name,omitempty" bson:"name,omitempty"`
	Namespace corev1.Namespace `json:"namespace,omitempty" bson:"namespace,omitempty"`
	Isolation string           `json:"isolation,omitempty" bson:"isolation,omitempty"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}
