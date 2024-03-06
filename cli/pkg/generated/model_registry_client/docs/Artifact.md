# Artifact

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
**ArtifactType** | **string** |  | [default to "doc-artifact"]
**ModelFormatName** | Pointer to **string** | Name of the model format. | [optional] 
**StorageKey** | Pointer to **string** | Storage secret name. | [optional] 
**StoragePath** | Pointer to **string** | Path for model in storage provided by &#x60;storageKey&#x60;. | [optional] 
**ModelFormatVersion** | Pointer to **string** | Version of the model format. | [optional] 
**ServiceAccountName** | Pointer to **string** | Name of the service account with storage secret. | [optional] 

## Methods

### NewArtifact

`func NewArtifact(artifactType string, ) *Artifact`

NewArtifact instantiates a new Artifact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactWithDefaults

`func NewArtifactWithDefaults() *Artifact`

NewArtifactWithDefaults instantiates a new Artifact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomProperties

`func (o *Artifact) GetCustomProperties() map[string]MetadataValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *Artifact) GetCustomPropertiesOk() (*map[string]MetadataValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *Artifact) SetCustomProperties(v map[string]MetadataValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *Artifact) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *Artifact) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Artifact) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Artifact) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Artifact) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *Artifact) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *Artifact) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *Artifact) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *Artifact) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetUri

`func (o *Artifact) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *Artifact) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *Artifact) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *Artifact) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetState

`func (o *Artifact) GetState() ArtifactState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Artifact) GetStateOk() (*ArtifactState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Artifact) SetState(v ArtifactState)`

SetState sets State field to given value.

### HasState

`func (o *Artifact) HasState() bool`

HasState returns a boolean if a field has been set.

### GetName

`func (o *Artifact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Artifact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Artifact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Artifact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *Artifact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Artifact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Artifact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Artifact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreateTimeSinceEpoch

`func (o *Artifact) GetCreateTimeSinceEpoch() string`

GetCreateTimeSinceEpoch returns the CreateTimeSinceEpoch field if non-nil, zero value otherwise.

### GetCreateTimeSinceEpochOk

`func (o *Artifact) GetCreateTimeSinceEpochOk() (*string, bool)`

GetCreateTimeSinceEpochOk returns a tuple with the CreateTimeSinceEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeSinceEpoch

`func (o *Artifact) SetCreateTimeSinceEpoch(v string)`

SetCreateTimeSinceEpoch sets CreateTimeSinceEpoch field to given value.

### HasCreateTimeSinceEpoch

`func (o *Artifact) HasCreateTimeSinceEpoch() bool`

HasCreateTimeSinceEpoch returns a boolean if a field has been set.

### GetLastUpdateTimeSinceEpoch

`func (o *Artifact) GetLastUpdateTimeSinceEpoch() string`

GetLastUpdateTimeSinceEpoch returns the LastUpdateTimeSinceEpoch field if non-nil, zero value otherwise.

### GetLastUpdateTimeSinceEpochOk

`func (o *Artifact) GetLastUpdateTimeSinceEpochOk() (*string, bool)`

GetLastUpdateTimeSinceEpochOk returns a tuple with the LastUpdateTimeSinceEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimeSinceEpoch

`func (o *Artifact) SetLastUpdateTimeSinceEpoch(v string)`

SetLastUpdateTimeSinceEpoch sets LastUpdateTimeSinceEpoch field to given value.

### HasLastUpdateTimeSinceEpoch

`func (o *Artifact) HasLastUpdateTimeSinceEpoch() bool`

HasLastUpdateTimeSinceEpoch returns a boolean if a field has been set.

### GetArtifactType

`func (o *Artifact) GetArtifactType() string`

GetArtifactType returns the ArtifactType field if non-nil, zero value otherwise.

### GetArtifactTypeOk

`func (o *Artifact) GetArtifactTypeOk() (*string, bool)`

GetArtifactTypeOk returns a tuple with the ArtifactType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactType

`func (o *Artifact) SetArtifactType(v string)`

SetArtifactType sets ArtifactType field to given value.


### GetModelFormatName

`func (o *Artifact) GetModelFormatName() string`

GetModelFormatName returns the ModelFormatName field if non-nil, zero value otherwise.

### GetModelFormatNameOk

`func (o *Artifact) GetModelFormatNameOk() (*string, bool)`

GetModelFormatNameOk returns a tuple with the ModelFormatName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelFormatName

`func (o *Artifact) SetModelFormatName(v string)`

SetModelFormatName sets ModelFormatName field to given value.

### HasModelFormatName

`func (o *Artifact) HasModelFormatName() bool`

HasModelFormatName returns a boolean if a field has been set.

### GetStorageKey

`func (o *Artifact) GetStorageKey() string`

GetStorageKey returns the StorageKey field if non-nil, zero value otherwise.

### GetStorageKeyOk

`func (o *Artifact) GetStorageKeyOk() (*string, bool)`

GetStorageKeyOk returns a tuple with the StorageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageKey

`func (o *Artifact) SetStorageKey(v string)`

SetStorageKey sets StorageKey field to given value.

### HasStorageKey

`func (o *Artifact) HasStorageKey() bool`

HasStorageKey returns a boolean if a field has been set.

### GetStoragePath

`func (o *Artifact) GetStoragePath() string`

GetStoragePath returns the StoragePath field if non-nil, zero value otherwise.

### GetStoragePathOk

`func (o *Artifact) GetStoragePathOk() (*string, bool)`

GetStoragePathOk returns a tuple with the StoragePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePath

`func (o *Artifact) SetStoragePath(v string)`

SetStoragePath sets StoragePath field to given value.

### HasStoragePath

`func (o *Artifact) HasStoragePath() bool`

HasStoragePath returns a boolean if a field has been set.

### GetModelFormatVersion

`func (o *Artifact) GetModelFormatVersion() string`

GetModelFormatVersion returns the ModelFormatVersion field if non-nil, zero value otherwise.

### GetModelFormatVersionOk

`func (o *Artifact) GetModelFormatVersionOk() (*string, bool)`

GetModelFormatVersionOk returns a tuple with the ModelFormatVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelFormatVersion

`func (o *Artifact) SetModelFormatVersion(v string)`

SetModelFormatVersion sets ModelFormatVersion field to given value.

### HasModelFormatVersion

`func (o *Artifact) HasModelFormatVersion() bool`

HasModelFormatVersion returns a boolean if a field has been set.

### GetServiceAccountName

`func (o *Artifact) GetServiceAccountName() string`

GetServiceAccountName returns the ServiceAccountName field if non-nil, zero value otherwise.

### GetServiceAccountNameOk

`func (o *Artifact) GetServiceAccountNameOk() (*string, bool)`

GetServiceAccountNameOk returns a tuple with the ServiceAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountName

`func (o *Artifact) SetServiceAccountName(v string)`

SetServiceAccountName sets ServiceAccountName field to given value.

### HasServiceAccountName

`func (o *Artifact) HasServiceAccountName() bool`

HasServiceAccountName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


