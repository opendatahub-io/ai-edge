# InferenceService

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomProperties** | Pointer to [**map[string]MetadataValue**](MetadataValue.md) | User provided custom properties which are not defined by its type. | [optional] 
**Description** | Pointer to **string** | An optional description about the resource. | [optional] 
**ExternalID** | Pointer to **string** | The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance. | [optional] 
**Name** | Pointer to **string** | The client provided name of the artifact. This field is optional. If set, it must be unique among all the artifacts of the same artifact type within a database instance and cannot be changed once set. | [optional] 
**Id** | Pointer to **string** | Output only. The unique server generated id of the resource. | [optional] [readonly] 
**CreateTimeSinceEpoch** | Pointer to **string** | Output only. Create time of the resource in millisecond since epoch. | [optional] [readonly] 
**LastUpdateTimeSinceEpoch** | Pointer to **string** | Output only. Last update time of the resource since epoch in millisecond since epoch. | [optional] [readonly] 
**ModelVersionId** | Pointer to **string** | ID of the &#x60;ModelVersion&#x60; to serve. If it&#39;s unspecified, then the latest &#x60;ModelVersion&#x60; by creation order will be served. | [optional] 
**Runtime** | Pointer to **string** | Model runtime. | [optional] 
**DesiredState** | Pointer to [**InferenceServiceState**](InferenceServiceState.md) |  | [optional] [default to INFERENCESERVICESTATE_DEPLOYED]
**RegisteredModelId** | **string** | ID of the &#x60;RegisteredModel&#x60; to serve. | 
**ServingEnvironmentId** | **string** | ID of the parent &#x60;ServingEnvironment&#x60; for this &#x60;InferenceService&#x60; entity. | 

## Methods

### NewInferenceService

`func NewInferenceService(registeredModelId string, servingEnvironmentId string, ) *InferenceService`

NewInferenceService instantiates a new InferenceService object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInferenceServiceWithDefaults

`func NewInferenceServiceWithDefaults() *InferenceService`

NewInferenceServiceWithDefaults instantiates a new InferenceService object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomProperties

`func (o *InferenceService) GetCustomProperties() map[string]MetadataValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *InferenceService) GetCustomPropertiesOk() (*map[string]MetadataValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *InferenceService) SetCustomProperties(v map[string]MetadataValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *InferenceService) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *InferenceService) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InferenceService) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InferenceService) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InferenceService) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *InferenceService) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *InferenceService) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *InferenceService) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *InferenceService) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetName

`func (o *InferenceService) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InferenceService) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InferenceService) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InferenceService) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *InferenceService) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InferenceService) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InferenceService) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InferenceService) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreateTimeSinceEpoch

`func (o *InferenceService) GetCreateTimeSinceEpoch() string`

GetCreateTimeSinceEpoch returns the CreateTimeSinceEpoch field if non-nil, zero value otherwise.

### GetCreateTimeSinceEpochOk

`func (o *InferenceService) GetCreateTimeSinceEpochOk() (*string, bool)`

GetCreateTimeSinceEpochOk returns a tuple with the CreateTimeSinceEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeSinceEpoch

`func (o *InferenceService) SetCreateTimeSinceEpoch(v string)`

SetCreateTimeSinceEpoch sets CreateTimeSinceEpoch field to given value.

### HasCreateTimeSinceEpoch

`func (o *InferenceService) HasCreateTimeSinceEpoch() bool`

HasCreateTimeSinceEpoch returns a boolean if a field has been set.

### GetLastUpdateTimeSinceEpoch

`func (o *InferenceService) GetLastUpdateTimeSinceEpoch() string`

GetLastUpdateTimeSinceEpoch returns the LastUpdateTimeSinceEpoch field if non-nil, zero value otherwise.

### GetLastUpdateTimeSinceEpochOk

`func (o *InferenceService) GetLastUpdateTimeSinceEpochOk() (*string, bool)`

GetLastUpdateTimeSinceEpochOk returns a tuple with the LastUpdateTimeSinceEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimeSinceEpoch

`func (o *InferenceService) SetLastUpdateTimeSinceEpoch(v string)`

SetLastUpdateTimeSinceEpoch sets LastUpdateTimeSinceEpoch field to given value.

### HasLastUpdateTimeSinceEpoch

`func (o *InferenceService) HasLastUpdateTimeSinceEpoch() bool`

HasLastUpdateTimeSinceEpoch returns a boolean if a field has been set.

### GetModelVersionId

`func (o *InferenceService) GetModelVersionId() string`

GetModelVersionId returns the ModelVersionId field if non-nil, zero value otherwise.

### GetModelVersionIdOk

`func (o *InferenceService) GetModelVersionIdOk() (*string, bool)`

GetModelVersionIdOk returns a tuple with the ModelVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersionId

`func (o *InferenceService) SetModelVersionId(v string)`

SetModelVersionId sets ModelVersionId field to given value.

### HasModelVersionId

`func (o *InferenceService) HasModelVersionId() bool`

HasModelVersionId returns a boolean if a field has been set.

### GetRuntime

`func (o *InferenceService) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *InferenceService) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *InferenceService) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *InferenceService) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetDesiredState

`func (o *InferenceService) GetDesiredState() InferenceServiceState`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *InferenceService) GetDesiredStateOk() (*InferenceServiceState, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *InferenceService) SetDesiredState(v InferenceServiceState)`

SetDesiredState sets DesiredState field to given value.

### HasDesiredState

`func (o *InferenceService) HasDesiredState() bool`

HasDesiredState returns a boolean if a field has been set.

### GetRegisteredModelId

`func (o *InferenceService) GetRegisteredModelId() string`

GetRegisteredModelId returns the RegisteredModelId field if non-nil, zero value otherwise.

### GetRegisteredModelIdOk

`func (o *InferenceService) GetRegisteredModelIdOk() (*string, bool)`

GetRegisteredModelIdOk returns a tuple with the RegisteredModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredModelId

`func (o *InferenceService) SetRegisteredModelId(v string)`

SetRegisteredModelId sets RegisteredModelId field to given value.


### GetServingEnvironmentId

`func (o *InferenceService) GetServingEnvironmentId() string`

GetServingEnvironmentId returns the ServingEnvironmentId field if non-nil, zero value otherwise.

### GetServingEnvironmentIdOk

`func (o *InferenceService) GetServingEnvironmentIdOk() (*string, bool)`

GetServingEnvironmentIdOk returns a tuple with the ServingEnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServingEnvironmentId

`func (o *InferenceService) SetServingEnvironmentId(v string)`

SetServingEnvironmentId sets ServingEnvironmentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


