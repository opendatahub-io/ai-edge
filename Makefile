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
# make -s -e GIT_REPO_URL="https\://github.com/opendatahub-io/ai-edge" -e GIT_BRANCH=my-git-branch -e TEST_NAMESPACE=my-test-namespace test-acm-bike-rental-app-generate
# make -s -e GIT_REPO_URL="https\://github.com/opendatahub-io/ai-edge" -e GIT_BRANCH=my-git-branch -e TEST_NAMESPACE=my-test-namespace test-acm-tensorflow-housing-app-generate
test-acm-%-generate: test/acm/%/kustomization.yaml
ifndef GIT_REPO_URL
	$(error GIT_REPO_URL is undefined)
endif
ifndef GIT_BRANCH
	$(error GIT_BRANCH is undefined)
endif
ifndef TEST_NAMESPACE
	$(error TEST_NAMESPACE is undefined)
endif
	kustomize build test/acm/$(subst -generate,,$(subst test-acm-,,$@))/ | sed -e "s|https://github.com/opendatahub-io/ai-edge|$(GIT_REPO_URL)|g" -e "s|my-git-branch|$(GIT_BRANCH)|g" -e "s|my-test-namespace|$(TEST_NAMESPACE)|g"

# Dummy target to allow onboarding to openshift-ci
test:
	echo "The tests will be run here"
