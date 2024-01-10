.PHONY: install install/observability-core install/observability-edge test-acm-%-generate test

install: install/observability-core install/observability-edge

# Setup the secret required for observability then generate and apply the files
install/observability-core: acm/odh-core/acm-observability/kustomization.yaml
	DOCKER_CONFIG_JSON=`oc extract secret/pull-secret -n openshift-config --to=-`
	oc create secret generic multiclusterhub-operator-pull-secret -n open-cluster-management-observability --from-literal=.dockerconfigjson="$$DOCKER_CONFIG_JSON" --type=kubernetes.io/dockerconfigjson
	kustomize build acm/odh-core/acm-observability | kubectl apply -f -

# Create the whitelist for model metrics
install/observability-edge: acm/odh-edge/acm-observability/kustomization.yaml
	kustomize build acm/odh-edge/acm-observability | kubectl apply -f -

# Generate app manifests using ACM GitOps flow, from custom GitHub org/branch, to custom namespace
# Example invocations:
# make -s -e GIT_REPO_URL="https\://github.com/opendatahub-io/ai-edge" GIT_BRANCH=my-git-branch CUSTOM_PREFIX=custom- CUSTOM_APP_NAMESPACE=custom-bike-rental-app test-acm-bike-rental-app-generate
# make -s -e GIT_REPO_URL="https\://github.com/opendatahub-io/ai-edge" GIT_BRANCH=my-git-branch CUSTOM_PREFIX=custom- CUSTOM_APP_NAMESPACE=custom-tensorflow-housing test-acm-tensorflow-housing-app-generate
test-acm-%-generate: test/acm/%/kustomization.yaml
ifndef GIT_REPO_URL
	$(error GIT_REPO_URL is undefined)
endif
ifndef GIT_BRANCH
	$(error GIT_BRANCH is undefined)
endif
ifndef CUSTOM_PREFIX
	$(error CUSTOM_PREFIX is undefined)
endif
ifndef CUSTOM_APP_NAMESPACE
	$(error CUSTOM_APP_NAMESPACE is undefined)
endif
	kustomize build test/acm/$(subst -generate,,$(subst test-acm-,,$@))/ | sed -e "s|https://github.com/opendatahub-io/ai-edge|$(GIT_REPO_URL)|g" -e "s|my-git-branch|$(GIT_BRANCH)|g" -e "s|custom-prefix-|$(CUSTOM_PREFIX)|g" -e "s|custom-app-namespace|$(CUSTOM_APP_NAMESPACE)|g"

GO=go
GOFLAGS=""

test:
	@(./test/shell-pipeline-tests/openvino-bike-rentals/pipelines-test-openvino-bike-rentals.sh)
	@(./test/shell-pipeline-tests/tensorflow-housing/pipelines-test-tensorflow-housing.sh)
	@(cd test/e2e-tests/tests && ${GO} test)
