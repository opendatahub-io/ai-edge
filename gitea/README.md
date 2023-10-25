### Gitea in cluster git server for GitOps workflow
You can deploy a [Gitea](https://about.gitea.com/) server in your cluster instead of GitHub to provide a self contained GitOps workflow environment for your specific use case.

The [gitea-operator](https://github.com/rhpds/gitea-operator) can be used to manage the Gitea server installation in the cluster.  It will simplify the setup so that you can create a minimal [Gitea](https://github.com/rhpds/gitea-operator#migrating-repositories-for-created-users) CR to configure and install the Gitea server.

1. Install the [OpenShift Lifecycle Manager](https://docs.openshift.com/container-platform/4.13/operators/understanding/olm/olm-understanding-olm.html) `CatalogSource` and `Subscription` to deploy the `gitea-operator` in the cluster
   ```bash
   oc apply -k gitea/operator
   ```

1. Wait for the gitea-operator installation to complete and the `gitea.pfe-rhpds.com` CRD is available on the `odh-core` cluster
   ```bash
   $ oc get crd gitea.pfe.rhpds.com
   NAME                  CREATED AT
   gitea.pfe.rhpds.com   2023-08-25T03:00:13Z
   ```

1. Create the Gitea CustomResource to deploy the server with an admin user
   ```bash
   oc apply -k gitea/server
   ```

1. Once complete, there will be a gitea application deployed in the `gitea` namespace on the cluster.
   You can login to the gitea server on the route in the `gitea` namespace using the credentials specifed
   in the [gitea](server/gitea.yaml)
   ```bash
   GITEA_SERVER_URL="http://$(oc get route -n gitea  gitea-ai-edge -o jsonpath='{.spec.host}')"
   ```

   Open a browser to `${GITEA_SERVER_URL}` OR `git clone` the repo locally to customize the manifests for your use case
   ```bash
   git clone ${GITEA_SERVER_URL}/edge-user-1/ai-edge-gitops
   ```
