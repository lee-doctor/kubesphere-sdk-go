# V1alpha2GenericPlacementFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterSelector** | Pointer to [**V1LabelSelector**](V1LabelSelector.md) |  | [optional] 
**Clusters** | Pointer to [**[]V1alpha2GenericClusterReference**](V1alpha2GenericClusterReference.md) |  | [optional] 

## Methods

### NewV1alpha2GenericPlacementFields

`func NewV1alpha2GenericPlacementFields() *V1alpha2GenericPlacementFields`

NewV1alpha2GenericPlacementFields instantiates a new V1alpha2GenericPlacementFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2GenericPlacementFieldsWithDefaults

`func NewV1alpha2GenericPlacementFieldsWithDefaults() *V1alpha2GenericPlacementFields`

NewV1alpha2GenericPlacementFieldsWithDefaults instantiates a new V1alpha2GenericPlacementFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterSelector

`func (o *V1alpha2GenericPlacementFields) GetClusterSelector() V1LabelSelector`

GetClusterSelector returns the ClusterSelector field if non-nil, zero value otherwise.

### GetClusterSelectorOk

`func (o *V1alpha2GenericPlacementFields) GetClusterSelectorOk() (*V1LabelSelector, bool)`

GetClusterSelectorOk returns a tuple with the ClusterSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSelector

`func (o *V1alpha2GenericPlacementFields) SetClusterSelector(v V1LabelSelector)`

SetClusterSelector sets ClusterSelector field to given value.

### HasClusterSelector

`func (o *V1alpha2GenericPlacementFields) HasClusterSelector() bool`

HasClusterSelector returns a boolean if a field has been set.

### GetClusters

`func (o *V1alpha2GenericPlacementFields) GetClusters() []V1alpha2GenericClusterReference`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *V1alpha2GenericPlacementFields) GetClustersOk() (*[]V1alpha2GenericClusterReference, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *V1alpha2GenericPlacementFields) SetClusters(v []V1alpha2GenericClusterReference)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *V1alpha2GenericPlacementFields) HasClusters() bool`

HasClusters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


