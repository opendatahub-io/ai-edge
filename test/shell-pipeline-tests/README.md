# Shell Pipeline Tests

This directory contains 2 shell pipeline tests:
* openvino-bike-rentals - OpenVINO version using the bike rentals model
* tensorflow-housing - TensorFlow version using the housing model

Both tests currently run [build-container-image-pipeline](../../pipelines/tekton/build-container-image-pipeline)
and the [test-mlflow-image-pipeline](../../pipelines/tekton/test-mlflow-image-pipeline). The tests will be switched in the near future to the full [aiedge-e2e](../../pipelines/tekton/aiedge-e2e) version of the pipeline.
After that, [GitOps pipeline](../../pipelines/tekton/gitops-update-pipeline) tests will be added as well.

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

After the credentials are configured, you can run the pipeline tests using:

```shell
ARTIFACT_DIR=./artifacts CUSTOM_AWS_SECRET_PATH=./secrets CUSTOM_IMAGE_REGISTRY_SECRET_PATH=./secrets ./openvino-bike-rentals/pipelines-test-openvino-bike-rentals.sh
```
and
```shell
ARTIFACT_DIR=./artifacts CUSTOM_AWS_SECRET_PATH=./secrets CUSTOM_IMAGE_REGISTRY_SECRET_PATH=./secrets ./tensorflow-housing/pipelines-test-tensorflow-housing.sh
```

This would put all the logs into the `$PWD/artifacts` directory and it also expects all the credential files to be stored under the `$PWD/secrets` directory.
