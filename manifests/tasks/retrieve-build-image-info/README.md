# `retrieve-build-image-info`

This task returns more detailed info about a model that has just been built and builds a url.txt file with all image tags to be pushed to

## Parameters
* **namespace**: The namespace where the model was built
* **model-name**: The name of the model
* **model-version**: The version of the model built
* **buildah-sha**: The built image digest
* **pipeline-run-uid**: The pipeline run id that was run to build the model
* **candidate-image-tag-reference**: The image tag references used when testing the image
* **target-image-tag-references**: The image tag references used for the final built image

## Workspaces
* **images-url**: workspace where url.txt file is created

## Results
* **model-name**: The name of the model
* **model-version**: The version of the model
* **image-size-bytes**: The size of the image in bytes
* **image-creation-time**: The date and time the image was created at
* **buildah-version**: The version of buildah used to build the image
* **image-digest-reference**: The fully qualified image digest reference of the image
* **target-image-tag-references**: The fully qualified image reference that the image was pushed to (e.g. registry.example.com/my-org/ai-model:1.0-1)
