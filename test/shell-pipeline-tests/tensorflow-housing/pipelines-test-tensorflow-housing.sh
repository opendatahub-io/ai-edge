#!/usr/bin/env bash
REPO_ROOT_DIR=$(dirname ${BASH_SOURCE[0]})/../../..
PIPELINES_DIR="$REPO_ROOT_DIR/pipelines"

source "$REPO_ROOT_DIR"/test/shell-pipeline-tests/common.sh

NAMESPACE="pipeline-test-tensorflow-housing"
oc delete project "$NAMESPACE" --ignore-not-found --timeout=60s
oc new-project "$NAMESPACE"

echo "Waiting for OpenShift Pipelines operator to be fully installed"
waitForOpResult 60 "True" "N/A" "oc get tektonconfig -n openshift-operators config -o jsonpath={.status.conditions[?\(@.type==\'Ready\'\)].status}"

##### BUILD CONTAINER IMAGE PIPELINE #####
BUILD_CONTAINER_IMAGE_PIPELINE_PATH="$PIPELINES_DIR"/tekton/build-container-image-pipeline

AWS_SECRET_PATH_TEMPLATE="$BUILD_CONTAINER_IMAGE_PIPELINE_PATH"/aws-env.yaml
AWS_SECRET_PATH="$BUILD_CONTAINER_IMAGE_PIPELINE_PATH"/aws-env-overridden.yaml

createS3Secret "$AWS_SECRET_PATH_TEMPLATE" "$AWS_SECRET_PATH"

oc create -f "$AWS_SECRET_PATH"

## oc apply -k pipelines
oc apply -k "$BUILD_CONTAINER_IMAGE_PIPELINE_PATH"/

## prepare parameters
cp "$BUILD_CONTAINER_IMAGE_PIPELINE_PATH"/build-container-image-pipelinerun-tensorflow-housing.yaml "$BUILD_CONTAINER_IMAGE_PIPELINE_PATH"/build-container-image-pipelinerun-tensorflow-housing-overridden.yaml
sed -i "s|value: rhoai-edge-models|value: rhoai-edge-models-ci|" "$BUILD_CONTAINER_IMAGE_PIPELINE_PATH"/build-container-image-pipelinerun-tensorflow-housing-overridden.yaml

## oc create pipeline run
oc create -f "$BUILD_CONTAINER_IMAGE_PIPELINE_PATH"/build-container-image-pipelinerun-tensorflow-housing-overridden.yaml
sleep 5 # Just to have the startTime field available

PIPELINE_RUN_NAME=$(oc get pipelinerun --sort-by={.status.startTime} -o=custom-columns=NAME:.metadata.name | grep "build.*housing" | tail -n 1)

if [[ $PIPELINE_RUN_NAME == "" ]]; then
  echo "Could not find any pipeline run"
  exit 1
fi

## wait for the result
waitForOpResult 200 "True" "False" "oc get pipelinerun $PIPELINE_RUN_NAME -o jsonpath={.status.conditions[?\(@.type==\'Succeeded\'\)].status}"
PIPELINE_RUN_RESULT=$?

saveArtifacts "$PIPELINE_RUN_NAME"

if [[ $PIPELINE_RUN_RESULT != 0 ]]; then
  echo "Build pipeline failed, aborting further tests"
  exit 1
fi


##### TEST MLFLOW IMAGE PIPELINE #####
TEST_MLFLOW_IMAGE_PIPELINE_PATH="$PIPELINES_DIR"/tekton/test-mlflow-image-pipeline

AI_EDGE_QUAY_SECRET_OPENSHIFT_CI_PATH="${CUSTOM_QUAY_SECRET_PATH:-/secrets/ai-edge-quay}"
oc create secret generic rhoai-edge-openshift-ci-secret --from-file=.dockerconfigjson="$AI_EDGE_QUAY_SECRET_OPENSHIFT_CI_PATH"/dockerconfigjson --type=kubernetes.io/dockerconfigjson --dry-run=client -o yaml | oc apply -f -
oc secret link pipeline rhoai-edge-openshift-ci-secret

## oc apply -k pipelines
oc apply -k "$TEST_MLFLOW_IMAGE_PIPELINE_PATH"/

## oc create pipeline run
oc create -f "$TEST_MLFLOW_IMAGE_PIPELINE_PATH"/test-mlflow-image-pipelinerun-tensorflow-housing.yaml
sleep 5 # Just to have the startTime field available

PIPELINE_RUN_NAME=$(oc get pipelinerun --sort-by={.status.startTime} -o=custom-columns=NAME:.metadata.name | grep "test.*housing" | tail -n 1)

if [[ $PIPELINE_RUN_NAME == "" ]]; then
  echo "Could not find any pipeline run"
  exit 1
fi

## wait for the result
waitForOpResult 200 "True" "False" "oc get pipelinerun $PIPELINE_RUN_NAME -o jsonpath={.status.conditions[?\(@.type==\'Succeeded\'\)].status}"
PIPELINE_RUN_RESULT=$?

saveArtifacts "$PIPELINE_RUN_NAME"

if [[ $PIPELINE_RUN_RESULT != 0 ]]; then
  echo "Test pipeline failed, aborting further tests"
  exit 1
fi

echo "All pipelines finished successfully"
