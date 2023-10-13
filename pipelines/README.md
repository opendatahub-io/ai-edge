# Pipelines Setup

## Models

This repository contains the following trained example MLflow models: [bike-rentals-auto-ml](models/bike-rentals-auto-ml/) and [tensorflow-housing](models/tensorflow-housing/):

```plaintext
bike-rentals-auto-ml/
├── conda.yaml
├── MLmodel
├── model.pkl
├── python_env.yaml
└── requirements.txt

tensorflow-housing/
├── conda.yaml
├── MLmodel
├── model.pkl
├── python_env.yaml
├── requirements.txt
└── tf2model/
    ├── saved_model.pb
    └── ...
```

They are also referenced by these names in the example pipeline YAML files.

## Prerequisites

- OpenShift cluster with [OpenShift Pipelines Operator](https://docs.openshift.com/container-platform/4.13/cicd/pipelines/installing-pipelines.html) installed
- OpenShift project / namespace. E.g.  `oc new-project model-to-edge`
- A repository on [Quay.io](https://quay.io/)
- S3 bucket for storing the models
- A clone of this repository

## Deploy build pipeline for AI runtime container images

### Setup S3 bucket and credentials

Create an S3 bucket and upload the directories with the models:

![S3 models example](../.github/images/S3-models.png)

Fill information about access to your S3 bucket in a copy of [`aws-env.yaml`](tekton/build-container-image-pipeline/aws-env.yaml).
If you don't have or know your access key, generate one in AWS account's Security credentials > Access keys.

Then store the credentials in an OpenShift secret:

```bash
cp tekton/build-container-image-pipeline/aws-env.yaml tekton/build-container-image-pipeline/aws-env-real.yaml
vi tekton/build-container-image-pipeline/aws-env-real.yaml
oc create -f tekton/build-container-image-pipeline/aws-env-real.yaml
```

### Deploy and run the build pipeline

Update the `aws-bucket-name` parameter value from its default `rhoai-edge-models` in
[`build-container-image-pipelinerun-bike-rentals.yaml`](tekton/build-container-image-pipeline/build-container-image-pipelinerun-bike-rentals.yaml)
and/or
[`build-container-image-pipelinerun-tensorflow-housing.yaml`](tekton/build-container-image-pipeline/build-container-image-pipelinerun-tensorflow-housing.yaml)
to match your S3 bucket name.

Then create the pipeline(s) to build the container image with AI runtime:

```bash
oc apply -k tekton/build-container-image-pipeline/
oc create -f tekton/build-container-image-pipeline/build-container-image-pipelinerun-bike-rentals.yaml
# and / or
oc create -f tekton/build-container-image-pipeline/build-container-image-pipelinerun-tensorflow-housing.yaml
```

Check what objects were created and what pipelines executed either in OpenShift Console
in Pipelines > Pipelines, Storage > PersistentVolumeClaims,
and Builds > ImageStreams,
or with CLI find the names of the objects with
```
oc get persistentvolumeclaim
oc get pipeline.tekton.dev
oc get task.tekton.dev
oc get pipelinerun.tekton.dev
oc get imagestream
```
and then run `oc describe` on them, for example
```
oc describe pipelinerun.tekton.dev/build-container-image-bike-rentals-66q8n
# or
oc describe imagestream/tensorflow-housing
```

To get the information about the built container image with the runtime and model
in one of the Image Streams, try
```
oc get -o json imagestream/tensorflow-housing | jq -r '.status.tags[0].items[0].dockerImageReference'
```

In the next steps we will automate testing the behaviour of a container
started from the newly built container image, and push the image to dedicated
container image repository if the test passes.

## Deploy Test MLflow Container image pipeline

### Quay Repository and Robot Permissions

- In your Quay namespace:
  - Create repositories `tensorflow-housing` and/or `bike-rentals-auto-ml`.
  - Add a robot account to push images and set write Permissions for the robot account on the repositories.
  - Download the Kubernetes Secret of the robot account and store it in a YAML file.
- Inspect the file with the pull secret and note the name of the secret, or edit it.
- Create the secret and link it to the `pipeline` Service Account that was created by the Red Hat OpenShift Pipelines operator using a Tekton Config. E.g.:

```bash
oc apply -f <downloaddir>/rhoai-edge-build-secret.yml
oc secret link pipeline rhoai-edge-build-pull-secret
```

Check that the secret is listed in the `Mountable secrets` of
```bash
oc describe sa/pipeline
```

### Deploy and run the test pipeline

Update the `target-imagerepo` parameter value from its default `rhoai-edge` in
[`test-mlflow-image-pipelinerun-tensorflow-housing.yaml`](tekton/test-mlflow-image-pipeline/test-mlflow-image-pipelinerun-tensorflow-housing.yaml)
and/or
[`test-mlflow-image-pipelinerun-bike-rental.yaml`](tekton/test-mlflow-image-pipeline/test-mlflow-image-pipelinerun-bike-rental.yaml)
to match the name of your Quay namespace.

Then create the pipeline(s) to test the container images behaviour using
test data in [`tensorflow-housing-test-data-cm.yaml`](tekton/test-mlflow-image-pipeline/tensorflow-housing-test-data-cm.yaml)
and/or [`bike-rentals-test-data-cm.yaml`](tekton/test-mlflow-image-pipeline/bike-rentals-test-data-cm.yaml):

```bash
oc apply -k tekton/test-mlflow-image-pipeline/
oc create -f tekton/test-mlflow-image-pipeline/test-mlflow-image-pipelinerun-tensorflow-housing.yaml
# and/or
oc create -f tekton/test-mlflow-image-pipeline/test-mlflow-image-pipelinerun-bike-rental.yaml
```

Check what pipeline was created and run either in OpenShift Console in Pipelines > Pipelines,
or using the CLI to find the names of the objects by issuing commands like
```bash
oc get pipeline
oc get pipelinerun
```
and then running `oc describe` on the reported objects.

If you add
```
   - name: upon-end
     value: keep
```
to the PipelineRun's `params`, the applications that got created for testing will not get deleted and you can inspect them in OpenShift Console for example in Topology > select the application and its Resources and Pod logs, or with CLI find the names of the objects with

```bash
oc get deployment
oc get all
```
and then run `oc describe` or `oc logs` on them, for example
```bash
oc describe deployment/bike-rentals-auto-ml-1
# or
oc logs deployment/tensorflow-housing-1
```

If the Pipeline Runs passed, the last Cluster Task `skopeo-copy` copied
the built container image(s) to your Quay repository. Check in Quay's WebUI
or with `podman pull` that the pushed container image has the same SHA-256
checksum as the one shown in the Image Stream.

## Deploy the GitOps pipeline

We have so far worked with the upstream source repo.

In this last pipeline we will use a clone of that repo to show how the identifier of the newly built and tested container image
can be recorded in the repository via a pull request. For that we will need a clone or copy of this repository, in GitHub or
in your Git server.

### Git Repository and Credentials

- Clone/Mirror this repository on your Git server
- Update the `git*` parameters in
  [`gitops-update-pipelinerun-bike-rentals.yaml`](tekton/gitops-update-pipeline/example-pipelineruns/gitops-update-pipelinerun-bike-rentals.yaml)
  and/or
  [`gitops-update-pipelinerun-tensorflow-housing.yaml`](tekton/gitops-update-pipeline/example-pipelineruns/gitops-update-pipelinerun-tensorflow-housing.yaml)
  to match location of your repository and the target branch for the pull request.
  The defaults assume `https://github.com/username/ai-edge-gitops` and `main`.
- Update the [`example-git-credentials-secret.yaml`](tekton/gitops-update-pipeline/example-pipelineruns/example-git-credentials-secret.yaml)
  with your repository information and credentials.
  For GitHub, the token can be generated at Settings > Developer Settings > Personal access tokens > Fine-grained tokens
  and it should have Read access to metadata and Read and Write access to code and pull requests permissions to the repository you use.

### Deploy and run the GitOps pipeline

The `gitops-update-pipeline` will fetch information about the last successfuly built and tested container image for the given model
from the PipelineRun of the above Pipelines, and record information about that image in your git repo.

```bash
oc apply -k tekton/gitops-update-pipeline/
oc apply -f tekton/gitops-update-pipeline/example-pipelineruns/example-git-credentials-secret.yaml
oc create -f tekton/gitops-update-pipeline/example-pipelineruns/gitops-update-pipelinerun-tensorflow-housing.yaml
# and/or
oc create -f tekton/gitops-update-pipeline/example-pipelineruns/gitops-update-pipelinerun-bike-rentals.yaml
```

After the Pipeline Run(s) finish, check your git repository -- there should be a pull request with an update of the respective
`acm/odh-edge/apps/*/kustomization.yaml` file with the SHA-256 of the new container image that got built, tested, and pushed to Quay
in previous steps.
