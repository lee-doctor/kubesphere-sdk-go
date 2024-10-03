# V1TopologySpreadConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LabelSelector** | Pointer to [**V1LabelSelector**](V1LabelSelector.md) |  | [optional] 
**MaxSkew** | **int32** | MaxSkew describes the degree to which pods may be unevenly distributed. It&#39;s the maximum permitted difference between the number of matching pods in any two topology domains of a given topology type. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 1/1/0:  | 
**TopologyKey** | **string** | TopologyKey is the key of node labels. Nodes that have a label with this key and identical values are considered to be in the same topology. We consider each &lt;key, value&gt; as a \&quot;bucket\&quot;, and try to put balanced number of pods into each bucket. It&#39;s a required field. | 
**WhenUnsatisfiable** | **string** | WhenUnsatisfiable indicates how to deal with a pod if it doesn&#39;t satisfy the spread constraint. - DoNotSchedule (default) tells the scheduler not to schedule it - ScheduleAnyway tells the scheduler to still schedule it It&#39;s considered as \&quot;Unsatisfiable\&quot; if and only if placing incoming pod on any topology violates \&quot;MaxSkew\&quot;. For example, in a 3-zone cluster, MaxSkew is set to 1, and pods with the same labelSelector spread as 3/1/1:  | 

## Methods

### NewV1TopologySpreadConstraint

`func NewV1TopologySpreadConstraint(maxSkew int32, topologyKey string, whenUnsatisfiable string, ) *V1TopologySpreadConstraint`

NewV1TopologySpreadConstraint instantiates a new V1TopologySpreadConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1TopologySpreadConstraintWithDefaults

`func NewV1TopologySpreadConstraintWithDefaults() *V1TopologySpreadConstraint`

NewV1TopologySpreadConstraintWithDefaults instantiates a new V1TopologySpreadConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabelSelector

`func (o *V1TopologySpreadConstraint) GetLabelSelector() V1LabelSelector`

GetLabelSelector returns the LabelSelector field if non-nil, zero value otherwise.

### GetLabelSelectorOk

`func (o *V1TopologySpreadConstraint) GetLabelSelectorOk() (*V1LabelSelector, bool)`

GetLabelSelectorOk returns a tuple with the LabelSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelSelector

`func (o *V1TopologySpreadConstraint) SetLabelSelector(v V1LabelSelector)`

SetLabelSelector sets LabelSelector field to given value.

### HasLabelSelector

`func (o *V1TopologySpreadConstraint) HasLabelSelector() bool`

HasLabelSelector returns a boolean if a field has been set.

### GetMaxSkew

`func (o *V1TopologySpreadConstraint) GetMaxSkew() int32`

GetMaxSkew returns the MaxSkew field if non-nil, zero value otherwise.

### GetMaxSkewOk

`func (o *V1TopologySpreadConstraint) GetMaxSkewOk() (*int32, bool)`

GetMaxSkewOk returns a tuple with the MaxSkew field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSkew

`func (o *V1TopologySpreadConstraint) SetMaxSkew(v int32)`

SetMaxSkew sets MaxSkew field to given value.


### GetTopologyKey

`func (o *V1TopologySpreadConstraint) GetTopologyKey() string`

GetTopologyKey returns the TopologyKey field if non-nil, zero value otherwise.

### GetTopologyKeyOk

`func (o *V1TopologySpreadConstraint) GetTopologyKeyOk() (*string, bool)`

GetTopologyKeyOk returns a tuple with the TopologyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyKey

`func (o *V1TopologySpreadConstraint) SetTopologyKey(v string)`

SetTopologyKey sets TopologyKey field to given value.


### GetWhenUnsatisfiable

`func (o *V1TopologySpreadConstraint) GetWhenUnsatisfiable() string`

GetWhenUnsatisfiable returns the WhenUnsatisfiable field if non-nil, zero value otherwise.

### GetWhenUnsatisfiableOk

`func (o *V1TopologySpreadConstraint) GetWhenUnsatisfiableOk() (*string, bool)`

GetWhenUnsatisfiableOk returns a tuple with the WhenUnsatisfiable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhenUnsatisfiable

`func (o *V1TopologySpreadConstraint) SetWhenUnsatisfiable(v string)`

SetWhenUnsatisfiable sets WhenUnsatisfiable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


