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
    #TODO return code
}

function saveArtifacts() {
    ## Backup all Pipeline Runs, Task Runs
    local PIPELINE_RUN_NAME=$1
    local LOGS_DIR="${ARTIFACT_DIR}/$PIPELINE_RUN_NAME"
    mkdir -p "$LOGS_DIR"
    echo "Archiving YAML definitions and logs for '$PIPELINE_RUN_NAME' to '$LOGS_DIR'"
    oc get pipeline -o yaml > "${LOGS_DIR}"/pipelines.txt
    oc get pipelinerun $PIPELINE_RUN_NAME -o yaml > "${LOGS_DIR}"/pipelineRuns.txt
    oc get task -o yaml > "${LOGS_DIR}"/tasks.txt
    oc get taskrun -l "tekton.dev/pipelineRun=$PIPELINE_RUN_NAME" -o yaml > "${LOGS_DIR}"/taskRuns.txt
    oc logs -l "tekton.dev/pipelineRun=$PIPELINE_RUN_NAME" --all-containers --prefix --tail=-1 > "${LOGS_DIR}"/logs.txt
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

    sed -i "s|{{ YOUR_AWS_ACCESS_KEY }}|${AWS_ACCESS_KEY}|" "$AWS_SECRET_PATH"
    sed -i "s|{{ YOUR_AWS_SECRET_KEY }}|${AWS_SECRET_KEY}|" "$AWS_SECRET_PATH"
    sed -i "s|{{ S3_ENDPOINT__https://example.amazonaws.com/ }}|https://s3.us-west-1.amazonaws.com|" "$AWS_SECRET_PATH"
    sed -i "s|{{ S3_REGION__us-west-1 }}|us-west-1|" "$AWS_SECRET_PATH"
}
