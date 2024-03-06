# BaseArtifactUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomProperties** | Pointer to [**map[string]MetadataValue**](MetadataValue.md) | User provided custom properties which are not defined by its type. | [optional] 
**Description** | Pointer to **string** | An optional description about the resource. | [optional] 
**ExternalID** | Pointer to **string** | The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance. | [optional] 
**Uri** | Pointer to **string** | The uniform resource identifier of the physical artifact. May be empty if there is no physical artifact. | [optional] 
**State** | Pointer to [**ArtifactState**](ArtifactState.md) |  | [optional] [default to ARTIFACTSTATE_UNKNOWN]

## Methods

### NewBaseArtifactUpdate

`func NewBaseArtifactUpdate() *BaseArtifactUpdate`

NewBaseArtifactUpdate instantiates a new BaseArtifactUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseArtifactUpdateWithDefaults

`func NewBaseArtifactUpdateWithDefaults() *BaseArtifactUpdate`

NewBaseArtifactUpdateWithDefaults instantiates a new BaseArtifactUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomProperties

`func (o *BaseArtifactUpdate) GetCustomProperties() map[string]MetadataValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *BaseArtifactUpdate) GetCustomPropertiesOk() (*map[string]MetadataValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *BaseArtifactUpdate) SetCustomProperties(v map[string]MetadataValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *BaseArtifactUpdate) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *BaseArtifactUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseArtifactUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseArtifactUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseArtifactUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *BaseArtifactUpdate) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *BaseArtifactUpdate) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *BaseArtifactUpdate) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *BaseArtifactUpdate) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetUri

`func (o *BaseArtifactUpdate) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *BaseArtifactUpdate) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *BaseArtifactUpdate) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *BaseArtifactUpdate) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetState

`func (o *BaseArtifactUpdate) GetState() ArtifactState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BaseArtifactUpdate) GetStateOk() (*ArtifactState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BaseArtifactUpdate) SetState(v ArtifactState)`

SetState sets State field to given value.

### HasState

`func (o *BaseArtifactUpdate) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


