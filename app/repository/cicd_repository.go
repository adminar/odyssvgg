package repository

import (
	"encoding/base64"

	"github.com/Riyoukou/odyssey/app/model"
	"github.com/Riyoukou/odyssey/pkg/logger"
)

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
