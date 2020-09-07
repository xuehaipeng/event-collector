package pkg

import (
	"flag"
	"fmt"
	"github.com/xuehaipeng/event-collector/pkg/generated/clientset/versioned"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"testing"
)

var kubeconfig string
var config *rest.Config

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "path to Kubernetes config file")
	flag.Parse()
	var err error

	if kubeconfig == "" {
		fmt.Println("using in-cluster configuration")
		config, err = rest.InClusterConfig()
	} else {
		fmt.Println("using configuration from '%s'", kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}

	if err != nil {
		panic(err)
	}
}

func TestAbc(t *testing.T) {
	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		fmt.Println(err)
	}
	alpha1 := clientset.RedisV1alpha1()
	fmt.Println(alpha1)
}
