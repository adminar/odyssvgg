package model

import (
	"gorm.io/datatypes"
)

// repository_model
type ClusterTable struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	APIServer   string `json:"api_server"`
	Config      string `json:"config"`
	Region      string `json:"region"`
	Version     string `json:"version"`
	Description string `json:"description"`
}

func (ClusterTable) TableName() string {
	return "clusters"
}

type ProjectTable struct {
	ID       int64          `json:"id"`
	Name     string         `json:"name"`
	Env      datatypes.JSON `json:"env"`
	Clusters datatypes.JSON `json:"clusters"`
}

func (ProjectTable) TableName() string {
	return "projects"
}

type EnvTable struct {
	ID           int64          `json:"id"`
	Name         string         `json:"name"`
	ProjectName  string         `json:"project_name"`
	Type         string         `json:"type"`
	NamespaceMap datatypes.JSON `json:"namespace_map"`
}

type NamespaceMap struct {
	Cluster   string `json:"cluster"`
	Namespace string `json:"namespace"`
}

func (EnvTable) TableName() string {
	return "envs"
}
