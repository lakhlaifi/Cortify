package models

type ConfigMap struct {
	ID       string        `json:"_id,omitempty" bson:"_id,omitempty"`
	Metadata ConfigMapMeta `json:"metadata,omitempty" bson:"metadata,omitempty"`
	Data     ConfigMapData `json:"data,omitempty" bson:"data,omitempty"`
}

type ConfigMapMeta struct {
	Name      string   `json:"name,omitempty" bson:"name,omitempty"`
	Namespace string   `json:"namespace,omitempty" bson:"namespace,omitempty"`
	Labels    []string `json:"labels,omitempty" bson:"labels,omitempty"`
}

type ConfigMapData struct {
	ConfigData []string `json:"configdata" bson:"configdata,omitempty"`
}

// ConfigMapModel define the model structure
type ConfigMapModel struct{}
