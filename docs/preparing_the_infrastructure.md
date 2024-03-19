# Preparing the Infrastructure

## Prerequisites
The following accounts are required to run the proof of concept

- A GitHub account or [Gitea installation](gitea/README.md).
- An AWS account with access to [Amazon S3](https://s3.console.aws.amazon.com/).
- A [Quay](https://www.quay.io) Registry account.

> [!NOTE]
> The rest of this documentation assumes usage of ACM for cluster and application management, and we highly recommend
> this approach.
> If you choose to use a different approach, you'll find some information on how to deploy the inference container
> application using standalone ArgoCD instances on the edge clusters, in the dedicated
> [Bring Your Own Cluster Management document](./byo-mgmt-gitops.md).

## Components

The following table lists the components used in the proof of concept, their versions, and whether they are used in the
core or near edge clusters.

| Components                                      | Version          | Core Cluster | Near Edge Cluster(s) |
|-------------------------------------------------|:-----------------|:------------:|:--------------------:|
| Red Hat OpenShift clusters                      | 4.12 or higher   |     Yes      |         Yes          |
| Red Hat OpenShift Pipelines                     | 1.11 or higher   |     Yes      |          No          |
| Advanced Cluster Management for Kubernetes      | 2.9 or higher    |     Yes      |          No          |
| OpenShift GitOps                                | 1.10.1 or higher |     Yes      |         Yes          |
| Open Data Hub (optional)                        | 1.x or 2.x       |     Yes      |          No          |


## Installation of OpenShift Clusters

The installation of OpenShift clusters is out of scope of this repository. At least one OpenShift cluster is required to
act as the core cluster and at least one OpenShift cluster is required to act as the near edge cluster.

### Edge Development Environment Using Red Hat OpenShift Local

For development purposes, you can use [Red Hat OpenShift
Local](https://access.redhat.com/documentation/en-us/red_hat_openshift_local/2.32/html/getting_started_guide/index) to
create a local OpenShift cluster and use it as the edge cluster. This way you can test that the management and
deployment of AI/ML inference workloads from the core cluster to the edge cluster works as expected.

To use Red Hat OpenShift Local you need:
- A local machine with at least 18GB of RAM and 6 CPU cores available only for the OpenShift Local cluster.
- A Red Hat OpenShift cluster as the core cluster which can be reached from the OpenShift Local cluster.

  Follow the instructions in the [Getting Started Guide](https://access.redhat.com/documentation/en-us/red_hat_openshift_local/2.32/html/getting_started_guide/index) to install Red Hat OpenShift Local.

> [!IMPORTANT]
> 1. OpenShift Local comes with three presets that represent managed container runtimes: `openshift`, `microshift`, and
>    `podman`. We will use the `openshift` preset to create the local OpenShift cluster.
> 1. Make sure you start the OpenShift Local cluster with at least 18GB of RAM and 6 CPU cores via
>    `crc start --memory 18432 --cpus 6`

Once the OpenShift Local cluster is up and running, you can use it as the edge cluster in the proof of concept and
continue with the installation of the components in the next sections.

## Installation of Red Hat OpenShift Pipelines

Building an inference service container image with a trained model can be done using Red Hat OpenShift Pipelines. Follow
the instructions in [pipelines/README.md](../pipelines/README.md) to install Red Hat OpenShift Pipelines and configure
the core cluster to use it to build the inference service container images.

## Installation of Red Hat Advanced Cluster Management for Kubernetes (ACM)

> [!IMPORTANT]
> The following setup steps need to be done as an admin user with `cluster-admin` privileges.
We will describe how to do them using the OpenShift Console but there may be other ways to achieve the same results.

1. On the ACM hub cluster, install the Advanced Cluster Management for Kubernetes operator.
   In the OpenShift Console of the cluster, find and install the operator
   in Menu > Operators > OperatorHub.
1. Create the MultiClusterHub.
   In the OpenShift Console in Menu > Operators > Installed Operators > Advanced Cluster Management for Kubernetes >
   the MultiClusterHub tab, use the "Create MultiClusterHub" button to create its configuration.
1. Refresh the OpenShift Console page.
1. Create a cluster set for your near edge clusters.
   In the OpenShift Console this can be done in All Clusters menu > Infrastructure > Clusters > Cluster sets tab.
   This repository assumes the name `poc-near-edge` in `clusterSets`
   in [`acm/registration/near-edge/base/near-edge.yaml`](../acm/registration/near-edge/base/near-edge.yaml)
   so update that YAML file if you use a different name.
1. Import the near edge cluster(s).
   In the OpenShift Console this can be done in All Clusters menu > Infrastructure > Clusters > Cluster list tab.
   In the Cluster set popup menu, select the cluster set you created for the near edge clusters.
   If you use the Import mode of "Enter your server URL and API token for the existing cluster",
   obtain the address of the server and the token on the other cluster's OpenShift Console
   at the top-right corner user menu > Copy login command, for an admin user.

> [!IMPORTANT]
> To import the OpenShift Local clusters you must use the "Run import commands manually" import mode.

## Installation of Red Hat OpenShift GitOps

OpenShift GitOps will be used by ACM as the GitOps engine to manage the application, using the Pull Controller. This
will require the OpenShift GitOps operator to be installed both on the ACM hub cluster, and all of the managed clusters
at the edge. After the Red Hat OpenShift GitOps Operator is installed, it automatically sets up a ready-to-use Argo CD
instance that is available in the `openshift-gitops` namespace, and an Argo CD icon is displayed in the console toolbar.
To install the OpenShift GitOps operator, follow the steps below.

1. On the ACM hub cluster, install the OpenShift GitOps operator.
   In the OpenShift Console of the cluster, find and install the operator in Menu > Operators > OperatorHub.
1. On each managed near-edge cluster, install the OpenShift GitOps operator.
   In the OpenShift Console of the cluster, find and install the operator in Menu > Operators > OperatorHub.
   these will play the role of near edge location(s) managed from the central ACM hub.
