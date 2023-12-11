package tests

import (
	"github.com/opendatahub-io/ai-edge/test/e2e-tests/oc"
	support "github.com/opendatahub-io/ai-edge/test/e2e-tests/support"
	"testing"
)

const (
	ContainerImagePipeline    = "../../../pipelines/tekton/build-container-image-pipeline/"
	ContainerImagePipelineRun = "../../../pipelines/tekton/build-container-image-pipeline/build-container-image-pipelinerun-bike-rentals.yaml"
)

func Test_ConnectingToCluster(t *testing.T) {
	_, err := support.ClusterClientSet()
	if err != nil {
		t.Fatal(err.Error())
	}
}

func Test_ContainerImagePipelineKustomizeApply(t *testing.T) {
	clientSet, err := support.ClusterClientSet()
	if err != nil {
		t.Fatal(err.Error())
	}

	ctx := support.CreateContext()

	if false {
		support.CreateTestNamespace(ctx, clientSet)
		defer support.DeleteTestNamespace(ctx, clientSet)
	}

	err = oc.Apply(oc.ApplyOptions{
		Kustomize: ContainerImagePipeline,
	})

	if err != nil {
		t.Fatal(err.Error())
	}
}

func Test_ContainerImagePipelineRuns(t *testing.T) {
	clientSet, err := support.ClusterClientSet()
	if err != nil {
		t.Fatal(err.Error())
	}

	ctx := support.CreateContext()

	if false {
		support.CreateTestNamespace(ctx, clientSet)
		defer support.DeleteTestNamespace(ctx, clientSet)
	}

	err = oc.Apply(oc.ApplyOptions{
		Kustomize: ContainerImagePipeline,
	})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = oc.Create(oc.CreateOptions{
		Filename: ContainerImagePipelineRun,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
}
