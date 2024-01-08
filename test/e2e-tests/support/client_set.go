package support

import (
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	v1 "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	knativetest "knative.dev/pkg/test"
	"os"
	"path/filepath"
)

type Clients struct {
	Kubernetes  *kubernetes.Clientset
	Pipeline    v1.PipelineInterface
	PipelineRun v1.PipelineRunInterface
}

func CreateClients(namespace string) (Clients, error) {
	clients := Clients{}

	clientSet, err := ClusterClientSet()
	if err != nil {
		return clients, err
	}

	config, err := knativetest.BuildClientConfig(knativetest.Flags.Kubeconfig, knativetest.Flags.Cluster)
	if err != nil {
		return clients, err
	}

	tektonClientSet, err := versioned.NewForConfig(config)

	clients.Kubernetes = clientSet
	clients.Pipeline = tektonClientSet.TektonV1().Pipelines(namespace)
	clients.PipelineRun = tektonClientSet.TektonV1().PipelineRuns(namespace)

	return clients, nil
}

func ClusterClientSet() (*kubernetes.Clientset, error) {
	kubeConfig := os.Getenv("KUBECONFIG")
	if kubeConfig == "" {
		kubeConfig = getDefaultKubeConfigLocation()
	}

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func getDefaultKubeConfigLocation() string {
	return filepath.Join(homedir.HomeDir(), ".kube", "config")
}
