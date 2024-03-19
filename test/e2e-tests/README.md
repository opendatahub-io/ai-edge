# AI Edge testing

Place for testing and to verify the MLOps pipelines are working correctly. These tests use the `k8s.io/client-go` package to interact with the cluster. Using `oc login` to log into the cluster as normal should mean little setup is needed to run the tests.

A local install of the Go compiler is needed to run the tests. Go version `1.21` is required to run the tests.

## Setup
- Log into the target cluster using `oc login`. This will update your default `kubeconfig` for the tests to use
- Create a S3 bucket with the models in the root directory

The following enviroment varaibles are required to run the test setup and the tests themselves. If any of these are not set then the tests will not run. Read [here](../../pipelines/README.md#ai-edge-end-to-end-pipeline) for more context on how to set these up.

- `AWS_SECRET_ACCESS_KEY` - Secret from AWS
- `AWS_ACCESS_KEY_ID` - Access key from AWS
- `S3_REGION` - Region of the bucket used to store the model
- `S3_ENDPOINT` - Endpint of the bucket
- `IMAGE_REGISTRY_USERNAME` - quay.io username
- `IMAGE_REGISTRY_PASSWORD` - quay.io password

The following enviroment varaibles are optional. You may still need to set them for the tests to pass depending on your setup. Read [here](../../pipelines/README.md#ai-edge-end-to-end-pipeline) for more context on how to set these up.
- `SELF_SIGNED_CERT` - Path self signed cert to be used in pipeline


```bash
make go-test-setup
```

## Run tests locally

The following enviroment varaibles are required to run the tests. If any of these are not set then the tests will not run. Read [here](../../pipelines/README.md#ai-edge-end-to-end-pipeline) for more context on how to set these up.

- `S3_BUCKET`	- Name of S3 bucket that contains the models
- `NAMESPACE`	- Cluster namespace that tests are run in
- `TARGET_IMAGE_TAGS_JSON`	- JSON array of image tags that the final image will be pushed to. E.g. '["quay.io/user/model-name:e2e-test"]'

The following enviroment varaibles are optional. You may still need to set them for the tests to pass depending on your setup. Read [here](../../pipelines/README.md#ai-edge-end-to-end-pipeline) for more context on how to set these up.
- `SELF_SIGNED_CERT` - Path self signed cert to be used in pipeline
- `GO` - Custom Go installation path that is not used in `PATH`

```bash
make go-test
```
Set the go binary used for testing that is in your `PATH`
```bash
make go-test
```

## CI/CD with Github Actions
Not yet implemented
