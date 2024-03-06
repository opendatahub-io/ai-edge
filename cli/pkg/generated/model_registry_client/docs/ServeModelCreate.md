# ServeModelCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastKnownState** | Pointer to [**ExecutionState**](ExecutionState.md) |  | [optional] [default to EXECUTIONSTATE_UNKNOWN]
**CustomProperties** | Pointer to [**map[string]MetadataValue**](MetadataValue.md) | User provided custom properties which are not defined by its type. | [optional] 
**Description** | Pointer to **string** | An optional description about the resource. | [optional] 
**ExternalID** | Pointer to **string** | The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance. | [optional] 
**Name** | Pointer to **string** | The client provided name of the artifact. This field is optional. If set, it must be unique among all the artifacts of the same artifact type within a database instance and cannot be changed once set. | [optional] 
**ModelVersionId** | **string** | ID of the &#x60;ModelVersion&#x60; that was served in &#x60;InferenceService&#x60;. | 

## Methods

### NewServeModelCreate

`func NewServeModelCreate(modelVersionId string, ) *ServeModelCreate`

NewServeModelCreate instantiates a new ServeModelCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServeModelCreateWithDefaults

`func NewServeModelCreateWithDefaults() *ServeModelCreate`

NewServeModelCreateWithDefaults instantiates a new ServeModelCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastKnownState

`func (o *ServeModelCreate) GetLastKnownState() ExecutionState`

GetLastKnownState returns the LastKnownState field if non-nil, zero value otherwise.

### GetLastKnownStateOk

`func (o *ServeModelCreate) GetLastKnownStateOk() (*ExecutionState, bool)`

GetLastKnownStateOk returns a tuple with the LastKnownState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKnownState

`func (o *ServeModelCreate) SetLastKnownState(v ExecutionState)`

SetLastKnownState sets LastKnownState field to given value.

### HasLastKnownState

`func (o *ServeModelCreate) HasLastKnownState() bool`

HasLastKnownState returns a boolean if a field has been set.

### GetCustomProperties

`func (o *ServeModelCreate) GetCustomProperties() map[string]MetadataValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ServeModelCreate) GetCustomPropertiesOk() (*map[string]MetadataValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ServeModelCreate) SetCustomProperties(v map[string]MetadataValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ServeModelCreate) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *ServeModelCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServeModelCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServeModelCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServeModelCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *ServeModelCreate) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *ServeModelCreate) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *ServeModelCreate) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *ServeModelCreate) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetName

`func (o *ServeModelCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServeModelCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServeModelCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServeModelCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModelVersionId

`func (o *ServeModelCreate) GetModelVersionId() string`

GetModelVersionId returns the ModelVersionId field if non-nil, zero value otherwise.

### GetModelVersionIdOk

`func (o *ServeModelCreate) GetModelVersionIdOk() (*string, bool)`

GetModelVersionIdOk returns a tuple with the ModelVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersionId

`func (o *ServeModelCreate) SetModelVersionId(v string)`

SetModelVersionId sets ModelVersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


