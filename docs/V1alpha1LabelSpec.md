# V1alpha1LabelSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackgroundColor** | Pointer to **string** |  | [optional] 
**Clusters** | Pointer to **[]string** |  | [optional] 
**Key** | **string** |  | 
**Value** | **string** |  | 

## Methods

### NewV1alpha1LabelSpec

`func NewV1alpha1LabelSpec(key string, value string, ) *V1alpha1LabelSpec`

NewV1alpha1LabelSpec instantiates a new V1alpha1LabelSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1LabelSpecWithDefaults

`func NewV1alpha1LabelSpecWithDefaults() *V1alpha1LabelSpec`

NewV1alpha1LabelSpecWithDefaults instantiates a new V1alpha1LabelSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackgroundColor

`func (o *V1alpha1LabelSpec) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *V1alpha1LabelSpec) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *V1alpha1LabelSpec) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *V1alpha1LabelSpec) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### GetClusters

`func (o *V1alpha1LabelSpec) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *V1alpha1LabelSpec) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *V1alpha1LabelSpec) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *V1alpha1LabelSpec) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetKey

`func (o *V1alpha1LabelSpec) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *V1alpha1LabelSpec) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *V1alpha1LabelSpec) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *V1alpha1LabelSpec) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *V1alpha1LabelSpec) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *V1alpha1LabelSpec) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


