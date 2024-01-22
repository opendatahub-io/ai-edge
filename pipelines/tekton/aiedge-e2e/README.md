# AI Edge End to End Pipline

### Description
End to end pipeline that supports a workflow to Fetch -> Build -> Test -> Push an immutable model container image to an image registry

* Fetch - Fetch a pre-trained [OpenVino compatible model](https://docs.openvino.ai/2023.2/Supported_Model_Formats.html) from S3 or a Git repo
* Build - Package and build the model image and it's dependencies into an immutable container image
* Test - Run a smoke test against a Pod running the immutable container to verify the exposed endpoint is available for inferencing
* Push - Push the immutable container image to an image registry supported

### Prerequisites
* [OpenShift Pipelines](https://docs.openshift.com/pipelines/1.12/install_config/installing-pipelines.html#op-installing-pipelines-operator-in-web-console_installing-pipelines)
* S3 Credentials required to access the pre-trained model
* Pre-trained model stored in S3 or Git that is compatible with [OpenVino Model Server](https://docs.openvino.ai/2023.2/Supported_Model_Formats.html) 
  * Test data that can be used to verify the inferencing endpoint of the model container is working as intended
* Credentials required to push to the destination image registry ([Quay](https://docs.quay.io/glossary/robot-accounts.html))
* OpenShift user with project admin permissions to a [Data Science Project](https://access.redhat.com/documentation/en-us/red_hat_openshift_data_science/1/html-single/getting_started_with_red_hat_openshift_data_science/index#creating-a-data-science-project_get-started) or OpenShift namespace where the Pipeline is running
* [OpenShift CLI](https://docs.openshift.com/container-platform/4.13/cli_reference/openshift_cli/getting-started-cli.html)
* [kustomize](https://kubectl.docs.kubernetes.io/guides/introduction/kustomize/)


#### Quay Repository and Robot Permissions
- In your image registry namespace ([Quay](https://quay.io)):
  - Add a robot account to push images and set write Permissions for the robot account on the repositories. ([Quay](https://access.redhat.com/documentation/en-us/red_hat_quay/3.10/html/use_red_hat_quay/use-quay-manage-repo))
  - Download the Kubernetes Secret of the robot account and store it in a YAML file.
- Inspect the file with the pull secret and note the name of the secret, or edit it.
- Create the secret and apply it. This will allow you to use the Quay Robot kubernetes secret directly as the dockerconfig workspace. E.g.:

```bash
oc apply -f <downloaddir>/rhoai-edge-build-secret.yml
```

#### Data for testing the model inferencing endpoint
To verify that that model container is working successfully, we include a test-model-rest-svc task, that will send data to the model inferencing endpoint and verify that expected output is returned.  You will need to create a ConfigMap with two files `data.json`, the jsondata payload for your model and `output.json`, the expected json output of your model

```
# Example ConfigMap
kind: ConfigMap
apiVersion: v1
metadata:
  name: my-model-test-data

data:
  data.json: |+
    {"dataframe_split": {"columns":[ "day", "mnth", "year", "season","holiday", "weekday", "workingday", "weathersit", "temp", "hum", "windspeed" ], "data":[[ 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11 ]]}}

  output.json: '{"predictions": [331]}'
```

Once the test data configmap is created, it can be included as the `test-data` workspace in the [PipelineRun](aiedge-e2e.pipelinerun.yaml)
```
apiVersion: tekton.dev/v1beta1
kind: PipelineRun
spec:
  ...
  workspaces:
  - configMap:
      name: my-model-test-data
    name: test-data
```

### Deploy the Pipeline
From the user's Data Science Projects namespace where the Pipeline will be running

1. Deploy the Tekton Tasks and Pipeline to the namespace
```
# From the root folder where this README is located
oc apply -k .
```

1. Create a copy of the file(s) below to include the required credentials for accessing any models and image registries required for the Pipeline to run successfully:
   * S3 Storage - [credentials-s3.secret.yaml.template](templates/credentials-s3.secret.yaml.template) to include your credentials required to access any model stored in S3 then apply it to the server
   * Image Registry - [credentials-image-registry.yaml.template](templates/credentials-image-registry.yaml.template) to include the username and password with write access to the image repository

   ```
   $ cp examples/credentials-s3.secret.yaml.template credentials-s3.secret.yaml
   $ cp examples/credentials-image-registry.secret.yaml.template credentials-image-registry.secret.yaml

   # Edit the credentials file
   # Login to the OpenShift cluster and add the credentials it to the server
   $ oc apply -f credentials-s3.secret.yaml -f credentials-image-registry
   ```

### Run the Pipeline

Update the `s3-bucket-name` parameter value from its default `rhoai-edge-models` in
[`aiedge-e2e.pipelinerun.yaml`](pipelines/tekton/aiedge-e2e/aiedge-e2e.pipelinerun.yaml) to match your S3 bucket name.

#### For Git fetch
Update the `git-model-repo` parameter with the repository url, the `modelRelativePath` parameter to the model files path and the `git-revision` parameter for the version/branch of the repository in [`aiedge-e2e.pipelinerun.yaml`](pipelines/tekton/aiedge-e2e/aiedge-e2e.pipelinerun.yaml).

#### View the PipelineRun results

If the PipelineRun completes successfully, you can see the results in the OpenShift Console by going to Pipelines > PipelineRuns > (Select your PipelineRun and scroll down).

You can also click on the "YAML" tab in the PipelineRun and scroll down to the `pipelineResults` section, it will look something like this:

<details><summary>Typical E2E pipeline results</summary>

```yaml
  pipelineResults:
    - name: git-model-fetched-commit
      value: f89e2bc61e6ff7f539c52fded3a2bdc991a55b7c
    - name: git-model-fetched-url
      value: 'https://github.com/piotrpdev/ai-edge.git'
    - name: git-model-fetched-commit-epoch
      value: '1702404808'
    - name: git-containerfile-fetched-commit
      value: f89e2bc61e6ff7f539c52fded3a2bdc991a55b7c
    - name: git-containerfile-fetched-url
      value: 'https://github.com/piotrpdev/ai-edge'
    - name: git-containerfile-fetched-commit-epoch
      value: '1702404808'
    - name: model-files-size
      value: |
        2084
    - name: model-files-list
      value: |
        MLmodel
        README.md
        conda.yaml
        convert_csv_to_json.py
        dataset.csv
        python_env.yaml
        requirements.txt
        tf2model/
    - name: internal-registry-url
      value: >
        image-registry.openshift-image-registry.svc:5000/pplaczek-pipeline-dev/tensorflow-housing:1
    - name: target-registry-url
      value: quay.io/pplaczek/tensorflow-housing
    - name: internal-image-url
      value: >-
        image-registry.openshift-image-registry.svc:5000/pplaczek-pipeline-dev/tensorflow-housing@sha256:0cc9a636c2f18b0f15224a234995ecd27a3dc2e5eb7ffefc8ecfd72c099da31f
    - name: target-image-url
      value: >-
        quay.io/pplaczek/tensorflow-housing:1-fa9e93e0-c66c-4075-9333-61769420f102
    - name: internal-image-size
      value: '126507187'
    - name: buildah-sha
      value: 'sha256:0cc9a636c2f18b0f15224a234995ecd27a3dc2e5eb7ffefc8ecfd72c099da31f'
    - name: model-name
      value: tensorflow-housing
    - name: model-version
      value: '1'
    - name: internal-image-created-at
      value: '2023-12-12T18:39:25Z'
    - name: internal-image-buildah-version
      value: 1.24.2
```

</details>