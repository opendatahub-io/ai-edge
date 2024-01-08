package tests

import (
	"github.com/opendatahub-io/ai-edge/test/e2e-tests/support"
	"golang.org/x/net/context"
	k8sApi "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
	"time"
)

func CreateContext() context.Context {
	return context.Background()
}

func Test_PipelineExists(t *testing.T) {
	ctx := CreateContext()

	clients, err := support.CreateClients("test-namespace")
	if err != nil {
		t.Fatal(err.Error())
	}

	list, err := clients.Pipeline.List(ctx, k8sApi.ListOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	expectedPipelineNames := []string{"build-container-image"}
	if len(list.Items) != len(expectedPipelineNames) {
		t.Fatalf("expected %v pipelines but %v existed on the cluster", len(expectedPipelineNames), len(list.Items))
	}

OUTER:
	for _, name := range expectedPipelineNames {
		for _, pipeline := range list.Items {
			if pipeline.Name == name {
				continue OUTER
			}
		}

		t.Fatalf("no pipeline had the name %v", name)
	}
}

func Test_PipelineRunsCompleted(t *testing.T) {
	ctx := CreateContext()

	clients, err := support.CreateClients("test-namespace")
	if err != nil {
		t.Fatal(err.Error())
	}

CheckRuns:
	for {
		list, err := clients.PipelineRun.List(ctx, k8sApi.ListOptions{})
		if err != nil {
			t.Fatal(err.Error())
		}

		for _, pipelineRun := range list.Items {
			for _, condition := range pipelineRun.Status.Conditions {
				if condition.Reason == "Failed" {
					t.Fatalf("pipelineRun \"%v\" failed with message \"%v\"", pipelineRun.Name, condition.Message)
					return
				}

				if condition.Reason == "Running" {
					time.Sleep(time.Second * 5)
					continue CheckRuns // it is still running so restart the loop and wait 5 seconds
				}
			}
		}

		break // none failed and none were still running so all pipeline runs have complete
	}
}

// func waitForPipelineRunsToFinish(pipelineRuns *v1.PipelineRunList) error {
// 	return nil
// }
