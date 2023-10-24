# Open Data Hub / AI Edge Use Cases

This repository contains artifacts to show and support Open Data Hub use cases.

For the purpose of a "near edge" Proof of Concept, the edge environment is represented by separate OpenShift cluster(s),
managed from a core OpenShift cluster using Red Hat Advanced Cluster Management (ACM),
based on the Open Cluster Management project.

| Components                                              | Version        |
|---------------------------------------------------------|----------------|
| OpenShift clusters (at least two) with admin access     | 4.12 or higher |
| GitHub account or [Gitea installation](gitea/README.md) | [github.com](https://github.com/) |
| AWS account with access to S3                           | [s3.console.aws.amazon.com](https://s3.console.aws.amazon.com/) |
| Red Hat OpenShift Pipelines                             | 1.11 or higher |
| Quay Registry account                                   | [quay.io](https://quay.io/) |
| Advanced Cluster Management for Kubernetes              | 2.8            |
| Open Data Hub                                           | 2.x            |

## Proof of Concept Edge use case with ACM

The main objective is to showcase that a user can take a trained model,
use pipelines to package it with all the dependencies into a container image,
and deploy it at the near edge location(s) (represented by ACM-managed clusters) in a centralized way.

The goal is to develop & incubate features in Open Data Hub that supports delivering and observing
models running in Edge environments.  As this work continues, Open Data Hub will support the
end-to-end workflow of building, training and pushing models to distributed environments within the 
Open Data Hub user experience.

### Developing and training a model

This step is out of scope of this Proof of Concept repository,
as this repository already contains trained models in the [pipelines/models/](pipelines/models/) directory.

If you wish to develop and train different models,
Jupyter notebooks provided by [Open Data Hub](https://opendatahub.io/) (ODH)
To install ODH, admin privileges to install operators from OpenShift OperatorHub are required.

Working and deploying your own models might require bigger changes
to the definition and configuration of the pipelines and ACM setup below,
so you might want to start with the pre-built models first.

### Building container image(s) with a model and serving runtime

In [pipelines/README.md](pipelines/README.md) we show how to take the trained models,
store them in a S3 bucket,
build container image(s) with the model and serving runtime using OpenShift Pipelines,
push the container image(s) to an image registry accessible by the near edge cluster(s),
and update a clone of this repository with a pull request,
configuring `acm/odh-edge/apps/*/kustomization.yaml` with location and digest (SHA-256) of the built images.

You can skip this step if you do not wish to rebuild the container images.
If you use the default configuration as shown in this git repository,
you will use already built container images from the https://quay.io/organization/rhoai-edge repositories.

### Serving to near edge remote clusters

We assume that you have admin access to an OpenShift cluster where you can install and configure
the Red Hat Advanced Cluster Management operator and which will serve as the central ACM hub.
We also assume you have admin access to other OpenShift cluster(s) that you will import to the ACM hub cluster,
these will play the role of near edge location(s) managed from the central ACM hub.

#### ACM setup

The following setup steps need to be done as an admin user with `cluster-admin` privileges.
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
   in [`acm/registration/near-edge/base/near-edge.yaml`](acm/registration/near-edge/base/near-edge.yaml)
   so update that YAML file if you use a different name.
1. Import the near edge cluster(s).
   In the OpenShift Console this can be done in All Clusters menu > Infrastructure > Clusters > Cluster list tab.
   In the Cluster set popup menu, select the cluster set you created for the near edge clusters.
   If you use the Import mode of "Enter your server URL and API token for the existing cluster",
   obtain the address of the server and the token on the other cluster's OpenShift Console 
   at the top-right corner user menu > Copy login command, for an admin user.

#### Create namespaces for applications

For each of the deployed models-turned-into-application, we will create a separate namespace/project on the ACM hub cluster.
The ACM setup in this repository expects namespaces `tensorflow-housing-app` and `bike-rental-app`.

As a regular user in the OpenShift Console, in the local-cluster > Developer menu, Project popup, use the "Create Project" button;
alternatively, run
```bash
oc new-project tensorflow-housing-app
oc new-project bike-rental-app
```

#### Bind namespaces to the cluster set

The newly created namespaces need to be bound to the cluster set that contains the near edge clusters.

On the ACM hub cluster, in the OpenShift Console in the All Clusters menu > Infrastructure > Clusters > Cluster sets tab,
click the three-dot button at the right of the cluster set, select Edit namespace bindings, and add the namespaces you've created.

This is the last step that needs to be done as an admin user with `cluster-admin` privileges.

#### Deploy the applications to remote clusters

With the remote near edge clusters now in the cluster set and that cluster set bound to the application namespaces,
it is now possible to deploy the applications to remote clusters through
ACM [Channels](https://access.redhat.com/documentation/en-us/red_hat_advanced_cluster_management_for_kubernetes/2.8/html-single/applications/index#channels),
[Subscriptions](https://access.redhat.com/documentation/en-us/red_hat_advanced_cluster_management_for_kubernetes/2.8/html-single/applications/index#subscriptions),
[Placements](https://access.redhat.com/documentation/en-us/red_hat_advanced_cluster_management_for_kubernetes/2.8/html-single/applications/index#placement-rules),
and [Applications](https://access.redhat.com/documentation/en-us/red_hat_advanced_cluster_management_for_kubernetes/2.8/html-single/applications/index#applications).

We will create the objects in the respective namespaces from the `acm/registration/` directory structure.

However, it is important to note that the actual application configuration, stored in `acm/odh-edge/`,
does not come from our local git repository checkout.
Instead, [`acm/registration/near-edge/overlays/tensorflow-housing-app/kustomization.yaml`](acm/registration/near-edge/overlays/tensorflow-housing-app/kustomization.yaml)
and [`acm/registration/near-edge/overlays/bike-rental-app/kustomization.yaml`](acm/registration/near-edge/overlays/bike-rental-app/kustomization.yaml)
contain source URLs of the git repositories that will become configurations of the ACM Channels,
and the applications in those remote git repositories will be deployed.
By default, [github.com/opendatahub-io/ai-edge](https://github.com/opendatahub-io/ai-edge) is configured;
edit the URLs to match your repositories.

Then run as a regular (non-admin) user
```bash
oc apply -k acm/registration/
```

#### View the application deployments

In the OpenShift Console of the ACM hub cluster in All Clusters > Applications, filter for Subscriptions.
You should see `tensorflow-housing-app-application` and `bike-rental-app-application` with Git resources.
Clicking on the Subscription name and then on the Topology tab, you can see what objects were created on the remote cluster(s).

Everything should be shown green. If it is not, click the icon of the faulty object and check the displayed information for debugging clues.

### Observability setup

* Edge cluster(s)
  * Login to the edge cluster using an account with cluster-admin privileges
  * Enable [monitoring for user-defined projects](https://access.redhat.com/documentation/en-us/openshift_container_platform/4.13/html/monitoring/enabling-monitoring-for-user-defined-projects) in OpenShift clusters
    * `oc -n openshift-monitoring edit configmap cluster-monitoring-config`
    * Set variable `enableUserWorkload` to `true`
  * Edit contents of [thanos-secret](acm/odh-core/acm-observability/secrets/thanos.yaml) file.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## Development

See [DEVELOPMENT.md](DEVELOPMENT.md).
