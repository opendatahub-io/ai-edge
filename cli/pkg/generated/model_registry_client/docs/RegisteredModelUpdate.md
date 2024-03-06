# RegisteredModelUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomProperties** | Pointer to [**map[string]MetadataValue**](MetadataValue.md) | User provided custom properties which are not defined by its type. | [optional] 
**Description** | Pointer to **string** | An optional description about the resource. | [optional] 
**ExternalID** | Pointer to **string** | The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance. | [optional] 
**State** | Pointer to [**RegisteredModelState**](RegisteredModelState.md) |  | [optional] [default to REGISTEREDMODELSTATE_LIVE]

## Methods

### NewRegisteredModelUpdate

`func NewRegisteredModelUpdate() *RegisteredModelUpdate`

NewRegisteredModelUpdate instantiates a new RegisteredModelUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisteredModelUpdateWithDefaults

`func NewRegisteredModelUpdateWithDefaults() *RegisteredModelUpdate`

NewRegisteredModelUpdateWithDefaults instantiates a new RegisteredModelUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomProperties

`func (o *RegisteredModelUpdate) GetCustomProperties() map[string]MetadataValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *RegisteredModelUpdate) GetCustomPropertiesOk() (*map[string]MetadataValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *RegisteredModelUpdate) SetCustomProperties(v map[string]MetadataValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *RegisteredModelUpdate) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *RegisteredModelUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RegisteredModelUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RegisteredModelUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RegisteredModelUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *RegisteredModelUpdate) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *RegisteredModelUpdate) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *RegisteredModelUpdate) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *RegisteredModelUpdate) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetState

`func (o *RegisteredModelUpdate) GetState() RegisteredModelState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RegisteredModelUpdate) GetStateOk() (*RegisteredModelState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RegisteredModelUpdate) SetState(v RegisteredModelState)`

SetState sets State field to given value.

### HasState

`func (o *RegisteredModelUpdate) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


