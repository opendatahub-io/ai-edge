# Bring Your Own Cluster Management

If you don't have a centralized ACM managing your edge cluster(s), you can follow the steps in this document to deploy
the inference service container image that was built using the MLOps Pipeline.

## Install OpenShift Pipelines / ArgoCD

Just like in the case with ACM, ArgoCD will need to be installed on each target edge cluster that the GitOps-managed
inference service container will be deployed on.
As a user with the ability to install operators from the OperatorHub on the target edge cluster(s), find and install the
OpenShift GitOps operator in the OpenShift Console, in Menu > Operators > OperatorHub.

## Ensure namespace in GitOps repo has correct label

In your GitOps repo, where the manifests for your application are stored, make sure that the namespace definition has
the correct label so that OpenShift GitOps can manage resources in it once it creates it.

To allow the default configuration of OpenShift GitOps to deploy into the namespace, it will need to have the following
label:

``` yaml
argocd.argoproj.io/managed-by: openshift-gitops
```

For more information on this label, see the [OpenShift GitOps docs][OpenShift GitOps docs: Deploying resources to a
different namespace].

## Create ArgoCD Application on edge clusters

Once the operator has been installed from the earlier step, the `Application` CRD will be available on the edge
cluster(s), allowing the creation of Application CRs in the GitOps namespace.

On each edge cluster, create an ArgoCD Application CR to point the ArgoCD instance at the location in the GitOps repo
where the manifests are located.

Examples using `kustomize` can be found in the [byo-mgmt](../byo-mgmt) directory.

### Examples

In the root of this repository, the following can be run to roll out the example inference container applications that
exist in this project:

#### Tensorflow housing example application

```
oc apply -k byo-mgmt/registration/near-edge/overlays/tensorflow-housing-app/
```

#### Bike rental example application

```
oc apply -k byo-mgmt/registration/near-edge/overlays/bike-rental-app/
```

## Observability

On each of the edge clusters, you can enable [monitoring for user-defined projects]:
* `oc -n openshift-monitoring edit configmap cluster-monitoring-config`
* Set variable `enableUserWorkload` to `true`.

If you forward metrics from each edge cluster to a central location, you can find a list of example metrics from the
OpenVINO and Seldon model servers from our examples in
[metrics_list](../acm/odh-core/acm-observability/files/metrics_list.yaml) that you may wish to forward.


[OpenShift GitOps docs: Deploying resources to a different namespace]: https://docs.openshift.com/gitops/1.11/argocd_instance/setting-up-argocd-instance.html#gitops-deploy-resources-different-namespaces_setting-up-argocd-instance
[monitoring for user-defined projects]: https://access.redhat.com/documentation/en-us/openshift_container_platform/4.14/html/monitoring/enabling-monitoring-for-user-defined-projects
