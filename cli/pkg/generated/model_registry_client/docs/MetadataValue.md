# MetadataValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntValue** | Pointer to **string** |  | [optional] 
**DoubleValue** | Pointer to **float64** |  | [optional] 
**StringValue** | Pointer to **string** |  | [optional] 
**StructValue** | Pointer to **string** | Base64 encoded bytes for struct value | [optional] 
**Type** | Pointer to **string** | url describing proto value | [optional] 
**ProtoValue** | Pointer to **string** | Base64 encoded bytes for proto value | [optional] 
**BoolValue** | Pointer to **bool** |  | [optional] 

## Methods

### NewMetadataValue

`func NewMetadataValue() *MetadataValue`

NewMetadataValue instantiates a new MetadataValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetadataValueWithDefaults

`func NewMetadataValueWithDefaults() *MetadataValue`

NewMetadataValueWithDefaults instantiates a new MetadataValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntValue

`func (o *MetadataValue) GetIntValue() string`

GetIntValue returns the IntValue field if non-nil, zero value otherwise.

### GetIntValueOk

`func (o *MetadataValue) GetIntValueOk() (*string, bool)`

GetIntValueOk returns a tuple with the IntValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntValue

`func (o *MetadataValue) SetIntValue(v string)`

SetIntValue sets IntValue field to given value.

### HasIntValue

`func (o *MetadataValue) HasIntValue() bool`

HasIntValue returns a boolean if a field has been set.

### GetDoubleValue

`func (o *MetadataValue) GetDoubleValue() float64`

GetDoubleValue returns the DoubleValue field if non-nil, zero value otherwise.

### GetDoubleValueOk

`func (o *MetadataValue) GetDoubleValueOk() (*float64, bool)`

GetDoubleValueOk returns a tuple with the DoubleValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoubleValue

`func (o *MetadataValue) SetDoubleValue(v float64)`

SetDoubleValue sets DoubleValue field to given value.

### HasDoubleValue

`func (o *MetadataValue) HasDoubleValue() bool`

HasDoubleValue returns a boolean if a field has been set.

### GetStringValue

`func (o *MetadataValue) GetStringValue() string`

GetStringValue returns the StringValue field if non-nil, zero value otherwise.

### GetStringValueOk

`func (o *MetadataValue) GetStringValueOk() (*string, bool)`

GetStringValueOk returns a tuple with the StringValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValue

`func (o *MetadataValue) SetStringValue(v string)`

SetStringValue sets StringValue field to given value.

### HasStringValue

`func (o *MetadataValue) HasStringValue() bool`

HasStringValue returns a boolean if a field has been set.

### GetStructValue

`func (o *MetadataValue) GetStructValue() string`

GetStructValue returns the StructValue field if non-nil, zero value otherwise.

### GetStructValueOk

`func (o *MetadataValue) GetStructValueOk() (*string, bool)`

GetStructValueOk returns a tuple with the StructValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructValue

`func (o *MetadataValue) SetStructValue(v string)`

SetStructValue sets StructValue field to given value.

### HasStructValue

`func (o *MetadataValue) HasStructValue() bool`

HasStructValue returns a boolean if a field has been set.

### GetType

`func (o *MetadataValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetadataValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetadataValue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MetadataValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetProtoValue

`func (o *MetadataValue) GetProtoValue() string`

GetProtoValue returns the ProtoValue field if non-nil, zero value otherwise.

### GetProtoValueOk

`func (o *MetadataValue) GetProtoValueOk() (*string, bool)`

GetProtoValueOk returns a tuple with the ProtoValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoValue

`func (o *MetadataValue) SetProtoValue(v string)`

SetProtoValue sets ProtoValue field to given value.

### HasProtoValue

`func (o *MetadataValue) HasProtoValue() bool`

HasProtoValue returns a boolean if a field has been set.

### GetBoolValue

`func (o *MetadataValue) GetBoolValue() bool`

GetBoolValue returns the BoolValue field if non-nil, zero value otherwise.

### GetBoolValueOk

`func (o *MetadataValue) GetBoolValueOk() (*bool, bool)`

GetBoolValueOk returns a tuple with the BoolValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoolValue

`func (o *MetadataValue) SetBoolValue(v bool)`

SetBoolValue sets BoolValue field to given value.

### HasBoolValue

`func (o *MetadataValue) HasBoolValue() bool`

HasBoolValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


