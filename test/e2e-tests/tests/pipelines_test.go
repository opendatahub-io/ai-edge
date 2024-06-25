package tests

import (
	"fmt"
	"github.com/opendatahub-io/ai-edge/test/e2e-tests/support"
	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
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
	TestStartTime time.Time
)

// init is called before any test is run, and it is called once.
// This is used to build then apply the kustomize config for each pipeline
func init() {

	TestStartTime = time.Now() // used to track when pipeline runs were created

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

	createdRun, err := config.Clients.PipelineRun.Create(ctx, &pipelineRun, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = WaitForAllPipelineRunsToComplete(ctx, []string{createdRun.Name}, config)
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
	support.SetPipelineRunParam("git-containerfile-repo", support.NewStringParamValue(config.GitFetchConfig.ContainerFileRepo), &pipelineRun)
	support.SetPipelineRunParam("git-containerfile-revision", support.NewStringParamValue(config.GitFetchConfig.ContainerFileRevision), &pipelineRun)
	support.SetPipelineRunParam("containerRelativePath", support.NewStringParamValue(config.GitFetchConfig.ContainerRelativePath), &pipelineRun)
	support.SetPipelineRunParam("git-model-repo", support.NewStringParamValue(config.GitFetchConfig.ModelRepo), &pipelineRun)
	support.SetPipelineRunParam("modelRelativePath", support.NewStringParamValue(config.GitFetchConfig.ModelRelativePath), &pipelineRun)
	support.SetPipelineRunParam("git-model-revision", support.NewStringParamValue(config.GitFetchConfig.ModelRevision), &pipelineRun)
	support.SetPipelineRunParam("model-dir", support.NewStringParamValue(config.GitFetchConfig.ModelDir), &pipelineRun)

	createdRun, err := config.Clients.PipelineRun.Create(ctx, &pipelineRun, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = WaitForAllPipelineRunsToComplete(ctx, []string{createdRun.Name}, config)
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

	if !config.GitOpsConfig.Enabled {
		t.Skipf("skipping %v, Gitops pipeline not configured", t.Name())
	}

	if !config.GitFetchConfig.Enabled && !config.S3FetchConfig.Enabled {
		t.Skipf("skipping %v, neither git_fetch or s3_fetch was not enabled", t.Name())
	}

	clients, err := support.CreateClients(config.Namespace)
	if err != nil {
		t.Fatal(err.Error())
	}

	gitURL, err := support.ParseGitURL(config.GitOpsConfig.Repo)
	if err != nil {
		t.Fatal(err.Error())
	}

	pipelineRunFilesAndLabels := []struct {
		pipelineRun string
		label       string
	}{
		{pipelineRun: GitOpsUpdateTensorflowHousingPipelineRunRelativePath, label: "model-name=tensorflow-housing"},
		{pipelineRun: GitOpsUpdateBikeRentalsPipelineRunRelativePath, label: "model-name=bike-rentals-auto-ml"},
	}

	var createPipelineRuns []string

	for _, p := range pipelineRunFilesAndLabels {
		pipelineRun, err := support.ReadFileAsPipelineRun(p.pipelineRun)
		if err != nil {
			t.Fatal(err.Error())
		}

		support.SetPipelineRunParam("gitServer", support.NewStringParamValue(gitURL.Server), &pipelineRun)
		support.SetPipelineRunParam("gitApiServer", support.NewStringParamValue(config.GitOpsConfig.ApiServer), &pipelineRun)
		support.SetPipelineRunParam("gitOrgName", support.NewStringParamValue(gitURL.OrgName), &pipelineRun)
		support.SetPipelineRunParam("gitRepoName", support.NewStringParamValue(gitURL.RepoName), &pipelineRun)
		support.SetPipelineRunParam("gitRepoBranchBase", support.NewStringParamValue(config.GitOpsConfig.Branch), &pipelineRun)

		// we need to get the results from the pipeline run created by the ml ops pipeline, we also need
		// to get the correct one that is for this model. Using the correct label and ensuring it was
		// created after the tests have started then we can ensure we are getting the correct one
		completedPipelineRuns, err := config.Clients.PipelineRun.List(ctx, metav1.ListOptions{
			LabelSelector: p.label,
		})
		if err != nil {
			t.Fatal(err.Error())
		}

		var targetPipelineRun *v1.PipelineRun
		for i := 0; i < len(completedPipelineRuns.Items); i += 1 {
			if completedPipelineRuns.Items[i].CreationTimestamp.After(TestStartTime) {
				targetPipelineRun = &completedPipelineRuns.Items[i]
			}
		}

		if targetPipelineRun == nil {
			t.Fatal(fmt.Errorf("no pipeline run found with label %v that was created during testing", p.label))
		}

		imageDigest, err := support.GetResultValueFromPipelineRun("buildah-sha", targetPipelineRun)
		if err != nil {
			t.Fatal(err.Error())
		}

		imageTagReferences, err := support.GetResultValueFromPipelineRun("target-image-tag-references", targetPipelineRun)
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

		createPipelineRuns = append(createPipelineRuns, createdRun.Name)
	}

	err = WaitForAllPipelineRunsToComplete(ctx, createPipelineRuns, config)
	if err != nil {
		t.Fatal(err.Error())
	}
}

func CreateContext() context.Context {
	return context.Background()
}

func WaitForAllPipelineRunsToComplete(ctx context.Context, pipelineRunNames []string, config *support.Config) error {
	callback := func() (bool, error) {
		var pipelineRuns []v1.PipelineRun

		// need to list each one as you can't use FieldSelector wtih an or operator
		// ideally you could do.. metadata.name=X||metadata.name=Y
		for _, name := range pipelineRunNames {
			list, err := config.Clients.PipelineRun.List(ctx, metav1.ListOptions{
				Limit:         1,
				FieldSelector: fmt.Sprintf("metadata.name=%v", name),
			})

			if err != nil {
				return false, err
			}

			if len(list.Items) != 1 {
				return false, fmt.Errorf("pipelineRun \"%v\" could not be found when trying to get it's status", name)
			}

			pipelineRuns = append(pipelineRuns, list.Items[0])
		}

		pipelineStillRunning := false

		// checking the status of each run:
		// any of failed or cancelled -> return with error
		// some have finished but not all -> continue
		// all of finished -> return with no error
		for _, pipelineRun := range pipelineRuns {
			// conditions may be empty when pipeline run is first created
			// this means it is counted as still running
			if len(pipelineRun.Status.Conditions) == 0 {
				pipelineStillRunning = true
				break
			}

			for _, condition := range pipelineRun.Status.Conditions {
				switch condition.Reason {
				case "Cancelled":
					return true, fmt.Errorf("pipelineRun \"%v\" was cancelled while running", pipelineRun.Name)
				case "Failed":
					return true, fmt.Errorf("pipelineRun \"%v\" failed with message \"%v\"", pipelineRun.Name, condition.Message)
				case "Running":
					// not returning here allows for checking all existing runs to see if they have failed
					pipelineStillRunning = true
					break
				case "Completed":
					return true, nil

				}
			}
		}

		return !pipelineStillRunning, nil
	}

	return support.WaitFor(30*time.Minute, 10*time.Second, callback)
}
