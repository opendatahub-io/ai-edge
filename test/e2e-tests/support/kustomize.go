package support

import (
	"context"
	"encoding/json"
	"fmt"
	pipepinev1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
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

func CreateObjectsFromResourceMap(ctx context.Context, clients *Clients, resourceMap resmap.ResMap, namespace string) error {
	for _, rsc := range resourceMap.Resources() {
		kind := rsc.GetKind()
		switch kind {
		case "ConfigMap":
			{
				var configMap corev1.ConfigMap
				err := ResourceToType(rsc, &configMap)
				if err != nil {
					return err
				}

				yaml, err := rsc.AsYAML()
				if err != nil {
					return err
				}

				_, err = clients.Kubernetes.CoreV1().ConfigMaps(namespace).Patch(ctx, configMap.Name, types.ApplyPatchType, yaml, metav1.PatchOptions{
					FieldManager: "Apply",
				})
				if err != nil {
					return err
				}
			}
		case "Task":
			{
				var task pipepinev1.Task
				err := ResourceToType(rsc, &task)
				if err != nil {
					return err
				}

				_, err = clients.Task.Get(ctx, task.Name, metav1.GetOptions{})
				if err == nil {
					return nil
				}

				_, err = clients.Task.Create(ctx, &task, metav1.CreateOptions{})
				if err != nil {
					return err
				}
			}
		case "Pipeline":
			{
				var pipeline pipepinev1.Pipeline
				err := ResourceToType(rsc, &pipeline)
				if err != nil {
					return err
				}

				_, err = clients.Pipeline.Get(ctx, pipeline.Name, metav1.GetOptions{})
				if err == nil {
					return nil
				}

				_, err = clients.Pipeline.Create(ctx, &pipeline, metav1.CreateOptions{})
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
