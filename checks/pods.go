package checks

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Pods runs all pod checks
func Pods() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}
	// Loop over all the pods in all the namespaces
	for _, pod := range pods.Items {
		// Loop over all the containers within a pod
		for _, container := range pod.Status.ContainerStatuses {

			// If a container is in a bad state need to alert
			if container.State.Waiting != nil && container.State.Waiting.Reason == "CrashLoopBackOff" {
				fmt.Println("Warning: The following pod is in a CrashLoopBackOff")
				fmt.Println("Container Name:", container.Name)
				fmt.Println("Container Restarts:", container.RestartCount)
			}
		}
	}
}
