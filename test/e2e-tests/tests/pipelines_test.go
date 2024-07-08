package tests

import (
	"fmt"
	"github.com/opendatahub-io/ai-edge/test/e2e-tests/support"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
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

var (
	MLOpsTensorflowHousingCreatePipelineRunName string = ""
	MLOpsBikeRentalsCreatePipelineRunName       string = ""
)

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

	if !config.S3FetchConfig.Enabled {
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
	support.SetPipelineRunParam("git-containerfile-repo", support.NewStringParamValue(config.GitContainerFileRepo), &pipelineRun)
	support.SetPipelineRunParam("git-containerfile-revision", support.NewStringParamValue(config.GitContainerFileRevision), &pipelineRun)
	support.SetPipelineRunParam("containerRelativePath", support.NewStringParamValue(config.ContainerRelativePath), &pipelineRun)

	createdRun, err := config.Clients.PipelineRun.Create(ctx, &pipelineRun, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	MLOpsBikeRentalsCreatePipelineRunName = createdRun.Name

	err = WaitForAllPipelineRunsToComplete(ctx, createdRun.Name, config)
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

	if !config.GitFetchConfig.Enabled {
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
	support.SetPipelineRunParam("git-containerfile-repo", support.NewStringParamValue(config.GitContainerFileRepo), &pipelineRun)
	support.SetPipelineRunParam("git-containerfile-revision", support.NewStringParamValue(config.GitContainerFileRevision), &pipelineRun)
	support.SetPipelineRunParam("containerRelativePath", support.NewStringParamValue(config.ContainerRelativePath), &pipelineRun)
	support.SetPipelineRunParam("git-model-repo", support.NewStringParamValue(config.GitFetchConfig.ModelRepo), &pipelineRun)
	support.SetPipelineRunParam("modelRelativePath", support.NewStringParamValue(config.GitFetchConfig.ModelRelativePath), &pipelineRun)
	support.SetPipelineRunParam("git-model-revision", support.NewStringParamValue(config.GitFetchConfig.ModelRevision), &pipelineRun)
	support.SetPipelineRunParam("model-dir", support.NewStringParamValue(config.GitFetchConfig.ModelDir), &pipelineRun)

	createdRun, err := config.Clients.PipelineRun.Create(ctx, &pipelineRun, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	// saved so gitops test knows which pipeline run to get the results from
	MLOpsTensorflowHousingCreatePipelineRunName = createdRun.Name

	err = WaitForAllPipelineRunsToComplete(ctx, createdRun.Name, config)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func Test_GitOpsUpdatePipeline_S3Fetch(t *testing.T) {
	ctx := CreateContext()

	config, err := support.GetConfig()
	if err != nil {
		panic(err.Error())
	}

	if !config.GitOpsConfig.Enabled {
		t.Skipf("skipping %v, Gitops pipeline not configured", t.Name())
	}

	if !config.S3FetchConfig.Enabled {
		t.Skipf("skipping %v, s3_fetch was not enabled", t.Name())
	}

	clients, err := support.CreateClients(config.Namespace)
	if err != nil {
		t.Fatal(err.Error())
	}

	gitURL, err := support.ParseGitURL(config.GitOpsConfig.Repo)
	if err != nil {
		t.Fatal(err.Error())
	}

	pipelineRun, err := support.ReadFileAsPipelineRun(GitOpsUpdateBikeRentalsPipelineRunRelativePath)
	if err != nil {
		t.Fatal(err.Error())
	}

	support.SetPipelineRunParam("gitServer", support.NewStringParamValue(gitURL.Server), &pipelineRun)
	support.SetPipelineRunParam("gitApiServer", support.NewStringParamValue(config.GitOpsConfig.ApiServer), &pipelineRun)
	support.SetPipelineRunParam("gitOrgName", support.NewStringParamValue(gitURL.OrgName), &pipelineRun)
	support.SetPipelineRunParam("gitRepoName", support.NewStringParamValue(gitURL.RepoName), &pipelineRun)
	support.SetPipelineRunParam("gitRepoBranchBase", support.NewStringParamValue(config.GitOpsConfig.Branch), &pipelineRun)

	// we need to get the results from the pipeline run created by the ml ops pipeline
	mlopsPipelineRun, err := config.Clients.PipelineRun.Get(ctx, MLOpsBikeRentalsCreatePipelineRunName, metav1.GetOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	imageDigest, err := support.GetResultValueFromPipelineRun("buildah-sha", mlopsPipelineRun)
	if err != nil {
		t.Fatal(err.Error())
	}

	imageTagReferences, err := support.GetResultValueFromPipelineRun("target-image-tag-references", mlopsPipelineRun)
	if err != nil {
		t.Fatal(err.Error())
	}

	// there may be edge cases here where this fails
	registryRepo := strings.Split(imageTagReferences.StringVal, " ")[0]

	support.SetPipelineRunParam("image-digest", support.NewStringParamValue(imageDigest.StringVal), &pipelineRun)
	support.SetPipelineRunParam("image-registry-repo", support.NewStringParamValue(registryRepo), &pipelineRun)

	createdRun, err := clients.PipelineRun.Create(ctx, &pipelineRun, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = WaitForAllPipelineRunsToComplete(ctx, createdRun.Name, config)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func Test_GitOpsUpdatePipeline_GitFetch(t *testing.T) {
	ctx := CreateContext()

	config, err := support.GetConfig()
	if err != nil {
		panic(err.Error())
	}

	if !config.GitOpsConfig.Enabled {
		t.Skipf("skipping %v, Gitops pipeline not configured", t.Name())
	}

	if !config.GitFetchConfig.Enabled {
		t.Skipf("skipping %v, git_fetch was not enabled", t.Name())
	}

	clients, err := support.CreateClients(config.Namespace)
	if err != nil {
		t.Fatal(err.Error())
	}

	gitURL, err := support.ParseGitURL(config.GitOpsConfig.Repo)
	if err != nil {
		t.Fatal(err.Error())
	}

	pipelineRun, err := support.ReadFileAsPipelineRun(GitOpsUpdateTensorflowHousingPipelineRunRelativePath)
	if err != nil {
		t.Fatal(err.Error())
	}

	support.SetPipelineRunParam("gitServer", support.NewStringParamValue(gitURL.Server), &pipelineRun)
	support.SetPipelineRunParam("gitApiServer", support.NewStringParamValue(config.GitOpsConfig.ApiServer), &pipelineRun)
	support.SetPipelineRunParam("gitOrgName", support.NewStringParamValue(gitURL.OrgName), &pipelineRun)
	support.SetPipelineRunParam("gitRepoName", support.NewStringParamValue(gitURL.RepoName), &pipelineRun)
	support.SetPipelineRunParam("gitRepoBranchBase", support.NewStringParamValue(config.GitOpsConfig.Branch), &pipelineRun)

	// we need to get the results from the pipeline run created by the ml ops pipeline
	mlopsPipelineRun, err := config.Clients.PipelineRun.Get(ctx, MLOpsTensorflowHousingCreatePipelineRunName, metav1.GetOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	imageDigest, err := support.GetResultValueFromPipelineRun("buildah-sha", mlopsPipelineRun)
	if err != nil {
		t.Fatal(err.Error())
	}

	imageTagReferences, err := support.GetResultValueFromPipelineRun("target-image-tag-references", mlopsPipelineRun)
	if err != nil {
		t.Fatal(err.Error())
	}

	// there may be edge cases here where this fails
	registryRepo := strings.Split(imageTagReferences.StringVal, " ")[0]

	support.SetPipelineRunParam("image-digest", support.NewStringParamValue(imageDigest.StringVal), &pipelineRun)
	support.SetPipelineRunParam("image-registry-repo", support.NewStringParamValue(registryRepo), &pipelineRun)

	createdRun, err := clients.PipelineRun.Create(ctx, &pipelineRun, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = WaitForAllPipelineRunsToComplete(ctx, createdRun.Name, config)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func CreateContext() context.Context {
	return context.Background()
}

func WaitForAllPipelineRunsToComplete(ctx context.Context, pipelineRunName string, config *support.Config) error {
	callback := func() (bool, error) {
		pipelineRun, err := config.Clients.PipelineRun.Get(ctx, pipelineRunName, metav1.GetOptions{})
		if err != nil {
			return false, err
		}

		// conditions may be empty when pipeline run is first created
		// this means it is counted as still running
		if len(pipelineRun.Status.Conditions) == 0 {
			return false, nil
		}

		for _, condition := range pipelineRun.Status.Conditions {
			switch condition.Reason {
			case "Cancelled":
				return true, fmt.Errorf("pipelineRun \"%v\" was cancelled while running", pipelineRun.Name)
			case "Failed":
				return true, fmt.Errorf("pipelineRun \"%v\" failed with message \"%v\"", pipelineRun.Name, condition.Message)
			case "Running":
				return false, nil
			case "Completed":
				return true, nil
			case "Succeeded":
				return true, nil
			default:
				panic(fmt.Sprintf("unknown status condition while waiting for \"%v\" pipeline run to finish: %v", pipelineRunName, condition.Reason))
			}
		}

		return false, nil
	}

	return support.WaitFor(30*time.Minute, 10*time.Second, callback)
}
