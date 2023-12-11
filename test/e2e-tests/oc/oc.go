package oc

import (
	v1 "k8s.io/api/core/v1"
	"os/exec"
)

type ApplyOptions struct {
	Filename  string
	Kustomize string
}

type CreateOptions struct {
	Filename string
}

func Project(namespace *v1.Namespace) error {
	cmd := exec.Command("oc", "project", namespace.Namespace)
	return cmd.Run()
}

func Apply(options ApplyOptions) error {
	optionsArgs := makeApplyArguments(&options)
	cmd := exec.Command("oc", optionsArgs...)
	return cmd.Run()
}

func makeApplyArguments(options *ApplyOptions) []string {
	args := []string{"apply"}

	if options.Filename != "" {
		args = append(args, "-f", options.Filename)
	}

	if options.Kustomize != "" {
		args = append(args, "-k", options.Kustomize)
	}

	return args
}

func Create(options CreateOptions) error {
	optionsArgs := makeCreateArguments(&options)
	cmd := exec.Command("oc", optionsArgs...)
	return cmd.Run()
}

func makeCreateArguments(options *CreateOptions) []string {
	args := []string{"create"}

	if options.Filename != "" {
		args = append(args, "-f", options.Filename)
	}

	return args
}
