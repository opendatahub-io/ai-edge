# RegisteredModel

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
**State** | Pointer to [**RegisteredModelState**](RegisteredModelState.md) |  | [optional] [default to REGISTEREDMODELSTATE_LIVE]

## Methods

### NewRegisteredModel

`func NewRegisteredModel() *RegisteredModel`

NewRegisteredModel instantiates a new RegisteredModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegisteredModelWithDefaults

`func NewRegisteredModelWithDefaults() *RegisteredModel`

NewRegisteredModelWithDefaults instantiates a new RegisteredModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomProperties

`func (o *RegisteredModel) GetCustomProperties() map[string]MetadataValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *RegisteredModel) GetCustomPropertiesOk() (*map[string]MetadataValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *RegisteredModel) SetCustomProperties(v map[string]MetadataValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *RegisteredModel) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *RegisteredModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RegisteredModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RegisteredModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RegisteredModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *RegisteredModel) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *RegisteredModel) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *RegisteredModel) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *RegisteredModel) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetName

`func (o *RegisteredModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RegisteredModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RegisteredModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RegisteredModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *RegisteredModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RegisteredModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RegisteredModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RegisteredModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreateTimeSinceEpoch

`func (o *RegisteredModel) GetCreateTimeSinceEpoch() string`

GetCreateTimeSinceEpoch returns the CreateTimeSinceEpoch field if non-nil, zero value otherwise.

### GetCreateTimeSinceEpochOk

`func (o *RegisteredModel) GetCreateTimeSinceEpochOk() (*string, bool)`

GetCreateTimeSinceEpochOk returns a tuple with the CreateTimeSinceEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeSinceEpoch

`func (o *RegisteredModel) SetCreateTimeSinceEpoch(v string)`

SetCreateTimeSinceEpoch sets CreateTimeSinceEpoch field to given value.

### HasCreateTimeSinceEpoch

`func (o *RegisteredModel) HasCreateTimeSinceEpoch() bool`

HasCreateTimeSinceEpoch returns a boolean if a field has been set.

### GetLastUpdateTimeSinceEpoch

`func (o *RegisteredModel) GetLastUpdateTimeSinceEpoch() string`

GetLastUpdateTimeSinceEpoch returns the LastUpdateTimeSinceEpoch field if non-nil, zero value otherwise.

### GetLastUpdateTimeSinceEpochOk

`func (o *RegisteredModel) GetLastUpdateTimeSinceEpochOk() (*string, bool)`

GetLastUpdateTimeSinceEpochOk returns a tuple with the LastUpdateTimeSinceEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimeSinceEpoch

`func (o *RegisteredModel) SetLastUpdateTimeSinceEpoch(v string)`

SetLastUpdateTimeSinceEpoch sets LastUpdateTimeSinceEpoch field to given value.

### HasLastUpdateTimeSinceEpoch

`func (o *RegisteredModel) HasLastUpdateTimeSinceEpoch() bool`

HasLastUpdateTimeSinceEpoch returns a boolean if a field has been set.

### GetState

`func (o *RegisteredModel) GetState() RegisteredModelState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RegisteredModel) GetStateOk() (*RegisteredModelState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RegisteredModel) SetState(v RegisteredModelState)`

SetState sets State field to given value.

### HasState

`func (o *RegisteredModel) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


