package support

import (
	"encoding/json"
	"os"
)

const (
	ConfigPath = "../config.json"
)

var (
	config *Config = nil
)

type Config struct {
	Namespace                string   `json:"namespace"`
	ImageRegistryUsername    string   `json:"image_registry_username"`
	ImageRegistryPassword    string   `json:"image_registry_password"`
	TargetImageTags          []string `json:"target_image_tags"`
	GitContainerFileRepo     string   `json:"git_container_file_repo"`
	GitContainerFileRevision string   `json:"git_container_file_revision"`
	ContainerRelativePath    string   `json:"container_relative_path"`

	GitFetchConfig GitFetchConfig `json:"git_fetch"`
	S3FetchConfig  S3FetchConfig  `json:"s3_fetch"`
	GitOpsConfig   GitOpsConfig   `json:"gitops"`

	Clients *Clients
}

type GitFetchConfig struct {
	Enabled           bool   `json:"enabled"`
	ModelRepo         string `json:"model_repo"`
	ModelRelativePath string `json:"model_relative_path"`
	ModelRevision     string `json:"model_revision"`
	ModelDir          string `json:"model_dir"`
	Username          string `json:"username"`
	Token             string `json:"token"`
	SelfSignedCert    string `json:"self_signed_cert"`
}

type S3FetchConfig struct {
	Enabled        bool   `json:"enabled"`
	AWSSecret      string `json:"aws_secret"`
	AWSAccess      string `json:"aws_access"`
	Region         string `json:"region"`
	Endpoint       string `json:"endpoint"`
	BucketName     string `json:"bucket_name"`
	SelfSignedCert string `json:"self_signed_cert"`
}

type GitOpsConfig struct {
	Enabled   bool   `json:"enabled"`
	Token     string `json:"token"`
	Username  string `json:"username"`
	Repo      string `json:"repo"`
	ApiServer string `json:"api_server"`
	Branch    string `json:"branch"`
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

	clients, err := CreateClients(config.Namespace)
	if err != nil {
		return nil, err
	}

	config.Clients = &clients

	return config, err
}
