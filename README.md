# Open Data Hub / AI Edge Use Cases
This should skip the PR check.

This repository contains artifacts to show and support Open Data Hub and Red Hat OpenShift Data Science Edge use cases.

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
| OpenShift GitOps                                        | 1.10.1 or higher |
| Open Data Hub (optional)                                | 1.x or 2.x     |

## Proof of Concept Edge use case with ACM

The main objective is to showcase that a user can take a trained model,
use pipelines to package it with all the dependencies into a container image,
and deploy it at the near edge location(s) (represented by ACM-managed clusters) in a centralized way.

### Developing and training a model

This step is out of scope of this Proof of Concept repository,
as this repository already contains trained models in the [pipelines/models/](pipelines/models/) directory.

If you wish to develop and train different models,
Jupyter notebooks provided by [Open Data Hub](https://opendatahub.io/) (ODH)
or [Red Hat OpenShift Data Science](https://www.redhat.com/en/technologies/cloud-computing/openshift/openshift-data-science) (RHODS) can be used.
To install ODH or RHODS operators, admin privileges in the OpenShift cluster are needed.

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

#### OpenShift GitOps setup

OpenShift GitOps will be used by ACM as the GitOps engine to manage the application, using the Pull Controller.
This will require the OpenShift GitOps operator to be installed both on the ACM hub cluster, and all of the managed clusters at the edge.

1. On the ACM hub cluster, install the OpenShift GitOps operator.
   In the OpenShift Console of the cluster, find and install the operator in Menu > Operators > OperatorHub.
1. On each managed cluster, install the OpenShift GitOps operator.
   In the OpenShift Console of the cluster, find and install the operator in Menu > Operators > OperatorHub.

#### Deploy the applications to remote clusters

With the remote near edge clusters now in the cluster set and that cluster set bound to the `openshift-gitops` namespace,
it is now possible to deploy the applications to those clusters using the Argo CD Pull model integration in
[ACM GitOps](https://access.redhat.com/documentation/en-us/red_hat_advanced_cluster_management_for_kubernetes/2.9/html/gitops/index).

We will create the objects in the respective namespaces from the `acm/registration/` directory structure.

However, it is important to note that the actual application configuration, stored in `acm/odh-edge/`,
does not come from our local git repository checkout.
Instead, [`acm/registration/near-edge/overlays/tensorflow-housing-app/kustomization.yaml`](acm/registration/near-edge/overlays/tensorflow-housing-app/kustomization.yaml)
and [`acm/registration/near-edge/overlays/bike-rental-app/kustomization.yaml`](acm/registration/near-edge/overlays/bike-rental-app/kustomization.yaml)
contain source URLs of the git repositories that will become configurations of the ACM Channels,
and the applications in those remote git repositories will be deployed.
By default, [github.com/opendatahub-io/ai-edge](https://github.com/opendatahub-io/ai-edge) is configured;
edit the URLs to match your repositories.

Then, as a user with permissions to create/update the following resource types in the `openshift-gitops` namespace on the ACM hub cluster:

``` text
ApplicationSets
GitOpsClusters
ManagedClusterSetBindings
PlacementBindings
Placements
Policies
```

run the following command:
```bash
oc apply -k acm/registration/
```

##### Credentials for private repositories

If the GitOps repository that Argo CD on the edge clusters will deploy resources from requires some form of credentials,
then these credentials will need to be provided to each cluster via a Secret.
This can be done by specifying the Secret inside an ACM Policy resource on the ACM hub cluster in such a way that ACM will instruct the edge clusters to create it.
The following example illustrates how this may be done:
```sh
cat << EOF | oc apply -f -
apiVersion: policy.open-cluster-management.io/v1
kind: Policy
metadata:
  name: namespace-policy
  namespace: openshift-gitops
spec:
  disabled: false
  policy-templates:
    - objectDefinition:
        apiVersion: policy.open-cluster-management.io/v1
        kind: ConfigurationPolicy
        metadata:
          name: ensure-namespace-exists
        spec:
          object-templates:
            - complianceType: musthave
              objectDefinition:
                apiVersion: v1
                kind: Secret
                metadata:
                  name: first-repo
                  namespace: openshift-gitops
                  labels:
                    argocd.argoproj.io/secret-type: repository
                stringData:
                  type: git
                  url: https://github.com/argoproj/private-repo
          pruneObjectBehavior: DeleteIfCreated
          severity: low
          remediationAction: enforce
EOF
```

#### View the application deployments
##### ACM

In the OpenShift Console of the ACM hub cluster in All Clusters > Applications, search for the application name.
You should see a few results results: `<app name>-appset`, `<cluster name>-<app name>`, and `<app name>-1`.
Clicking on the `<cluster name>-<app name>` entry, and then navigating to the Topology tab, you can see what objects were created on the remote cluster.

Everything should be shown green. If it is not, click the icon of the faulty object and check the displayed information for debugging clues.

##### Argo CD

The Argo CD provided by Openshift GitOps has a console on each near edge cluster, showing detailed information on each application that it manages on that cluster.
On the particular edge cluster, open the console by navigating to the domain specified by the `openshift-gitops-cluster` Route in the `openshift-gitops` namespace on that cluster.
Once logged-in to the Argo console, search for the application name, and then select the entry named in the format `<cluster name>-<application name>` to see more information.

### Observability setup

* Edge cluster(s)
  * Login to the edge cluster using an account with cluster-admin privileges
  * Enable [monitoring for user-defined projects](https://access.redhat.com/documentation/en-us/openshift_container_platform/4.13/html/monitoring/enabling-monitoring-for-user-defined-projects) in OpenShift clusters
    * `oc -n openshift-monitoring edit configmap cluster-monitoring-config`
    * Set variable `enableUserWorkload` to `true`
* Core/Hub cluster
  * Edit contents of [thanos-secret](acm/odh-core/acm-observability/secrets/thanos.yaml) file.
  * Install the ACM observability stack by running `make install`

### Using local models in pipelines

In `pipelines/model-upload/` you can upload a local
model file to be used in our pipelines. This is done by uploading a model to a PVC
and copying that model to our pipeline's workspace for use while it is running.

Upload model to PVC:
```bash
make MODEL_PATH="PATH_TO_A_FILE" NAME=my-model create
```
You should get a final output showing details of the upload
```
PVC name: model-upload-pvc
Size: 1G
Model path in pod: /workspace/model-upload-pvc/model_dir/model.model
```
You can set the `SIZE` and `PVC` values aswell
```bash
make MODEL_PATH="PATH_TO_A_FILE" NAME=my-model SIZE=1G PVC=my-new-PVC create
```

You can then use the [copy-model-from-pvc](pipelines/tekton/aiedge-e2e/tasks/copy-model-from-pvc.yaml) task to
copy the model from the `model-workspace`, which can be set to the PVC created in the last step, to the `buildah-cache` workspace,
which is then used by the `buildah` task in the pipeline. [copy-model-from-pvc](pipelines/tekton/aiedge-e2e/tasks/copy-model-from-pvc.yaml) task
is not included in the pipeline by default, you have to add it yourself. 

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## Development

See [DEVELOPMENT.md](DEVELOPMENT.md).
