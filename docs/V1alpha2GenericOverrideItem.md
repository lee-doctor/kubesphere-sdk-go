# V1alpha2GenericOverrideItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterName** | **string** |  | 
**ClusterOverrides** | Pointer to [**[]V1alpha2ClusterOverride**](V1alpha2ClusterOverride.md) |  | [optional] 

## Methods

### NewV1alpha2GenericOverrideItem

`func NewV1alpha2GenericOverrideItem(clusterName string, ) *V1alpha2GenericOverrideItem`

NewV1alpha2GenericOverrideItem instantiates a new V1alpha2GenericOverrideItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2GenericOverrideItemWithDefaults

`func NewV1alpha2GenericOverrideItemWithDefaults() *V1alpha2GenericOverrideItem`

NewV1alpha2GenericOverrideItemWithDefaults instantiates a new V1alpha2GenericOverrideItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterName

`func (o *V1alpha2GenericOverrideItem) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *V1alpha2GenericOverrideItem) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *V1alpha2GenericOverrideItem) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetClusterOverrides

`func (o *V1alpha2GenericOverrideItem) GetClusterOverrides() []V1alpha2ClusterOverride`

GetClusterOverrides returns the ClusterOverrides field if non-nil, zero value otherwise.

### GetClusterOverridesOk

`func (o *V1alpha2GenericOverrideItem) GetClusterOverridesOk() (*[]V1alpha2ClusterOverride, bool)`

GetClusterOverridesOk returns a tuple with the ClusterOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterOverrides

`func (o *V1alpha2GenericOverrideItem) SetClusterOverrides(v []V1alpha2ClusterOverride)`

SetClusterOverrides sets ClusterOverrides field to given value.

### HasClusterOverrides

`func (o *V1alpha2GenericOverrideItem) HasClusterOverrides() bool`

HasClusterOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


