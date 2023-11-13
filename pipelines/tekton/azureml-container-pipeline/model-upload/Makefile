.PHONY: create
create:
	@oc process -f local-model-pvc.yaml -p PVC=${PVC} -p SIZE=${SIZE} | oc create -f -
	@oc process -f local-model-to-pvc-pod.yaml -p PVC=${PVC} | oc create -f -
	@oc get pod -w local-model-to-pvc-pod

.PHONY: upload_model
upload_model:
	@oc exec local-model-to-pvc-pod -- mkdir workspace/${PVC}/model_dir/pipelines/models/ -p
	@oc cp ${MODEL_PATH} local-model-to-pvc-pod:/workspace/${PVC}/model_dir/pipelines/models/${NAME}

.PHONY: clean
clean:
	@oc process -f local-model-to-pvc-pod.yaml -p PVC=${PVC} | oc delete -f -
