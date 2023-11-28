# AI Edge testing

Place for testing and to verify the MLOps pipelines are working correctly. These tests use the `k8s.io/client-go` package to interact with the cluster. Using `oc login` to log into the cluster as normal should mean no setup is needed to run the tests.

A locally install of the Go compiler is needed to run the tests. Go version `1.21` is required to run the tests.

## Run tests locally
```bash
make test
```
Set the go binary used for testing that is in your `PATH`
```bash
make GO=go1.19 test
```

## CI/CD with Github Actions
Not yet implemented
