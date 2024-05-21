package support

import (
	"context"
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	applyconfigv1 "k8s.io/client-go/applyconfigurations/core/v1"
	"os"
	"sigs.k8s.io/yaml"
	"strings"
)

func RunSetup(ctx context.Context, config *Config) error {
	// image registry is always needed, read the file,
	// replace values and then apply
	{
		bytes, err := os.ReadFile(ImageRegistryCredentialsTemplatePath)
		if err != nil {
			return err
		}

		secret := applyconfigv1.SecretApplyConfiguration{}
		err = yaml.Unmarshal(bytes, &secret)
		if err != nil {
			return err
		}

		secret.StringData["username"] = config.ImageRegistryUsername
		secret.StringData["password"] = config.ImageRegistryPassword

		_, err = config.Clients.Kubernetes.CoreV1().Secrets(config.Namespace).Apply(ctx, &secret, metav1.ApplyOptions{FieldManager: "Apply"})
		if err != nil {
			return err
		}

		err = LinkSecretToServiceAccount(ctx, config, "pipeline", *secret.Name)
		if err != nil {
			return err
		}
	}

	// S3 config has been set, load the credential template file
	// and fill in the values in the config, then apply
	if config.S3FetchConfig.Enabled {
		bytes, err := os.ReadFile(S3CredentialsTemplatePath)
		if err != nil {
			return err
		}

		secret := applyconfigv1.SecretApplyConfiguration{}
		err = yaml.Unmarshal(bytes, &secret)
		if err != nil {
			return err
		}

		storageConfigString := secret.StringData["s3-storage-config"]
		storageConfigString = strings.ReplaceAll(storageConfigString, "{{ AWS_ACCESS_KEY_ID }}", config.S3FetchConfig.AWSAccess)
		storageConfigString = strings.ReplaceAll(storageConfigString, "{{ AWS_SECRET_ACCESS_KEY }}", config.S3FetchConfig.AWSSecret)
		storageConfigString = strings.ReplaceAll(storageConfigString, "{{ S3_ENDPOINT }}", config.S3FetchConfig.Endpoint)
		storageConfigString = strings.ReplaceAll(storageConfigString, "{{ S3_REGION }}", config.S3FetchConfig.Region)
		secret.StringData["s3-storage-config"] = storageConfigString

		_, err = config.Clients.Kubernetes.CoreV1().Secrets(config.Namespace).Apply(ctx, &secret, metav1.ApplyOptions{FieldManager: "Apply"})
		if err != nil {
			return err
		}
	}

	// Git config has been set, load the credential template file
	// and fill in the values in the config, then apply
	if config.GitFetchConfig.Enabled {
		bytes, err := os.ReadFile(GitCredentialsTemplatePath)
		if err != nil {
			return err
		}

		secret := applyconfigv1.SecretApplyConfiguration{}
		err = yaml.Unmarshal(bytes, &secret)
		if err != nil {
			return err
		}

		secret.StringData["token"] = config.GitOpsConfig.Token
		secret.StringData[".git-credentials"] = fmt.Sprintf("https://%v:%v@github.com", config.GitOpsConfig.Username, config.GitOpsConfig.Token)

		_, err = config.Clients.Kubernetes.CoreV1().Secrets(config.Namespace).Apply(ctx, &secret, metav1.ApplyOptions{FieldManager: "Apply"})
		if err != nil {
			return err
		}
	}

	return nil
}

func LinkSecretToServiceAccount(ctx context.Context, config *Config, serviceAccountName string, secretName string) error {
	serviceAccount, err := config.Clients.Kubernetes.CoreV1().ServiceAccounts(config.Namespace).Get(context.TODO(), serviceAccountName, metav1.GetOptions{})
	if err != nil {
		return err
	}

	for _, secret := range serviceAccount.Secrets {
		if secret.Name == secretName {
			return nil
		}
	}

	serviceAccount.Secrets = append(serviceAccount.Secrets, corev1.ObjectReference{Name: secretName})

	_, err = config.Clients.Kubernetes.CoreV1().ServiceAccounts(config.Namespace).Update(ctx, serviceAccount, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}
