package support

import (
	"fmt"
	pipelinev1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	v1 "k8s.io/api/core/v1"
	"os"
	"sigs.k8s.io/yaml"
)

func MountConfigMapAsWorkspaceToPipelineRun(configMapName string, workspaceName string, pipelineRun *pipelinev1.PipelineRun) {
	pipelineRun.Spec.Workspaces = append(pipelineRun.Spec.Workspaces, pipelinev1.WorkspaceBinding{
		Name: workspaceName,
		ConfigMap: &v1.ConfigMapVolumeSource{
			LocalObjectReference: v1.LocalObjectReference{Name: configMapName},
		},
	})
}

func MountSecretAsWorkspaceToPipelineRun(secretName string, workspaceName string, pipelineRun *pipelinev1.PipelineRun) {
	pipelineRun.Spec.Workspaces = append(pipelineRun.Spec.Workspaces, pipelinev1.WorkspaceBinding{
		Name: workspaceName,
		Secret: &v1.SecretVolumeSource{
			SecretName: secretName,
		},
	})
}

func SetPipelineRunParam(name string, value pipelinev1.ParamValue, pipelineRun *pipelinev1.PipelineRun) {
	for index := range pipelineRun.Spec.Params {
		param := &pipelineRun.Spec.Params[index]
		if param.Name == name {
			param.Value = value
		}
	}
}

func NewStringParamValue(value string) pipelinev1.ParamValue {
	return pipelinev1.ParamValue{
		Type:      pipelinev1.ParamTypeString,
		StringVal: value,
	}
}

func NewArrayParamValue(value []string) pipelinev1.ParamValue {
	return pipelinev1.ParamValue{
		Type:     pipelinev1.ParamTypeArray,
		ArrayVal: value,
	}
}

func NewObjectParamValue(value map[string]string) pipelinev1.ParamValue {
	return pipelinev1.ParamValue{
		Type:      pipelinev1.ParamTypeObject,
		ObjectVal: value,
	}
}

func ReadFileAsPipelineRun(path string) (pipelinev1.PipelineRun, error) {
	var pipelineRun pipelinev1.PipelineRun

	bytes, err := os.ReadFile(path)
	if err != nil {
		return pipelineRun, err
	}

	err = yaml.Unmarshal(bytes, &pipelineRun)
	if err != nil {
		return pipelineRun, err
	}

	return pipelineRun, nil
}

func GetResultValueFromPipelineRun(resultName string, pipelineRun *pipelinev1.PipelineRun) (pipelinev1.ResultValue, error) {
	for _, result := range pipelineRun.Status.Results {
		if result.Name == resultName {
			return result.Value, nil
		}
	}

	return pipelinev1.ResultValue{}, fmt.Errorf("no result with name %v in pipeline run %v", resultName, pipelineRun.Name)
}
