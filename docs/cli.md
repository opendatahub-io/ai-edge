# ODH Edge CLI

The ODH Edge CLI is a command line interface that allows you to interact with the ODH Edge platform. It is a tool that
allows you to manage your ODH Edge models and images.

## Overview

The ODH Edge CLI uses the Model Registry as the source of truth for all models and images. The Model Registry is a
central repository that stores all the metadata about the models and images that are available on the ODH Edge platform.
The ODH Edge CLI allows you to interact with the Model Registry to perform various operations such as listing and adding
models.

Since building a model image is done via OpenShift Pipelines, all the pipeline parameters that are required to build the
model image can be stored in the Model Registry. The ODH Edge CLI allows you to interact with the Model Registry to
store and retrieve these pipeline parameters.

> [!NOTE]
> The Model Registry stores the metadata about the models and images, but it does not store the actual model artifacts.
> The model artifacts are stored in a separate storage location that is accessible to the ODH Edge platform such as an
> S3 bucket or a Git repository.


## Building ODH Edge CLI

To build the ODH Edge CLI, run the following commands from the root of the repository: 

```bash
cd cli
make cli-build
```

This will create a binary called `odh` in your current directory. You can move this binary to a location in your PATH to
make it easier to run the ODH Edge CLI from anywhere.

## Environment Setup

### Prerequisites

- Kubectl access to the OpenShift core cluster.
- Go (version 1.20 or 1.21, version 1.22 is not supported).
- [Model registry operator](https://github.com/opendatahub-io/model-registry-operator/releases/tag/v0.1.2) version 0.1.2 locally.

### Installing Model Registry

To install the Model Registry, follow the following steps:

1. Change to the `model-registry-operator-0.1.2` directory:

    ```bash
    cd model-registry-operator-0.1.2
    ```
2. Install the CRDs into the cluster:

    ```bash
    make install
    ```
3. Deploy the Model Registry operator:

    ```bash
    make deploy
    ```
4. Create a Model Registry namespace:

    ```bash
   oc create ns odh-model-registry
    ```
5. Edit the `config/samples/mysql/modelregistry_v1alpha1_modelregistry.yaml` and change the `spec.rest.serviceRoute` to `enabled`.

6. Create a Model Registry instance:

    ```bash
   oc -n odh-model-registry apply -k config/samples/mysql
    ```

7. Wait for the Model Registry pods to be ready:

    ```bash
    oc -n odh-model-registry get pods -w
    ```

8. Create a model-registry route:

    ```bash
    oc -n odh-model-registry expose service modelregistry-sample --port http-api
    ```
9. Get the route:

    ```bash
    oc -n odh-model-registry get route
    ```

## Usage

The following flags are available globally for all commands:

- `-m, --model-registry-url`: The URL of the Model Registry. This is the URL of the Model Registry service that is
  running in the OpenShift cluster.
- `k, --kubeconfig` : The path to the kubeconfig file. This is the kubeconfig file that is used to access the OpenShift 
  cluster.

### Listing Models

To list all the models that are available in the Model Registry, run the following command:

```bash
odh -m <model-registry-url> models
```

### Adding a Model

Adding a model to the model registry means adding the metadata about the model to the Model Registry. This metadata
includes the model name, a model version and the pipeline parameters that are required to build the model image.

#### Preparing The Pipeline Parameters

The parameters can be provided via a yaml file. The yaml file should have the following structure:

```yaml
params:
  - name: <parameter-name>
    value: <parameter-value>
  - name: <parameter-name>
    value: <parameter-value>
```

This repository provides an example pipeline parameters file in the `cli/examples/params.yaml` file. You can use this
file as a template to create your own pipeline parameters file.

After you have created the pipeline parameters file, you can add the model to the Model Registry by running the
following command:

```bash
odh models add -m model-name -d model-description [-v model-version] [-m model-registry-url] [-p parameters-file]
```

This should print out the message `Model added successfully` if the model was added successfully.

To verify that the model was added successfully, you can list all the models in the Model Registry by running the
following command:

```bash
odh models [-m model-registry-url]
```

### Listing Images

> [!NOTE]
> A model image is represented by a registered model and a specific version of the model. The model image is the actual
> image that is built from the model and is used to deploy the model on the ODH Edge platform.
> The model image is stored in a container registry such as Quay.io but only the metadata about the model image is stored
> in the Model Registry.

To list all the model images that are available in the Model Registry, run the following command:

```bash
odh images [-m model-registry-url]
```

### Updating Build Parameters

To update the build parameters for a model version, you can run the following command:

```bash
odh images update -i model-id -v version [-m model-registry-url] [-p params-file]"
```

This command will update the build parameters for the specified model version. The parameters can be provided via a yaml
file. For more information on the parameters file, see the section on adding a model.

### Building a Model Image

To build a model image, the CLI uses the OpenShift Pipelines to build the model image. The pipeline parameters that are
required to build the model image are stored in the Model Registry. The ODH Edge CLI allows you to create a pipeline run
using the pipeline parameters that are stored in the Model Registry. To build a model image, run the following command:

```bash
odh images build -i model-id -v version [-m model-registry-url] [-n namespace] [-k kubeconfig]
```

### Viewing a Model Image details

To view a model image's details, you can run describe command which will show you image details along with a list of parameters. The command is as follows:

```bash
odh images describe -i model-id -v version [-m model-registry-url] [-p params-file]
```
