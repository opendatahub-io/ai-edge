package support

import (
	"fmt"
	"os"
)

const (
	S3BucketNameEnvKey         = "S3_BUCKET"
	TargetImageNamespaceEnvKey = "TARGET_IMAGE_NAMESPACE"
)

var (
	options *Options = nil
)

type Options struct {
	S3BucketName     string // required
	RegistryRepoName string // required
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

	if options.S3BucketName = os.Getenv(S3BucketNameEnvKey); os.Getenv(S3BucketNameEnvKey) == "" {
		return options, fmt.Errorf("env variable %v not set, but is required to run tests", S3BucketNameEnvKey)
	}

	if options.RegistryRepoName = os.Getenv(TargetImageNamespaceEnvKey); os.Getenv(TargetImageNamespaceEnvKey) == "" {
		return options, fmt.Errorf("env variable %v not set, but is required to run tests", TargetImageNamespaceEnvKey)
	}

	return options, nil
}
