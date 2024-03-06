# InferenceServiceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomProperties** | Pointer to [**map[string]MetadataValue**](MetadataValue.md) | User provided custom properties which are not defined by its type. | [optional] 
**Description** | Pointer to **string** | An optional description about the resource. | [optional] 
**ExternalID** | Pointer to **string** | The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance. | [optional] 
**ModelVersionId** | Pointer to **string** | ID of the &#x60;ModelVersion&#x60; to serve. If it&#39;s unspecified, then the latest &#x60;ModelVersion&#x60; by creation order will be served. | [optional] 
**Runtime** | Pointer to **string** | Model runtime. | [optional] 
**DesiredState** | Pointer to [**InferenceServiceState**](InferenceServiceState.md) |  | [optional] [default to INFERENCESERVICESTATE_DEPLOYED]

## Methods

### NewInferenceServiceUpdate

`func NewInferenceServiceUpdate() *InferenceServiceUpdate`

NewInferenceServiceUpdate instantiates a new InferenceServiceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInferenceServiceUpdateWithDefaults

`func NewInferenceServiceUpdateWithDefaults() *InferenceServiceUpdate`

NewInferenceServiceUpdateWithDefaults instantiates a new InferenceServiceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomProperties

`func (o *InferenceServiceUpdate) GetCustomProperties() map[string]MetadataValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *InferenceServiceUpdate) GetCustomPropertiesOk() (*map[string]MetadataValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *InferenceServiceUpdate) SetCustomProperties(v map[string]MetadataValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *InferenceServiceUpdate) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *InferenceServiceUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InferenceServiceUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InferenceServiceUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InferenceServiceUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *InferenceServiceUpdate) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *InferenceServiceUpdate) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *InferenceServiceUpdate) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *InferenceServiceUpdate) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetModelVersionId

`func (o *InferenceServiceUpdate) GetModelVersionId() string`

GetModelVersionId returns the ModelVersionId field if non-nil, zero value otherwise.

### GetModelVersionIdOk

`func (o *InferenceServiceUpdate) GetModelVersionIdOk() (*string, bool)`

GetModelVersionIdOk returns a tuple with the ModelVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersionId

`func (o *InferenceServiceUpdate) SetModelVersionId(v string)`

SetModelVersionId sets ModelVersionId field to given value.

### HasModelVersionId

`func (o *InferenceServiceUpdate) HasModelVersionId() bool`

HasModelVersionId returns a boolean if a field has been set.

### GetRuntime

`func (o *InferenceServiceUpdate) GetRuntime() string`

GetRuntime returns the Runtime field if non-nil, zero value otherwise.

### GetRuntimeOk

`func (o *InferenceServiceUpdate) GetRuntimeOk() (*string, bool)`

GetRuntimeOk returns a tuple with the Runtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntime

`func (o *InferenceServiceUpdate) SetRuntime(v string)`

SetRuntime sets Runtime field to given value.

### HasRuntime

`func (o *InferenceServiceUpdate) HasRuntime() bool`

HasRuntime returns a boolean if a field has been set.

### GetDesiredState

`func (o *InferenceServiceUpdate) GetDesiredState() InferenceServiceState`

GetDesiredState returns the DesiredState field if non-nil, zero value otherwise.

### GetDesiredStateOk

`func (o *InferenceServiceUpdate) GetDesiredStateOk() (*InferenceServiceState, bool)`

GetDesiredStateOk returns a tuple with the DesiredState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredState

`func (o *InferenceServiceUpdate) SetDesiredState(v InferenceServiceState)`

SetDesiredState sets DesiredState field to given value.

### HasDesiredState

`func (o *InferenceServiceUpdate) HasDesiredState() bool`

HasDesiredState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


