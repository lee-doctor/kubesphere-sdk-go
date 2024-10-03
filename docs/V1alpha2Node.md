# V1alpha2Node

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adjacency** | Pointer to **[]string** |  | [optional] 
**Children** | Pointer to [**[]V1alpha2NodeSummaryGroup**](V1alpha2NodeSummaryGroup.md) |  | [optional] 
**Connections** | Pointer to [**[]V1alpha2ConnectionsSummary**](V1alpha2ConnectionsSummary.md) |  | [optional] 
**Controls** | [**[]V1alpha2ControlInstance**](V1alpha2ControlInstance.md) |  | 
**Id** | **string** |  | 
**Label** | **string** |  | 
**LabelMinor** | **string** |  | 
**Metadata** | Pointer to [**[]V1alpha2MetadataRow**](V1alpha2MetadataRow.md) |  | [optional] 
**Metrics** | Pointer to [**[]V1alpha2MetricRow**](V1alpha2MetricRow.md) |  | [optional] 
**Parents** | Pointer to [**[]V1alpha2Parent**](V1alpha2Parent.md) |  | [optional] 
**Pseudo** | Pointer to **bool** |  | [optional] 
**Rank** | **string** |  | 
**Shape** | Pointer to **string** |  | [optional] 
**Stack** | Pointer to **bool** |  | [optional] 
**Tables** | Pointer to [**[]V1alpha2Table**](V1alpha2Table.md) |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha2Node

`func NewV1alpha2Node(controls []V1alpha2ControlInstance, id string, label string, labelMinor string, rank string, ) *V1alpha2Node`

NewV1alpha2Node instantiates a new V1alpha2Node object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2NodeWithDefaults

`func NewV1alpha2NodeWithDefaults() *V1alpha2Node`

NewV1alpha2NodeWithDefaults instantiates a new V1alpha2Node object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjacency

`func (o *V1alpha2Node) GetAdjacency() []string`

GetAdjacency returns the Adjacency field if non-nil, zero value otherwise.

### GetAdjacencyOk

`func (o *V1alpha2Node) GetAdjacencyOk() (*[]string, bool)`

GetAdjacencyOk returns a tuple with the Adjacency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjacency

`func (o *V1alpha2Node) SetAdjacency(v []string)`

SetAdjacency sets Adjacency field to given value.

### HasAdjacency

`func (o *V1alpha2Node) HasAdjacency() bool`

HasAdjacency returns a boolean if a field has been set.

### GetChildren

`func (o *V1alpha2Node) GetChildren() []V1alpha2NodeSummaryGroup`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *V1alpha2Node) GetChildrenOk() (*[]V1alpha2NodeSummaryGroup, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *V1alpha2Node) SetChildren(v []V1alpha2NodeSummaryGroup)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *V1alpha2Node) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetConnections

`func (o *V1alpha2Node) GetConnections() []V1alpha2ConnectionsSummary`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *V1alpha2Node) GetConnectionsOk() (*[]V1alpha2ConnectionsSummary, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *V1alpha2Node) SetConnections(v []V1alpha2ConnectionsSummary)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *V1alpha2Node) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetControls

`func (o *V1alpha2Node) GetControls() []V1alpha2ControlInstance`

GetControls returns the Controls field if non-nil, zero value otherwise.

### GetControlsOk

`func (o *V1alpha2Node) GetControlsOk() (*[]V1alpha2ControlInstance, bool)`

GetControlsOk returns a tuple with the Controls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControls

`func (o *V1alpha2Node) SetControls(v []V1alpha2ControlInstance)`

SetControls sets Controls field to given value.


### GetId

`func (o *V1alpha2Node) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1alpha2Node) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1alpha2Node) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *V1alpha2Node) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *V1alpha2Node) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *V1alpha2Node) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLabelMinor

`func (o *V1alpha2Node) GetLabelMinor() string`

GetLabelMinor returns the LabelMinor field if non-nil, zero value otherwise.

### GetLabelMinorOk

`func (o *V1alpha2Node) GetLabelMinorOk() (*string, bool)`

GetLabelMinorOk returns a tuple with the LabelMinor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelMinor

`func (o *V1alpha2Node) SetLabelMinor(v string)`

SetLabelMinor sets LabelMinor field to given value.


### GetMetadata

`func (o *V1alpha2Node) GetMetadata() []V1alpha2MetadataRow`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1alpha2Node) GetMetadataOk() (*[]V1alpha2MetadataRow, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1alpha2Node) SetMetadata(v []V1alpha2MetadataRow)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1alpha2Node) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetMetrics

`func (o *V1alpha2Node) GetMetrics() []V1alpha2MetricRow`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *V1alpha2Node) GetMetricsOk() (*[]V1alpha2MetricRow, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *V1alpha2Node) SetMetrics(v []V1alpha2MetricRow)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *V1alpha2Node) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetParents

`func (o *V1alpha2Node) GetParents() []V1alpha2Parent`

GetParents returns the Parents field if non-nil, zero value otherwise.

### GetParentsOk

`func (o *V1alpha2Node) GetParentsOk() (*[]V1alpha2Parent, bool)`

GetParentsOk returns a tuple with the Parents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParents

`func (o *V1alpha2Node) SetParents(v []V1alpha2Parent)`

SetParents sets Parents field to given value.

### HasParents

`func (o *V1alpha2Node) HasParents() bool`

HasParents returns a boolean if a field has been set.

### GetPseudo

`func (o *V1alpha2Node) GetPseudo() bool`

GetPseudo returns the Pseudo field if non-nil, zero value otherwise.

### GetPseudoOk

`func (o *V1alpha2Node) GetPseudoOk() (*bool, bool)`

GetPseudoOk returns a tuple with the Pseudo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPseudo

`func (o *V1alpha2Node) SetPseudo(v bool)`

SetPseudo sets Pseudo field to given value.

### HasPseudo

`func (o *V1alpha2Node) HasPseudo() bool`

HasPseudo returns a boolean if a field has been set.

### GetRank

`func (o *V1alpha2Node) GetRank() string`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *V1alpha2Node) GetRankOk() (*string, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *V1alpha2Node) SetRank(v string)`

SetRank sets Rank field to given value.


### GetShape

`func (o *V1alpha2Node) GetShape() string`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *V1alpha2Node) GetShapeOk() (*string, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *V1alpha2Node) SetShape(v string)`

SetShape sets Shape field to given value.

### HasShape

`func (o *V1alpha2Node) HasShape() bool`

HasShape returns a boolean if a field has been set.

### GetStack

`func (o *V1alpha2Node) GetStack() bool`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *V1alpha2Node) GetStackOk() (*bool, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *V1alpha2Node) SetStack(v bool)`

SetStack sets Stack field to given value.

### HasStack

`func (o *V1alpha2Node) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetTables

`func (o *V1alpha2Node) GetTables() []V1alpha2Table`

GetTables returns the Tables field if non-nil, zero value otherwise.

### GetTablesOk

`func (o *V1alpha2Node) GetTablesOk() (*[]V1alpha2Table, bool)`

GetTablesOk returns a tuple with the Tables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTables

`func (o *V1alpha2Node) SetTables(v []V1alpha2Table)`

SetTables sets Tables field to given value.

### HasTables

`func (o *V1alpha2Node) HasTables() bool`

HasTables returns a boolean if a field has been set.

### GetTag

`func (o *V1alpha2Node) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *V1alpha2Node) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *V1alpha2Node) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *V1alpha2Node) HasTag() bool`

HasTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


