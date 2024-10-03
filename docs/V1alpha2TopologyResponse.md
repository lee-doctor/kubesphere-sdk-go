# V1alpha2TopologyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nodes** | [**map[string]V1alpha2NodeSummary**](V1alpha2NodeSummary.md) |  | 

## Methods

### NewV1alpha2TopologyResponse

`func NewV1alpha2TopologyResponse(nodes map[string]V1alpha2NodeSummary, ) *V1alpha2TopologyResponse`

NewV1alpha2TopologyResponse instantiates a new V1alpha2TopologyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2TopologyResponseWithDefaults

`func NewV1alpha2TopologyResponseWithDefaults() *V1alpha2TopologyResponse`

NewV1alpha2TopologyResponseWithDefaults instantiates a new V1alpha2TopologyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodes

`func (o *V1alpha2TopologyResponse) GetNodes() map[string]V1alpha2NodeSummary`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V1alpha2TopologyResponse) GetNodesOk() (*map[string]V1alpha2NodeSummary, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V1alpha2TopologyResponse) SetNodes(v map[string]V1alpha2NodeSummary)`

SetNodes sets Nodes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


