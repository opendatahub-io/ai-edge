# ModelArtifactCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomProperties** | Pointer to [**map[string]MetadataValue**](MetadataValue.md) | User provided custom properties which are not defined by its type. | [optional] 
**Description** | Pointer to **string** | An optional description about the resource. | [optional] 
**ExternalID** | Pointer to **string** | The external id that come from the clients’ system. This field is optional. If set, it must be unique among all resources within a database instance. | [optional] 
**Uri** | Pointer to **string** | The uniform resource identifier of the physical artifact. May be empty if there is no physical artifact. | [optional] 
**State** | Pointer to [**ArtifactState**](ArtifactState.md) |  | [optional] [default to ARTIFACTSTATE_UNKNOWN]
**Name** | Pointer to **string** | The client provided name of the artifact. This field is optional. If set, it must be unique among all the artifacts of the same artifact type within a database instance and cannot be changed once set. | [optional] 
**ModelFormatName** | Pointer to **string** | Name of the model format. | [optional] 
**StorageKey** | Pointer to **string** | Storage secret name. | [optional] 
**StoragePath** | Pointer to **string** | Path for model in storage provided by &#x60;storageKey&#x60;. | [optional] 
**ModelFormatVersion** | Pointer to **string** | Version of the model format. | [optional] 
**ServiceAccountName** | Pointer to **string** | Name of the service account with storage secret. | [optional] 

## Methods

### NewModelArtifactCreate

`func NewModelArtifactCreate() *ModelArtifactCreate`

NewModelArtifactCreate instantiates a new ModelArtifactCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelArtifactCreateWithDefaults

`func NewModelArtifactCreateWithDefaults() *ModelArtifactCreate`

NewModelArtifactCreateWithDefaults instantiates a new ModelArtifactCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomProperties

`func (o *ModelArtifactCreate) GetCustomProperties() map[string]MetadataValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ModelArtifactCreate) GetCustomPropertiesOk() (*map[string]MetadataValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ModelArtifactCreate) SetCustomProperties(v map[string]MetadataValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ModelArtifactCreate) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *ModelArtifactCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelArtifactCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelArtifactCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelArtifactCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *ModelArtifactCreate) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *ModelArtifactCreate) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *ModelArtifactCreate) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *ModelArtifactCreate) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetUri

`func (o *ModelArtifactCreate) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ModelArtifactCreate) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ModelArtifactCreate) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ModelArtifactCreate) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetState

`func (o *ModelArtifactCreate) GetState() ArtifactState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ModelArtifactCreate) GetStateOk() (*ArtifactState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ModelArtifactCreate) SetState(v ArtifactState)`

SetState sets State field to given value.

### HasState

`func (o *ModelArtifactCreate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetName

`func (o *ModelArtifactCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelArtifactCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelArtifactCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelArtifactCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetModelFormatName

`func (o *ModelArtifactCreate) GetModelFormatName() string`

GetModelFormatName returns the ModelFormatName field if non-nil, zero value otherwise.

### GetModelFormatNameOk

`func (o *ModelArtifactCreate) GetModelFormatNameOk() (*string, bool)`

GetModelFormatNameOk returns a tuple with the ModelFormatName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelFormatName

`func (o *ModelArtifactCreate) SetModelFormatName(v string)`

SetModelFormatName sets ModelFormatName field to given value.

### HasModelFormatName

`func (o *ModelArtifactCreate) HasModelFormatName() bool`

HasModelFormatName returns a boolean if a field has been set.

### GetStorageKey

`func (o *ModelArtifactCreate) GetStorageKey() string`

GetStorageKey returns the StorageKey field if non-nil, zero value otherwise.

### GetStorageKeyOk

`func (o *ModelArtifactCreate) GetStorageKeyOk() (*string, bool)`

GetStorageKeyOk returns a tuple with the StorageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageKey

`func (o *ModelArtifactCreate) SetStorageKey(v string)`

SetStorageKey sets StorageKey field to given value.

### HasStorageKey

`func (o *ModelArtifactCreate) HasStorageKey() bool`

HasStorageKey returns a boolean if a field has been set.

### GetStoragePath

`func (o *ModelArtifactCreate) GetStoragePath() string`

GetStoragePath returns the StoragePath field if non-nil, zero value otherwise.

### GetStoragePathOk

`func (o *ModelArtifactCreate) GetStoragePathOk() (*string, bool)`

GetStoragePathOk returns a tuple with the StoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePath

`func (o *ModelArtifactCreate) SetStoragePath(v string)`

SetStoragePath sets StoragePath field to given value.

### HasStoragePath

`func (o *ModelArtifactCreate) HasStoragePath() bool`

HasStoragePath returns a boolean if a field has been set.

### GetModelFormatVersion

`func (o *ModelArtifactCreate) GetModelFormatVersion() string`

GetModelFormatVersion returns the ModelFormatVersion field if non-nil, zero value otherwise.

### GetModelFormatVersionOk

`func (o *ModelArtifactCreate) GetModelFormatVersionOk() (*string, bool)`

GetModelFormatVersionOk returns a tuple with the ModelFormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelFormatVersion

`func (o *ModelArtifactCreate) SetModelFormatVersion(v string)`

SetModelFormatVersion sets ModelFormatVersion field to given value.

### HasModelFormatVersion

`func (o *ModelArtifactCreate) HasModelFormatVersion() bool`

HasModelFormatVersion returns a boolean if a field has been set.

### GetServiceAccountName

`func (o *ModelArtifactCreate) GetServiceAccountName() string`

GetServiceAccountName returns the ServiceAccountName field if non-nil, zero value otherwise.

### GetServiceAccountNameOk

`func (o *ModelArtifactCreate) GetServiceAccountNameOk() (*string, bool)`

GetServiceAccountNameOk returns a tuple with the ServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountName

`func (o *ModelArtifactCreate) SetServiceAccountName(v string)`

SetServiceAccountName sets ServiceAccountName field to given value.

### HasServiceAccountName

`func (o *ModelArtifactCreate) HasServiceAccountName() bool`

HasServiceAccountName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


