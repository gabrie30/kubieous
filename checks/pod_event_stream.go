package checks

// Other is for experiments
import (
	"fmt"
	"time"

	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

// PodEventStream is for listening to pod events
func PodEventStream() {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	watchlist := cache.NewListWatchFromClient(clientset.Core().RESTClient(), "pods", v1.NamespaceDefault,
		fields.Everything())
	_, controller := cache.NewInformer(
		watchlist,
		&v1.Pod{},
		time.Second*0,
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				fmt.Printf("add: %s \n", obj)
			},
			DeleteFunc: func(obj interface{}) {
				fmt.Printf("delete: %s \n", obj)
			},
			UpdateFunc: func(oldObj, newObj interface{}) {
				fmt.Printf("old: %s, new: %s \n", oldObj, newObj)
			},
		},
	)
	stop := make(chan struct{})
	go controller.Run(stop)
}
