/*
Copyright © 2024 Open Data Hub Authors

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
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// clientSet is a kubernetes clientset that can be used to interact with the kubernetes API
var clientSet kubernetes.Interface
var kubeconfig string
var modelRegistryURL string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "odh-cli",
	Short: "Manage Open Data Hub resources from the command line.",
	Long: `Manage Open Data Hub resources from the command line.

This application is a tool to perform various operations on Open Data Hub.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Get the pods in the "default" namespace
		namespaces, err := clientSet.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
		if err != nil {
			panic(err.Error())
		}

		// Print the pod names
		for _, ns := range namespaces.Items {
			cmd.Printf("Namespace: %s\n", ns.Name)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match

	if kubeconfig = viper.GetString("KUBECONFIG"); kubeconfig == "" {
		kubeconfig = fmt.Sprintf("%s/.kube/config", os.Getenv("HOME"))
	}

	if modelRegistryURL = viper.GetString("MODEL_REGISTRY_URL"); modelRegistryURL == "" {
		modelRegistryURL = "http://localhost:8080"
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// Create a new clientset which includes all the API schemas
	clientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
}
