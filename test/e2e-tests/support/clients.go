package support

import (
	"github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	pipelinev1 "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	knativetest "knative.dev/pkg/test"
)

type Clients struct {
	Kubernetes  *kubernetes.Clientset
	Task        pipelinev1.TaskInterface
	Pipeline    pipelinev1.PipelineInterface
	PipelineRun pipelinev1.PipelineRunInterface
}

func CreateClients(namespace string) (Clients, error) {
	clients := Clients{}

	config, err := knativetest.BuildClientConfig("", "")
	if err != nil {
		return clients, err
	}

	clientSet, err := ClusterClientSet(config)
	if err != nil {
		return clients, err
	}

	tektonClientSet, err := versioned.NewForConfig(config)

	clients.Kubernetes = clientSet
	clients.Pipeline = tektonClientSet.TektonV1().Pipelines(namespace)
	clients.PipelineRun = tektonClientSet.TektonV1().PipelineRuns(namespace)
	clients.Task = tektonClientSet.TektonV1().Tasks(namespace)

	return clients, nil
}

func ClusterClientSet(config *rest.Config) (*kubernetes.Clientset, error) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}
