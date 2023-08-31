
install: install/observability-core

test: test/kustomize

# Setup the secret required for observability then generate and apply the files
install/observability-core:
	DOCKER_CONFIG_JSON=`oc extract secret/pull-secret -n openshift-config --to=-`
	oc create secret generic multiclusterhub-operator-pull-secret -n open-cluster-management-observability --from-literal=.dockerconfigjson="$$DOCKER_CONFIG_JSON" --type=kubernetes.io/dockerconfigjson
	kustomize build acm/odh-core/acm-observability | kubectl apply -f

# Create the whitelist for model metrics
install/observability-edge:
	kustomize build acm/odh-edge/acm-observability | kubectl apply -f

# Test all kustomize files build sucessfully
test/kustomize:
	@for dir in `find . -name kustomization.yaml -exec dirname {} \;`; do \
		echo "Testing $$dir"; \
		cd $$dir; \
		kustomize build . > /dev/null; \
		cd - > /dev/null; \
	done

# Print a visual tree of all of the kustomization.yaml files in the repo
tree/kustomize:
	@tree -fP kustomization.yaml

