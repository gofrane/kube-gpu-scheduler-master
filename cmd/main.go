// Example showing how to patch kubernetes resources.
package main

import (
	"fmt"

	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"PHCGS/pkg/kube-gpu-scheduler/client"
	"k8s.io/client-go/kubernetes"
)

var (
	//  Leave blank for the default context in your kube config.
	context = ""


)



func main() {
	//  Get the local kube config.
	fmt.Printf("Connecting to Kubernetes Context %v\n", context)
var clientset *kubernetes.Clientset
clientset,_= client.CreateClient(context)
	//  Scale our replication controller.
	fmt.Println("clienset", clientset)






}
