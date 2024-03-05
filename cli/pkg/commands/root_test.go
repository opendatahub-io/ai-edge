/*
Copyright 2023 KStreamer Authors

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

package commands

import (
	"bytes"
	"context"
	"testing"

	"github.com/spf13/cobra"
	"gotest.tools/v3/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

// TestRootCmd is a basic test that will be updated later once we have some real functionality to test
func TestRootCmd(t *testing.T) {
	// Create a buffer to hold the output
	var buf bytes.Buffer

	cobra.OnInitialize(
		func() {
			// Create a fake clientset
			clientSet = fake.NewSimpleClientset()

			// Add a couple of namespaces to simulate existing data
			ns1 := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "test-namespace-1"}}
			ns2 := &corev1.Namespace{ObjectMeta: metav1.ObjectMeta{Name: "test-namespace-2"}}
			clientSet.CoreV1().Namespaces().Create(context.TODO(), ns1, metav1.CreateOptions{})
			clientSet.CoreV1().Namespaces().Create(context.TODO(), ns2, metav1.CreateOptions{})
		},
	)

	rootCmd.SetOut(&buf)
	rootCmd.Execute()
	assert.Equal(t, "Namespace: test-namespace-1\nNamespace: test-namespace-2\n", buf.String())
}
