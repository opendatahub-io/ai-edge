package support

import (
	"context"
	"encoding/json"
	"fmt"
	pipepinev1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/kustomize/api/krusty"
	"sigs.k8s.io/kustomize/api/resmap"
	"sigs.k8s.io/kustomize/api/resource"
	"sigs.k8s.io/kustomize/kyaml/filesys"
)

func ResourceToType[T any](resource *resource.Resource, t *T) error {
	bytes, err := resource.MarshalJSON()
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &t)
	if err != nil {
		return err
	}

	return nil
}

func KustomizeBuild(path string) (resmap.ResMap, error) {
	options := krusty.MakeDefaultOptions()
	k := krusty.MakeKustomizer(options)
	fs := filesys.FileSystemOrOnDisk{
		FileSystem: nil,
	}

	resourceMap, err := k.Run(fs, path)
	if err != nil {
		return nil, err
	}

	return resourceMap, nil
}

func CreateObjectsFromResourceMap(ctx context.Context, clients *Clients, resourceMap resmap.ResMap) error {
	for _, resrc := range resourceMap.Resources() {
		kind := resrc.GetKind()
		switch kind {
		case "ConfigMap":
			{
				var configMap corev1.ConfigMap
				err := ResourceToType(resrc, &configMap)
				if err != nil {
					return err
				}

				_, err = clients.Kubernetes.CoreV1().ConfigMaps("test-namespace").Create(ctx, &configMap, metav1.CreateOptions{})
				if err != nil {
					return err
				}
			}
		case "Task":
			{
				var task pipepinev1.Task
				err := ResourceToType(resrc, &task)
				if err != nil {
					return err
				}

				_, err = clients.Task.Create(ctx, &task, metav1.CreateOptions{})
				if err != nil {
					return err
				}
			}
		case "Pipeline":
			{
				var pipeline pipepinev1.Pipeline
				err := ResourceToType(resrc, &pipeline)
				if err != nil {
					return err
				}

				_, err = clients.Pipeline.Create(ctx, &pipeline, metav1.CreateOptions{})
				if err != nil {
					return err
				}
			}
		case "PipelineRun":
			{
				var pipelineRun pipepinev1.PipelineRun
				err := ResourceToType(resrc, &pipelineRun)
				if err != nil {
					return err
				}

				_, err = clients.PipelineRun.Create(ctx, &pipelineRun, metav1.CreateOptions{})
				if err != nil {
					return err
				}
			}
		default:
			{
				return fmt.Errorf("object kind '%v' cannot be created - not supported", kind)
			}
		}
	}

	return nil
}
