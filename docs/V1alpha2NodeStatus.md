# V1alpha2NodeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthyNodes** | **int32** | the number of healthy nodes | 
**TotalNodes** | **int32** | total number of nodes | 

## Methods

### NewV1alpha2NodeStatus

`func NewV1alpha2NodeStatus(healthyNodes int32, totalNodes int32, ) *V1alpha2NodeStatus`

NewV1alpha2NodeStatus instantiates a new V1alpha2NodeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2NodeStatusWithDefaults

`func NewV1alpha2NodeStatusWithDefaults() *V1alpha2NodeStatus`

NewV1alpha2NodeStatusWithDefaults instantiates a new V1alpha2NodeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthyNodes

`func (o *V1alpha2NodeStatus) GetHealthyNodes() int32`

GetHealthyNodes returns the HealthyNodes field if non-nil, zero value otherwise.

### GetHealthyNodesOk

`func (o *V1alpha2NodeStatus) GetHealthyNodesOk() (*int32, bool)`

GetHealthyNodesOk returns a tuple with the HealthyNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyNodes

`func (o *V1alpha2NodeStatus) SetHealthyNodes(v int32)`

SetHealthyNodes sets HealthyNodes field to given value.


### GetTotalNodes

`func (o *V1alpha2NodeStatus) GetTotalNodes() int32`

GetTotalNodes returns the TotalNodes field if non-nil, zero value otherwise.

### GetTotalNodesOk

`func (o *V1alpha2NodeStatus) GetTotalNodesOk() (*int32, bool)`

GetTotalNodesOk returns a tuple with the TotalNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalNodes

`func (o *V1alpha2NodeStatus) SetTotalNodes(v int32)`

SetTotalNodes sets TotalNodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


