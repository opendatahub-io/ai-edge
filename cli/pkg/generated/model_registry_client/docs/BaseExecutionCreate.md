# BaseExecutionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastKnownState** | Pointer to [**ExecutionState**](ExecutionState.md) |  | [optional] [default to EXECUTIONSTATE_UNKNOWN]
**CustomProperties** | Pointer to [**map[string]MetadataValue**](MetadataValue.md) | User provided custom properties which are not defined by its type. | [optional] 
**Description** | Pointer to **string** | An optional description about the resource. | [optional] 
**ExternalID** | Pointer to **string** | The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance. | [optional] 
**Name** | Pointer to **string** | The client provided name of the artifact. This field is optional. If set, it must be unique among all the artifacts of the same artifact type within a database instance and cannot be changed once set. | [optional] 

## Methods

### NewBaseExecutionCreate

`func NewBaseExecutionCreate() *BaseExecutionCreate`

NewBaseExecutionCreate instantiates a new BaseExecutionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseExecutionCreateWithDefaults

`func NewBaseExecutionCreateWithDefaults() *BaseExecutionCreate`

NewBaseExecutionCreateWithDefaults instantiates a new BaseExecutionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastKnownState

`func (o *BaseExecutionCreate) GetLastKnownState() ExecutionState`

GetLastKnownState returns the LastKnownState field if non-nil, zero value otherwise.

### GetLastKnownStateOk

`func (o *BaseExecutionCreate) GetLastKnownStateOk() (*ExecutionState, bool)`

GetLastKnownStateOk returns a tuple with the LastKnownState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKnownState

`func (o *BaseExecutionCreate) SetLastKnownState(v ExecutionState)`

SetLastKnownState sets LastKnownState field to given value.

### HasLastKnownState

`func (o *BaseExecutionCreate) HasLastKnownState() bool`

HasLastKnownState returns a boolean if a field has been set.

### GetCustomProperties

`func (o *BaseExecutionCreate) GetCustomProperties() map[string]MetadataValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *BaseExecutionCreate) GetCustomPropertiesOk() (*map[string]MetadataValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *BaseExecutionCreate) SetCustomProperties(v map[string]MetadataValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *BaseExecutionCreate) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *BaseExecutionCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseExecutionCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseExecutionCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseExecutionCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *BaseExecutionCreate) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *BaseExecutionCreate) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *BaseExecutionCreate) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *BaseExecutionCreate) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetName

`func (o *BaseExecutionCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseExecutionCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseExecutionCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseExecutionCreate) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


