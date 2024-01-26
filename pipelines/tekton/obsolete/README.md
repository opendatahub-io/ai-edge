# Pipelines Setup

## Models

This repository contains the following trained examples.

- [bike-rentals-auto-ml](models/bike-rentals-auto-ml/) is using MLFlow format and can run in [Seldon MLServer](https://github.com/SeldonIO/MLServer).
- [tensorflow-housing](models/tensorflow-housing/) is using MLFlow format and wraps a TensorFlow model. It can run in [Seldon MLServer](https://github.com/SeldonIO/MLServer), but can also run in [OVMS](https://github.com/openvinotoolkit/model_server) by loading the [tf2model](models/tensorflow-housing/tf2model) artifacts.
- [MNIST](models/onnx-mnist) is using ONNX format that can run on [OVMS](https://github.com/openvinotoolkit/model_server).
- [Face Detection](models/tensorflow-facedetection) is using OpenVino IR propietary format and would run only on  [OVMS](https://github.com/openvinotoolkit/model_server).
- [Iris](models/lightgbm-iris) is using Booster format which can run on [Seldon MLServer](https://github.com/SeldonIO/MLServer).
- [Mushrooms](models/lightgbm-mushrooms) is using Booster format which can run on [Seldon MLServer](https://github.com/SeldonIO/MLServer).

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
    
onnx-mnist/
├── 1
│   ├── mnist.onnx
│   └── schema
│       └── schema.json
└── README.md

tensorflow-facedetection/
├── 1
│   ├── face-detection-retail-0004.bin
│   └── face-detection-retail-0004.xml
└── README.md

lightgbm-iris/
├── iris-lightgbm.bst
├── model-settings.json
├── README.md
└── settings.json

lightgbm-mushrooms/
├── model-settings.json
├── mushroom-lightgbm.bst
├── README.md
└── settings.json

```

## Prerequisites

- OpenShift cluster with [OpenShift Pipelines Operator](https://docs.openshift.com/container-platform/4.13/cicd/pipelines/installing-pipelines.html) installed
  - To install the operator, you need to log in as an admin user with `cluster-admins` privileges.
  - All the following steps can and should be done as a regular user with no extra privileges.
- OpenShift project / namespace. E.g.  `oc new-project model-to-edge`
- A repository on [Quay.io](https://quay.io/)
- S3 bucket for storing the models or Git repository storing the model files
- A clone of this repository

## Deploy build pipeline for AI runtime container images

### Setup S3 bucket and credentials

Create an S3 bucket and upload the directories with the models:

![S3 models example](../../../.github/images/S3-models.png)

Fill information about access to your S3 bucket in a copy of [`aws-env.yaml`](build-container-image-pipeline/aws-env.yaml).
If you don't have or know your access key, generate one in AWS account's Security credentials > Access keys.

Then store the credentials in an OpenShift secret:

```bash
cp build-container-image-pipeline/aws-env.yaml build-container-image-pipeline/aws-env-real.yaml
vi build-container-image-pipeline/aws-env-real.yaml
oc create -f build-container-image-pipeline/aws-env-real.yaml
```

### Deploy and run the build pipeline

#### For S3 fetch
Update the `aws-bucket-name` parameter value from its default `rhoai-edge-models` in
[`build-container-image-pipelinerun-bike-rentals.yaml`](build-container-image-pipeline/build-container-image-pipelinerun-bike-rentals.yaml)
and/or
[`build-container-image-pipelinerun-tensorflow-housing.yaml`](build-container-image-pipeline/build-container-image-pipelinerun-tensorflow-housing.yaml)
to match your S3 bucket name.

#### For Git fetch
Update the `git-model-repo` parameter with the repository url, the `modelRelativePath` parameter to the model files path and the `git-revision` parameter for the version/branch of the repository in [`build-container-image-pipelinerun-bike-rentals.yaml`](build-container-image-pipeline/build-container-image-pipelinerun-bike-rentals.yaml).

Then create the pipeline(s) to build the container image with AI runtime:

```bash
oc apply -k build-container-image-pipeline/
oc create -f build-container-image-pipeline/build-container-image-pipelinerun-bike-rentals.yaml
# and / or
oc create -f build-container-image-pipeline/build-container-image-pipelinerun-tensorflow-housing.yaml
```

Check what objects were created and what pipelines executed either in OpenShift Console
in Pipelines > Pipelines, Storage > PersistentVolumeClaims,
and Builds > ImageStreams,
or with CLI find the names of the objects with
```bash
oc get persistentvolumeclaim
oc get pipeline.tekton.dev
oc get task.tekton.dev
oc get pipelinerun.tekton.dev
oc get imagestream
```
and then run `oc describe` on them, for example
```bash
oc describe pipelinerun.tekton.dev/build-container-image-bike-rentals-66q8n
# or
oc describe imagestream/tensorflow-housing
```

To get the information about the built container image with the runtime and model
in one of the Image Streams, try
```bash
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
[`test-mlflow-image-pipelinerun-tensorflow-housing.yaml`](test-mlflow-image-pipeline/test-mlflow-image-pipelinerun-tensorflow-housing.yaml)
and/or
[`test-mlflow-image-pipelinerun-bike-rental.yaml`](test-mlflow-image-pipeline/test-mlflow-image-pipelinerun-bike-rental.yaml)
to match the name of your Quay namespace.

Then create the pipeline(s) to test the container images behaviour using
test data in [`tensorflow-housing-test-data-cm.yaml`](test-mlflow-image-pipeline/tensorflow-housing-test-data-cm.yaml)
and/or [`bike-rentals-test-data-cm.yaml`](test-mlflow-image-pipeline/bike-rentals-test-data-cm.yaml):

```bash
oc apply -k test-mlflow-image-pipeline/
oc create -f test-mlflow-image-pipeline/test-mlflow-image-pipelinerun-tensorflow-housing.yaml
# and/or
oc create -f test-mlflow-image-pipeline/test-mlflow-image-pipelinerun-bike-rental.yaml
```

Check what pipeline was created and run either in OpenShift Console in Pipelines > Pipelines,
or using the CLI to find the names of the objects by issuing commands like
```bash
oc get pipeline
oc get pipelinerun
```
and then running `oc describe` on the reported objects.

If you add
```yaml
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

Continue with the GitOps pipeline in the [original documentation](../../README.md#deploy-and-run-the-gitops-pipeline).