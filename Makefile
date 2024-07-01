# Read any custom variables overrides from a local.vars.mk file.  This will only be read if it exists in the
# same directory as this Makefile.  Variables can be specified in the standard format supported by
# GNU Make since include process Makefiles
# Standard variables override would include anything you would pass at runtime that is different from the defaults specified in this file
MAKE_ENV_FILE = local.vars.mk
-include $(MAKE_ENV_FILE)

.PHONY: setup-observability test-acm-%-generate test go-test

# Setup the secret required for observability then generate and apply the files
setup-observability: acm/odh-core/acm-observability/kustomization.yaml
	oc apply -k acm/odh-core/acm-observability

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
	oc kustomize test/acm/$(subst -generate,,$(subst test-acm-,,$@))/ | sed -e "s|https://github.com/opendatahub-io/ai-edge|$(GIT_REPO_URL)|g" -e "s|my-git-branch|$(GIT_BRANCH)|g" -e "s|custom-prefix-|$(CUSTOM_PREFIX)|g" -e "s|custom-app-namespace|$(CUSTOM_APP_NAMESPACE)|g"

GO=go
GOFLAGS=""

go-test:
	(cd test/e2e-tests/tests && ${GO} test -timeout 60m -shuffle off)

test:
	${MAKE} -C cli cli-test
	@(./test/shell-pipeline-tests/seldon-bike-rentals/pipelines-test-seldon-bike-rentals.sh)
	@(./test/shell-pipeline-tests/openvino-tensorflow-housing/pipelines-test-openvino-tensorflow-housing.sh)

# This is a generic target to forward any cli-* targets to the cli Makefile
cli-%:
	${MAKE} -C cli $@
