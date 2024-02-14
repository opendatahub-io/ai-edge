.PHONY: install install/observability-core install/observability-edge test-acm-%-generate test go-test go-test-setup

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

# requires:
#	AWS_SECRET_ACCESS_KEY		- Secret from AWS
#	AWS_ACCESS_KEY_ID		- Access key from AWS
#	S3_REGION			- Region of the bucket used to store the model
#	S3_ENDPOINT			- Endpint of the bucket
#	IMAGE_REGISTRY_USERNAME		- quay.io username
#	IMAGE_REGISTRY_PASSWORD		- quay.io password
go-test-setup:
ifndef AWS_SECRET_ACCESS_KEY
	$(error AWS_SECRET_ACCESS_KEY is undefined)
endif
ifndef AWS_ACCESS_KEY_ID
	$(error AWS_ACCESS_KEY_ID is undefined)
endif
ifndef S3_REGION
	$(error S3_REGION is undefined)
endif
ifndef S3_ENDPOINT
	$(error S3_ENDPOINT is undefined)
endif
ifndef IMAGE_REGISTRY_USERNAME
	$(error IMAGE_REGISTRY_USERNAME is undefined)
endif
ifndef IMAGE_REGISTRY_PASSWORD
	$(error IMAGE_REGISTRY_PASSWORD is undefined)
endif
	@sed -e "s#{{ AWS_SECRET_ACCESS_KEY }}#${AWS_SECRET_ACCESS_KEY}#g" -e "s#{{ AWS_ACCESS_KEY_ID }}#${AWS_ACCESS_KEY_ID}#g" -e "s#{{ S3_REGION }}#${S3_REGION}#g" -e "s#{{ S3_ENDPOINT }}#${S3_ENDPOINT}#g" pipelines/tekton/aiedge-e2e/templates/credentials-s3.secret.yaml.template | oc create -f -
	@sed -e "s#{{ IMAGE_REGISTRY_USERNAME }}#${IMAGE_REGISTRY_USERNAME}#g" -e "s#{{ IMAGE_REGISTRY_PASSWORD }}#${IMAGE_REGISTRY_PASSWORD}#g" pipelines/tekton/aiedge-e2e/templates/credentials-image-registry.secret.yaml.template | oc create -f -
	oc secret link pipeline credentials-image-registry

# requires:
#	S3_BUCKET		- Name of S3 bucket that has the model
#	TARGET_IMAGE_NAMESPACE	- Image registry namespace that the built model will be pushed to a repository in
#	NAMESPACE		- Cluster namespace that tests are run in
go-test:
ifndef S3_BUCKET
	$(error S3_BUCKET is undefined)
endif
ifndef TARGET_IMAGE_NAMESPACE
	$(error TARGET_IMAGE_NAMESPACE is undefined)
endif
ifndef NAMESPACE
	$(error NAMESPACE is undefined)
endif
	(cd test/e2e-tests/tests && S3_BUCKET=${S3_BUCKET} TARGET_IMAGE_NAMESPACE=${TARGET_IMAGE_NAMESPACE} NAMESPACE=${NAMESPACE} ${GO} test -timeout 30m)

test:
	@(./test/shell-pipeline-tests/seldon-bike-rentals/pipelines-test-seldon-bike-rentals.sh)
	@(./test/shell-pipeline-tests/openvino-tensorflow-housing/pipelines-test-openvino-tensorflow-housing.sh)
