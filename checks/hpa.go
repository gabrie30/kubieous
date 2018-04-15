package checks

import (
	"fmt"

	"k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// HPA calls checks for horizontal pod autoscalers
func HPA() {
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

	// hpas, err := clientset.AutoscalingV1().HorizontalPodAutoscalers("").List(metav1.ListOptions{})
	// hpaClient := clientset.AutoscalingV1().RESTClient
	// meta.NewAccessor().Namespace(data.Object)

	hpaWatch, _ := clientset.AutoscalingV1().HorizontalPodAutoscalers("").Watch(metav1.ListOptions{})

	for {
		select {
		case event := <-hpaWatch.ResultChan():
			handleEventz(event)
		}
	}
}

func handleEventz(event watch.Event) {
	hpa := event.Object.(*v1.HorizontalPodAutoscaler)

	if hpa.Status.CurrentReplicas > hpa.Status.DesiredReplicas {
		fmt.Println(hpa.ObjectMeta.Namespace + " is scaling DOWN")
		fmt.Println("Current VS Desired Replicas: ", hpa.Status.CurrentReplicas, hpa.Status.DesiredReplicas)
		fmt.Println("Current CPU utilization: ", *hpa.Status.CurrentCPUUtilizationPercentage)
	} else if hpa.Status.CurrentReplicas < hpa.Status.DesiredReplicas {
		fmt.Println(hpa.ObjectMeta.Namespace + " is scaling up")
		fmt.Println("Current VS Desired Replicas: ", hpa.Status.CurrentReplicas, hpa.Status.DesiredReplicas)
		fmt.Println("Current CPU utilization: ", *hpa.Status.CurrentCPUUtilizationPercentage)
	}
}
