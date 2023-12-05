package tests

import (
	"context"
	support "github.com/opendatahub-io/ai-edge/test/e2e-tests/support"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestGetAllPods(t *testing.T) {
	clientSet, err := support.ClusterClientSet()
	if err != nil {
		t.Fatal(err.Error())
	}

	_, err = clientSet.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}
}
