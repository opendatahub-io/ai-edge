# `test-model-rest-svc`

This task will send data to a testing container with the model inferencing endpoint and verify that expected output is returned. The `test-data` workspace is required to store the required test data used for the test

## Parameters
* **service-name**: The name of the service to be tested against
* **test-endpoint**: The endpoint of the service that will be tested against

## Workspaces
* **test-data**: A workspace that contains test data to be used. The files expected are data.json, the jsondata payload for your model, and output.json, the expected json output for that input payload

## Results
