package client

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/kubernetes"
	"github.com/logrus"
	"k8s.io/client-go/tools/clientcmd"
)

func CreateClient(pathTOCfg string) (*kubernetes.Clientset, error) {

	var config *rest.Config
	var err error

	if pathTOCfg==""{
		logrus.Info("using in cluster config")
		config,err=rest.InClusterConfig()
		//in cluster access

	}else {
		logrus.Info("Using out of cluster config")
		config,err=clientcmd.BuildConfigFromFlags("",pathTOCfg)
	}
	if err !=nil {
		return nil,err

	}


	return kubernetes.NewForConfig(config)
}
