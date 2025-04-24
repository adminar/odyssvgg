package repository

import (
	"encoding/base64"
	"fmt"
)

func CreateK8SCluster(cluster model.K8SCluster) error {
	// 检查是否已存在相同 name 和 api_server 的集群
	var existing model.K8SCluster
	if err := DB.Where("name = ? AND api_server = ?", k8sCluster.Name, k8sCluster.APIServer).First(&existing).Error; err == nil {
		return fmt.Errorf("k8s cluster with name %s and api_server %s already exists", k8sCluster.Name, k8sCluster.APIServer)
	}

	// 编码 kubeconfig 内容
	k8sCluster.Config = base64.StdEncoding.EncodeToString([]byte(k8sCluster.Config))
	// 创建记录
	if err := DB.Create(&k8sCluster).Error; err != nil {
		return fmt.Errorf("failed to create k8s cluster: %w", err)
	}

	return nil
}
