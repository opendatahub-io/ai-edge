# `check-model-and-containerfile-exists`

This Task can be used to check if the model files fetched exist and the containerfile cloned from git is also present

## Parameters
* **model-name**: The name of the model to be checked
* **containerfilePath**: A path from the root of the orignial git repo cloned to the containerfile to be checked

## Workspaces
* **workspace**: The workspace that contains the downloaded model

## Results
* **model-files-size**: Total size of the model files in MB
* **model-files-list**: Space separated list of files that are within the model folder
