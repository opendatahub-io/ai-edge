# MetadataProtoValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | url describing proto value | [optional] 
**ProtoValue** | Pointer to **string** | Base64 encoded bytes for proto value | [optional] 

## Methods

### NewMetadataProtoValue

`func NewMetadataProtoValue() *MetadataProtoValue`

NewMetadataProtoValue instantiates a new MetadataProtoValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataProtoValueWithDefaults

`func NewMetadataProtoValueWithDefaults() *MetadataProtoValue`

NewMetadataProtoValueWithDefaults instantiates a new MetadataProtoValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MetadataProtoValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetadataProtoValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetadataProtoValue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MetadataProtoValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProtoValue

`func (o *MetadataProtoValue) GetProtoValue() string`

GetProtoValue returns the ProtoValue field if non-nil, zero value otherwise.

### GetProtoValueOk

`func (o *MetadataProtoValue) GetProtoValueOk() (*string, bool)`

GetProtoValueOk returns a tuple with the ProtoValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoValue

`func (o *MetadataProtoValue) SetProtoValue(v string)`

SetProtoValue sets ProtoValue field to given value.

### HasProtoValue

`func (o *MetadataProtoValue) HasProtoValue() bool`

HasProtoValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


