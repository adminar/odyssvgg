package model

import (
	"time"

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
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	ProjectName string `json:"project_name"`
	Type        string `json:"type"`
	Namespace   string `json:"namespace"`
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

type BuildRecord struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Env         string    `json:"env"`
	Tag         string    `json:"tag"`
	Status      string    `json:"status"`
	ProjectName string    `json:"project_name"`
	BuildUser   string    `json:"build_user"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
}

func (BuildRecord) TableName() string {
	return "build_records"
}

type ApiBuildRecord struct {
	Describe    string                      `json:"describe"`
	Env         string                      `json:"env"`
	ProjectName string                      `json:"project_name"`
	Services    []ApiCICDBuildRecordService `json:"services"`
	Tag         string                      `json:"tag"`
	BuildUser   string                      `json:"build_user"`
	Name        string                      `json:"name"`
}

type ApiCICDBuildRecordService struct {
	ServiceName string `json:"service_name"`
	Branch      string `json:"branch"`
}

type BuildServiceRecord struct {
	ID              int64  `json:"id"`
	ServiceName     string `json:"service_name"`
	ProjectName     string `json:"project_name"`
	Env             string `json:"env"`
	Image           string `json:"image"`
	BuildRecordName string `json:"build_record_name"`
	BuildURL        string `json:"build_url"`
	Status          string `json:"status"`
	Branch          string `json:"branch"`
	BuildID         int64  `json:"build_id"`
}

func (BuildServiceRecord) TableName() string {
	return "build_service_records"
}

type DeployRecord struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	Env             string    `json:"env"`
	ProjectName     string    `json:"project_name"`
	DeployUser      string    `json:"deploy_user"`
	BuildRecordName string    `json:"build_record_name"`
	Status          string    `json:"status"`
	Tag             string    `json:"tag"`
	ClusterNames    string    `json:"cluster_names"`
	CreatedAt       time.Time `json:"created_at"`
	Description     string    `json:"description"`
}

func (DeployRecord) TableName() string {
	return "deploy_records"
}

type DeployServiceRecord struct {
	ID               int64  `json:"id"`
	ServiceName      string `json:"service_name"`
	ProjectName      string `json:"project_name"`
	Env              string `json:"env"`
	DeployRecordName string `json:"deploy_record_name"`
	ClusterName      string `json:"cluster_name"`
	Status           string `json:"status"`
	Image            string `json:"image"`
}

func (DeployServiceRecord) TableName() string {
	return "deploy_service_records"
}
