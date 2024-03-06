# ModelArtifact

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
**ArtifactType** | **string** |  | [default to "model-artifact"]
**ModelFormatName** | Pointer to **string** | Name of the model format. | [optional] 
**StorageKey** | Pointer to **string** | Storage secret name. | [optional] 
**StoragePath** | Pointer to **string** | Path for model in storage provided by &#x60;storageKey&#x60;. | [optional] 
**ModelFormatVersion** | Pointer to **string** | Version of the model format. | [optional] 
**ServiceAccountName** | Pointer to **string** | Name of the service account with storage secret. | [optional] 

## Methods

### NewModelArtifact

`func NewModelArtifact(artifactType string, ) *ModelArtifact`

NewModelArtifact instantiates a new ModelArtifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelArtifactWithDefaults

`func NewModelArtifactWithDefaults() *ModelArtifact`

NewModelArtifactWithDefaults instantiates a new ModelArtifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomProperties

`func (o *ModelArtifact) GetCustomProperties() map[string]MetadataValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ModelArtifact) GetCustomPropertiesOk() (*map[string]MetadataValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ModelArtifact) SetCustomProperties(v map[string]MetadataValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ModelArtifact) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *ModelArtifact) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelArtifact) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelArtifact) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelArtifact) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *ModelArtifact) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *ModelArtifact) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *ModelArtifact) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *ModelArtifact) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetUri

`func (o *ModelArtifact) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ModelArtifact) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ModelArtifact) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ModelArtifact) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetState

`func (o *ModelArtifact) GetState() ArtifactState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ModelArtifact) GetStateOk() (*ArtifactState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ModelArtifact) SetState(v ArtifactState)`

SetState sets State field to given value.

### HasState

`func (o *ModelArtifact) HasState() bool`

HasState returns a boolean if a field has been set.

### GetName

`func (o *ModelArtifact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelArtifact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelArtifact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelArtifact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ModelArtifact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelArtifact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelArtifact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelArtifact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreateTimeSinceEpoch

`func (o *ModelArtifact) GetCreateTimeSinceEpoch() string`

GetCreateTimeSinceEpoch returns the CreateTimeSinceEpoch field if non-nil, zero value otherwise.

### GetCreateTimeSinceEpochOk

`func (o *ModelArtifact) GetCreateTimeSinceEpochOk() (*string, bool)`

GetCreateTimeSinceEpochOk returns a tuple with the CreateTimeSinceEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeSinceEpoch

`func (o *ModelArtifact) SetCreateTimeSinceEpoch(v string)`

SetCreateTimeSinceEpoch sets CreateTimeSinceEpoch field to given value.

### HasCreateTimeSinceEpoch

`func (o *ModelArtifact) HasCreateTimeSinceEpoch() bool`

HasCreateTimeSinceEpoch returns a boolean if a field has been set.

### GetLastUpdateTimeSinceEpoch

`func (o *ModelArtifact) GetLastUpdateTimeSinceEpoch() string`

GetLastUpdateTimeSinceEpoch returns the LastUpdateTimeSinceEpoch field if non-nil, zero value otherwise.

### GetLastUpdateTimeSinceEpochOk

`func (o *ModelArtifact) GetLastUpdateTimeSinceEpochOk() (*string, bool)`

GetLastUpdateTimeSinceEpochOk returns a tuple with the LastUpdateTimeSinceEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimeSinceEpoch

`func (o *ModelArtifact) SetLastUpdateTimeSinceEpoch(v string)`

SetLastUpdateTimeSinceEpoch sets LastUpdateTimeSinceEpoch field to given value.

### HasLastUpdateTimeSinceEpoch

`func (o *ModelArtifact) HasLastUpdateTimeSinceEpoch() bool`

HasLastUpdateTimeSinceEpoch returns a boolean if a field has been set.

### GetArtifactType

`func (o *ModelArtifact) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *ModelArtifact) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *ModelArtifact) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.


### GetModelFormatName

`func (o *ModelArtifact) GetModelFormatName() string`

GetModelFormatName returns the ModelFormatName field if non-nil, zero value otherwise.

### GetModelFormatNameOk

`func (o *ModelArtifact) GetModelFormatNameOk() (*string, bool)`

GetModelFormatNameOk returns a tuple with the ModelFormatName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelFormatName

`func (o *ModelArtifact) SetModelFormatName(v string)`

SetModelFormatName sets ModelFormatName field to given value.

### HasModelFormatName

`func (o *ModelArtifact) HasModelFormatName() bool`

HasModelFormatName returns a boolean if a field has been set.

### GetStorageKey

`func (o *ModelArtifact) GetStorageKey() string`

GetStorageKey returns the StorageKey field if non-nil, zero value otherwise.

### GetStorageKeyOk

`func (o *ModelArtifact) GetStorageKeyOk() (*string, bool)`

GetStorageKeyOk returns a tuple with the StorageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageKey

`func (o *ModelArtifact) SetStorageKey(v string)`

SetStorageKey sets StorageKey field to given value.

### HasStorageKey

`func (o *ModelArtifact) HasStorageKey() bool`

HasStorageKey returns a boolean if a field has been set.

### GetStoragePath

`func (o *ModelArtifact) GetStoragePath() string`

GetStoragePath returns the StoragePath field if non-nil, zero value otherwise.

### GetStoragePathOk

`func (o *ModelArtifact) GetStoragePathOk() (*string, bool)`

GetStoragePathOk returns a tuple with the StoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePath

`func (o *ModelArtifact) SetStoragePath(v string)`

SetStoragePath sets StoragePath field to given value.

### HasStoragePath

`func (o *ModelArtifact) HasStoragePath() bool`

HasStoragePath returns a boolean if a field has been set.

### GetModelFormatVersion

`func (o *ModelArtifact) GetModelFormatVersion() string`

GetModelFormatVersion returns the ModelFormatVersion field if non-nil, zero value otherwise.

### GetModelFormatVersionOk

`func (o *ModelArtifact) GetModelFormatVersionOk() (*string, bool)`

GetModelFormatVersionOk returns a tuple with the ModelFormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelFormatVersion

`func (o *ModelArtifact) SetModelFormatVersion(v string)`

SetModelFormatVersion sets ModelFormatVersion field to given value.

### HasModelFormatVersion

`func (o *ModelArtifact) HasModelFormatVersion() bool`

HasModelFormatVersion returns a boolean if a field has been set.

### GetServiceAccountName

`func (o *ModelArtifact) GetServiceAccountName() string`

GetServiceAccountName returns the ServiceAccountName field if non-nil, zero value otherwise.

### GetServiceAccountNameOk

`func (o *ModelArtifact) GetServiceAccountNameOk() (*string, bool)`

GetServiceAccountNameOk returns a tuple with the ServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountName

`func (o *ModelArtifact) SetServiceAccountName(v string)`

SetServiceAccountName sets ServiceAccountName field to given value.

### HasServiceAccountName

`func (o *ModelArtifact) HasServiceAccountName() bool`

HasServiceAccountName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


