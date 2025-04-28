package utils

import (
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func CreateKubernetesClientset(kubernetesConfig string) *kubernetes.Clientset {
	config, err := clientcmd.RESTConfigFromKubeConfig([]byte(kubernetesConfig))
	if err != nil {
		log.Fatalf("Failed to create REST config: %v", err)
	}

	// 创建 Clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}
