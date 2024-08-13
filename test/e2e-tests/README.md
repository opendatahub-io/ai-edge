# AI Edge testing

Place for testing and to verify the MLOps pipelines are working correctly. These tests use the `k8s.io/client-go` package to interact with the cluster. Using `oc login` to log into the cluster as normal should mean little setup is needed to run the tests.

A local install of the Go compiler is needed to run the tests. Go version `1.21` is required to run the tests.

## Setup
- Log into the target cluster using `oc login`. This will update your default `kubeconfig` for the tests to use

## Run tests locally

The e2e-tests use a `config.json` to read values passed to it. in the `e2e-tests` directory copy the `template.config.json` to `config.json`. You can now fill in the fields in the `config.json`.

```bash
cp template.config.json config.json
```

The structure of the `config.json` is in four sections, the top level fields, `git_fetch`, `s3_fetch` and `gitops`. All fields at the top level are required.

- `namespace` - Cluster namespace that tests are run in
- `image_registry_username` - quay.io username
- `image_registry_password` - quay.io password
- `target_image_tags` - JSON array of image tags that the final image will be pushed to. E.g. '["quay.io/user/model-name:e2e-test"]'
- `git_container_file_repo` - Git repo containing the container file
- `git_container_file_revision` - Git branch in the container file repo
- `container_relative_path` - Relative path from the root of the container file repo to where the container file is

After the top level fields each sub object is used for a type of test. Setting `enabled` to `true` in each of these will tell the test suite to use those values in that object and to run those tests.

These are all the fields in `git_fetch`

- `model_repo` - Git repo of the model
- `model_relative_path` - Relative path from the root of the model repo to where the model is 
- `model_revision` - Branch of the model repo
- `model_dir` - Sub-directory of the model in the model folder 
- `username` - (optional) Used for when git repo is private. This is the username associated with the private repo, when set the `token` field must also be set
- `token` - (optional) Used for when git repo is private. This is the token associated with the user who is the owner of the private repo, when set the `username` field must also be set, [see info here](../../pipelines/README.md#git-repository-and-credentials)
- `self_signed_cert` - (optional) path to a self signed cert to connect to access the repo

These are all the fields in `s3_fetch`

- `aws_secret` - AWS secret key
- `aws_access` - AWS access key
- `region` - AWS region of the bucket used
- `endpoint` - Endpoint of the bucket used
- `bucket_name` - Name of the bucket 
- `self_signed_cert` - (optional) path to a self signed cert to connect to access the bucket

These are all the fields in `gitops`

- `token` - Auth token used by the git ops pipeline to make a pull request, [see info here](../../pipelines/README.md#git-repository-and-credentials)
- `username` - Username linked to the `GIT_TOKEN`
- `repo` - Git repo URL used to make a pull request in the git ops pipeline (https://github.com/org/repo)
- `api_server` - Git API server (api.github.com)
- `branch` - Base branch used for pull request in git ops pipeline

Now run the `e2e-tests` with:

```bash
make go-test
```
Set the go binary used for testing that is in your `PATH`
```bash
make GO=go1.19 go-test
```

If you want to re-run the tests in the same namespace you can just re-run the `go-test` target. The tests fail if there are **any** failed pipeline runs. Therefore if you have a failed run and want to re-test make sure to delete any that have failed.

## CI/CD with Github Actions
Not yet implemented
