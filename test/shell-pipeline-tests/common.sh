#!/usr/bin/env bash

function waitForOpResult() {
    max_retry=$1
    shift
    expected=$1
    shift
    unexpected=$1
    shift
    counter=0
    cmd=$1
    echo "waitForOpResult waiting for command '$cmd' to finish with expected result '$expected' or unexpected result '$unexpected'"
    res=$(eval "$cmd")
    until [ "$res" == "$expected" ]
    do
       [[ counter -eq $max_retry ]] && echo "Failed! waitForOpResult running command '$cmd' and waiting for expected output '$expected' reached max retry count '$max_retry'." >&2 && return 1
       [[ "$res" == "$unexpected" ]] && echo "Failed! waitForOpResult running command '$cmd' and waiting for expected output '$expected' finished with unexpected result '$res'." >&2 && return 1
       echo "Waiting for another try"
       sleep 5
       ((counter++))
       echo "Trying again. Try #$counter out of $max_retry" >&2
       res=$(eval "$cmd")
    done
    echo "waitForOpResult running command '$cmd' finished with expected result '$res'"
}

function saveArtifacts() {
    ## Backup all Pipeline Runs, Task Runs
    local PIPELINE_RUN_NAME=$1
    local LOGS_DIR="${ARTIFACT_DIR}/$PIPELINE_RUN_NAME" # ARTIFACT_DIR is an env var on OpenShift-CI
    mkdir -p "$LOGS_DIR"
    echo "Archiving YAML definitions and logs for '$PIPELINE_RUN_NAME' to '$LOGS_DIR'"
    oc get pipeline -o yaml > "${LOGS_DIR}"/pipelines.txt
    oc get pipelinerun $PIPELINE_RUN_NAME -o yaml > "${LOGS_DIR}"/pipelineRuns.txt
    oc get task -o yaml > "${LOGS_DIR}"/tasks.txt
    oc get taskrun -l "tekton.dev/pipelineRun=$PIPELINE_RUN_NAME" -o yaml > "${LOGS_DIR}"/taskRuns.txt
    oc logs -l "tekton.dev/pipelineRun=$PIPELINE_RUN_NAME" --all-containers --prefix --tail=-1 > "${LOGS_DIR}"/pipelineLogs.txt
    oc get deployment -o yaml > "${LOGS_DIR}"/deployments.txt
    oc logs -l '!tekton.dev/pipelineRun' --all-containers --prefix --tail=-1 > "${LOGS_DIR}"/deploymentLogs.txt
    # https://access.redhat.com/solutions/4725511
    oc get events -o custom-columns="LAST SEEN:{lastTimestamp},FIRST SEEN:{firstTimestamp},COUNT:{count},TYPE:{type},REASON:{reason},KIND:{involvedObject.kind},NAME:{involvedObject.name},SOURCE:{source.component},MESSAGE:{message}" --sort-by={lastTimestamp} > "${LOGS_DIR}"/events.txt
}

function createS3Secret() {
    local AWS_SECRET_PATH_TEMPLATE=$1
    local AWS_SECRET_PATH=$2

    local AI_EDGE_AWS_VAULT_OPENSHIFT_CI_SECRET_PATH
    local AWS_ACCESS_KEY
    local AWS_SECRET_KEY

    AI_EDGE_AWS_VAULT_OPENSHIFT_CI_SECRET_PATH="${CUSTOM_AWS_SECRET_PATH:-/secrets/ai-edge-aws}"
    AWS_ACCESS_KEY=$(cat "$AI_EDGE_AWS_VAULT_OPENSHIFT_CI_SECRET_PATH"/accessKey)
    AWS_SECRET_KEY=$(cat "$AI_EDGE_AWS_VAULT_OPENSHIFT_CI_SECRET_PATH"/secretAccessKey)

    cp "$AWS_SECRET_PATH_TEMPLATE" "$AWS_SECRET_PATH"

    sed -i "s|{{ AWS_ACCESS_KEY_ID }}|${AWS_ACCESS_KEY}|" "$AWS_SECRET_PATH"
    sed -i "s|{{ AWS_SECRET_ACCESS_KEY }}|${AWS_SECRET_KEY}|" "$AWS_SECRET_PATH"
    sed -i "s|{{ S3_ENDPOINT }}|https://s3.us-west-1.amazonaws.com|" "$AWS_SECRET_PATH"
    sed -i "s|{{ S3_REGION }}|us-west-1|" "$AWS_SECRET_PATH"
}

function createImageRegistrySecret() {
    local -r IMAGE_REGISTRY_SECRET_PATH_TEMPLATE=$1
    local -r IMAGE_REGISTRY_SECRET_PATH=$2

    local -r AI_EDGE_IMAGE_REGISTRY_OPENSHIFT_CI_SECRET_PATH="${CUSTOM_IMAGE_REGISTRY_SECRET_PATH:-/secrets/ai-edge-quay}"
    local -r USERNAME=$(cat "$AI_EDGE_IMAGE_REGISTRY_OPENSHIFT_CI_SECRET_PATH"/username)
    local -r PASSWORD=$(cat "$AI_EDGE_IMAGE_REGISTRY_OPENSHIFT_CI_SECRET_PATH"/password)

    cp "$IMAGE_REGISTRY_SECRET_PATH_TEMPLATE" "$IMAGE_REGISTRY_SECRET_PATH"

    sed -i "s|{{ IMAGE_REGISTRY_USERNAME }}|${USERNAME}|" "$IMAGE_REGISTRY_SECRET_PATH"
    sed -i "s|{{ IMAGE_REGISTRY_PASSWORD }}|${PASSWORD}|" "$IMAGE_REGISTRY_SECRET_PATH"
}

function createGitCredentialsSecret() {
    local GIT_CREDENTIALS_SECRET_PATH_TEMPLATE=$1
    local GIT_CREDENTIALS_SECRET_PATH=$2

    local GH_TOKEN
    GH_TOKEN=$(getGitHubToken)

    cp "$GIT_CREDENTIALS_SECRET_PATH_TEMPLATE" "$GIT_CREDENTIALS_SECRET_PATH"

    sed -i "s|{github_pat_1234567890ABCDAPI_TOKEN}|$GH_TOKEN|" "$GIT_CREDENTIALS_SECRET_PATH"
    sed -i "s|{username}|ods-qe-test|" "$GIT_CREDENTIALS_SECRET_PATH"
}

function usePRBranchInPipelineRunIfPRCheck() {
  local -r PIPELINE_RUN_FILE_PATH=$1

  if [ -n "$PULL_NUMBER" ] && [ "$REPO_NAME" == "ai-edge" ]; then
    sed -i "s|value: \"main\"|value: \"pull/$PULL_NUMBER/head\"|" "$PIPELINE_RUN_FILE_PATH"
  fi
}

function assertContains() {
  local -r STRING=$1
  local -r SUBSTRING=$2

  [[ -z $SUBSTRING ]] && echo -e "Assertion failed!\n '$STRING' contains '$SUBSTRING' but it is empty, so the assertion would otherwise yield a false positive result" && return 1;
  [[ "$STRING" =~ $SUBSTRING ]] || { echo -e "Assertion failed!\n '$STRING' does not contain '$SUBSTRING'" && return 1; }
}

assertGitHubPullRequest() {
  local -r PIPELINE_RUN_NAME=$1
  local -r PR_URL=$2

  local -r PR_INFO_PATH="pr-info-${PIPELINE_RUN_NAME}.txt"
  local -r PR_DIFF_PATH="pr-diff-${PIPELINE_RUN_NAME}.txt"

  gh pr view "$PR_URL" > "$PR_INFO_PATH"
  gh pr diff "$PR_URL" > "$PR_DIFF_PATH"

  # Expected values
  PIPELINE_RUN_UID=$(oc get pipelinerun "$PIPELINE_RUN_NAME" -o jsonpath={.metadata.uid})
  IMAGE_REGISTRY=$(oc get pipelinerun "$PIPELINE_RUN_NAME" -o jsonpath={.spec.params[?\(@.name==\'image-registry-repo\'\)].value})
  NEW_DIGEST=$(oc get pipelinerun "$PIPELINE_RUN_NAME" -o jsonpath={.spec.params[?\(@.name==\'image-digest\'\)].value})
  BASE_REF_NAME=$(oc get pipelinerun "$PIPELINE_RUN_NAME" -o jsonpath={.spec.params[?\(@.name==\'gitRepoBranchBase\'\)].value})
  GIT_SERVER=$(oc get pipelinerun "$PIPELINE_RUN_NAME" -o jsonpath={.spec.params[?\(@.name==\'gitServer\'\)].value})
  GIT_ORG_NAME=$(oc get pipelinerun "$PIPELINE_RUN_NAME" -o jsonpath={.spec.params[?\(@.name==\'gitOrgName\'\)].value})
  GIT_REPO_NAME=$(oc get pipelinerun "$PIPELINE_RUN_NAME" -o jsonpath={.spec.params[?\(@.name==\'gitRepoName\'\)].value})
  GIT_URL="$GIT_SERVER/$GIT_ORG_NAME/$GIT_REPO_NAME"

  # Actual values
  PR_PIPELINE_RUN_NAME=$(grep "PipelineRun Name" "$PR_INFO_PATH" )
  PR_PIPELINE_RUN_UID=$(grep "PipelinRun UID" "$PR_INFO_PATH")
  PR_IMAGE_REGISTRY=$(grep "Image registry" "$PR_INFO_PATH")
  PR_NEW_DIGEST=$(grep "New Digest" "$PR_INFO_PATH")
  PR_BASE_REF_NAME=$(gh pr view "$PR_URL" --json baseRefName -q ".baseRefName")
  PR_HEAD_REF_NAME=$(gh pr view "$PR_URL" --json headRefName -q ".headRefName")
  PR_DIFF_DIGEST=$(grep "+ *digest" "$PR_DIFF_PATH")
  PR_DIFF_NEW_NAME=$(grep "newName" "$PR_DIFF_PATH")

  # Assertions
  echo "Running Pull Request Assertions"
  assertExitCode=0
  assertContains "$PR_PIPELINE_RUN_NAME" "$PIPELINE_RUN_NAME"; (( assertExitCode = assertExitCode || $? ))
  assertContains "$PR_PIPELINE_RUN_UID" "$PIPELINE_RUN_UID"; (( assertExitCode = assertExitCode || $? ))
  assertContains "$PR_IMAGE_REGISTRY" "$IMAGE_REGISTRY"; (( assertExitCode = assertExitCode || $? ))
  assertContains "$PR_NEW_DIGEST" "$NEW_DIGEST"; (( assertExitCode = assertExitCode || $? ))
  assertContains "$PR_BASE_REF_NAME" "$BASE_REF_NAME"; (( assertExitCode = assertExitCode || $? ))
  assertContains "$PR_URL" "$GIT_URL"; (( assertExitCode = assertExitCode || $? ))
  assertContains "$PR_HEAD_REF_NAME" "pipeline_$PIPELINE_RUN_UID"; (( assertExitCode = assertExitCode || $? ))
  assertContains "$PR_DIFF_DIGEST" "$NEW_DIGEST"; (( assertExitCode = assertExitCode || $? ))
  assertContains "$PR_DIFF_NEW_NAME" "$IMAGE_REGISTRY"; (( assertExitCode = assertExitCode || $? ))

  return "$assertExitCode"
}

function getGitHubToken() {
  AI_EDGE_GITHUB_VAULT_OPENSHIFT_CI_SECRET_PATH="${CUSTOM_GIT_CREDENTIALS_SECRET_PATH:-/secrets/ai-edge-github}"
  GH_TOKEN=$(cat "$AI_EDGE_GITHUB_VAULT_OPENSHIFT_CI_SECRET_PATH"/token)
  echo "$GH_TOKEN"
}

function installGithubCLI() {
  echo "Installing GitHub CLI"
  wget -nv https://github.com/cli/cli/releases/download/v2.47.0/gh_2.47.0_linux_amd64.tar.gz
  tar -xf gh_2.47.0_linux_amd64.tar.gz
  export PATH="$(pwd)/gh_2.47.0_linux_amd64/bin:$PATH"
  echo "Using GitHub CLI from $(which gh)"
}
