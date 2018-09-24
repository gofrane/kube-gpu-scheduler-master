package main

import (
	"fmt"

	"kube-gpu-scheduler-master/pkg/kube-gpu-scheduler/client"
	nodeutil "kube-gpu-scheduler-master/pkg/kube-gpu-scheduler/node"
	podutil "kube-gpu-scheduler-master/pkg/kube-gpu-scheduler/pod"
	resourceutil "kube-gpu-scheduler-master/pkg/kube-gpu-scheduler/resource"
)

func main() {
	clientset, erro := client.CreateClient()
	if erro != nil {
		fmt.Println(erro)
	}
	stopChannel := make(chan struct{})
	nodes, err := nodeutil.ReadyNodes(clientset, stopChannel)
	if err != nil {
		fmt.Println(err)
	}
	narm := resourceutil.NamespaceResourceMap{}
	ncm := resourceutil.NamespaceCPUMap{}
	nmm := resourceutil.NamespaceMemoryMap{}
	ngm := resourceutil.NamespaceGPUMap{}
	for _, node := range nodes {
		pods, err := podutil.ListPodsOnANode(clientset, node)
		if err != nil {
			fmt.Println(err)
		}
		narm = resourceutil.ResourceListerforPod(pods, narm, ncm, nmm, ngm)
	}
	fmt.Println("Resource Map:",narm)
	fmt.Println("CPU Map:",ncm)
	fmt.Println("GPU Map:", ngm)
	fmt.Println("Memory Map:", nmm)
}
