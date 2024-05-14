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

	GitConfig GitConfig `json:"git"`
	S3Config  S3Config  `json:"s3"`

	Clients    *Clients
	GitEnabled bool
	S3Enabled  bool
}

type GitConfig struct {
	Token          string `json:"TOKEN"`
	Username       string `json:"USERNAME"`
	Repo           string `json:"REPO"`
	ApiServer      string `json:"API_SERVER"`
	Branch         string `json:"BRANCH"`
	SelfSignedCert string `json:"SELF_SIGNED_CERT"`
}

type S3Config struct {
	AWSSecret      string `json:"AWS_SECRET"`
	AWSAccess      string `json:"AWS_ACCESS"`
	Region         string `json:"REGION"`
	Endpoint       string `json:"ENDPOINT"`
	BucketName     string `json:"BUCKET_NAME"`
	SelfSignedCert string `json:"SELF_SIGNED_CERT"`
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

	if config.S3Config != (S3Config{}) {
		config.S3Enabled = true
	}

	if config.GitConfig != (GitConfig{}) {
		config.GitEnabled = true
	}

	clients, err := CreateClients(config.Namespace)
	if err != nil {
		return nil, err
	}

	config.Clients = &clients

	return config, err
}
