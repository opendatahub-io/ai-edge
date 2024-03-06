# ModelVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RegisteredModelID** | **string** | ID of the &#x60;RegisteredModel&#x60; to which this version belongs. | 
**CustomProperties** | Pointer to [**map[string]MetadataValue**](MetadataValue.md) | User provided custom properties which are not defined by its type. | [optional] 
**Description** | Pointer to **string** | An optional description about the resource. | [optional] 
**ExternalID** | Pointer to **string** | The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance. | [optional] 
**Name** | Pointer to **string** | The client provided name of the artifact. This field is optional. If set, it must be unique among all the artifacts of the same artifact type within a database instance and cannot be changed once set. | [optional] 
**State** | Pointer to [**ModelVersionState**](ModelVersionState.md) |  | [optional] [default to MODELVERSIONSTATE_LIVE]
**Author** | Pointer to **string** | Name of the author. | [optional] 
**Id** | Pointer to **string** | Output only. The unique server generated id of the resource. | [optional] [readonly] 
**CreateTimeSinceEpoch** | Pointer to **string** | Output only. Create time of the resource in millisecond since epoch. | [optional] [readonly] 
**LastUpdateTimeSinceEpoch** | Pointer to **string** | Output only. Last update time of the resource since epoch in millisecond since epoch. | [optional] [readonly] 

## Methods

### NewModelVersion

`func NewModelVersion(registeredModelID string, ) *ModelVersion`

NewModelVersion instantiates a new ModelVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelVersionWithDefaults

`func NewModelVersionWithDefaults() *ModelVersion`

NewModelVersionWithDefaults instantiates a new ModelVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegisteredModelID

`func (o *ModelVersion) GetRegisteredModelID() string`

GetRegisteredModelID returns the RegisteredModelID field if non-nil, zero value otherwise.

### GetRegisteredModelIDOk

`func (o *ModelVersion) GetRegisteredModelIDOk() (*string, bool)`

GetRegisteredModelIDOk returns a tuple with the RegisteredModelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredModelID

`func (o *ModelVersion) SetRegisteredModelID(v string)`

SetRegisteredModelID sets RegisteredModelID field to given value.


### GetCustomProperties

`func (o *ModelVersion) GetCustomProperties() map[string]MetadataValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ModelVersion) GetCustomPropertiesOk() (*map[string]MetadataValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ModelVersion) SetCustomProperties(v map[string]MetadataValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ModelVersion) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *ModelVersion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelVersion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelVersion) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelVersion) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *ModelVersion) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *ModelVersion) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *ModelVersion) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *ModelVersion) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetName

`func (o *ModelVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelVersion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelVersion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *ModelVersion) GetState() ModelVersionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ModelVersion) GetStateOk() (*ModelVersionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ModelVersion) SetState(v ModelVersionState)`

SetState sets State field to given value.

### HasState

`func (o *ModelVersion) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAuthor

`func (o *ModelVersion) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ModelVersion) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ModelVersion) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ModelVersion) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetId

`func (o *ModelVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelVersion) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreateTimeSinceEpoch

`func (o *ModelVersion) GetCreateTimeSinceEpoch() string`

GetCreateTimeSinceEpoch returns the CreateTimeSinceEpoch field if non-nil, zero value otherwise.

### GetCreateTimeSinceEpochOk

`func (o *ModelVersion) GetCreateTimeSinceEpochOk() (*string, bool)`

GetCreateTimeSinceEpochOk returns a tuple with the CreateTimeSinceEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeSinceEpoch

`func (o *ModelVersion) SetCreateTimeSinceEpoch(v string)`

SetCreateTimeSinceEpoch sets CreateTimeSinceEpoch field to given value.

### HasCreateTimeSinceEpoch

`func (o *ModelVersion) HasCreateTimeSinceEpoch() bool`

HasCreateTimeSinceEpoch returns a boolean if a field has been set.

### GetLastUpdateTimeSinceEpoch

`func (o *ModelVersion) GetLastUpdateTimeSinceEpoch() string`

GetLastUpdateTimeSinceEpoch returns the LastUpdateTimeSinceEpoch field if non-nil, zero value otherwise.

### GetLastUpdateTimeSinceEpochOk

`func (o *ModelVersion) GetLastUpdateTimeSinceEpochOk() (*string, bool)`

GetLastUpdateTimeSinceEpochOk returns a tuple with the LastUpdateTimeSinceEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimeSinceEpoch

`func (o *ModelVersion) SetLastUpdateTimeSinceEpoch(v string)`

SetLastUpdateTimeSinceEpoch sets LastUpdateTimeSinceEpoch field to given value.

### HasLastUpdateTimeSinceEpoch

`func (o *ModelVersion) HasLastUpdateTimeSinceEpoch() bool`

HasLastUpdateTimeSinceEpoch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


