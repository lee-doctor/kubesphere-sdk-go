# V1alpha2HealthStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KubesphereStatus** | [**[]V1alpha2ComponentStatus**](V1alpha2ComponentStatus.md) | kubesphere components status | 
**NodeStatus** | [**V1alpha2NodeStatus**](V1alpha2NodeStatus.md) |  | 

## Methods

### NewV1alpha2HealthStatus

`func NewV1alpha2HealthStatus(kubesphereStatus []V1alpha2ComponentStatus, nodeStatus V1alpha2NodeStatus, ) *V1alpha2HealthStatus`

NewV1alpha2HealthStatus instantiates a new V1alpha2HealthStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2HealthStatusWithDefaults

`func NewV1alpha2HealthStatusWithDefaults() *V1alpha2HealthStatus`

NewV1alpha2HealthStatusWithDefaults instantiates a new V1alpha2HealthStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKubesphereStatus

`func (o *V1alpha2HealthStatus) GetKubesphereStatus() []V1alpha2ComponentStatus`

GetKubesphereStatus returns the KubesphereStatus field if non-nil, zero value otherwise.

### GetKubesphereStatusOk

`func (o *V1alpha2HealthStatus) GetKubesphereStatusOk() (*[]V1alpha2ComponentStatus, bool)`

GetKubesphereStatusOk returns a tuple with the KubesphereStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubesphereStatus

`func (o *V1alpha2HealthStatus) SetKubesphereStatus(v []V1alpha2ComponentStatus)`

SetKubesphereStatus sets KubesphereStatus field to given value.


### GetNodeStatus

`func (o *V1alpha2HealthStatus) GetNodeStatus() V1alpha2NodeStatus`

GetNodeStatus returns the NodeStatus field if non-nil, zero value otherwise.

### GetNodeStatusOk

`func (o *V1alpha2HealthStatus) GetNodeStatusOk() (*V1alpha2NodeStatus, bool)`

GetNodeStatusOk returns a tuple with the NodeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStatus

`func (o *V1alpha2HealthStatus) SetNodeStatus(v V1alpha2NodeStatus)`

SetNodeStatus sets NodeStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


