# InferenceServiceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomProperties** | Pointer to [**map[string]MetadataValue**](MetadataValue.md) | User provided custom properties which are not defined by its type. | [optional] 
**Description** | Pointer to **string** | An optional description about the resource. | [optional] 
**ExternalID** | Pointer to **string** | The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance. | [optional] 
**Name** | Pointer to **string** | The client provided name of the artifact. This field is optional. If set, it must be unique among all the artifacts of the same artifact type within a database instance and cannot be changed once set. | [optional] 
**ModelVersionId** | Pointer to **string** | ID of the &#x60;ModelVersion&#x60; to serve. If it&#39;s unspecified, then the latest &#x60;ModelVersion&#x60; by creation order will be served. | [optional] 
**Runtime** | Pointer to **string** | Model runtime. | [optional] 
**DesiredState** | Pointer to [**InferenceServiceState**](InferenceServiceState.md) |  | [optional] [default to INFERENCESERVICESTATE_DEPLOYED]
**RegisteredModelId** | **string** | ID of the &#x60;RegisteredModel&#x60; to serve. | 
**ServingEnvironmentId** | **string** | ID of the parent &#x60;ServingEnvironment&#x60; for this &#x60;InferenceService&#x60; entity. | 

## Methods

### NewInferenceServiceCreate

`func NewInferenceServiceCreate(registeredModelId string, servingEnvironmentId string, ) *InferenceServiceCreate`

NewInferenceServiceCreate instantiates a new InferenceServiceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInferenceServiceCreateWithDefaults

`func NewInferenceServiceCreateWithDefaults() *InferenceServiceCreate`

NewInferenceServiceCreateWithDefaults instantiates a new InferenceServiceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomProperties

`func (o *InferenceServiceCreate) GetCustomProperties() map[string]MetadataValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *InferenceServiceCreate) GetCustomPropertiesOk() (*map[string]MetadataValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *InferenceServiceCreate) SetCustomProperties(v map[string]MetadataValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *InferenceServiceCreate) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *InferenceServiceCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InferenceServiceCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InferenceServiceCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InferenceServiceCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *InferenceServiceCreate) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *InferenceServiceCreate) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *InferenceServiceCreate) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *InferenceServiceCreate) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetName

`func (o *InferenceServiceCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InferenceServiceCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InferenceServiceCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InferenceServiceCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModelVersionId

`func (o *InferenceServiceCreate) GetModelVersionId() string`

GetModelVersionId returns the ModelVersionId field if non-nil, zero value otherwise.

### GetModelVersionIdOk

`func (o *InferenceServiceCreate) GetModelVersionIdOk() (*string, bool)`

GetModelVersionIdOk returns a tuple with the ModelVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersionId

`func (o *InferenceServiceCreate) SetModelVersionId(v string)`

SetModelVersionId sets ModelVersionId field to given value.

### HasModelVersionId

`func (o *InferenceServiceCreate) HasModelVersionId() bool`

HasModelVersionId returns a boolean if a field has been set.

### GetRuntime

`func (o *InferenceServiceCreate) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *InferenceServiceCreate) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *InferenceServiceCreate) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *InferenceServiceCreate) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetDesiredState

`func (o *InferenceServiceCreate) GetDesiredState() InferenceServiceState`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *InferenceServiceCreate) GetDesiredStateOk() (*InferenceServiceState, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *InferenceServiceCreate) SetDesiredState(v InferenceServiceState)`

SetDesiredState sets DesiredState field to given value.

### HasDesiredState

`func (o *InferenceServiceCreate) HasDesiredState() bool`

HasDesiredState returns a boolean if a field has been set.

### GetRegisteredModelId

`func (o *InferenceServiceCreate) GetRegisteredModelId() string`

GetRegisteredModelId returns the RegisteredModelId field if non-nil, zero value otherwise.

### GetRegisteredModelIdOk

`func (o *InferenceServiceCreate) GetRegisteredModelIdOk() (*string, bool)`

GetRegisteredModelIdOk returns a tuple with the RegisteredModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredModelId

`func (o *InferenceServiceCreate) SetRegisteredModelId(v string)`

SetRegisteredModelId sets RegisteredModelId field to given value.


### GetServingEnvironmentId

`func (o *InferenceServiceCreate) GetServingEnvironmentId() string`

GetServingEnvironmentId returns the ServingEnvironmentId field if non-nil, zero value otherwise.

### GetServingEnvironmentIdOk

`func (o *InferenceServiceCreate) GetServingEnvironmentIdOk() (*string, bool)`

GetServingEnvironmentIdOk returns a tuple with the ServingEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingEnvironmentId

`func (o *InferenceServiceCreate) SetServingEnvironmentId(v string)`

SetServingEnvironmentId sets ServingEnvironmentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


