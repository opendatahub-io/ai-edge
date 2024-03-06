# BaseExecutionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomProperties** | Pointer to [**map[string]MetadataValue**](MetadataValue.md) | User provided custom properties which are not defined by its type. | [optional] 
**Description** | Pointer to **string** | An optional description about the resource. | [optional] 
**ExternalID** | Pointer to **string** | The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance. | [optional] 
**LastKnownState** | Pointer to [**ExecutionState**](ExecutionState.md) |  | [optional] [default to EXECUTIONSTATE_UNKNOWN]

## Methods

### NewBaseExecutionUpdate

`func NewBaseExecutionUpdate() *BaseExecutionUpdate`

NewBaseExecutionUpdate instantiates a new BaseExecutionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseExecutionUpdateWithDefaults

`func NewBaseExecutionUpdateWithDefaults() *BaseExecutionUpdate`

NewBaseExecutionUpdateWithDefaults instantiates a new BaseExecutionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomProperties

`func (o *BaseExecutionUpdate) GetCustomProperties() map[string]MetadataValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *BaseExecutionUpdate) GetCustomPropertiesOk() (*map[string]MetadataValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *BaseExecutionUpdate) SetCustomProperties(v map[string]MetadataValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *BaseExecutionUpdate) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *BaseExecutionUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseExecutionUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseExecutionUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseExecutionUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *BaseExecutionUpdate) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *BaseExecutionUpdate) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *BaseExecutionUpdate) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *BaseExecutionUpdate) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetLastKnownState

`func (o *BaseExecutionUpdate) GetLastKnownState() ExecutionState`

GetLastKnownState returns the LastKnownState field if non-nil, zero value otherwise.

### GetLastKnownStateOk

`func (o *BaseExecutionUpdate) GetLastKnownStateOk() (*ExecutionState, bool)`

GetLastKnownStateOk returns a tuple with the LastKnownState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKnownState

`func (o *BaseExecutionUpdate) SetLastKnownState(v ExecutionState)`

SetLastKnownState sets LastKnownState field to given value.

### HasLastKnownState

`func (o *BaseExecutionUpdate) HasLastKnownState() bool`

HasLastKnownState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


