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
- OpenShift user with project admin permissions to a [Data Science Project](https://access.redhat.com/documentation/en-us/red_hat_openshift_data_science/1/html-single/getting_started_with_red_hat_openshift_data_science/index#creating-a-data-science-project_get-started) or OpenShift namespace where the Pipeline is running
- Credentials required to push to the destination image registry ([Quay](https://docs.quay.io/glossary/robot-accounts.html))
- Pre-trained model stored in S3 or Git that is compatible with [OpenVino Model Server](https://docs.openvino.ai/2023.2/Supported_Model_Formats.html) or [Seldon MLServer](https://mlserver.readthedocs.io/en/latest/runtimes/index.html)
  - Test data that can be used to verify the inferencing endpoint of the model container is working as intended
- S3 Credentials required to access the pre-trained model
- [OpenShift CLI](https://docs.openshift.com/container-platform/4.13/cli_reference/openshift_cli/getting-started-cli.html)
- [kustomize](https://kubectl.docs.kubernetes.io/guides/introduction/kustomize/)
- A clone of this repository

## AI Edge End to End Pipeline
[//]: # (Change to S3 Fetch/Git Pipeline once the split is done describing the differences)

> [!IMPORTANT]
> This pipeline is currently being split into 2 pipelines, which will differ just in the fetch stage (S3 or Git), otherwise they are identical.

### Description
End to end pipeline that supports a workflow to Fetch -> Build -> Test -> Push an immutable model container image to an image registry

* Fetch - Fetch a pre-trained [OpenVino](https://docs.openvino.ai/2023.2/Supported_Model_Formats.html) or [Seldon MLServer](https://mlserver.readthedocs.io/en/latest/runtimes/index.html) compatible model from S3 or a Git repo
* Build - Package and build the model image and it's dependencies into an immutable container image
* Test - Run a smoke test against a Pod running the immutable container to verify the exposed endpoint is available for inferencing
* Push - Push the immutable container image to an image registry supported

### Setup S3 bucket

Create an S3 bucket and upload the directories with the models:

![S3 models example](../.github/images/S3-models.png)

If you don't have or know your access key, generate one in AWS account's Security credentials > Access keys.

### Quay Repository and Robot Permissions
- In your image registry namespace ([Quay](https://quay.io)):
  - Add a robot account to push images and set write Permissions for the robot account on the repositories. ([Quay](https://access.redhat.com/documentation/en-us/red_hat_quay/3.10/html/use_red_hat_quay/use-quay-manage-repo))
  - Take note of the username and password of the robot account.
- Another option is to directly download the Kubernetes Secret of the robot account and store it in a YAML file.
  - Inspect the file with the pull secret and note the name of the secret, or edit it.
  - Create the secret and link it to the `pipeline` Service Account that was created by the Red Hat OpenShift Pipelines operator using a Tekton Config. E.g.:

  ```bash
  oc apply -f <downloaddir>/rhoai-edge-build-secret.yml
  oc secret link pipeline rhoai-edge-build-pull-secret
  ```

### Setup credentials
Create a copy of the file(s) below to include the required credentials for accessing any models and image registries required for the Pipeline to run successfully:
* S3 Storage - [credentials-s3.secret.yaml.template](tekton/aiedge-e2e/templates/credentials-s3.secret.yaml.template) to include your credentials required to access any model stored in S3 then apply it to the server
* Image Registry - [credentials-image-registry.secret.yaml.template](tekton/aiedge-e2e/templates/credentials-image-registry.secret.yaml.template) to include the username and password with write access to the image repository.
  This is needed only in case you noted username and password of the robot account. In case you directly downloaded and applied a Kubernetes Secret, this file can be skipped.
    ```bash
    # From the root folder where this README is located
    $ cp tekton/aiedge-e2e/templates/credentials-s3.secret.yaml.template credentials-s3.secret.yaml
    $ cp tekton/aiedge-e2e/templates/credentials-image-registry.secret.yaml.template credentials-image-registry.secret.yaml

    # Edit the credentials files with S3 and Robot credentials
    # Login to the OpenShift cluster and add the credentials to the server
    $ oc apply -f credentials-s3.secret.yaml -f credentials-image-registry.secret.yaml

    # Linking secret is needed only if the secret from the Robot account hasn't been already applied and linked in the previous step
    $ oc secret link pipeline credentials-image-registry
    ```

### Data for testing the model inferencing endpoint
To verify that that model container is working successfully, the pipeline invokes a Task `test-model-rest-svc` which will send data to a testing container with the model inferencing endpoint and verify that expected output is returned.
The Task expects a workspace `test-data` with files `data.json`, the jsondata payload for your model, and `output.json`, the expected json output for that input payload.

The example PipelineRun files ([OpenVino example using Git](tekton/aiedge-e2e/example-pipelineruns/aiedge-e2e.tensorflow-housing.pipelinerun.yaml),
[Seldon example using S3](tekton/aiedge-e2e/example-pipelineruns/s3-fetch.bike-rentals.pipelinerun.yaml)) demonstrate that approach, referencing the ConfigMap defined in
[tensorflow-housing-test-data-cm.yaml](tekton/aiedge-e2e/test-data/tensorflow-housing-test-data-cm.yaml) and
[bike-rentals-test-data-cm.yaml](tekton/aiedge-e2e/test-data/bike-rentals-test-data-cm.yaml), respectively.

If using your models, you will want to adjust these accordingly.

### Deploy the Pipeline
From the user's Data Science Projects namespace where the Pipeline will be running

1. Deploy the Tekton Tasks and Pipelines to the namespace
```bash
# From the repository's root folder
oc apply -k manifests/
```

### Run the Pipeline

Update the `s3-bucket-name` parameter value in your S3 Fetch PipelineRun file to match your S3 bucket name.
In [this example PipelineRun file](tekton/aiedge-e2e/example-pipelineruns/s3-fetch.bike-rentals.pipelinerun.yaml) it's set to a default of `rhoai-edge-models`.

#### For Git fetch

Update the `git-model-repo` parameter with the repository url, the `modelRelativePath` parameter to the model files path and the `git-revision` parameter for the version/branch of the repository in your PipelineRun file.
[This example PipelineRun file](tekton/aiedge-e2e/example-pipelineruns/aiedge-e2e.tensorflow-housing.pipelinerun.yaml) can be used as an example.

#### Other parameters

[//]: # (TODO Revisit after the split is done)
You may also want to change other parameters like:
* `model-name`
* `containerfileRelativePath` - to try a different Containerfile
* `fetch-model` - to switch between S3 and Git
* `test-endpoint` - endpoint of the running model server used for testing the inference
* `target-image-tag-references` - a list of image tag references in image repositories in image registries, that the image should be pushed to

Be sure to also use the correct config map with the test data.

#### Create a new PipelineRun
```bash
# From the root folder where this README is located
oc create -f tekton/aiedge-e2e/example-pipelineruns/s3-fetch.bike-rentals.pipelinerun.yaml
# and/or
oc create -f tekton/aiedge-e2e/example-pipelineruns/aiedge-e2e.tensorflow-housing.pipelinerun.yaml
```

> [!IMPORTANT]
> Since the `build-workspace-pv` workspace is used to share data between TaskRuns in a PipelineRun, a PersistentVolumeClaim type VolumeSource is required to fulfill it properly.
> We strongly recommend that this is fulfilled using the `volumeClaimTemplate` approach, rather than the `persistentVolumeClaim` approach.
> If you must use the `persistentVolumeClaim` approach to re-use an existing PersistentVolumeClaim, then you will likely hit issues if two PipelineRuns for the same model name are executed concurrently (and possibly other corner cases).
> See the Tekton documentation around [Using PersistentVolumeClaims as VolumeSource](https://tekton.dev/docs/pipelines/workspaces/#using-persistentvolumeclaims-as-volumesource).

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
oc describe pipelinerun.tekton.dev/aiedge-e2e-66q8n
# or
oc describe imagestream/tensorflow-housing
```

To get the information about the built container image with the runtime and model
in one of the Image Streams, try
```bash
oc get -o json imagestream/tensorflow-housing | jq -r '.status.tags[0].items[0].dockerImageReference'
```

If you add
```yaml
   - name: upon-end
     value: keep
```
to the PipelineRun's `params`, the applications that got created for testing will not get deleted and you can inspect them in OpenShift Console.
For example in Topology > select the application and its Resources and Pod logs, or with CLI find the names of the objects with

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

Another option is to set the `upon-end` parameter to `stop`, which will scale down the deployment to 0 replicas.

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
- Create a copy of the [`example-git-credentials-secret.yaml.template`](tekton/gitops-update-pipeline/templates/example-git-credentials-secret.yaml.template) and update it with your repository information and credentials.
  For GitHub, the token can be generated at Settings > Developer Settings > Personal access tokens > Fine-grained tokens
  and it should have Read access to metadata and Read and Write access to code and pull requests permissions to the repository you use.

### Deploy and run the GitOps pipeline

Update the `image-registry-repo` and `image-digest` parameters corresponding to the image that the GitOps repository should be updated to.
Then execute the following commands to create the Pipeline and start a PipelineRun.

```bash
# From the repository's root folder
oc apply -k manifests/

# From the root folder where this README is located
cp tekton/gitops-update-pipeline/templates/example-git-credentials-secret.yaml.template example-git-credentials-secret.yaml
# Edit the credentials and add the credentials to the server
oc apply -f example-git-credentials-secret.yaml

oc create -f tekton/gitops-update-pipeline/example-pipelineruns/gitops-update-pipelinerun-tensorflow-housing.yaml
# and/or
oc create -f tekton/gitops-update-pipeline/example-pipelineruns/gitops-update-pipelinerun-bike-rentals.yaml
```

After the Pipeline Run(s) finish, check your git repository -- there should be a pull request with an update of the respective
`acm/odh-edge/apps/*/kustomization.yaml` file with the SHA-256 of the new container image that got built, tested, and pushed to Quay
in previous steps.

#### View the PipelineRun results

If the PipelineRun completes successfully, you can see the results in the OpenShift Console by going to Pipelines > PipelineRuns > (Select your PipelineRun and scroll down).

You can also click on the "YAML" tab in the PipelineRun and scroll down to the `pipelineResults` section, it will look something like this:

```yaml
  pipelineResults:
    - name: pr-url
      value: 'https://github.com/opendatahub-io/ai-edge/pull/5'
```
