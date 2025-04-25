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

type ServiceTable struct {
	ID              int64          `json:"id"`
	Name            string         `json:"name"`
	ProjectName     string         `json:"project_name"`
	EnvName         string         `json:"env_name"`
	CodeLibraryName string         `json:"code_library_name"`
	DeployMap       datatypes.JSON `json:"deploy_map"`
}

func (ServiceTable) TableName() string {
	return "services"
}

type CodeLibraryTable struct {
	ID             int64  `json:"id"`
	Name           string `json:"name"`
	ProjectName    string `json:"project_name"`
	URL            string `json:"url"`
	CodeSourceName string `json:"code_source_name"`
}

func (CodeLibraryTable) TableName() string {
	return "code_libraries"
}

type CodeSourceTable struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	Type         string `json:"type"`
	PrivateToken string `json:"private_token"`
}

func (CodeSourceTable) TableName() string {
	return "code_sources"
}
