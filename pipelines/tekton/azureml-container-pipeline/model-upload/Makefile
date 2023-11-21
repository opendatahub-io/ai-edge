# Used for uploading local model files to the cluster through a 
# PVC. Use create to create the pod and PVC needed, then use
# upload model to upload a local file the PVC that the pod
# is mounted to

# Creates a PVC and a pod which mounts the PVC create
# SIZE - size of the storage used in the PVC created
# PVC - name of the PVC 
.PHONY: create
create:
	@oc process -f local-model-pvc-template.yaml -p PVC=${PVC} -p SIZE=${SIZE} | oc create -f -
	@oc process -f local-model-to-pvc-pod-template.yaml -p PVC=${PVC} | oc create -f -
	@oc get pod -w local-model-to-pvc-pod

# Uploads a file to a PVC   
# MODEL_PATH - local path to model file to upload 
# PVC - name of the PVC
# NAME - name of the file when uploaded to the pod
.PHONY: upload_model
upload_model:
	@oc exec local-model-to-pvc-pod -- mkdir workspace/${PVC}/model_dir/pipelines/models/ -p
	@oc cp ${MODEL_PATH} local-model-to-pvc-pod:/workspace/${PVC}/model_dir/pipelines/models/${NAME}

# deletes the pod created from the template
# PVC - name of he PVC that is mounted to the pod
.PHONY: clean
clean:
	@oc process -f local-model-to-pvc-pod.yaml -p PVC=${PVC} | oc delete -f -
