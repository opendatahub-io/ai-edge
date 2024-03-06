# BaseArtifact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomProperties** | Pointer to [**map[string]MetadataValue**](MetadataValue.md) | User provided custom properties which are not defined by its type. | [optional] 
**Description** | Pointer to **string** | An optional description about the resource. | [optional] 
**ExternalID** | Pointer to **string** | The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance. | [optional] 
**Uri** | Pointer to **string** | The uniform resource identifier of the physical artifact. May be empty if there is no physical artifact. | [optional] 
**State** | Pointer to [**ArtifactState**](ArtifactState.md) |  | [optional] [default to ARTIFACTSTATE_UNKNOWN]
**Name** | Pointer to **string** | The client provided name of the artifact. This field is optional. If set, it must be unique among all the artifacts of the same artifact type within a database instance and cannot be changed once set. | [optional] 
**Id** | Pointer to **string** | Output only. The unique server generated id of the resource. | [optional] [readonly] 
**CreateTimeSinceEpoch** | Pointer to **string** | Output only. Create time of the resource in millisecond since epoch. | [optional] [readonly] 
**LastUpdateTimeSinceEpoch** | Pointer to **string** | Output only. Last update time of the resource since epoch in millisecond since epoch. | [optional] [readonly] 
**ArtifactType** | **string** |  | 

## Methods

### NewBaseArtifact

`func NewBaseArtifact(artifactType string, ) *BaseArtifact`

NewBaseArtifact instantiates a new BaseArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseArtifactWithDefaults

`func NewBaseArtifactWithDefaults() *BaseArtifact`

NewBaseArtifactWithDefaults instantiates a new BaseArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomProperties

`func (o *BaseArtifact) GetCustomProperties() map[string]MetadataValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *BaseArtifact) GetCustomPropertiesOk() (*map[string]MetadataValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *BaseArtifact) SetCustomProperties(v map[string]MetadataValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *BaseArtifact) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *BaseArtifact) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseArtifact) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseArtifact) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseArtifact) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *BaseArtifact) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *BaseArtifact) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *BaseArtifact) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *BaseArtifact) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetUri

`func (o *BaseArtifact) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *BaseArtifact) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *BaseArtifact) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *BaseArtifact) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetState

`func (o *BaseArtifact) GetState() ArtifactState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BaseArtifact) GetStateOk() (*ArtifactState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BaseArtifact) SetState(v ArtifactState)`

SetState sets State field to given value.

### HasState

`func (o *BaseArtifact) HasState() bool`

HasState returns a boolean if a field has been set.

### GetName

`func (o *BaseArtifact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseArtifact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseArtifact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseArtifact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *BaseArtifact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseArtifact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseArtifact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseArtifact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreateTimeSinceEpoch

`func (o *BaseArtifact) GetCreateTimeSinceEpoch() string`

GetCreateTimeSinceEpoch returns the CreateTimeSinceEpoch field if non-nil, zero value otherwise.

### GetCreateTimeSinceEpochOk

`func (o *BaseArtifact) GetCreateTimeSinceEpochOk() (*string, bool)`

GetCreateTimeSinceEpochOk returns a tuple with the CreateTimeSinceEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeSinceEpoch

`func (o *BaseArtifact) SetCreateTimeSinceEpoch(v string)`

SetCreateTimeSinceEpoch sets CreateTimeSinceEpoch field to given value.

### HasCreateTimeSinceEpoch

`func (o *BaseArtifact) HasCreateTimeSinceEpoch() bool`

HasCreateTimeSinceEpoch returns a boolean if a field has been set.

### GetLastUpdateTimeSinceEpoch

`func (o *BaseArtifact) GetLastUpdateTimeSinceEpoch() string`

GetLastUpdateTimeSinceEpoch returns the LastUpdateTimeSinceEpoch field if non-nil, zero value otherwise.

### GetLastUpdateTimeSinceEpochOk

`func (o *BaseArtifact) GetLastUpdateTimeSinceEpochOk() (*string, bool)`

GetLastUpdateTimeSinceEpochOk returns a tuple with the LastUpdateTimeSinceEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimeSinceEpoch

`func (o *BaseArtifact) SetLastUpdateTimeSinceEpoch(v string)`

SetLastUpdateTimeSinceEpoch sets LastUpdateTimeSinceEpoch field to given value.

### HasLastUpdateTimeSinceEpoch

`func (o *BaseArtifact) HasLastUpdateTimeSinceEpoch() bool`

HasLastUpdateTimeSinceEpoch returns a boolean if a field has been set.

### GetArtifactType

`func (o *BaseArtifact) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *BaseArtifact) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *BaseArtifact) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


