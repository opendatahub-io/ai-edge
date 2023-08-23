# ODH AI Edge Use Cases
Artifacts in support of ODH Edge use cases that integration with Red Hat Advanced Cluster Management(Open Cluster Management)

| Components                           | Version |
|--------------------------------------|---------|
| OpenShift                            | 4.13    |
| Open Data Hub                        | 2.x     |
| Red Hat Advanced Cluster Management  | 2.8     |
| OpenShift Pipelines                  |         |
| Quay Registry                        | 2.8     |


## Proof of Concept Edge use case with ACM

The main objective is to showcase that a user can take a trained model, use a pipeline to package it with all the dependencies and deploy it at the near edge location(s) in a centralized way.

### Infrastructure Configuration
1. Provision OpenShift Cluster
1. Configure the default Identity Provider
1. Install Red Hat Advanced Cluster Management
1. Register the clusters
   * Core - Cluster host the ODH Core components that will be used in the MLOps Engineer workflow to train, build and push the model
   * Near Edge - Cluster that will host the running model at the edge.  This is the target environment after a new model is available for use
1. Deploy Open Data Hub to the Core cluster and register any configurations to support pushing models to the edge cluster
   * GitOps repos
   * Image repos
1. Manage the edge cluster environments to support deployment of models and support for monitoring
   * Configure ACM Observability
   * Deploy the Model container

### MLOps Engineer workflows
1. Develop the model in an ODH Jupyter notebook
1. Build the model from the notebook using Data Science Pipelines
1. Push the model to the image registry accessible by the near edge cluster(s)
1. Update the GitOps config for the near edge cluster

### Observability setup

* Core cluster
   *  Login to the core cluster and run `make install/observability-core` to setup acm-observability on the core cluster.
* Edge cluster(s)
   * Login to edge cluster
   * Enable userWorkloadMonitoring
      * `oc edit cm cluster-monitoring-config`
      * Set variable `enableUserWorkload` to `true`
   *  Run `make install/observability-edge` to create the configmap required for metric whitelisting.
