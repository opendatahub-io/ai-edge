# \ModelRegistryServiceAPI

All URIs are relative to *https://localhost:8080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironmentInferenceService**](ModelRegistryServiceAPI.md#CreateEnvironmentInferenceService) | **Post** /api/model_registry/v1alpha1/serving_environments/{servingenvironmentId}/inference_services | Create a InferenceService in ServingEnvironment
[**CreateInferenceService**](ModelRegistryServiceAPI.md#CreateInferenceService) | **Post** /api/model_registry/v1alpha1/inference_services | Create a InferenceService
[**CreateInferenceServiceServe**](ModelRegistryServiceAPI.md#CreateInferenceServiceServe) | **Post** /api/model_registry/v1alpha1/inference_services/{inferenceserviceId}/serves | Create a ServeModel action in a InferenceService
[**CreateModelArtifact**](ModelRegistryServiceAPI.md#CreateModelArtifact) | **Post** /api/model_registry/v1alpha1/model_artifacts | Create a ModelArtifact
[**CreateModelVersion**](ModelRegistryServiceAPI.md#CreateModelVersion) | **Post** /api/model_registry/v1alpha1/model_versions | Create a ModelVersion
[**CreateModelVersionArtifact**](ModelRegistryServiceAPI.md#CreateModelVersionArtifact) | **Post** /api/model_registry/v1alpha1/model_versions/{modelversionId}/artifacts | Create an Artifact in a ModelVersion
[**CreateRegisteredModel**](ModelRegistryServiceAPI.md#CreateRegisteredModel) | **Post** /api/model_registry/v1alpha1/registered_models | Create a RegisteredModel
[**CreateRegisteredModelVersion**](ModelRegistryServiceAPI.md#CreateRegisteredModelVersion) | **Post** /api/model_registry/v1alpha1/registered_models/{registeredmodelId}/versions | Create a ModelVersion in RegisteredModel
[**CreateServingEnvironment**](ModelRegistryServiceAPI.md#CreateServingEnvironment) | **Post** /api/model_registry/v1alpha1/serving_environments | Create a ServingEnvironment
[**FindInferenceService**](ModelRegistryServiceAPI.md#FindInferenceService) | **Get** /api/model_registry/v1alpha1/inference_service | Get an InferenceServices that matches search parameters.
[**FindModelArtifact**](ModelRegistryServiceAPI.md#FindModelArtifact) | **Get** /api/model_registry/v1alpha1/model_artifact | Get a ModelArtifact that matches search parameters.
[**FindModelVersion**](ModelRegistryServiceAPI.md#FindModelVersion) | **Get** /api/model_registry/v1alpha1/model_version | Get a ModelVersion that matches search parameters.
[**FindRegisteredModel**](ModelRegistryServiceAPI.md#FindRegisteredModel) | **Get** /api/model_registry/v1alpha1/registered_model | Get a RegisteredModel that matches search parameters.
[**FindServingEnvironment**](ModelRegistryServiceAPI.md#FindServingEnvironment) | **Get** /api/model_registry/v1alpha1/serving_environment | Find ServingEnvironment
[**GetEnvironmentInferenceServices**](ModelRegistryServiceAPI.md#GetEnvironmentInferenceServices) | **Get** /api/model_registry/v1alpha1/serving_environments/{servingenvironmentId}/inference_services | List All ServingEnvironment&#39;s InferenceServices
[**GetInferenceService**](ModelRegistryServiceAPI.md#GetInferenceService) | **Get** /api/model_registry/v1alpha1/inference_services/{inferenceserviceId} | Get a InferenceService
[**GetInferenceServiceModel**](ModelRegistryServiceAPI.md#GetInferenceServiceModel) | **Get** /api/model_registry/v1alpha1/inference_services/{inferenceserviceId}/model | Get InferenceService&#39;s RegisteredModel
[**GetInferenceServiceServes**](ModelRegistryServiceAPI.md#GetInferenceServiceServes) | **Get** /api/model_registry/v1alpha1/inference_services/{inferenceserviceId}/serves | List All InferenceService&#39;s ServeModel actions
[**GetInferenceServiceVersion**](ModelRegistryServiceAPI.md#GetInferenceServiceVersion) | **Get** /api/model_registry/v1alpha1/inference_services/{inferenceserviceId}/version | Get InferenceService&#39;s ModelVersion
[**GetInferenceServices**](ModelRegistryServiceAPI.md#GetInferenceServices) | **Get** /api/model_registry/v1alpha1/inference_services | List All InferenceServices
[**GetModelArtifact**](ModelRegistryServiceAPI.md#GetModelArtifact) | **Get** /api/model_registry/v1alpha1/model_artifacts/{modelartifactId} | Get a ModelArtifact
[**GetModelArtifacts**](ModelRegistryServiceAPI.md#GetModelArtifacts) | **Get** /api/model_registry/v1alpha1/model_artifacts | List All ModelArtifacts
[**GetModelVersion**](ModelRegistryServiceAPI.md#GetModelVersion) | **Get** /api/model_registry/v1alpha1/model_versions/{modelversionId} | Get a ModelVersion
[**GetModelVersionArtifacts**](ModelRegistryServiceAPI.md#GetModelVersionArtifacts) | **Get** /api/model_registry/v1alpha1/model_versions/{modelversionId}/artifacts | List all artifacts associated with the &#x60;ModelVersion&#x60;
[**GetModelVersions**](ModelRegistryServiceAPI.md#GetModelVersions) | **Get** /api/model_registry/v1alpha1/model_versions | List All ModelVersions
[**GetRegisteredModel**](ModelRegistryServiceAPI.md#GetRegisteredModel) | **Get** /api/model_registry/v1alpha1/registered_models/{registeredmodelId} | Get a RegisteredModel
[**GetRegisteredModelVersions**](ModelRegistryServiceAPI.md#GetRegisteredModelVersions) | **Get** /api/model_registry/v1alpha1/registered_models/{registeredmodelId}/versions | List All RegisteredModel&#39;s ModelVersions
[**GetRegisteredModels**](ModelRegistryServiceAPI.md#GetRegisteredModels) | **Get** /api/model_registry/v1alpha1/registered_models | List All RegisteredModels
[**GetServingEnvironment**](ModelRegistryServiceAPI.md#GetServingEnvironment) | **Get** /api/model_registry/v1alpha1/serving_environments/{servingenvironmentId} | Get a ServingEnvironment
[**GetServingEnvironments**](ModelRegistryServiceAPI.md#GetServingEnvironments) | **Get** /api/model_registry/v1alpha1/serving_environments | List All ServingEnvironments
[**UpdateInferenceService**](ModelRegistryServiceAPI.md#UpdateInferenceService) | **Patch** /api/model_registry/v1alpha1/inference_services/{inferenceserviceId} | Update a InferenceService
[**UpdateModelArtifact**](ModelRegistryServiceAPI.md#UpdateModelArtifact) | **Patch** /api/model_registry/v1alpha1/model_artifacts/{modelartifactId} | Update a ModelArtifact
[**UpdateModelVersion**](ModelRegistryServiceAPI.md#UpdateModelVersion) | **Patch** /api/model_registry/v1alpha1/model_versions/{modelversionId} | Update a ModelVersion
[**UpdateRegisteredModel**](ModelRegistryServiceAPI.md#UpdateRegisteredModel) | **Patch** /api/model_registry/v1alpha1/registered_models/{registeredmodelId} | Update a RegisteredModel
[**UpdateServingEnvironment**](ModelRegistryServiceAPI.md#UpdateServingEnvironment) | **Patch** /api/model_registry/v1alpha1/serving_environments/{servingenvironmentId} | Update a ServingEnvironment



## CreateEnvironmentInferenceService

> InferenceService CreateEnvironmentInferenceService(ctx, servingenvironmentId).InferenceServiceCreate(inferenceServiceCreate).Execute()

Create a InferenceService in ServingEnvironment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	servingenvironmentId := "servingenvironmentId_example" // string | A unique identifier for a `ServingEnvironment`.
	inferenceServiceCreate := *openapiclient.NewInferenceServiceCreate("RegisteredModelId_example", "ServingEnvironmentId_example") // InferenceServiceCreate | A new `InferenceService` to be created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.CreateEnvironmentInferenceService(context.Background(), servingenvironmentId).InferenceServiceCreate(inferenceServiceCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.CreateEnvironmentInferenceService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEnvironmentInferenceService`: InferenceService
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.CreateEnvironmentInferenceService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servingenvironmentId** | **string** | A unique identifier for a &#x60;ServingEnvironment&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentInferenceServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inferenceServiceCreate** | [**InferenceServiceCreate**](InferenceServiceCreate.md) | A new &#x60;InferenceService&#x60; to be created. | 

### Return type

[**InferenceService**](InferenceService.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInferenceService

> InferenceService CreateInferenceService(ctx).InferenceServiceCreate(inferenceServiceCreate).Execute()

Create a InferenceService



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	inferenceServiceCreate := *openapiclient.NewInferenceServiceCreate("RegisteredModelId_example", "ServingEnvironmentId_example") // InferenceServiceCreate | A new `InferenceService` to be created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.CreateInferenceService(context.Background()).InferenceServiceCreate(inferenceServiceCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.CreateInferenceService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInferenceService`: InferenceService
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.CreateInferenceService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateInferenceServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **inferenceServiceCreate** | [**InferenceServiceCreate**](InferenceServiceCreate.md) | A new &#x60;InferenceService&#x60; to be created. | 

### Return type

[**InferenceService**](InferenceService.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateInferenceServiceServe

> ServeModel CreateInferenceServiceServe(ctx, inferenceserviceId).ServeModelCreate(serveModelCreate).Execute()

Create a ServeModel action in a InferenceService



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	inferenceserviceId := "inferenceserviceId_example" // string | A unique identifier for a `InferenceService`.
	serveModelCreate := *openapiclient.NewServeModelCreate("ModelVersionId_example") // ServeModelCreate | A new `ServeModel` to be associated with the `InferenceService`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.CreateInferenceServiceServe(context.Background(), inferenceserviceId).ServeModelCreate(serveModelCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.CreateInferenceServiceServe``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateInferenceServiceServe`: ServeModel
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.CreateInferenceServiceServe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inferenceserviceId** | **string** | A unique identifier for a &#x60;InferenceService&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInferenceServiceServeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **serveModelCreate** | [**ServeModelCreate**](ServeModelCreate.md) | A new &#x60;ServeModel&#x60; to be associated with the &#x60;InferenceService&#x60;. | 

### Return type

[**ServeModel**](ServeModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateModelArtifact

> ModelArtifact CreateModelArtifact(ctx).ModelArtifactCreate(modelArtifactCreate).Execute()

Create a ModelArtifact



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	modelArtifactCreate := *openapiclient.NewModelArtifactCreate() // ModelArtifactCreate | A new `ModelArtifact` to be created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.CreateModelArtifact(context.Background()).ModelArtifactCreate(modelArtifactCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.CreateModelArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateModelArtifact`: ModelArtifact
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.CreateModelArtifact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateModelArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelArtifactCreate** | [**ModelArtifactCreate**](ModelArtifactCreate.md) | A new &#x60;ModelArtifact&#x60; to be created. | 

### Return type

[**ModelArtifact**](ModelArtifact.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateModelVersion

> ModelVersion CreateModelVersion(ctx).ModelVersionCreate(modelVersionCreate).Execute()

Create a ModelVersion



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	modelVersionCreate := *openapiclient.NewModelVersionCreate("RegisteredModelID_example") // ModelVersionCreate | A new `ModelVersion` to be created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.CreateModelVersion(context.Background()).ModelVersionCreate(modelVersionCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.CreateModelVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateModelVersion`: ModelVersion
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.CreateModelVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateModelVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **modelVersionCreate** | [**ModelVersionCreate**](ModelVersionCreate.md) | A new &#x60;ModelVersion&#x60; to be created. | 

### Return type

[**ModelVersion**](ModelVersion.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateModelVersionArtifact

> Artifact CreateModelVersionArtifact(ctx, modelversionId).Artifact(artifact).Execute()

Create an Artifact in a ModelVersion



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	modelversionId := "modelversionId_example" // string | A unique identifier for a `ModelVersion`.
	artifact := openapiclient.Artifact{DocArtifact: openapiclient.NewDocArtifact("ArtifactType_example")} // Artifact | A new or existing `Artifact` to be associated with the `ModelVersion`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.CreateModelVersionArtifact(context.Background(), modelversionId).Artifact(artifact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.CreateModelVersionArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateModelVersionArtifact`: Artifact
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.CreateModelVersionArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelversionId** | **string** | A unique identifier for a &#x60;ModelVersion&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateModelVersionArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **artifact** | [**Artifact**](Artifact.md) | A new or existing &#x60;Artifact&#x60; to be associated with the &#x60;ModelVersion&#x60;. | 

### Return type

[**Artifact**](Artifact.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRegisteredModel

> RegisteredModel CreateRegisteredModel(ctx).RegisteredModelCreate(registeredModelCreate).Execute()

Create a RegisteredModel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	registeredModelCreate := *openapiclient.NewRegisteredModelCreate() // RegisteredModelCreate | A new `RegisteredModel` to be created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.CreateRegisteredModel(context.Background()).RegisteredModelCreate(registeredModelCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.CreateRegisteredModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRegisteredModel`: RegisteredModel
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.CreateRegisteredModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegisteredModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **registeredModelCreate** | [**RegisteredModelCreate**](RegisteredModelCreate.md) | A new &#x60;RegisteredModel&#x60; to be created. | 

### Return type

[**RegisteredModel**](RegisteredModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRegisteredModelVersion

> ModelVersion CreateRegisteredModelVersion(ctx, registeredmodelId).ModelVersion(modelVersion).Execute()

Create a ModelVersion in RegisteredModel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	registeredmodelId := "registeredmodelId_example" // string | A unique identifier for a `RegisteredModel`.
	modelVersion := *openapiclient.NewModelVersion("RegisteredModelID_example") // ModelVersion | A new `ModelVersion` to be created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.CreateRegisteredModelVersion(context.Background(), registeredmodelId).ModelVersion(modelVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.CreateRegisteredModelVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRegisteredModelVersion`: ModelVersion
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.CreateRegisteredModelVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registeredmodelId** | **string** | A unique identifier for a &#x60;RegisteredModel&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegisteredModelVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelVersion** | [**ModelVersion**](ModelVersion.md) | A new &#x60;ModelVersion&#x60; to be created. | 

### Return type

[**ModelVersion**](ModelVersion.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServingEnvironment

> ServingEnvironment CreateServingEnvironment(ctx).ServingEnvironmentCreate(servingEnvironmentCreate).Execute()

Create a ServingEnvironment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	servingEnvironmentCreate := *openapiclient.NewServingEnvironmentCreate() // ServingEnvironmentCreate | A new `ServingEnvironment` to be created.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.CreateServingEnvironment(context.Background()).ServingEnvironmentCreate(servingEnvironmentCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.CreateServingEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServingEnvironment`: ServingEnvironment
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.CreateServingEnvironment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServingEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **servingEnvironmentCreate** | [**ServingEnvironmentCreate**](ServingEnvironmentCreate.md) | A new &#x60;ServingEnvironment&#x60; to be created. | 

### Return type

[**ServingEnvironment**](ServingEnvironment.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindInferenceService

> InferenceService FindInferenceService(ctx).Name(name).ExternalID(externalID).ParentResourceID(parentResourceID).Execute()

Get an InferenceServices that matches search parameters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	name := "entity-name" // string | Name of entity to search. (optional)
	externalID := "10" // string | External ID of entity to search. (optional)
	parentResourceID := "10" // string | ID of the parent resource to use for search. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.FindInferenceService(context.Background()).Name(name).ExternalID(externalID).ParentResourceID(parentResourceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.FindInferenceService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindInferenceService`: InferenceService
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.FindInferenceService`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindInferenceServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name of entity to search. | 
 **externalID** | **string** | External ID of entity to search. | 
 **parentResourceID** | **string** | ID of the parent resource to use for search. | 

### Return type

[**InferenceService**](InferenceService.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindModelArtifact

> ModelArtifact FindModelArtifact(ctx).Name(name).ExternalID(externalID).ParentResourceID(parentResourceID).Execute()

Get a ModelArtifact that matches search parameters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	name := "entity-name" // string | Name of entity to search. (optional)
	externalID := "10" // string | External ID of entity to search. (optional)
	parentResourceID := "10" // string | ID of the parent resource to use for search. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.FindModelArtifact(context.Background()).Name(name).ExternalID(externalID).ParentResourceID(parentResourceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.FindModelArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindModelArtifact`: ModelArtifact
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.FindModelArtifact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindModelArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name of entity to search. | 
 **externalID** | **string** | External ID of entity to search. | 
 **parentResourceID** | **string** | ID of the parent resource to use for search. | 

### Return type

[**ModelArtifact**](ModelArtifact.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindModelVersion

> ModelVersion FindModelVersion(ctx).Name(name).ExternalID(externalID).ParentResourceID(parentResourceID).Execute()

Get a ModelVersion that matches search parameters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	name := "entity-name" // string | Name of entity to search. (optional)
	externalID := "10" // string | External ID of entity to search. (optional)
	parentResourceID := "10" // string | ID of the parent resource to use for search. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.FindModelVersion(context.Background()).Name(name).ExternalID(externalID).ParentResourceID(parentResourceID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.FindModelVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindModelVersion`: ModelVersion
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.FindModelVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindModelVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name of entity to search. | 
 **externalID** | **string** | External ID of entity to search. | 
 **parentResourceID** | **string** | ID of the parent resource to use for search. | 

### Return type

[**ModelVersion**](ModelVersion.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindRegisteredModel

> RegisteredModel FindRegisteredModel(ctx).Name(name).ExternalID(externalID).Execute()

Get a RegisteredModel that matches search parameters.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	name := "entity-name" // string | Name of entity to search. (optional)
	externalID := "10" // string | External ID of entity to search. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.FindRegisteredModel(context.Background()).Name(name).ExternalID(externalID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.FindRegisteredModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindRegisteredModel`: RegisteredModel
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.FindRegisteredModel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindRegisteredModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name of entity to search. | 
 **externalID** | **string** | External ID of entity to search. | 

### Return type

[**RegisteredModel**](RegisteredModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindServingEnvironment

> ServingEnvironment FindServingEnvironment(ctx).Name(name).ExternalID(externalID).Execute()

Find ServingEnvironment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	name := "entity-name" // string | Name of entity to search. (optional)
	externalID := "10" // string | External ID of entity to search. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.FindServingEnvironment(context.Background()).Name(name).ExternalID(externalID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.FindServingEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindServingEnvironment`: ServingEnvironment
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.FindServingEnvironment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindServingEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Name of entity to search. | 
 **externalID** | **string** | External ID of entity to search. | 

### Return type

[**ServingEnvironment**](ServingEnvironment.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentInferenceServices

> InferenceServiceList GetEnvironmentInferenceServices(ctx, servingenvironmentId).Name(name).ExternalID(externalID).PageSize(pageSize).OrderBy(orderBy).SortOrder(sortOrder).NextPageToken(nextPageToken).Execute()

List All ServingEnvironment's InferenceServices



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	servingenvironmentId := "servingenvironmentId_example" // string | A unique identifier for a `ServingEnvironment`.
	name := "entity-name" // string | Name of entity to search. (optional)
	externalID := "10" // string | External ID of entity to search. (optional)
	pageSize := "100" // string | Number of entities in each page. (optional)
	orderBy := openapiclient.OrderByField("CREATE_TIME") // OrderByField | Specifies the order by criteria for listing entities. (optional)
	sortOrder := openapiclient.SortOrder("ASC") // SortOrder | Specifies the sort order for listing entities, defaults to ASC. (optional)
	nextPageToken := "IkhlbGxvLCB3b3JsZC4i" // string | Token to use to retrieve next page of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.GetEnvironmentInferenceServices(context.Background(), servingenvironmentId).Name(name).ExternalID(externalID).PageSize(pageSize).OrderBy(orderBy).SortOrder(sortOrder).NextPageToken(nextPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.GetEnvironmentInferenceServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentInferenceServices`: InferenceServiceList
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.GetEnvironmentInferenceServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servingenvironmentId** | **string** | A unique identifier for a &#x60;ServingEnvironment&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentInferenceServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of entity to search. | 
 **externalID** | **string** | External ID of entity to search. | 
 **pageSize** | **string** | Number of entities in each page. | 
 **orderBy** | [**OrderByField**](OrderByField.md) | Specifies the order by criteria for listing entities. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | Specifies the sort order for listing entities, defaults to ASC. | 
 **nextPageToken** | **string** | Token to use to retrieve next page of results. | 

### Return type

[**InferenceServiceList**](InferenceServiceList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInferenceService

> InferenceService GetInferenceService(ctx, inferenceserviceId).Execute()

Get a InferenceService



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	inferenceserviceId := "inferenceserviceId_example" // string | A unique identifier for a `InferenceService`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.GetInferenceService(context.Background(), inferenceserviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.GetInferenceService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInferenceService`: InferenceService
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.GetInferenceService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inferenceserviceId** | **string** | A unique identifier for a &#x60;InferenceService&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInferenceServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**InferenceService**](InferenceService.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInferenceServiceModel

> RegisteredModel GetInferenceServiceModel(ctx, inferenceserviceId).Execute()

Get InferenceService's RegisteredModel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	inferenceserviceId := "inferenceserviceId_example" // string | A unique identifier for a `InferenceService`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.GetInferenceServiceModel(context.Background(), inferenceserviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.GetInferenceServiceModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInferenceServiceModel`: RegisteredModel
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.GetInferenceServiceModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inferenceserviceId** | **string** | A unique identifier for a &#x60;InferenceService&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInferenceServiceModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegisteredModel**](RegisteredModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInferenceServiceServes

> ServeModelList GetInferenceServiceServes(ctx, inferenceserviceId).Name(name).ExternalID(externalID).PageSize(pageSize).OrderBy(orderBy).SortOrder(sortOrder).NextPageToken(nextPageToken).Execute()

List All InferenceService's ServeModel actions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	inferenceserviceId := "inferenceserviceId_example" // string | A unique identifier for a `InferenceService`.
	name := "entity-name" // string | Name of entity to search. (optional)
	externalID := "10" // string | External ID of entity to search. (optional)
	pageSize := "100" // string | Number of entities in each page. (optional)
	orderBy := openapiclient.OrderByField("CREATE_TIME") // OrderByField | Specifies the order by criteria for listing entities. (optional)
	sortOrder := openapiclient.SortOrder("ASC") // SortOrder | Specifies the sort order for listing entities, defaults to ASC. (optional)
	nextPageToken := "IkhlbGxvLCB3b3JsZC4i" // string | Token to use to retrieve next page of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.GetInferenceServiceServes(context.Background(), inferenceserviceId).Name(name).ExternalID(externalID).PageSize(pageSize).OrderBy(orderBy).SortOrder(sortOrder).NextPageToken(nextPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.GetInferenceServiceServes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInferenceServiceServes`: ServeModelList
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.GetInferenceServiceServes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inferenceserviceId** | **string** | A unique identifier for a &#x60;InferenceService&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInferenceServiceServesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of entity to search. | 
 **externalID** | **string** | External ID of entity to search. | 
 **pageSize** | **string** | Number of entities in each page. | 
 **orderBy** | [**OrderByField**](OrderByField.md) | Specifies the order by criteria for listing entities. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | Specifies the sort order for listing entities, defaults to ASC. | 
 **nextPageToken** | **string** | Token to use to retrieve next page of results. | 

### Return type

[**ServeModelList**](ServeModelList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInferenceServiceVersion

> ModelVersion GetInferenceServiceVersion(ctx, inferenceserviceId).Execute()

Get InferenceService's ModelVersion



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	inferenceserviceId := "inferenceserviceId_example" // string | A unique identifier for a `InferenceService`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.GetInferenceServiceVersion(context.Background(), inferenceserviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.GetInferenceServiceVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInferenceServiceVersion`: ModelVersion
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.GetInferenceServiceVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inferenceserviceId** | **string** | A unique identifier for a &#x60;InferenceService&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInferenceServiceVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelVersion**](ModelVersion.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInferenceServices

> InferenceServiceList GetInferenceServices(ctx).PageSize(pageSize).OrderBy(orderBy).SortOrder(sortOrder).NextPageToken(nextPageToken).Execute()

List All InferenceServices



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	pageSize := "100" // string | Number of entities in each page. (optional)
	orderBy := openapiclient.OrderByField("CREATE_TIME") // OrderByField | Specifies the order by criteria for listing entities. (optional)
	sortOrder := openapiclient.SortOrder("ASC") // SortOrder | Specifies the sort order for listing entities, defaults to ASC. (optional)
	nextPageToken := "IkhlbGxvLCB3b3JsZC4i" // string | Token to use to retrieve next page of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.GetInferenceServices(context.Background()).PageSize(pageSize).OrderBy(orderBy).SortOrder(sortOrder).NextPageToken(nextPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.GetInferenceServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInferenceServices`: InferenceServiceList
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.GetInferenceServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetInferenceServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | Number of entities in each page. | 
 **orderBy** | [**OrderByField**](OrderByField.md) | Specifies the order by criteria for listing entities. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | Specifies the sort order for listing entities, defaults to ASC. | 
 **nextPageToken** | **string** | Token to use to retrieve next page of results. | 

### Return type

[**InferenceServiceList**](InferenceServiceList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelArtifact

> ModelArtifact GetModelArtifact(ctx, modelartifactId).Execute()

Get a ModelArtifact



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	modelartifactId := "modelartifactId_example" // string | A unique identifier for a `ModelArtifact`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.GetModelArtifact(context.Background(), modelartifactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.GetModelArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelArtifact`: ModelArtifact
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.GetModelArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelartifactId** | **string** | A unique identifier for a &#x60;ModelArtifact&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelArtifact**](ModelArtifact.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelArtifacts

> ModelArtifactList GetModelArtifacts(ctx).PageSize(pageSize).OrderBy(orderBy).SortOrder(sortOrder).NextPageToken(nextPageToken).Execute()

List All ModelArtifacts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	pageSize := "100" // string | Number of entities in each page. (optional)
	orderBy := openapiclient.OrderByField("CREATE_TIME") // OrderByField | Specifies the order by criteria for listing entities. (optional)
	sortOrder := openapiclient.SortOrder("ASC") // SortOrder | Specifies the sort order for listing entities, defaults to ASC. (optional)
	nextPageToken := "IkhlbGxvLCB3b3JsZC4i" // string | Token to use to retrieve next page of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.GetModelArtifacts(context.Background()).PageSize(pageSize).OrderBy(orderBy).SortOrder(sortOrder).NextPageToken(nextPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.GetModelArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelArtifacts`: ModelArtifactList
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.GetModelArtifacts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetModelArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | Number of entities in each page. | 
 **orderBy** | [**OrderByField**](OrderByField.md) | Specifies the order by criteria for listing entities. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | Specifies the sort order for listing entities, defaults to ASC. | 
 **nextPageToken** | **string** | Token to use to retrieve next page of results. | 

### Return type

[**ModelArtifactList**](ModelArtifactList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelVersion

> ModelVersion GetModelVersion(ctx, modelversionId).Execute()

Get a ModelVersion



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	modelversionId := "modelversionId_example" // string | A unique identifier for a `ModelVersion`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.GetModelVersion(context.Background(), modelversionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.GetModelVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelVersion`: ModelVersion
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.GetModelVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelversionId** | **string** | A unique identifier for a &#x60;ModelVersion&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelVersion**](ModelVersion.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelVersionArtifacts

> ArtifactList GetModelVersionArtifacts(ctx, modelversionId).Name(name).ExternalID(externalID).PageSize(pageSize).OrderBy(orderBy).SortOrder(sortOrder).NextPageToken(nextPageToken).Execute()

List all artifacts associated with the `ModelVersion`

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	modelversionId := "modelversionId_example" // string | A unique identifier for a `ModelVersion`.
	name := "entity-name" // string | Name of entity to search. (optional)
	externalID := "10" // string | External ID of entity to search. (optional)
	pageSize := "100" // string | Number of entities in each page. (optional)
	orderBy := openapiclient.OrderByField("CREATE_TIME") // OrderByField | Specifies the order by criteria for listing entities. (optional)
	sortOrder := openapiclient.SortOrder("ASC") // SortOrder | Specifies the sort order for listing entities, defaults to ASC. (optional)
	nextPageToken := "IkhlbGxvLCB3b3JsZC4i" // string | Token to use to retrieve next page of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.GetModelVersionArtifacts(context.Background(), modelversionId).Name(name).ExternalID(externalID).PageSize(pageSize).OrderBy(orderBy).SortOrder(sortOrder).NextPageToken(nextPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.GetModelVersionArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelVersionArtifacts`: ArtifactList
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.GetModelVersionArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelversionId** | **string** | A unique identifier for a &#x60;ModelVersion&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetModelVersionArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of entity to search. | 
 **externalID** | **string** | External ID of entity to search. | 
 **pageSize** | **string** | Number of entities in each page. | 
 **orderBy** | [**OrderByField**](OrderByField.md) | Specifies the order by criteria for listing entities. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | Specifies the sort order for listing entities, defaults to ASC. | 
 **nextPageToken** | **string** | Token to use to retrieve next page of results. | 

### Return type

[**ArtifactList**](ArtifactList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetModelVersions

> ModelVersionList GetModelVersions(ctx).PageSize(pageSize).OrderBy(orderBy).SortOrder(sortOrder).NextPageToken(nextPageToken).Execute()

List All ModelVersions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	pageSize := "100" // string | Number of entities in each page. (optional)
	orderBy := openapiclient.OrderByField("CREATE_TIME") // OrderByField | Specifies the order by criteria for listing entities. (optional)
	sortOrder := openapiclient.SortOrder("ASC") // SortOrder | Specifies the sort order for listing entities, defaults to ASC. (optional)
	nextPageToken := "IkhlbGxvLCB3b3JsZC4i" // string | Token to use to retrieve next page of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.GetModelVersions(context.Background()).PageSize(pageSize).OrderBy(orderBy).SortOrder(sortOrder).NextPageToken(nextPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.GetModelVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetModelVersions`: ModelVersionList
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.GetModelVersions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetModelVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | Number of entities in each page. | 
 **orderBy** | [**OrderByField**](OrderByField.md) | Specifies the order by criteria for listing entities. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | Specifies the sort order for listing entities, defaults to ASC. | 
 **nextPageToken** | **string** | Token to use to retrieve next page of results. | 

### Return type

[**ModelVersionList**](ModelVersionList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegisteredModel

> RegisteredModel GetRegisteredModel(ctx, registeredmodelId).Execute()

Get a RegisteredModel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	registeredmodelId := "registeredmodelId_example" // string | A unique identifier for a `RegisteredModel`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.GetRegisteredModel(context.Background(), registeredmodelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.GetRegisteredModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegisteredModel`: RegisteredModel
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.GetRegisteredModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registeredmodelId** | **string** | A unique identifier for a &#x60;RegisteredModel&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisteredModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegisteredModel**](RegisteredModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegisteredModelVersions

> ModelVersionList GetRegisteredModelVersions(ctx, registeredmodelId).Name(name).ExternalID(externalID).PageSize(pageSize).OrderBy(orderBy).SortOrder(sortOrder).NextPageToken(nextPageToken).Execute()

List All RegisteredModel's ModelVersions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	registeredmodelId := "registeredmodelId_example" // string | A unique identifier for a `RegisteredModel`.
	name := "entity-name" // string | Name of entity to search. (optional)
	externalID := "10" // string | External ID of entity to search. (optional)
	pageSize := "100" // string | Number of entities in each page. (optional)
	orderBy := openapiclient.OrderByField("CREATE_TIME") // OrderByField | Specifies the order by criteria for listing entities. (optional)
	sortOrder := openapiclient.SortOrder("ASC") // SortOrder | Specifies the sort order for listing entities, defaults to ASC. (optional)
	nextPageToken := "IkhlbGxvLCB3b3JsZC4i" // string | Token to use to retrieve next page of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.GetRegisteredModelVersions(context.Background(), registeredmodelId).Name(name).ExternalID(externalID).PageSize(pageSize).OrderBy(orderBy).SortOrder(sortOrder).NextPageToken(nextPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.GetRegisteredModelVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegisteredModelVersions`: ModelVersionList
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.GetRegisteredModelVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registeredmodelId** | **string** | A unique identifier for a &#x60;RegisteredModel&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisteredModelVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Name of entity to search. | 
 **externalID** | **string** | External ID of entity to search. | 
 **pageSize** | **string** | Number of entities in each page. | 
 **orderBy** | [**OrderByField**](OrderByField.md) | Specifies the order by criteria for listing entities. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | Specifies the sort order for listing entities, defaults to ASC. | 
 **nextPageToken** | **string** | Token to use to retrieve next page of results. | 

### Return type

[**ModelVersionList**](ModelVersionList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRegisteredModels

> RegisteredModelList GetRegisteredModels(ctx).PageSize(pageSize).OrderBy(orderBy).SortOrder(sortOrder).NextPageToken(nextPageToken).Execute()

List All RegisteredModels



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	pageSize := "100" // string | Number of entities in each page. (optional)
	orderBy := openapiclient.OrderByField("CREATE_TIME") // OrderByField | Specifies the order by criteria for listing entities. (optional)
	sortOrder := openapiclient.SortOrder("ASC") // SortOrder | Specifies the sort order for listing entities, defaults to ASC. (optional)
	nextPageToken := "IkhlbGxvLCB3b3JsZC4i" // string | Token to use to retrieve next page of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.GetRegisteredModels(context.Background()).PageSize(pageSize).OrderBy(orderBy).SortOrder(sortOrder).NextPageToken(nextPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.GetRegisteredModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegisteredModels`: RegisteredModelList
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.GetRegisteredModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisteredModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | Number of entities in each page. | 
 **orderBy** | [**OrderByField**](OrderByField.md) | Specifies the order by criteria for listing entities. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | Specifies the sort order for listing entities, defaults to ASC. | 
 **nextPageToken** | **string** | Token to use to retrieve next page of results. | 

### Return type

[**RegisteredModelList**](RegisteredModelList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServingEnvironment

> ServingEnvironment GetServingEnvironment(ctx, servingenvironmentId).Execute()

Get a ServingEnvironment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	servingenvironmentId := "servingenvironmentId_example" // string | A unique identifier for a `ServingEnvironment`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.GetServingEnvironment(context.Background(), servingenvironmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.GetServingEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServingEnvironment`: ServingEnvironment
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.GetServingEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servingenvironmentId** | **string** | A unique identifier for a &#x60;ServingEnvironment&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServingEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServingEnvironment**](ServingEnvironment.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServingEnvironments

> ServingEnvironmentList GetServingEnvironments(ctx).PageSize(pageSize).OrderBy(orderBy).SortOrder(sortOrder).NextPageToken(nextPageToken).Execute()

List All ServingEnvironments



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	pageSize := "100" // string | Number of entities in each page. (optional)
	orderBy := openapiclient.OrderByField("CREATE_TIME") // OrderByField | Specifies the order by criteria for listing entities. (optional)
	sortOrder := openapiclient.SortOrder("ASC") // SortOrder | Specifies the sort order for listing entities, defaults to ASC. (optional)
	nextPageToken := "IkhlbGxvLCB3b3JsZC4i" // string | Token to use to retrieve next page of results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.GetServingEnvironments(context.Background()).PageSize(pageSize).OrderBy(orderBy).SortOrder(sortOrder).NextPageToken(nextPageToken).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.GetServingEnvironments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetServingEnvironments`: ServingEnvironmentList
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.GetServingEnvironments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetServingEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageSize** | **string** | Number of entities in each page. | 
 **orderBy** | [**OrderByField**](OrderByField.md) | Specifies the order by criteria for listing entities. | 
 **sortOrder** | [**SortOrder**](SortOrder.md) | Specifies the sort order for listing entities, defaults to ASC. | 
 **nextPageToken** | **string** | Token to use to retrieve next page of results. | 

### Return type

[**ServingEnvironmentList**](ServingEnvironmentList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInferenceService

> InferenceService UpdateInferenceService(ctx, inferenceserviceId).InferenceServiceUpdate(inferenceServiceUpdate).Execute()

Update a InferenceService



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	inferenceserviceId := "inferenceserviceId_example" // string | A unique identifier for a `InferenceService`.
	inferenceServiceUpdate := *openapiclient.NewInferenceServiceUpdate() // InferenceServiceUpdate | Updated `InferenceService` information.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.UpdateInferenceService(context.Background(), inferenceserviceId).InferenceServiceUpdate(inferenceServiceUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.UpdateInferenceService``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateInferenceService`: InferenceService
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.UpdateInferenceService`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inferenceserviceId** | **string** | A unique identifier for a &#x60;InferenceService&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInferenceServiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inferenceServiceUpdate** | [**InferenceServiceUpdate**](InferenceServiceUpdate.md) | Updated &#x60;InferenceService&#x60; information. | 

### Return type

[**InferenceService**](InferenceService.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateModelArtifact

> ModelArtifact UpdateModelArtifact(ctx, modelartifactId).ModelArtifactUpdate(modelArtifactUpdate).Execute()

Update a ModelArtifact



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	modelartifactId := "modelartifactId_example" // string | A unique identifier for a `ModelArtifact`.
	modelArtifactUpdate := *openapiclient.NewModelArtifactUpdate() // ModelArtifactUpdate | Updated `ModelArtifact` information.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.UpdateModelArtifact(context.Background(), modelartifactId).ModelArtifactUpdate(modelArtifactUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.UpdateModelArtifact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateModelArtifact`: ModelArtifact
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.UpdateModelArtifact`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelartifactId** | **string** | A unique identifier for a &#x60;ModelArtifact&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateModelArtifactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelArtifactUpdate** | [**ModelArtifactUpdate**](ModelArtifactUpdate.md) | Updated &#x60;ModelArtifact&#x60; information. | 

### Return type

[**ModelArtifact**](ModelArtifact.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateModelVersion

> ModelVersion UpdateModelVersion(ctx, modelversionId).ModelVersion(modelVersion).Execute()

Update a ModelVersion



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	modelversionId := "modelversionId_example" // string | A unique identifier for a `ModelVersion`.
	modelVersion := *openapiclient.NewModelVersion("RegisteredModelID_example") // ModelVersion | Updated `ModelVersion` information.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.UpdateModelVersion(context.Background(), modelversionId).ModelVersion(modelVersion).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.UpdateModelVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateModelVersion`: ModelVersion
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.UpdateModelVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**modelversionId** | **string** | A unique identifier for a &#x60;ModelVersion&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateModelVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **modelVersion** | [**ModelVersion**](ModelVersion.md) | Updated &#x60;ModelVersion&#x60; information. | 

### Return type

[**ModelVersion**](ModelVersion.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRegisteredModel

> RegisteredModel UpdateRegisteredModel(ctx, registeredmodelId).RegisteredModelUpdate(registeredModelUpdate).Execute()

Update a RegisteredModel



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	registeredmodelId := "registeredmodelId_example" // string | A unique identifier for a `RegisteredModel`.
	registeredModelUpdate := *openapiclient.NewRegisteredModelUpdate() // RegisteredModelUpdate | Updated `RegisteredModel` information.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.UpdateRegisteredModel(context.Background(), registeredmodelId).RegisteredModelUpdate(registeredModelUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.UpdateRegisteredModel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRegisteredModel`: RegisteredModel
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.UpdateRegisteredModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registeredmodelId** | **string** | A unique identifier for a &#x60;RegisteredModel&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRegisteredModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **registeredModelUpdate** | [**RegisteredModelUpdate**](RegisteredModelUpdate.md) | Updated &#x60;RegisteredModel&#x60; information. | 

### Return type

[**RegisteredModel**](RegisteredModel.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateServingEnvironment

> ServingEnvironment UpdateServingEnvironment(ctx, servingenvironmentId).ServingEnvironmentUpdate(servingEnvironmentUpdate).Execute()

Update a ServingEnvironment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/opendatahub-io/ai-edge/cli/pkg/generated/model_registry_client"
)

func main() {
	servingenvironmentId := "servingenvironmentId_example" // string | A unique identifier for a `ServingEnvironment`.
	servingEnvironmentUpdate := *openapiclient.NewServingEnvironmentUpdate() // ServingEnvironmentUpdate | Updated `ServingEnvironment` information.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ModelRegistryServiceAPI.UpdateServingEnvironment(context.Background(), servingenvironmentId).ServingEnvironmentUpdate(servingEnvironmentUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ModelRegistryServiceAPI.UpdateServingEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateServingEnvironment`: ServingEnvironment
	fmt.Fprintf(os.Stdout, "Response from `ModelRegistryServiceAPI.UpdateServingEnvironment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**servingenvironmentId** | **string** | A unique identifier for a &#x60;ServingEnvironment&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateServingEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **servingEnvironmentUpdate** | [**ServingEnvironmentUpdate**](ServingEnvironmentUpdate.md) | Updated &#x60;ServingEnvironment&#x60; information. | 

### Return type

[**ServingEnvironment**](ServingEnvironment.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

