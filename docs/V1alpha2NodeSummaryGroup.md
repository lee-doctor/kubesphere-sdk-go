# V1alpha2NodeSummaryGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | [**[]V1alpha2Column**](V1alpha2Column.md) |  | 
**Id** | **string** |  | 
**Label** | **string** |  | 
**Nodes** | [**[]V1alpha2NodeSummary**](V1alpha2NodeSummary.md) |  | 
**TopologyId** | **string** |  | 

## Methods

### NewV1alpha2NodeSummaryGroup

`func NewV1alpha2NodeSummaryGroup(columns []V1alpha2Column, id string, label string, nodes []V1alpha2NodeSummary, topologyId string, ) *V1alpha2NodeSummaryGroup`

NewV1alpha2NodeSummaryGroup instantiates a new V1alpha2NodeSummaryGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2NodeSummaryGroupWithDefaults

`func NewV1alpha2NodeSummaryGroupWithDefaults() *V1alpha2NodeSummaryGroup`

NewV1alpha2NodeSummaryGroupWithDefaults instantiates a new V1alpha2NodeSummaryGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *V1alpha2NodeSummaryGroup) GetColumns() []V1alpha2Column`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *V1alpha2NodeSummaryGroup) GetColumnsOk() (*[]V1alpha2Column, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *V1alpha2NodeSummaryGroup) SetColumns(v []V1alpha2Column)`

SetColumns sets Columns field to given value.


### GetId

`func (o *V1alpha2NodeSummaryGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1alpha2NodeSummaryGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1alpha2NodeSummaryGroup) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *V1alpha2NodeSummaryGroup) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *V1alpha2NodeSummaryGroup) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *V1alpha2NodeSummaryGroup) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetNodes

`func (o *V1alpha2NodeSummaryGroup) GetNodes() []V1alpha2NodeSummary`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *V1alpha2NodeSummaryGroup) GetNodesOk() (*[]V1alpha2NodeSummary, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *V1alpha2NodeSummaryGroup) SetNodes(v []V1alpha2NodeSummary)`

SetNodes sets Nodes field to given value.


### GetTopologyId

`func (o *V1alpha2NodeSummaryGroup) GetTopologyId() string`

GetTopologyId returns the TopologyId field if non-nil, zero value otherwise.

### GetTopologyIdOk

`func (o *V1alpha2NodeSummaryGroup) GetTopologyIdOk() (*string, bool)`

GetTopologyIdOk returns a tuple with the TopologyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyId

`func (o *V1alpha2NodeSummaryGroup) SetTopologyId(v string)`

SetTopologyId sets TopologyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


