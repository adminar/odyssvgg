package repository

import (
	"encoding/base64"
	"fmt"

	"github.com/Riyoukou/odyssey/app/model"
	"github.com/Riyoukou/odyssey/pkg/logger"
)

// cicd_cluster
func FetchClusters() ([]model.ClusterTable, error) {
	var clusters []model.ClusterTable
	if err := DB.Find(&clusters).Error; err != nil {
		logger.Errorf("Failed to fetch clusters: %v", err)
		return nil, err
	}

	return clusters, nil
}

func CreateCluster(cluster model.ClusterTable) error {
	if err := DB.Where("name = ? AND api_server = ?", cluster.Name, cluster.APIServer).
		First(&model.ClusterTable{}).Error; err == nil {
		logger.Errorf("Cluster already exists: name=%s, api_server=%s", cluster.Name, cluster.APIServer)
		return err
	}

	cluster.Config = base64.StdEncoding.EncodeToString([]byte(cluster.Config))

	if err := DB.Create(&cluster).Error; err != nil {
		logger.Errorf("Failed to create cluster: %v", err)
		return err
	}

	return nil
}

func DeleteCluster(clusterID int64) error {
	if err := DB.Delete(&model.ClusterTable{}, clusterID).Error; err != nil {
		logger.Errorf("Failed to delete cluster: %v", err)
		return err
	}

	return nil
}

func GetClusterByName(name string) (model.ClusterTable, error) {
	var cluster model.ClusterTable
	if err := DB.Where("name = ?", name).
		First(&cluster).Error; err != nil {
		return model.ClusterTable{}, err
	}

	// Base64 解码 kubeconfig
	kubeconfigBytes, err := base64.StdEncoding.DecodeString(cluster.Config)
	if err != nil {
		return model.ClusterTable{}, fmt.Errorf("invalid kubeconfig encoding: %w", err)
	}

	cluster.Config = string(kubeconfigBytes)
	return cluster, nil
}

func UpdateCluster(cluster model.ClusterTable) error {
	if err := DB.Model(&model.ClusterTable{}).Where("id = ?", cluster.ID).Updates(cluster).Error; err != nil {
		logger.Errorf("Failed to update cluster: %v", err)
		return err
	}

	return nil
}

// cicd_project
func FetchProjects() ([]model.ProjectTable, error) {
	var projects []model.ProjectTable
	if err := DB.Find(&projects).Error; err != nil {
		logger.Errorf("Failed to fetch projects: %v", err)
		return nil, err
	}

	return projects, nil
}

func CreateProject(project model.ProjectTable) error {
	if err := DB.Where("name = ?", project.Name).
		First(&model.ProjectTable{}).Error; err == nil {
		logger.Errorf("Project already exists: name=%s", project.Name)
		return err
	}

	if err := DB.Create(&project).Error; err != nil {
		logger.Errorf("Failed to create project: %v", err)
		return err
	}

	return nil
}

func DeleteProject(projectID int64) error {
	if err := DB.Delete(&model.ProjectTable{}, projectID).Error; err != nil {
		logger.Errorf("Failed to delete project: %v", err)
		return err
	}

	return nil
}

func GetProjectByName(name string) (model.ProjectTable, error) {
	var project model.ProjectTable
	if err := DB.Where("name = ?", name).
		First(&project).Error; err != nil {
		return model.ProjectTable{}, err
	}

	return project, nil
}

func UpdateProject(project model.ProjectTable) error {
	if err := DB.Model(&model.ProjectTable{}).Where("id = ?", project.ID).Updates(project).Error; err != nil {
		logger.Errorf("Failed to update project: %v", err)
		return err
	}

	return nil
}

// cicd_service
func FetchServicesByProjectAndEnv(projectName, envName string) ([]model.ServiceTable, error) {
	var services []model.ServiceTable
	if err := DB.Find(&services, "project_name = ? AND env_name = ?", projectName, envName).Error; err != nil {
		logger.Errorf("Failed to fetch services: %v", err)
		return nil, err
	}

	return services, nil
}

func CreateService(service model.ServiceTable) error {
	if err := DB.Where("name = ? AND project_name = ? AND env_name = ?", service.Name, service.ProjectName, service.EnvName).
		First(&model.ServiceTable{}).Error; err == nil {
		logger.Errorf("Service already exists: name=%s project_name=%s env_name=%s", service.Name, service.ProjectName, service.EnvName)
		return err
	}

	if err := DB.Create(&service).Error; err != nil {
		logger.Errorf("Failed to create service: %v", err)
		return err
	}

	return nil
}

func DeleteService(serviceID int64) error {
	if err := DB.Delete(&model.ServiceTable{}, serviceID).Error; err != nil {
		logger.Errorf("Failed to delete service: %v", err)
		return err
	}

	return nil
}

func GetServiceByNameAndProjectByEnv(name, projectName, envName string) (model.ServiceTable, error) {
	var service model.ServiceTable
	if err := DB.Where("name = ? AND project_name = ? AND env_name = ?", name, projectName, envName).
		First(&service).Error; err != nil {
		return model.ServiceTable{}, err
	}

	return service, nil
}

func UpdateServiceByNameAndProjectByEnv(service model.ServiceTable) error {
	if err := DB.Model(&model.ServiceTable{}).Where("name = ? AND project_name = ? AND env_name = ?", service.Name, service.ProjectName, service.EnvName).Updates(service).Error; err != nil {
		logger.Errorf("Failed to update service: %v", err)
		return err
	}

	return nil
}

// code_library
func FetchCodeLibraries() ([]model.CodeLibraryTable, error) {
	var codeLibraries []model.CodeLibraryTable
	if err := DB.Find(&codeLibraries).Error; err != nil {
		logger.Errorf("Failed to fetch code libraries: %v", err)
		return nil, err
	}

	return codeLibraries, nil
}

func CreateCodeLibrary(codeLibrary model.CodeLibraryTable) error {
	if err := DB.Where("name = ? AND project_name = ?", codeLibrary.Name, codeLibrary.ProjectName).
		First(&model.CodeLibraryTable{}).Error; err == nil {
		logger.Errorf("Code library already exists: name=%s project_name=%s", codeLibrary.Name, codeLibrary.ProjectName)
		return err
	}
	if err := DB.Create(&codeLibrary).Error; err != nil {
		logger.Errorf("Failed to create code library: %v", err)
		return err
	}

	return nil
}

func DeleteCodeLibrary(codeLibraryID int64) error {
	if err := DB.Delete(&model.CodeLibraryTable{}, codeLibraryID).Error; err != nil {
		logger.Errorf("Failed to delete code library: %v", err)
		return err
	}

	return nil
}

func GetCodeLibraryByNameAndProject(name, projectName string) (model.CodeLibraryTable, error) {
	var codeLibrary model.CodeLibraryTable
	if err := DB.Where("name = ? AND project_name = ?", name, projectName).
		First(&codeLibrary).Error; err != nil {
		return model.CodeLibraryTable{}, err
	}

	return codeLibrary, nil
}

func UpdateCodeLibraryByNameAndProject(codeLibrary model.CodeLibraryTable) error {
	if err := DB.Model(&model.CodeLibraryTable{}).Where("name = ? AND project_name = ?", codeLibrary.Name, codeLibrary.ProjectName).Updates(codeLibrary).Error; err != nil {
		logger.Errorf("Failed to update code library: %v", err)
		return err
	}

	return nil
}

// code_source
func FetchCodeSources() ([]model.CodeSourceTable, error) {
	var codeSources []model.CodeSourceTable
	if err := DB.Find(&codeSources).Error; err != nil {
		logger.Errorf("Failed to fetch code sources: %v", err)
		return nil, err
	}

	return codeSources, nil
}

func CreateCodeSource(codeSource model.CodeSourceTable) error {
	if err := DB.Where("name = ?", codeSource.Name).
		First(&model.CodeSourceTable{}).Error; err == nil {
		logger.Errorf("Code source already exists: name=%s", codeSource.Name)
		return err
	}
	codeSource.PrivateToken = base64.StdEncoding.EncodeToString([]byte(codeSource.PrivateToken))
	if err := DB.Create(&codeSource).Error; err != nil {
		logger.Errorf("Failed to create code source: %v", err)
		return err
	}

	return nil
}

func DeleteCodeSource(codeSourceID int64) error {
	if err := DB.Delete(&model.CodeSourceTable{}, codeSourceID).Error; err != nil {
		logger.Errorf("Failed to delete code source: %v", err)
		return err
	}

	return nil
}

func GetCodeSourceByName(name string) (model.CodeSourceTable, error) {
	var codeSource model.CodeSourceTable
	if err := DB.Where("name = ?", name).
		First(&codeSource).Error; err != nil {
		return model.CodeSourceTable{}, err
	}
	// Base64 解码 kubeconfig
	privateToken, err := base64.StdEncoding.DecodeString(codeSource.PrivateToken)
	if err != nil {
		return model.CodeSourceTable{}, fmt.Errorf("invalid private token encoding: %w", err)
	}
	codeSource.PrivateToken = string(privateToken)
	return codeSource, nil
}

func UpdateCodeSourceByName(codeSource model.CodeSourceTable) error {
	if err := DB.Model(&model.CodeSourceTable{}).Where("name = ?", codeSource.Name).Updates(codeSource).Error; err != nil {
		logger.Errorf("Failed to update code source: %v", err)
		return err
	}

	return nil
}
