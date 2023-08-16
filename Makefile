
install: install/observability

# Setup the secret required for observability then generate and apply the files
install/observability:
	DOCKER_CONFIG_JSON=`oc extract secret/pull-secret -n openshift-config --to=-`
	oc create secret generic multiclusterhub-operator-pull-secret -n open-cluster-management-observability --from-literal=.dockerconfigjson="$DOCKER_CONFIG_JSON" --type=kubernetes.io/dockerconfigjson
	kustomize build acm/odh-core/acm-observability | kubectl apply -f
