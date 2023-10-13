# Open Data Hub / AI Edge Use Cases

This repository contains artifacts to show and support Open Data Hub and Red Hat OpenShift Data Science Edge use cases.

For the purpose of a "near edge" Proof of Concept, the edge environment is represented by separate OpenShift cluster(s),
managed from a core OpenShift cluster using Red Hat Advanced Cluster Management (ACM),
based on the Open Cluster Management project.

| Components                                          | Version        |
|-----------------------------------------------------|----------------|
| OpenShift clusters (at least two) with admin access | 4.12 or higher |
| GitHub account or Gitea installation                | [github.com](https://github.com/) |
| AWS account with access to S3                       | [s3.console.aws.amazon.com](https://s3.console.aws.amazon.com/) |
| Red Hat OpenShift Pipelines                         | 1.11 or higher |
| Quay Registry account                               | [quay.io](https://quay.io/) |
| Advanced Cluster Management for Kubernetes          | 2.8            |
| Open Data Hub (optional)                            | 1.x or 2.x     |

## Proof of Concept Edge use case with ACM

The main objective is to showcase that a user can take a trained model, use a pipeline to package it with all the dependencies into a container image, and deploy it at the near edge location(s) (represented by ACM-managed clusters) in a centralized way.

### Infrastructure Configuration

1. Provision OpenShift Cluster
1. Configure the default Identity Provider
1. Install Red Hat Advanced Cluster Management
1. Register the clusters

   [ACM Application](https://access.redhat.com/documentation/en-us/red_hat_advanced_cluster_management_for_kubernetes/2.8/html/applications/managing-applications) manifests are located in [acm/registration](acm/registration) to register and configure the target environments required for the AI at the Edge use cases.  The files can be applied to the ACM hub cluster manually:
   ```
   $  oc apply -k acm/registration/
   ```
   * Core - Cluster host the ODH Core components that will be used in the MLOps Engineer workflow to train, build and push the model.  This cluster is not required to be co-located with the ACM Hub but we group them together to simplify the use case
   * Near Edge - Cluster(s) that will host the running model at the edge.  This is the target environment after a new model is available for use
1. Deploy Open Data Hub to the Core cluster and register any configurations to support pushing models to the edge cluster
   * Open Data Hub - v2.x of the operator will be installed as part of the dependencies for the core cluster.  Manual installation of the [DataScienceCluster](https://github.com/opendatahub-io/opendatahub-operator#example-datasciencecluster) object will be required to deploy the Open Data Hub core stack to support the data science workflow
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

### Pipelines setup

See [pipelines/README.md](pipelines/README.md)

### Observability setup

* Edge cluster(s)
  * Login to the edge cluster using an account with cluster-admin privileges
  * Enable [monitoring for user-defined projects](https://access.redhat.com/documentation/en-us/openshift_container_platform/4.13/html/monitoring/enabling-monitoring-for-user-defined-projects) in OpenShift clusters
    * `oc -n openshift-monitoring edit configmap cluster-monitoring-config`
    * Set variable `enableUserWorkload` to `true`
  * Edit contents of [thanos-secret](acm/odh-core/acm-observability/secrets/thanos.yaml) file.

### Gitea in cluster git server for GitOps workflow
We are relying on the [gitea-operator](https://github.com/rhpds/gitea-operator) to manage the Gitea server installation in the cluster.  This will simplify the setup of Gitea so that we can create a minimal [Gitea](https://github.com/rhpds/gitea-operator#migrating-repositories-for-created-users) CR to configure and install the Gitea server.  The gitea-operator will be installed on the `odh-core` cluster as part of the ACM application rollout.

1. Wait for the gitea-operator installation to complete and the `gitea.pfe-rhpds.com` CRD is available on the `odh-core` cluster
   ```
   $ oc get crd gitea.pfe.rhpds.com
   NAME                  CREATED AT
   gitea.pfe.rhpds.com   2023-08-25T03:00:13Z
   ```

1. Create the Gitea CustomResource to deploy the server with an admin user
   ```
   cat <<EOF | oc apply -f -
   ---
   apiVersion: v1
   kind: Namespace
   metadata:
     name: gitea
   ---
   apiVersion: pfe.rhpds.com/v1
   kind: Gitea
   metadata:
     name: gitea-ai-edge
     namespace: gitea
   spec:
     # Create the admin user
     giteaAdminUser: admin-edge
     giteaAdminEmail: admin@ai-edge
     giteaAdminPassword: "opendatahub"

     # Create the gitea users accounts to access the cluster
     giteaCreateUsers: true
     giteaGenerateUserFormat: "edge-user-%d"
     giteaUserNumber: 3
     giteaUserPassword: "opendatahub"

     # Populate each gitea user org with a clone of the entries in the giteaRepositoriesList
     giteaMigrateRepositories: true
     giteaRepositoriesList:
     - repo: https://github.com/opendatahub-io/ai-edge.git
       name: ai-edge-gitops
       private: false
    EOF
   ```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## Development

See [DEVELOPMENT.md](DEVELOPMENT.md).
