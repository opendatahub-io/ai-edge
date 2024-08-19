# Shell Pipeline Tests

This directory contains 2 shell pipeline tests:
* seldon-bike-rentals - Seldon.io image using the bike rentals model
* openvino-tensorflow-housing - OpenVino image using the Tensorflow housing model

First test runs the [S3 Fetch](../../manifests/pipelines/s3-fetch-pipeline.yaml) pipeline, while the second one
runs both [Git Fetch](../../manifests/pipelines/git-fetch-pipeline.yaml) and [GitOps pipeline](../../manifests/pipelines/gitops-update-pipeline.yaml)

Scripts are primarily run in the OpenShift CI environment, so they make use of
OpenShift CI secrets. You need to configure these if you want to run it locally, see the next section.

## Local execution

For local execution, these environment variables need to be set:

* **ARTIFACT_DIR** - Directory where logs and yaml files from the namespace should be stored for easier debugging.
* **CUSTOM_AWS_SECRET_PATH** - Directory where credentials for the AWS S3 bucket are stored. S3 bucket is used as a source of the AI model. The directory should have 2 files:
  * accessKey - containing the access key, sometimes also called access key ID
  * secretAccessKey - containing the secret access key
* **CUSTOM_IMAGE_REGISTRY_SECRET_PATH** - Directory where credentials for the image repository (e.g. Quay) are stored. This repository is used to publish the image after it is tested. The pipeline uses [basic-auth](https://tekton.dev/docs/pipelines/auth/#configuring-basic-auth-authentication-for-docker) for authentication. The directory should contain the files:
  * username - containing the username of the account used to access the image registry
  * password - containing the password used to access the image registry
* **CUSTOM_GIT_CREDENTIALS_SECRET_PATH** - Directory where the GitHub-compatible forge token is stored. The directory should have 1 file:
  * token - containing the GitHub-compatible forge token, e.g. for GitHub specifically it will have the form of `github_pat_123...`

After the credentials are configured, you can run the pipeline tests using:

```shell
ARTIFACT_DIR=./artifacts CUSTOM_AWS_SECRET_PATH=./secrets CUSTOM_IMAGE_REGISTRY_SECRET_PATH=./secrets CUSTOM_GIT_CREDENTIALS_SECRET_PATH=./secrets ./seldon-bike-rentals/pipelines-test-seldon-bike-rentals.sh
```
and
```shell
ARTIFACT_DIR=./artifacts CUSTOM_AWS_SECRET_PATH=./secrets CUSTOM_IMAGE_REGISTRY_SECRET_PATH=./secrets CUSTOM_GIT_CREDENTIALS_SECRET_PATH=./secrets ./openvino-tensorflow-housing/pipelines-test-openvino-tensorflow-housing.sh
```

This would put all the logs into the `$PWD/artifacts` directory and it also expects all the credential files to be stored under the `$PWD/secrets` directory.

> [!NOTE]
> If you have made changes to Containerfiles or models used in the tests, change the Pipeline Run parameters accordingly, i.e. to fetch these files from your branch.
> This is done automatically if running in the OpenShift CI environment.
