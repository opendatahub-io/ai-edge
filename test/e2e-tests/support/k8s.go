package support

import (
	"fmt"
	"github.com/opendatahub-io/ai-edge/test/e2e-tests/oc"
	"golang.org/x/net/context"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	TestingNamespaceName = "testing-namespace"
)

func CreateContext() context.Context {
	return context.Background()
}

func CreateTestNamespace(ctx context.Context, clientSet *kubernetes.Clientset) {
	namespaceConfig := v12.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: TestingNamespaceName,
		},
	}

	namespace, err := clientSet.CoreV1().Namespaces().Create(ctx, &namespaceConfig, metav1.CreateOptions{})
	if err != nil {
		panic(fmt.Errorf("cannot create testing namespace: %v", err.Error()))
	}

	err = oc.Project(namespace)
	if err != nil {
		panic(fmt.Errorf("cannot set %v: %v", namespace.Name, err.Error()))
	}
}

func DeleteTestNamespace(ctx context.Context, clientSet *kubernetes.Clientset) {
	err := clientSet.CoreV1().Namespaces().Delete(ctx, TestingNamespaceName, metav1.DeleteOptions{})
	if err != nil {
		panic(fmt.Errorf("cannot delete testing namespace: %v", err.Error()))
	}
}
