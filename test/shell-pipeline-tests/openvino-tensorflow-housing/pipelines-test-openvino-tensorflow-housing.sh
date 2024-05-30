#!/usr/bin/env bash
REPO_ROOT_DIR=$(dirname ${BASH_SOURCE[0]})/../../..
PIPELINES_DIR="$REPO_ROOT_DIR/pipelines"

source "$REPO_ROOT_DIR"/test/shell-pipeline-tests/common.sh

NAMESPACE="pipeline-test-tensorflow-housing"
oc delete project "$NAMESPACE" --ignore-not-found --timeout=60s
oc new-project "$NAMESPACE"

echo "Waiting for OpenShift Pipelines operator to be fully installed"
waitForOpResult 60 "True" "N/A" "oc get tektonconfig -n openshift-operators config -o jsonpath={.status.conditions[?\(@.type==\'Ready\'\)].status}"
waitForOpResult 10 "pipeline" "N/A" "oc get serviceaccount -o=custom-columns=NAME:.metadata.name | grep pipeline"

##### AIEDGE E2E PIPELINE
AIEDGE_E2E_PIPELINE_DIR_PATH="$PIPELINES_DIR"/tekton/aiedge-e2e

AWS_SECRET_PATH_TEMPLATE="$AIEDGE_E2E_PIPELINE_DIR_PATH"/templates/credentials-s3.secret.yaml.template
AWS_SECRET_PATH="$AIEDGE_E2E_PIPELINE_DIR_PATH"/templates/credentials-s3.secret.yaml

createS3Secret "$AWS_SECRET_PATH_TEMPLATE" "$AWS_SECRET_PATH"

oc create -f "$AWS_SECRET_PATH"

IMAGE_REGISTRY_SECRET_PATH_TEMPLATE="$AIEDGE_E2E_PIPELINE_DIR_PATH"/templates/credentials-image-registry.secret.yaml.template
IMAGE_REGISTRY_SECRET_PATH="$AIEDGE_E2E_PIPELINE_DIR_PATH"/templates/credentials-image-registry.secret.yaml

createImageRegistrySecret "$IMAGE_REGISTRY_SECRET_PATH_TEMPLATE" "$IMAGE_REGISTRY_SECRET_PATH"

oc create -f "$IMAGE_REGISTRY_SECRET_PATH"
oc secret link pipeline credentials-image-registry

## oc apply -k pipelines
oc apply -k "$AIEDGE_E2E_PIPELINE_DIR_PATH"/

## prepare parameters
AIEDGE_E2E_PIPELINE_OVERRIDDEN_PATH="$AIEDGE_E2E_PIPELINE_DIR_PATH"/aiedge-e2e.pipelinerun-overridden.yaml
cp "$AIEDGE_E2E_PIPELINE_DIR_PATH"/aiedge-e2e.tensorflow-housing.pipelinerun.yaml "$AIEDGE_E2E_PIPELINE_OVERRIDDEN_PATH"
sed -i "s|value: \"delete\"|value: \"keep\"|" "$AIEDGE_E2E_PIPELINE_OVERRIDDEN_PATH"
usePRBranchInPipelineRunIfPRCheck "$AIEDGE_E2E_PIPELINE_OVERRIDDEN_PATH"

## oc create pipeline run
oc create -f "$AIEDGE_E2E_PIPELINE_OVERRIDDEN_PATH"
sleep 5 # Just to have the startTime field available

PIPELINE_RUN_NAME=$(oc get pipelinerun --sort-by={.status.startTime} -o=custom-columns=NAME:.metadata.name | grep "aiedge-e2e-.*" | tail -n 1)

if [[ $PIPELINE_RUN_NAME == "" ]]; then
  echo "Could not find any pipeline run"
  exit 1
fi

## wait for the result
waitForOpResult 200 "True" "False" "oc get pipelinerun $PIPELINE_RUN_NAME -o jsonpath={.status.conditions[?\(@.type==\'Succeeded\'\)].status}"
PIPELINE_RUN_RESULT=$?

saveArtifacts "$PIPELINE_RUN_NAME"

if [[ $PIPELINE_RUN_RESULT != 0 ]]; then
  echo "The aiedge-e2e pipeline failed"
  exit 1
else
  echo "The aiedge-e2e pipeline finished successfully"
fi


##### GITOPS UPDATE PIPELINE
GITOPS_UPDATE_PIPELINE_DIR_PATH="$PIPELINES_DIR"/tekton/gitops-update-pipeline

GIT_CREDENTIALS_SECRET_PATH_TEMPLATE="$GITOPS_UPDATE_PIPELINE_DIR_PATH"/example-pipelineruns/example-git-credentials-secret.yaml.template
GIT_CREDENTIALS_SECRET_PATH="$GITOPS_UPDATE_PIPELINE_DIR_PATH"/example-pipelineruns/example-git-credentials-secret.yaml

createGitCredentialsSecret "$GIT_CREDENTIALS_SECRET_PATH_TEMPLATE" "$GIT_CREDENTIALS_SECRET_PATH"

oc create -f "$GIT_CREDENTIALS_SECRET_PATH"

## oc apply -k pipelines
oc apply -k "$GITOPS_UPDATE_PIPELINE_DIR_PATH"/

## prepare parameters
GITOPS_UPDATE_PIPELINERUN_PATH="$GITOPS_UPDATE_PIPELINE_DIR_PATH"/example-pipelineruns/gitops-update-pipelinerun-tensorflow-housing.yaml
GITOPS_UPDATE_PIPELINERUN_OVERRIDDEN_PATH="$GITOPS_UPDATE_PIPELINE_DIR_PATH"/example-pipelineruns/gitops-update-pipelinerun-tensorflow-housing-overridden.yaml
cp "$GITOPS_UPDATE_PIPELINERUN_PATH" "$GITOPS_UPDATE_PIPELINERUN_OVERRIDDEN_PATH"
NEW_DIGEST=$(oc get pipelinerun "$PIPELINE_RUN_NAME" -o jsonpath={.status.results[?\(@.name==\'buildah-sha\'\)].value})

sed -i "s|value: username|value: redhat-rhods-qe|" "$GITOPS_UPDATE_PIPELINERUN_OVERRIDDEN_PATH"
sed -i "s|value: ai-edge-gitops|value: ai-edge-ci-test|" "$GITOPS_UPDATE_PIPELINERUN_OVERRIDDEN_PATH"
sed -i "s|value: sha256.*|value: ${NEW_DIGEST}|" "$GITOPS_UPDATE_PIPELINERUN_OVERRIDDEN_PATH"

## oc create pipeline run
oc create -f "$GITOPS_UPDATE_PIPELINERUN_OVERRIDDEN_PATH"
sleep 5 # Just to have the startTime field available

PIPELINE_RUN_NAME=$(oc get pipelinerun --sort-by={.status.startTime} -o=custom-columns=NAME:.metadata.name | grep "gitops-update-pipeline-tensorflow-housing-.*" | tail -n 1)

if [[ $PIPELINE_RUN_NAME == "" ]]; then
  echo "Could not find any pipeline run"
  exit 1
fi

## wait for the result
waitForOpResult 200 "True" "False" "oc get pipelinerun $PIPELINE_RUN_NAME -o jsonpath={.status.conditions[?\(@.type==\'Succeeded\'\)].status}"
PIPELINE_RUN_RESULT=$?

saveArtifacts "$PIPELINE_RUN_NAME"

if [[ $PIPELINE_RUN_RESULT != 0 ]]; then
  echo "The gitops-update-pipeline pipeline failed"
  exit 1
else
  echo "The gitops-update-pipeline pipeline finished successfully"
fi


##### Pull Request validation
installGithubCLI
export GH_TOKEN=$(getGitHubToken)

echo "Checking PR content"
PR_URL=$(oc get pipelinerun "$PIPELINE_RUN_NAME" -o jsonpath={.status.results[?\(@.name==\'pr-url\'\)].value})

echo "Pull Request URL is: $PR_URL"

assertGitHubPullRequest "$PIPELINE_RUN_NAME" "$PR_URL"
assertExitCode=$?

echo "Closing Pull Request"
gh pr close "$PR_URL" -d

if [[ $assertExitCode != 0 ]]; then
  echo "At least one assertion failed, check messages above for details"
  exit 1
else
  echo "All pipelines finished successfully and the PR check is correct"
fi
