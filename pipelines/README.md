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
- OpenShift project / namespace. E.g.  `oc new-project azureml-model-to-edge`
- A repository on [Quay.io](https://quay.io/)
- S3 bucket for storing the models
- A clone of this repository

## Deploy build pipeline for AI runtime container images

### Setup S3 bucket and credentials

Create an S3 bucket and upload the directories with the models:

![S3 models example](../.github/images/S3-models.png)

Fill information about access to your S3 bucket in a copy of [`aws-env.yaml`](tekton/azureml-container-pipeline/aws-env.yaml).
If you don't have or know your access key, generate one in AWS account's Security credentials > Access keys.

Then store the credentials in an OpenShift secret:

```bash
cp tekton/azureml-container-pipeline/aws-env.yaml tekton/azureml-container-pipeline/aws-env-real.yaml
vi tekton/azureml-container-pipeline/aws-env-real.yaml
oc create -f tekton/azureml-container-pipeline/aws-env-real.yaml
```

### Deploy and run the build pipeline

Update the `aws-bucket-name` parameter value from its default `rhoai-edge-models` in
[`azureml-container-pipelinerun-bike-rentals.yaml`](tekton/azureml-container-pipeline/azureml-container-pipelinerun-bike-rentals.yaml)
and/or
[`azureml-container-pipelinerun-tensorflow-housing.yaml`](tekton/azureml-container-pipeline/azureml-container-pipelinerun-tensorflow-housing.yaml)
to match your S3 bucket name.

Then create the pipeline(s) to build the container image with AI runtime:

```bash
oc apply -k tekton/azureml-container-pipeline/
oc create -f tekton/azureml-container-pipeline/azureml-container-pipelinerun-bike-rentals.yaml
# and / or
oc create -f tekton/azureml-container-pipeline/azureml-container-pipelinerun-tensorflow-housing.yaml
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
oc describe pipelinerun.tekton.dev/azureml-container-bike-rentals-66q8n
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

> **NOTE**
> Make sure to change the `target-imagerepo` parameter to match the name of your Quay namespace if using one of the provided `PipelineRun` files.

```bash
oc apply -k tekton/test-mlflow-image-pipeline/
oc create -f tekton/test-mlflow-image-pipeline/test-mlflow-image-pipelinerun-tensorflow-housing.yaml
```

## Deploy the GitOps pipeline

### Git Repository and Credentials

- Clone/Mirror this repository on your Git server
- Change the provided `gitops-git-user-` files to match your Git credentials

### Deploy the GitOps pipeline

```bash
oc apply -k tekton/gitops-update-pipeline/
```

#### Run the GitOps pipeline

The `tekton/gitops-update-pipeline/example-pipelineruns/` contains some examples that can be modified and used.
In these examples, notice that there is a template Secret file for the Git credentials that are referenced by different tasks.
Create an equivalent Secret with appropriate details for your environment, and change the parameter values in the PipelineRun definition to match.

``` bash
# Bike rentals app
cp pipelines/tekton/gitops-update-pipeline/example-pipelineruns/example-git-credentials-secret.yaml /tmp/gitea-edge-user-1-secret.yaml
$EDITOR /tmp/gitea-edge-user-1-secret.yaml # make changes for your specific environment
oc apply -f /tmp/gitea-edge-user-1-secret.yaml

oc create -f tekton/gitops-update-pipeline/example-pipelineruns/gitops-update-pipelinerun-bike-rentals.yaml

# Tensorflow housing app
cp pipelines/tekton/gitops-update-pipeline/example-pipelineruns/example-git-credentials-secret.yaml /tmp/gitea-edge-user-2-secret.yaml
$EDITOR /tmp/gitea-edge-user-2-secret.yaml # make changes for your specific environment
oc apply -f /tmp/gitea-edge-user-2-secret.yaml

oc create -f tekton/gitops-update-pipeline/example-pipelineruns/gitops-update-pipelinerun-tensorflow-housing.yaml
```
