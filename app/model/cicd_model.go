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

type ProjectsTable struct {
	ID       int64          `json:"id"`
	Name     string         `json:"name"`
	Env      string         `json:"env"`
	Clusters datatypes.JSON `json:"clusters"`
}

func (ProjectsTable) TableName() string {
	return "projects"
}
