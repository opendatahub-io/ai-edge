# ServeModelUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastKnownState** | Pointer to [**ExecutionState**](ExecutionState.md) |  | [optional] [default to EXECUTIONSTATE_UNKNOWN]
**CustomProperties** | Pointer to [**map[string]MetadataValue**](MetadataValue.md) | User provided custom properties which are not defined by its type. | [optional] 
**Description** | Pointer to **string** | An optional description about the resource. | [optional] 
**ExternalID** | Pointer to **string** | The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance. | [optional] 

## Methods

### NewServeModelUpdate

`func NewServeModelUpdate() *ServeModelUpdate`

NewServeModelUpdate instantiates a new ServeModelUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServeModelUpdateWithDefaults

`func NewServeModelUpdateWithDefaults() *ServeModelUpdate`

NewServeModelUpdateWithDefaults instantiates a new ServeModelUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastKnownState

`func (o *ServeModelUpdate) GetLastKnownState() ExecutionState`

GetLastKnownState returns the LastKnownState field if non-nil, zero value otherwise.

### GetLastKnownStateOk

`func (o *ServeModelUpdate) GetLastKnownStateOk() (*ExecutionState, bool)`

GetLastKnownStateOk returns a tuple with the LastKnownState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKnownState

`func (o *ServeModelUpdate) SetLastKnownState(v ExecutionState)`

SetLastKnownState sets LastKnownState field to given value.

### HasLastKnownState

`func (o *ServeModelUpdate) HasLastKnownState() bool`

HasLastKnownState returns a boolean if a field has been set.

### GetCustomProperties

`func (o *ServeModelUpdate) GetCustomProperties() map[string]MetadataValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ServeModelUpdate) GetCustomPropertiesOk() (*map[string]MetadataValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ServeModelUpdate) SetCustomProperties(v map[string]MetadataValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ServeModelUpdate) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *ServeModelUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServeModelUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServeModelUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServeModelUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *ServeModelUpdate) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *ServeModelUpdate) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *ServeModelUpdate) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *ServeModelUpdate) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


