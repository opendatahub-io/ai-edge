# `kserve-download-model`

This task is used to download a model folder from an S3 bucket. Credentials to allow the download of the model are stored in the `s3-secret` workspace.

## Parameters
* **model-name**: The name of the folder that contains the model files
* **s3-bucket-name**: The name of the S3 bucket to be downloaded from
* **model-relative-path**: The path from the root of the S3 bucket to the folder in which the model folder is located. Passing in an empty value means the model is stored at the root of the bucket

## Workspaces
* **workspace**: The workspace for the downloaded model
* **s3-secret**: The workspace containing the S3 credentials needed to download the model. A config map like this can be used
```json
{
  "type": "s3",
  "access_key_id": "ACCESSKEY",
  "secret_access_key": "SECRETKEY",
  "endpoint_url": "https://s3.us-west-2.amazonaws.com",
  "region": "us-east-1"
}
```
* **ssl-ca-directory** (optional): A workspace containing CA certificates, this will be used by the model download script to
verify the peer with when fetching over HTTPS.


## Results
* **s3-url**: The S3 URL used to download the model
