# ArtifactList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextPageToken** | **string** | Token to use to retrieve next page of results. | 
**PageSize** | **int32** | Maximum number of resources to return in the result. | 
**Size** | **int32** | Number of items in result list. | 
**Items** | Pointer to [**[]Artifact**](Artifact.md) | Array of &#x60;Artifact&#x60; entities. | [optional] 

## Methods

### NewArtifactList

`func NewArtifactList(nextPageToken string, pageSize int32, size int32, ) *ArtifactList`

NewArtifactList instantiates a new ArtifactList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArtifactListWithDefaults

`func NewArtifactListWithDefaults() *ArtifactList`

NewArtifactListWithDefaults instantiates a new ArtifactList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextPageToken

`func (o *ArtifactList) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *ArtifactList) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *ArtifactList) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.


### GetPageSize

`func (o *ArtifactList) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *ArtifactList) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *ArtifactList) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.


### GetSize

`func (o *ArtifactList) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ArtifactList) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ArtifactList) SetSize(v int32)`

SetSize sets Size field to given value.


### GetItems

`func (o *ArtifactList) GetItems() []Artifact`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ArtifactList) GetItemsOk() (*[]Artifact, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ArtifactList) SetItems(v []Artifact)`

SetItems sets Items field to given value.

### HasItems

`func (o *ArtifactList) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


