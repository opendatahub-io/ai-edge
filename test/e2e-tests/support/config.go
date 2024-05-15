package support

import (
	"encoding/json"
	"os"
)

const (
	S3CredentialsTemplatePath            = "../../../pipelines/tekton/aiedge-e2e/templates/credentials-s3.secret.yaml.template"
	ImageRegistryCredentialsTemplatePath = "../../../pipelines/tekton/aiedge-e2e/templates/credentials-image-registry.secret.yaml.template"
	GitCredentialsTemplatePath           = "../../../pipelines/tekton/gitops-update-pipeline/example-pipelineruns/example-git-credentials-secret.yaml.template"

	ConfigPath = "../config.json"
)

var (
	config *Config = nil
)

type Config struct {
	Namespace             string   `json:"NAMESPACE"`
	ImageRegistryUsername string   `json:"IMAGE_REGISTRY_USERNAME"`
	ImageRegistryPassword string   `json:"IMAGE_REGISTRY_PASSWORD"`
	TargetImageTags       []string `json:"TARGET_IMAGE_TAGS"`

	GitFetchConfig GitFetchConfig `json:"git-fetch"`
	S3FetchConfig  S3FetchConfig  `json:"s3-fetch"`
	GitOpsConfig   GitOpsConfig   `json:"gitops"`

	Clients         *Clients
	GitFetchEnabled bool
	S3FetchEnabled  bool
	GitOpsEnabled   bool
}

type GitFetchConfig struct {
	CONTAINERFILE_REPO      string `json:"CONTAINERFILE_REPO"`
	CONTAINERFILE_REVISION  string `json:"CONTAINERFILE_REVISION"`
	CONTAINER_RELATIVE_PATH string `json:"CONTAINER_RELATIVE_PATH"`
	MODEL_REPO              string `json:"MODEL_REPO"`
	MODEL_RELATIVE_PATH     string `json:"MODEL_RELATIVE_PATH"`
	MODEL_REVISION          string `json:"MODEL_REVISION"`
	MODEL_DIR               string `json:"MODEL_DIR"`
	SelfSignedCert          string `json:"SELF_SIGNED_CERT"`
}

type S3FetchConfig struct {
	AWSSecret      string `json:"AWS_SECRET"`
	AWSAccess      string `json:"AWS_ACCESS"`
	Region         string `json:"REGION"`
	Endpoint       string `json:"ENDPOINT"`
	BucketName     string `json:"BUCKET_NAME"`
	SelfSignedCert string `json:"SELF_SIGNED_CERT"`
}

type GitOpsConfig struct {
	Token     string `json:"TOKEN"`
	Username  string `json:"USERNAME"`
	Repo      string `json:"REPO"`
	ApiServer string `json:"API_SERVER"`
	Branch    string `json:"BRANCH"`
}

func GetConfig() (*Config, error) {
	if config != nil {
		return config, nil
	}

	config = &Config{}

	bytes, err := os.ReadFile(ConfigPath)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return config, err
	}

	if config.GitFetchConfig != (GitFetchConfig{}) {
		config.GitFetchEnabled = true
	}

	if config.S3FetchConfig != (S3FetchConfig{}) {
		config.S3FetchEnabled = true
	}

	if config.GitOpsConfig != (GitOpsConfig{}) {
		config.GitOpsEnabled = true
	}

	clients, err := CreateClients(config.Namespace)
	if err != nil {
		return nil, err
	}

	config.Clients = &clients

	return config, err
}
