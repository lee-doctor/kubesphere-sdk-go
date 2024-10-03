# V1alpha2ConnectionsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | [**[]V1alpha2Column**](V1alpha2Column.md) |  | 
**Connections** | [**[]V1alpha2Connection**](V1alpha2Connection.md) |  | 
**Id** | **string** |  | 
**Label** | **string** |  | 
**TopologyId** | **string** |  | 

## Methods

### NewV1alpha2ConnectionsSummary

`func NewV1alpha2ConnectionsSummary(columns []V1alpha2Column, connections []V1alpha2Connection, id string, label string, topologyId string, ) *V1alpha2ConnectionsSummary`

NewV1alpha2ConnectionsSummary instantiates a new V1alpha2ConnectionsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2ConnectionsSummaryWithDefaults

`func NewV1alpha2ConnectionsSummaryWithDefaults() *V1alpha2ConnectionsSummary`

NewV1alpha2ConnectionsSummaryWithDefaults instantiates a new V1alpha2ConnectionsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *V1alpha2ConnectionsSummary) GetColumns() []V1alpha2Column`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *V1alpha2ConnectionsSummary) GetColumnsOk() (*[]V1alpha2Column, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *V1alpha2ConnectionsSummary) SetColumns(v []V1alpha2Column)`

SetColumns sets Columns field to given value.


### GetConnections

`func (o *V1alpha2ConnectionsSummary) GetConnections() []V1alpha2Connection`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *V1alpha2ConnectionsSummary) GetConnectionsOk() (*[]V1alpha2Connection, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *V1alpha2ConnectionsSummary) SetConnections(v []V1alpha2Connection)`

SetConnections sets Connections field to given value.


### GetId

`func (o *V1alpha2ConnectionsSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1alpha2ConnectionsSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1alpha2ConnectionsSummary) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *V1alpha2ConnectionsSummary) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *V1alpha2ConnectionsSummary) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *V1alpha2ConnectionsSummary) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetTopologyId

`func (o *V1alpha2ConnectionsSummary) GetTopologyId() string`

GetTopologyId returns the TopologyId field if non-nil, zero value otherwise.

### GetTopologyIdOk

`func (o *V1alpha2ConnectionsSummary) GetTopologyIdOk() (*string, bool)`

GetTopologyIdOk returns a tuple with the TopologyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyId

`func (o *V1alpha2ConnectionsSummary) SetTopologyId(v string)`

SetTopologyId sets TopologyId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


