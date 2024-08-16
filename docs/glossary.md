# AI Edge Terminology

## Core

- The central OpenShift cluster containing the tools responsible for creating any artifacts required for the successful deployment of an Inference Application to Near Edge environments.
- There are no resources or network constraints expected in the core cluster as it is expected that it fully supports all workflows required for creating and verifying Inference Application container images.

## Near Edge

- This is a non-core distributed environment to run and serve AI/ML inference workloads in moderate yet constrained compute resources and network.
- For the purpose of this repository, the near edge environment is represented by separate OpenShift clusters disconnected from the core, the internet or both but may be managed from a core OpenShift cluster.

## Model Server

- A Model Server is responsible for hosting models as a service to "to return predictions based on data inputs that you provide through API calls."[^2]
- For any workflows under opendatahub-io/ai-edge, we will be focusing on using the Model Servers and serving runtimes supported by Open Data Hub

## Inference Application Container

- OCI compliant container image[^3] with the models included during the build process
- Support for container images where the model and model serving runtimes are stored together

## Model Registry

- A centralized repository for the models and their metadata and managing the model lifecycle and versions. Currently our pipelines only support S3 and Git as the source where models can be stored. 

## OCI Distribution Registry

- Open Container Initiative (OCI) compliant container registry where the model and other artifacts are stored and versioned ready to be deployed on production or staging environments.

## GitOps

- GitOps is an established configuration management pattern to store the configuration of your infrastructure configuration and workflow automation in a Git repository for reproducibility and version control.
- "GitOps uses Git repositories as a single source of truth to deliver infrastructure as code."[^1]

[^1]: [What is GitOps](https://www.redhat.com/en/topics/devops/what-is-gitops)
[^2]: [Red Hat OpenShift AI -> Serving models](https://access.redhat.com/documentation/en-us/red_hat_openshift_ai_self-managed/2.8/html/serving_models/about-model-serving_about-model-serving)
[^3]: [Open Container Initiative](https://opencontainers.org/)
