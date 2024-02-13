package support

import (
	"fmt"
	"os"
)

const (
	S3BucketNameEnvKey         = "S3_BUCKET"
	TargetImageNamespaceEnvKey = "TARGET_IMAGE_NAMESPACE"
	CLusterNamespaceEnvKey     = "NAMESPACE"
)

var (
	options *Options = nil
)

type Options struct {
	S3BucketName     string // required
	RegistryRepoName string // required
	ClusterNamespace string // required
}

func GetOptions() (*Options, error) {
	if options == nil {
		o, err := setOptions()
		if err != nil {
			return nil, err
		}

		options = o
	}

	return options, nil
}

func setOptions() (*Options, error) {
	if options != nil {
		return options, nil
	}

	var options = &Options{}

	if options.S3BucketName = os.Getenv(S3BucketNameEnvKey); options.S3BucketName == "" {
		return options, fmt.Errorf("env variable %v not set, but is required to run tests", S3BucketNameEnvKey)
	}

	if options.RegistryRepoName = os.Getenv(TargetImageNamespaceEnvKey); options.RegistryRepoName == "" {
		return options, fmt.Errorf("env variable %v not set, but is required to run tests", TargetImageNamespaceEnvKey)
	}

	if options.ClusterNamespace = os.Getenv(CLusterNamespaceEnvKey); options.ClusterNamespace == "" {
		return options, fmt.Errorf("env variable %v not set, but is required to run tests", CLusterNamespaceEnvKey)
	}

	return options, nil
}
