# V1AggregationRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterRoleSelectors** | Pointer to [**[]V1LabelSelector**](V1LabelSelector.md) | ClusterRoleSelectors holds a list of selectors which will be used to find ClusterRoles and create the rules. If any of the selectors match, then the ClusterRole&#39;s permissions will be added | [optional] 

## Methods

### NewV1AggregationRule

`func NewV1AggregationRule() *V1AggregationRule`

NewV1AggregationRule instantiates a new V1AggregationRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1AggregationRuleWithDefaults

`func NewV1AggregationRuleWithDefaults() *V1AggregationRule`

NewV1AggregationRuleWithDefaults instantiates a new V1AggregationRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterRoleSelectors

`func (o *V1AggregationRule) GetClusterRoleSelectors() []V1LabelSelector`

GetClusterRoleSelectors returns the ClusterRoleSelectors field if non-nil, zero value otherwise.

### GetClusterRoleSelectorsOk

`func (o *V1AggregationRule) GetClusterRoleSelectorsOk() (*[]V1LabelSelector, bool)`

GetClusterRoleSelectorsOk returns a tuple with the ClusterRoleSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterRoleSelectors

`func (o *V1AggregationRule) SetClusterRoleSelectors(v []V1LabelSelector)`

SetClusterRoleSelectors sets ClusterRoleSelectors field to given value.

### HasClusterRoleSelectors

`func (o *V1AggregationRule) HasClusterRoleSelectors() bool`

HasClusterRoleSelectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


