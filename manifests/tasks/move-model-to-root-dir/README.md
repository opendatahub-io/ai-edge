# `move-model-to-root-dir`

Used when using the git fetch method. This task moves the model folder which can be in any location in a cloned repo to the same directory as the git repo orignialy cloned. The cloned repo is then deleted from the workspace. This task can be used to match the behaviour of each fetch method to ensure the location of the model is consistant between tasks

## Parameters
* **model-name**: The name of the model folder that contians the model files
* **subdirectory**: The relative path from the workspace to the location of the git repo cloned
* **src-model-relative-path**: The relative path from the root of the git repo to the folder containing the model folder

## Workspaces
* **workspace**: The workspace for the downloaded model

## Results
