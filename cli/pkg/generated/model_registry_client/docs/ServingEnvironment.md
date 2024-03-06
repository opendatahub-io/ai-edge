# ServingEnvironment

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

## Methods

### NewServingEnvironment

`func NewServingEnvironment() *ServingEnvironment`

NewServingEnvironment instantiates a new ServingEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServingEnvironmentWithDefaults

`func NewServingEnvironmentWithDefaults() *ServingEnvironment`

NewServingEnvironmentWithDefaults instantiates a new ServingEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomProperties

`func (o *ServingEnvironment) GetCustomProperties() map[string]MetadataValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ServingEnvironment) GetCustomPropertiesOk() (*map[string]MetadataValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ServingEnvironment) SetCustomProperties(v map[string]MetadataValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ServingEnvironment) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *ServingEnvironment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServingEnvironment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServingEnvironment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServingEnvironment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *ServingEnvironment) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *ServingEnvironment) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *ServingEnvironment) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *ServingEnvironment) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetName

`func (o *ServingEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServingEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServingEnvironment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServingEnvironment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetId

`func (o *ServingEnvironment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServingEnvironment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServingEnvironment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServingEnvironment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreateTimeSinceEpoch

`func (o *ServingEnvironment) GetCreateTimeSinceEpoch() string`

GetCreateTimeSinceEpoch returns the CreateTimeSinceEpoch field if non-nil, zero value otherwise.

### GetCreateTimeSinceEpochOk

`func (o *ServingEnvironment) GetCreateTimeSinceEpochOk() (*string, bool)`

GetCreateTimeSinceEpochOk returns a tuple with the CreateTimeSinceEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTimeSinceEpoch

`func (o *ServingEnvironment) SetCreateTimeSinceEpoch(v string)`

SetCreateTimeSinceEpoch sets CreateTimeSinceEpoch field to given value.

### HasCreateTimeSinceEpoch

`func (o *ServingEnvironment) HasCreateTimeSinceEpoch() bool`

HasCreateTimeSinceEpoch returns a boolean if a field has been set.

### GetLastUpdateTimeSinceEpoch

`func (o *ServingEnvironment) GetLastUpdateTimeSinceEpoch() string`

GetLastUpdateTimeSinceEpoch returns the LastUpdateTimeSinceEpoch field if non-nil, zero value otherwise.

### GetLastUpdateTimeSinceEpochOk

`func (o *ServingEnvironment) GetLastUpdateTimeSinceEpochOk() (*string, bool)`

GetLastUpdateTimeSinceEpochOk returns a tuple with the LastUpdateTimeSinceEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTimeSinceEpoch

`func (o *ServingEnvironment) SetLastUpdateTimeSinceEpoch(v string)`

SetLastUpdateTimeSinceEpoch sets LastUpdateTimeSinceEpoch field to given value.

### HasLastUpdateTimeSinceEpoch

`func (o *ServingEnvironment) HasLastUpdateTimeSinceEpoch() bool`

HasLastUpdateTimeSinceEpoch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


