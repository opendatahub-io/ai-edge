# Used for uploading local model files to the cluster through a 
# pod to a PVC. Use the create target to create the pod and PVC
# needed and upload the file

SIZE=1G
PVC=model-upload-pvc

# Creates a PVC and a pod which mounts the PVC created
# then uploads the file given
# SIZE - size of the storage used in the PVC created
# PVC - name of the PVC 
# MODEL_PATH - local path to model file to upload 
# NAME - name of the file when uploaded to the pod
.PHONY: create
create:
# create pod and PVC and wait for pod to be ready
	@oc process -f local-model-pvc-template.yaml -p PVC=${PVC} -p SIZE=${SIZE} | oc create -f -
	@oc process -f local-model-to-pvc-pod-template.yaml -p PVC=${PVC} | oc create -f -
	@oc wait pod local-model-to-pvc-pod --for condition=Ready=True --timeout=60s

# upload model to the pod to the PVC
	@oc exec local-model-to-pvc-pod -- mkdir workspace/${PVC}/model_dir/ -p
	@oc cp ${MODEL_PATH} local-model-to-pvc-pod:/workspace/${PVC}/model_dir-${NAME}/${NAME}

# delete pod
	@oc delete pod local-model-to-pvc-pod

	@echo "PVC name: ${PVC}" 
	@echo "Size: ${SIZE}" 
	@echo "Model path in pod: /workspace/${PVC}/model_dir-${NAME}/${NAME}" 
