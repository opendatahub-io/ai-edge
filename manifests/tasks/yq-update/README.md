# `yq-update`

This task is used by the GitOps update Pipeline, to allow modifying files with yq

## Parameters
* **SCRIPT**: The yq script to execute. Can be multiple lines for complex tasks. `(Default: )`
* **image**: The yq image to use. `(Default: docker.io/mikefarah/yq:4.27.5@sha256:2be3626ed633fbe1fc33ee9343a1256a6be53334412b2251b9a859f8c145bb53)`
* **git-repo-path**: The path of the git repo directory inside the source workspace. SCRIPT will be run inside it. `(Default: )`
* **env-image-name**: The image name to be made available as the environment variable IMAGE_NAME within the task script. `(Default: )`
* **env-image-digest**: The image tag to be made available as the environment variable IMAGE_DIGEST within the task script. `(Default: )`

## Workspaces
* **source**: A workspace that contains the files which needs to be altered.

## Results
* **yq**: The result from your yq script. You can write to it using `$(results.yq.path)`.
