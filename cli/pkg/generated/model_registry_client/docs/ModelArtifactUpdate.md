# ModelArtifactUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomProperties** | Pointer to [**map[string]MetadataValue**](MetadataValue.md) | User provided custom properties which are not defined by its type. | [optional] 
**Description** | Pointer to **string** | An optional description about the resource. | [optional] 
**ExternalID** | Pointer to **string** | The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance. | [optional] 
**Uri** | Pointer to **string** | The uniform resource identifier of the physical artifact. May be empty if there is no physical artifact. | [optional] 
**State** | Pointer to [**ArtifactState**](ArtifactState.md) |  | [optional] [default to ARTIFACTSTATE_UNKNOWN]
**ModelFormatName** | Pointer to **string** | Name of the model format. | [optional] 
**StorageKey** | Pointer to **string** | Storage secret name. | [optional] 
**StoragePath** | Pointer to **string** | Path for model in storage provided by &#x60;storageKey&#x60;. | [optional] 
**ModelFormatVersion** | Pointer to **string** | Version of the model format. | [optional] 
**ServiceAccountName** | Pointer to **string** | Name of the service account with storage secret. | [optional] 

## Methods

### NewModelArtifactUpdate

`func NewModelArtifactUpdate() *ModelArtifactUpdate`

NewModelArtifactUpdate instantiates a new ModelArtifactUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelArtifactUpdateWithDefaults

`func NewModelArtifactUpdateWithDefaults() *ModelArtifactUpdate`

NewModelArtifactUpdateWithDefaults instantiates a new ModelArtifactUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomProperties

`func (o *ModelArtifactUpdate) GetCustomProperties() map[string]MetadataValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ModelArtifactUpdate) GetCustomPropertiesOk() (*map[string]MetadataValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ModelArtifactUpdate) SetCustomProperties(v map[string]MetadataValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ModelArtifactUpdate) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *ModelArtifactUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelArtifactUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelArtifactUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelArtifactUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *ModelArtifactUpdate) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *ModelArtifactUpdate) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *ModelArtifactUpdate) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *ModelArtifactUpdate) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetUri

`func (o *ModelArtifactUpdate) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ModelArtifactUpdate) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ModelArtifactUpdate) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ModelArtifactUpdate) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetState

`func (o *ModelArtifactUpdate) GetState() ArtifactState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ModelArtifactUpdate) GetStateOk() (*ArtifactState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ModelArtifactUpdate) SetState(v ArtifactState)`

SetState sets State field to given value.

### HasState

`func (o *ModelArtifactUpdate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetModelFormatName

`func (o *ModelArtifactUpdate) GetModelFormatName() string`

GetModelFormatName returns the ModelFormatName field if non-nil, zero value otherwise.

### GetModelFormatNameOk

`func (o *ModelArtifactUpdate) GetModelFormatNameOk() (*string, bool)`

GetModelFormatNameOk returns a tuple with the ModelFormatName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelFormatName

`func (o *ModelArtifactUpdate) SetModelFormatName(v string)`

SetModelFormatName sets ModelFormatName field to given value.

### HasModelFormatName

`func (o *ModelArtifactUpdate) HasModelFormatName() bool`

HasModelFormatName returns a boolean if a field has been set.

### GetStorageKey

`func (o *ModelArtifactUpdate) GetStorageKey() string`

GetStorageKey returns the StorageKey field if non-nil, zero value otherwise.

### GetStorageKeyOk

`func (o *ModelArtifactUpdate) GetStorageKeyOk() (*string, bool)`

GetStorageKeyOk returns a tuple with the StorageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageKey

`func (o *ModelArtifactUpdate) SetStorageKey(v string)`

SetStorageKey sets StorageKey field to given value.

### HasStorageKey

`func (o *ModelArtifactUpdate) HasStorageKey() bool`

HasStorageKey returns a boolean if a field has been set.

### GetStoragePath

`func (o *ModelArtifactUpdate) GetStoragePath() string`

GetStoragePath returns the StoragePath field if non-nil, zero value otherwise.

### GetStoragePathOk

`func (o *ModelArtifactUpdate) GetStoragePathOk() (*string, bool)`

GetStoragePathOk returns a tuple with the StoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePath

`func (o *ModelArtifactUpdate) SetStoragePath(v string)`

SetStoragePath sets StoragePath field to given value.

### HasStoragePath

`func (o *ModelArtifactUpdate) HasStoragePath() bool`

HasStoragePath returns a boolean if a field has been set.

### GetModelFormatVersion

`func (o *ModelArtifactUpdate) GetModelFormatVersion() string`

GetModelFormatVersion returns the ModelFormatVersion field if non-nil, zero value otherwise.

### GetModelFormatVersionOk

`func (o *ModelArtifactUpdate) GetModelFormatVersionOk() (*string, bool)`

GetModelFormatVersionOk returns a tuple with the ModelFormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelFormatVersion

`func (o *ModelArtifactUpdate) SetModelFormatVersion(v string)`

SetModelFormatVersion sets ModelFormatVersion field to given value.

### HasModelFormatVersion

`func (o *ModelArtifactUpdate) HasModelFormatVersion() bool`

HasModelFormatVersion returns a boolean if a field has been set.

### GetServiceAccountName

`func (o *ModelArtifactUpdate) GetServiceAccountName() string`

GetServiceAccountName returns the ServiceAccountName field if non-nil, zero value otherwise.

### GetServiceAccountNameOk

`func (o *ModelArtifactUpdate) GetServiceAccountNameOk() (*string, bool)`

GetServiceAccountNameOk returns a tuple with the ServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountName

`func (o *ModelArtifactUpdate) SetServiceAccountName(v string)`

SetServiceAccountName sets ServiceAccountName field to given value.

### HasServiceAccountName

`func (o *ModelArtifactUpdate) HasServiceAccountName() bool`

HasServiceAccountName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


