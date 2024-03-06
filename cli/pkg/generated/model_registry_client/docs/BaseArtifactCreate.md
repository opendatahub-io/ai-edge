# BaseArtifactCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomProperties** | Pointer to [**map[string]MetadataValue**](MetadataValue.md) | User provided custom properties which are not defined by its type. | [optional] 
**Description** | Pointer to **string** | An optional description about the resource. | [optional] 
**ExternalID** | Pointer to **string** | The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance. | [optional] 
**Uri** | Pointer to **string** | The uniform resource identifier of the physical artifact. May be empty if there is no physical artifact. | [optional] 
**State** | Pointer to [**ArtifactState**](ArtifactState.md) |  | [optional] [default to ARTIFACTSTATE_UNKNOWN]
**Name** | Pointer to **string** | The client provided name of the artifact. This field is optional. If set, it must be unique among all the artifacts of the same artifact type within a database instance and cannot be changed once set. | [optional] 

## Methods

### NewBaseArtifactCreate

`func NewBaseArtifactCreate() *BaseArtifactCreate`

NewBaseArtifactCreate instantiates a new BaseArtifactCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseArtifactCreateWithDefaults

`func NewBaseArtifactCreateWithDefaults() *BaseArtifactCreate`

NewBaseArtifactCreateWithDefaults instantiates a new BaseArtifactCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomProperties

`func (o *BaseArtifactCreate) GetCustomProperties() map[string]MetadataValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *BaseArtifactCreate) GetCustomPropertiesOk() (*map[string]MetadataValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *BaseArtifactCreate) SetCustomProperties(v map[string]MetadataValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *BaseArtifactCreate) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *BaseArtifactCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseArtifactCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseArtifactCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseArtifactCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *BaseArtifactCreate) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *BaseArtifactCreate) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *BaseArtifactCreate) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *BaseArtifactCreate) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetUri

`func (o *BaseArtifactCreate) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *BaseArtifactCreate) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *BaseArtifactCreate) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *BaseArtifactCreate) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetState

`func (o *BaseArtifactCreate) GetState() ArtifactState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BaseArtifactCreate) GetStateOk() (*ArtifactState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BaseArtifactCreate) SetState(v ArtifactState)`

SetState sets State field to given value.

### HasState

`func (o *BaseArtifactCreate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetName

`func (o *BaseArtifactCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseArtifactCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseArtifactCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseArtifactCreate) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


