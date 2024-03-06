# ModelVersionCreate

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

## Methods

### NewModelVersionCreate

`func NewModelVersionCreate(registeredModelID string, ) *ModelVersionCreate`

NewModelVersionCreate instantiates a new ModelVersionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelVersionCreateWithDefaults

`func NewModelVersionCreateWithDefaults() *ModelVersionCreate`

NewModelVersionCreateWithDefaults instantiates a new ModelVersionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegisteredModelID

`func (o *ModelVersionCreate) GetRegisteredModelID() string`

GetRegisteredModelID returns the RegisteredModelID field if non-nil, zero value otherwise.

### GetRegisteredModelIDOk

`func (o *ModelVersionCreate) GetRegisteredModelIDOk() (*string, bool)`

GetRegisteredModelIDOk returns a tuple with the RegisteredModelID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredModelID

`func (o *ModelVersionCreate) SetRegisteredModelID(v string)`

SetRegisteredModelID sets RegisteredModelID field to given value.


### GetCustomProperties

`func (o *ModelVersionCreate) GetCustomProperties() map[string]MetadataValue`

GetCustomProperties returns the CustomProperties field if non-nil, zero value otherwise.

### GetCustomPropertiesOk

`func (o *ModelVersionCreate) GetCustomPropertiesOk() (*map[string]MetadataValue, bool)`

GetCustomPropertiesOk returns a tuple with the CustomProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomProperties

`func (o *ModelVersionCreate) SetCustomProperties(v map[string]MetadataValue)`

SetCustomProperties sets CustomProperties field to given value.

### HasCustomProperties

`func (o *ModelVersionCreate) HasCustomProperties() bool`

HasCustomProperties returns a boolean if a field has been set.

### GetDescription

`func (o *ModelVersionCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelVersionCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelVersionCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelVersionCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExternalID

`func (o *ModelVersionCreate) GetExternalID() string`

GetExternalID returns the ExternalID field if non-nil, zero value otherwise.

### GetExternalIDOk

`func (o *ModelVersionCreate) GetExternalIDOk() (*string, bool)`

GetExternalIDOk returns a tuple with the ExternalID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalID

`func (o *ModelVersionCreate) SetExternalID(v string)`

SetExternalID sets ExternalID field to given value.

### HasExternalID

`func (o *ModelVersionCreate) HasExternalID() bool`

HasExternalID returns a boolean if a field has been set.

### GetName

`func (o *ModelVersionCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelVersionCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelVersionCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelVersionCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetState

`func (o *ModelVersionCreate) GetState() ModelVersionState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ModelVersionCreate) GetStateOk() (*ModelVersionState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ModelVersionCreate) SetState(v ModelVersionState)`

SetState sets State field to given value.

### HasState

`func (o *ModelVersionCreate) HasState() bool`

HasState returns a boolean if a field has been set.

### GetAuthor

`func (o *ModelVersionCreate) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ModelVersionCreate) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ModelVersionCreate) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ModelVersionCreate) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


