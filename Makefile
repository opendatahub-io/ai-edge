# Read any custom variables overrides from a local.vars.mk file.  This will only be read if it exists in the
# same directory as this Makefile.  Variables can be specified in the standard format supported by
# GNU Make since include process Makefiles
# Standard variables override would include anything you would pass at runtime that is different from the defaults specified in this file
MAKE_ENV_FILE = local.vars.mk
-include $(MAKE_ENV_FILE)

.PHONY: setup-observability test-acm-%-generate test go-test go-test-setup

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

# requires:
#   AWS_SECRET_ACCESS_KEY       - Secret from AWS
#   AWS_ACCESS_KEY_ID           - Access key from AWS
#   S3_REGION                   - Region of the bucket used to store the model
#   S3_ENDPOINT                 - Endpint of the bucket
#   IMAGE_REGISTRY_USERNAME     - quay.io username
#   IMAGE_REGISTRY_PASSWORD     - quay.io password
#   GIT_SELF_SIGNED_CERT        - Self-signed cert to be used in pipeline by git (optional)
#   S3_SELF_SIGNED_CERT         - Self-signed cert to be used in pipeline by the script that's used to download the model from S3-compatible object storage (optional)
#   GIT_TOKEN                   - Auth token used by the git ops pipeline to make a pull request
#   GIT_USERNAME                - Username linked to the GIT_TOKEN
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
ifdef GIT_SELF_SIGNED_CERT
	oc create configmap git-self-signed-cert --from-file=ca-bundle.crt=${GIT_SELF_SIGNED_CERT}
endif
ifdef S3_SELF_SIGNED_CERT
	oc create configmap s3-self-signed-cert --from-file=ca-bundle.crt=${S3_SELF_SIGNED_CERT}
endif
ifndef GIT_TOKEN
	$(error GIT_TOKEN is undefined)
endif
ifndef GIT_USERNAME
	$(error GIT_USERNAME is undefined)
endif
	# MLOps pipieline setup
	@sed \
	-e "s#{{ AWS_SECRET_ACCESS_KEY }}#${AWS_SECRET_ACCESS_KEY}#g" \
	-e "s#{{ AWS_ACCESS_KEY_ID }}#${AWS_ACCESS_KEY_ID}#g" \
	-e "s#{{ S3_REGION }}#${S3_REGION}#g" \
	-e "s#{{ S3_ENDPOINT }}#${S3_ENDPOINT}#g" \
	pipelines/tekton/aiedge-e2e/templates/credentials-s3.secret.yaml.template | oc apply -f -

	@sed \
	-e "s#{{ IMAGE_REGISTRY_USERNAME }}#${IMAGE_REGISTRY_USERNAME}#g"  \
	-e "s#{{ IMAGE_REGISTRY_PASSWORD }}#${IMAGE_REGISTRY_PASSWORD}#g" \
	pipelines/tekton/aiedge-e2e/templates/credentials-image-registry.secret.yaml.template | oc apply -f -

	oc secret link pipeline credentials-image-registry

	# GITOps pipeline setup
	@sed -e \
	"s#{username}#${GIT_USERNAME}#g" \
	-e "s#{github_pat_1234567890ABCDAPI_TOKEN}#${GIT_TOKEN}#g" \
	pipelines/tekton/gitops-update-pipeline/example-pipelineruns/example-git-credentials-secret.yaml | oc apply -f -


# requires:
#   S3_BUCKET               - Name of S3 bucket that contains the models
#   NAMESPACE               - Cluster namespace that tests are run in
#   TARGET_IMAGE_TAGS_JSON  - JSON array of image tags that the final image will be pushed to. E.g. '["quay.io/user/model-name:e2e-test"]'
#   GIT_SELF_SIGNED_CERT    - Self-signed cert to be used in pipeline by git (optional)
#   S3_SELF_SIGNED_CERT     - Self-signed cert to be used in pipeline by the script that's used to download the model from S3-compatible object storage (optional)
#   GIT_REPO                - Git repo URL used to make a pull request in the git ops pipeline (https://github.com/org/repo)
#   GIT_API_SERVER          - Git API server (api.github.com)
#   GIT_BRANCH              - Base branch used for pull request in git ops pipeline
go-test:
ifndef S3_BUCKET
	$(error S3_BUCKET is undefined)
endif
ifndef NAMESPACE
	$(error NAMESPACE is undefined)
endif
ifndef TARGET_IMAGE_TAGS_JSON
	$(error TARGET_IMAGE_TAGS_JSON is undefined)
endif
ifndef GIT_REPO
	$(error GIT_REPO is undefined)
endif
ifndef GIT_API_SERVER
	$(error GIT_API_SERVER is undefined)
endif
ifndef GIT_BRANCH
	$(error GIT_BRANCH is undefined)
endif
	(cd test/e2e-tests/tests && S3_BUCKET=${S3_BUCKET} \
	TARGET_IMAGE_TAGS_JSON=${TARGET_IMAGE_TAGS_JSON} \
	NAMESPACE=${NAMESPACE} \
	GIT_SELF_SIGNED_CERT=${GIT_SELF_SIGNED_CERT} \
	S3_SELF_SIGNED_CERT=${S3_SELF_SIGNED_CERT} \
	GIT_REPO=${GIT_REPO} \
	GIT_API_SERVER=${GIT_API_SERVER} \
	GIT_BRANCH=${GIT_BRANCH} \
	${GO} test -timeout 30m -shuffle off)

test:
	${MAKE} -C cli cli-test
	@(./test/shell-pipeline-tests/seldon-bike-rentals/pipelines-test-seldon-bike-rentals.sh)
	@(./test/shell-pipeline-tests/openvino-tensorflow-housing/pipelines-test-openvino-tensorflow-housing.sh)

# This is a generic target to forward any cli-* targets to the cli Makefile
cli-%:
	${MAKE} -C cli $@
