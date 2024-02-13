package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/opendatahub-io/ai-edge/test/e2e-tests/support"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	AIEdgeE2EPipelineDirectoryRelativePath = "../../../pipelines/tekton/aiedge-e2e"
	PipelineRunFileRelativePath            = AIEdgeE2EPipelineDirectoryRelativePath + "/aiedge-e2e.pipelinerun.yaml"
)

func CreateContext() context.Context {
	return context.Background()
}

func WaitForAllPipelineRunsToComplete(ctx context.Context, clients *support.Clients) error {
	callback := func() (bool, error) {
		list, err := clients.PipelineRun.List(ctx, metav1.ListOptions{})
		if err != nil {
			return false, err
		}

		pipelineStillRunning := false

		for _, pipelineRun := range list.Items {
			// conditions may be empty when pipeline run is first created
			// this means it is counted as still running
			if len(pipelineRun.Status.Conditions) == 0 {
				pipelineStillRunning = true
				break
			}

			for _, condition := range pipelineRun.Status.Conditions {
				if condition.Reason == "Failed" {
					return false, fmt.Errorf("pipelineRun \"%v\" failed with message \"%v\"", pipelineRun.Name, condition.Message)

				} else if condition.Reason == "Running" {
					pipelineStillRunning = true
					break
				}
			}
		}

		return !pipelineStillRunning, nil
	}

	return support.WaitFor(18*time.Minute, 10*time.Second, callback)
}

// init is called before any test is run, and it is called once.
// This is used to build then apply the kustomize config for the main pipeline
func init() {
	resourceMap, err := support.KustomizeBuild(AIEdgeE2EPipelineDirectoryRelativePath)
	if err != nil {
		panic(fmt.Sprintf("error while building kustomize : %v", err.Error()))
	}

	ctx := CreateContext()

	options, err := support.GetOptions()
	if err != nil {
		panic(err.Error())
	}

	clients, err := support.CreateClients(options.ClusterNamespace)
	if err != nil {
		panic(fmt.Sprintf("error while creating client : %v", err.Error()))
	}

	err = support.CreateObjectsFromResourceMap(ctx, &clients, resourceMap, options.ClusterNamespace)
	if err != nil {
		panic(fmt.Errorf("error while creating objects from kustomize resources : %v", err.Error()))
	}
}

func Test_DefaultPipelineRun(t *testing.T) {
	ctx := CreateContext()

	options, err := support.GetOptions()
	if err != nil {
		t.Fatal(err.Error())
	}

	clients, err := support.CreateClients(options.ClusterNamespace)
	if err != nil {
		t.Fatal(err.Error())
	}

	pipelineRun, err := support.ReadFileAsPipelineRun(PipelineRunFileRelativePath)
	if err != nil {
		t.Fatal(err.Error())
	}

	// setting these values to the options passed in as env vars
	support.SetPipelineRunParam("s3-bucket-name", support.NewStringParamValue(options.S3BucketName), &pipelineRun)
	support.SetPipelineRunParam("target-imagerepo", support.NewStringParamValue(options.RegistryRepoName), &pipelineRun)

	_, err = clients.PipelineRun.Create(ctx, &pipelineRun, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = WaitForAllPipelineRunsToComplete(ctx, &clients)
	if err != nil {
		t.Fatal(err.Error())
	}
}
