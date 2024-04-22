/*
Copyright 2024. Open Data Hub Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package edgeclient

import (
	"context"
	"fmt"

	tektonv1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	tektonclientset "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/client-go/tools/clientcmd"
)

// CreatePipelineRun creates a tekton PipelineRun to build a model container image from a model version artifact.
func (c *Client) CreatePipelineRun(
	modelName string,
	modelVersion string,
	namespace string,
	parameters map[string]interface{},
) (*PipelineRunSummary, error) {
	var s3SecretName string
	var testDataConfigMapName string

	if s3s, ok := parameters["s3SecretName"]; !ok {
		return nil, fmt.Errorf("s3SecretName pipeline parameter is required")
	} else if s3s, ok := s3s.(string); !ok {
		return nil, fmt.Errorf("s3SecretName pipeline parameter must be a string")
	} else {
		s3SecretName = s3s
	}

	if tdc, ok := parameters["testDataConfigMapName"]; !ok {
		return nil, fmt.Errorf("testDataConfigMapName pipeline parameter is required")
	} else if tdc, ok := tdc.(string); !ok {
		return nil, fmt.Errorf("testDataConfigMapName pipeline parameter must be a string")
	} else {
		testDataConfigMapName = tdc
	}

	params, err := toTektonParams(modelName, modelVersion, parameters)
	if err != nil {
		return nil, fmt.Errorf("failed to convert parameters to tekton params: %w", err)
	}

	pipelineRun := newPipelineRunObject(modelName, modelVersion, namespace, params, s3SecretName, testDataConfigMapName)

	config, _ := clientcmd.BuildConfigFromFlags("", c.kubeconfig)
	tektonClient, _ := tektonclientset.NewForConfig(config)

	createdPipelineRun, err := tektonClient.TektonV1().PipelineRuns(namespace).Create(
		context.Background(), pipelineRun, metav1.CreateOptions{},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create pipeline run: %w", err)
	}
	return &PipelineRunSummary{
		Name:      createdPipelineRun.GetName(),
		Namespace: createdPipelineRun.GetObjectMeta().GetNamespace(),
	}, nil
}

func newPipelineRunObject(
	modelName string,
	modelVersion string,
	namespace string,
	params tektonv1.Params,
	s3SecretName string,
	testDataConfigMapName string,
) *tektonv1.PipelineRun {
	pipelineRun := &tektonv1.PipelineRun{
		ObjectMeta: metav1.ObjectMeta{
			Namespace:    namespace,
			GenerateName: "aiedge-e2e-" + modelName + "-",
			Labels: map[string]string{
				"tekton.dev/pipeline": "aiedge-e2e",
				"model-name":          modelName,
				"model-version":       modelVersion,
			},
		},
		Spec: tektonv1.PipelineRunSpec{
			TaskRunTemplate: tektonv1.PipelineTaskRunTemplate{
				ServiceAccountName: "pipeline",
			},
			Params: params,
			PipelineRef: &tektonv1.PipelineRef{
				Name: "aiedge-e2e",
			},
			Workspaces: []tektonv1.WorkspaceBinding{
				{
					Name: "build-workspace-pv",
					VolumeClaimTemplate: &corev1.PersistentVolumeClaim{
						Spec: corev1.PersistentVolumeClaimSpec{
							AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteOnce},
							Resources: corev1.VolumeResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceStorage: resource.MustParse("1Gi"),
								},
							},
						},
					},
				},
				{
					Name: "s3-secret",
					Secret: &corev1.SecretVolumeSource{
						SecretName: s3SecretName,
					},
				},
				{
					Name: "test-data",
					ConfigMap: &corev1.ConfigMapVolumeSource{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: testDataConfigMapName,
						},
					},
				},
			},
		},
	}
	return pipelineRun
}

// GetPipelineRuns returns a list of PipelineRuns for a given model version.
func (c *Client) GetPipelineRuns(namespace string, modelNames ...string) (*PipelineRunList, error) {
	if len(modelNames) == 0 || namespace == "" {
		return nil, fmt.Errorf("modelNames and namespace are required")
	}

	config, _ := clientcmd.BuildConfigFromFlags("", "/home/abd4lla/.kube/config")
	tektonClient, _ := tektonclientset.NewForConfig(config)

	modelReq, err := labels.NewRequirement("model-name", selection.In, modelNames)
	if err != nil {
		return nil, fmt.Errorf("failed to create label requirement: %w", err)
	}
	// versionReq, err := labels.NewRequirement("model-version", selection.Equals, []string{modelVersion})
	if err != nil {
		return nil, fmt.Errorf("failed to create label requirement: %w", err)
	}

	labelsSelector := labels.NewSelector().Add(*modelReq)

	tktnPipelineRunList, err := tektonClient.TektonV1().PipelineRuns(namespace).List(
		context.Background(), metav1.ListOptions{LabelSelector: labelsSelector.String()},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to list pipeline runs: %w", err)
	}
	pipelineRunsMap := make(map[string]map[string]*PipelineRunSummary)

	for _, pr := range tktnPipelineRunList.Items {
		model := pr.Labels["model-name"]
		version := pr.Labels["model-version"]
		if model == "" || version == "" {
			continue
		}
		if _, ok := pipelineRunsMap[model]; !ok {
			pipelineRunsMap[model] = make(map[string]*PipelineRunSummary)
			pipelineRunsMap[model][version] = &PipelineRunSummary{
				Name:           pr.GetName(),
				Namespace:      pr.GetObjectMeta().GetNamespace(),
				StartTime:      pr.Status.StartTime,
				CompletionTime: pr.Status.CompletionTime,
				Status:         fmt.Sprintf("%s", pr.Status.Conditions[0].Reason),
				Message:        fmt.Sprintf("%s", pr.Status.Conditions[0].Message),
			}
		} else {
			// {"m1":
			// 	"v1": PipelineRunSummary{nthoesnuah},
			// }
			if _, ok := pipelineRunsMap[model][version]; !ok {
				pipelineRunsMap[model][version] = &PipelineRunSummary{
					Name:           pr.GetName(),
					Namespace:      pr.GetObjectMeta().GetNamespace(),
					StartTime:      pr.Status.StartTime,
					CompletionTime: pr.Status.CompletionTime,
					Status:         fmt.Sprintf("%s", pr.Status.Conditions[0].Reason),
					Message:        fmt.Sprintf("%s", pr.Status.Conditions[0].Message),
				}
				// We want to keep the latest pipeline run
			} else if pipelineRunsMap[model][version].StartTime.Before(pr.Status.StartTime) {
				pipelineRunsMap[model][version] = &PipelineRunSummary{
					Name:           pr.GetName(),
					Namespace:      pr.GetObjectMeta().GetNamespace(),
					StartTime:      pr.Status.StartTime,
					CompletionTime: pr.Status.CompletionTime,
					Status:         fmt.Sprintf("%s", pr.Status.Conditions[0].Reason),
					Message:        fmt.Sprintf("%s", pr.Status.Conditions[0].Message),
				}
			}
		}
	}

	return &PipelineRunList{PipelineRuns: pipelineRunsMap}, nil
}

func toTektonParams(modelName, modelVersion string, parameters map[string]interface{}) (tektonv1.Params, error) {
	params := tektonv1.Params{
		{
			Name:  "model-name",
			Value: *tektonv1.NewStructuredValues(modelName),
		},
		{
			Name:  "model-version",
			Value: *tektonv1.NewStructuredValues(modelVersion),
		},
	}
	for k, v := range parameters {
		var pv *tektonv1.ParamValue
		if s, ok := v.(string); ok {
			pv = tektonv1.NewStructuredValues(s)
		} else if sv, ok := v.([]interface{}); ok {
			pv = &tektonv1.ParamValue{Type: tektonv1.ParamTypeArray}
			for _, s := range sv {
				if ss, ok := s.(string); ok {
					pv.ArrayVal = append(pv.ArrayVal, ss)
				} else {
					return nil, fmt.Errorf("paramater %s has unsupported type %T", k, s)
				}
			}
		} else if sv, ok := v.([]string); ok {
			pv = &tektonv1.ParamValue{Type: tektonv1.ParamTypeArray}
			for _, s := range sv {
				pv.ArrayVal = append(pv.ArrayVal, s)
			}
		} else {
			return nil, fmt.Errorf("paramater %s has unsupported type %T", k, v)
		}

		params = append(
			params, tektonv1.Param{Name: k, Value: *pv},
		)
	}
	return params, nil
}
