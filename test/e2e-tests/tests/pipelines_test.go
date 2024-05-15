package tests

import (
	"fmt"
	"github.com/opendatahub-io/ai-edge/test/e2e-tests/support"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
	"time"
)

const (
	AIEdgeE2EPipelineDirectoryRelativePath            = "../../../pipelines/tekton/aiedge-e2e"
	AIEdgeE2EBikeRentalsPipelineRunRelativePath       = AIEdgeE2EPipelineDirectoryRelativePath + "/aiedge-e2e.bike-rentals.pipelinerun.yaml"
	AIEdgeE2ETensorflowHousingPipelineRunRelativePath = AIEdgeE2EPipelineDirectoryRelativePath + "/aiedge-e2e.tensorflow-housing.pipelinerun.yaml"

	GitOpsUpdatePipelineDirectoryRelativePath            = "../../../pipelines/tekton/gitops-update-pipeline"
	GitOpsUpdateBikeRentalsPipelineRunRelativePath       = GitOpsUpdatePipelineDirectoryRelativePath + "/example-pipelineruns/gitops-update-pipelinerun-bike-rentals.yaml"
	GitOpsUpdateTensorflowHousingPipelineRunRelativePath = GitOpsUpdatePipelineDirectoryRelativePath + "/example-pipelineruns/gitops-update-pipelinerun-tensorflow-housing.yaml"
)

func CreateContext() context.Context {
	return context.Background()
}

func WaitForAllPipelineRunsToComplete(ctx context.Context, config *support.Config) error {
	callback := func() (bool, error) {
		list, err := config.Clients.PipelineRun.List(ctx, metav1.ListOptions{})
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
// This is used to build then apply the kustomize config for each pipeline
func init() {
	ctx := CreateContext()

	config, err := support.GetConfig()
	if err != nil {
		panic(err.Error())
	}

	err = support.RunSetup(ctx, config)
	if err != nil {
		panic(err.Error())
	}

	kustomizePaths := []string{AIEdgeE2EPipelineDirectoryRelativePath, GitOpsUpdatePipelineDirectoryRelativePath}

	for _, path := range kustomizePaths {
		resourceMap, err := support.KustomizeBuild(path)
		if err != nil {
			panic(fmt.Sprintf("error while building kustomize : %v", err.Error()))
		}

		err = support.CreateObjectsFromResourceMap(ctx, config.Clients, resourceMap, config.Namespace)
		if err != nil {
			panic(fmt.Errorf("error while creating objects from kustomize resources : %v", err.Error()))
		}
	}

}

func Test_MLOpsPipeline_S3Fetch(t *testing.T) {
	ctx := CreateContext()

	config, err := support.GetConfig()
	if err != nil {
		panic(err.Error())
	}

	if !config.S3FetchEnabled {
		t.Skipf("skipping %v, S3 is not enabled by the given configuration", t.Name())
	}

	pipelineRun, err := support.ReadFileAsPipelineRun(AIEdgeE2EBikeRentalsPipelineRunRelativePath)
	if err != nil {
		t.Fatal(err.Error())
	}

	if config.S3FetchConfig.SelfSignedCert != "" {
		support.MountConfigMapAsWorkspaceToPipelineRun("s3-self-signed-cert", "s3-ssl-cert", &pipelineRun)
	}

	support.SetPipelineRunParam("s3-bucket-name", support.NewStringParamValue(config.S3FetchConfig.BucketName), &pipelineRun)
	support.SetPipelineRunParam("target-image-tag-references", support.NewArrayParamValue(config.TargetImageTags), &pipelineRun)

	_, err = config.Clients.PipelineRun.Create(ctx, &pipelineRun, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = WaitForAllPipelineRunsToComplete(ctx, config)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func Test_MLOpsPipeline_GitFetch(t *testing.T) {
	ctx := CreateContext()

	config, err := support.GetConfig()
	if err != nil {
		panic(err.Error())
	}

	if !config.GitFetchEnabled {
		t.Skipf("skipping %v, Git is not enabled by the given configuration", t.Name())
	}

	pipelineRun, err := support.ReadFileAsPipelineRun(AIEdgeE2ETensorflowHousingPipelineRunRelativePath)
	if err != nil {
		t.Fatal(err.Error())
	}

	if config.GitFetchConfig.SelfSignedCert != "" {
		support.MountConfigMapAsWorkspaceToPipelineRun("git-self-signed-cert", "git-ssl-cert", &pipelineRun)
	}

	support.SetPipelineRunParam("target-image-tag-references", support.NewArrayParamValue(config.TargetImageTags), &pipelineRun)
	support.SetPipelineRunParam("git-containerfile-repo", support.NewStringParamValue(config.GitFetchConfig.CONTAINERFILE_REPO), &pipelineRun)
	support.SetPipelineRunParam("git-containerfile-revision", support.NewStringParamValue(config.GitFetchConfig.CONTAINERFILE_REVISION), &pipelineRun)
	support.SetPipelineRunParam("containerRelativePath", support.NewStringParamValue(config.GitFetchConfig.CONTAINER_RELATIVE_PATH), &pipelineRun)
	support.SetPipelineRunParam("git-model-repo", support.NewStringParamValue(config.GitFetchConfig.MODEL_REPO), &pipelineRun)
	support.SetPipelineRunParam("modelRelativePath", support.NewStringParamValue(config.GitFetchConfig.MODEL_RELATIVE_PATH), &pipelineRun)
	support.SetPipelineRunParam("git-model-revision", support.NewStringParamValue(config.GitFetchConfig.MODEL_REVISION), &pipelineRun)
	support.SetPipelineRunParam("model-dir", support.NewStringParamValue(config.GitFetchConfig.MODEL_DIR), &pipelineRun)

	_, err = config.Clients.PipelineRun.Create(ctx, &pipelineRun, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = WaitForAllPipelineRunsToComplete(ctx, config)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func Test_GitOpsUpdatePipeline(t *testing.T) {
	ctx := CreateContext()

	config, err := support.GetConfig()
	if err != nil {
		panic(err.Error())
	}

	if !config.GitOpsEnabled {
		t.Skipf("skipping %v, Gitops pipeline not configured", t.Name())
	}

	if !config.GitFetchEnabled && !config.S3FetchEnabled {
		t.Skipf("skipping %v, neither fetch methods were not configured", t.Name())
	}

	clients, err := support.CreateClients(config.Namespace)
	if err != nil {
		t.Fatal(err.Error())
	}

	// the pipeline runs we can run here are based on the fetch models used by the mlops pipeline
	// because this pipeline depends on the finished pipeline run of the mlops pipeline
	// -> https://issues.redhat.com/browse/RHOAIENG-4982
	// there is also a possibilty to split these based on the fetch types like the mlops pipeline
	var pipelineRunPaths []string
	if config.GitFetchEnabled {
		pipelineRunPaths = append(pipelineRunPaths, GitOpsUpdateTensorflowHousingPipelineRunRelativePath)
	}

	if config.S3FetchEnabled {
		pipelineRunPaths = append(pipelineRunPaths, GitOpsUpdateBikeRentalsPipelineRunRelativePath)
	}

	gitURL, err := support.ParseGitURL(config.GitOpsConfig.Repo)
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, path := range pipelineRunPaths {
		pipelineRun, err := support.ReadFileAsPipelineRun(path)
		if err != nil {
			t.Fatal(err.Error())
		}

		support.SetPipelineRunParam("gitServer", support.NewStringParamValue(gitURL.Server), &pipelineRun)
		support.SetPipelineRunParam("gitApiServer", support.NewStringParamValue(config.GitOpsConfig.ApiServer), &pipelineRun)
		support.SetPipelineRunParam("gitOrgName", support.NewStringParamValue(gitURL.OrgName), &pipelineRun)
		support.SetPipelineRunParam("gitRepoName", support.NewStringParamValue(gitURL.RepoName), &pipelineRun)
		support.SetPipelineRunParam("gitRepoBranchBase", support.NewStringParamValue(config.GitOpsConfig.Branch), &pipelineRun)

		_, err = clients.PipelineRun.Create(ctx, &pipelineRun, metav1.CreateOptions{})
		if err != nil {
			t.Fatal(err.Error())
		}
	}

	err = WaitForAllPipelineRunsToComplete(ctx, config)
	if err != nil {
		t.Fatal(err.Error())
	}
}
