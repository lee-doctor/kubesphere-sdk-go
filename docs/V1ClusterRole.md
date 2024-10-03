# V1ClusterRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationRule** | Pointer to [**V1AggregationRule**](V1AggregationRule.md) |  | [optional] 
**ApiVersion** | Pointer to **string** | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources | [optional] 
**Kind** | Pointer to **string** | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds | [optional] 
**Metadata** | Pointer to [**V1ObjectMeta**](V1ObjectMeta.md) |  | [optional] 
**Rules** | [**[]V1PolicyRule**](V1PolicyRule.md) | Rules holds all the PolicyRules for this ClusterRole | 

## Methods

### NewV1ClusterRole

`func NewV1ClusterRole(rules []V1PolicyRule, ) *V1ClusterRole`

NewV1ClusterRole instantiates a new V1ClusterRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ClusterRoleWithDefaults

`func NewV1ClusterRoleWithDefaults() *V1ClusterRole`

NewV1ClusterRoleWithDefaults instantiates a new V1ClusterRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationRule

`func (o *V1ClusterRole) GetAggregationRule() V1AggregationRule`

GetAggregationRule returns the AggregationRule field if non-nil, zero value otherwise.

### GetAggregationRuleOk

`func (o *V1ClusterRole) GetAggregationRuleOk() (*V1AggregationRule, bool)`

GetAggregationRuleOk returns a tuple with the AggregationRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationRule

`func (o *V1ClusterRole) SetAggregationRule(v V1AggregationRule)`

SetAggregationRule sets AggregationRule field to given value.

### HasAggregationRule

`func (o *V1ClusterRole) HasAggregationRule() bool`

HasAggregationRule returns a boolean if a field has been set.

### GetApiVersion

`func (o *V1ClusterRole) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *V1ClusterRole) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *V1ClusterRole) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *V1ClusterRole) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetKind

`func (o *V1ClusterRole) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *V1ClusterRole) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *V1ClusterRole) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *V1ClusterRole) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *V1ClusterRole) GetMetadata() V1ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *V1ClusterRole) GetMetadataOk() (*V1ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *V1ClusterRole) SetMetadata(v V1ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *V1ClusterRole) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRules

`func (o *V1ClusterRole) GetRules() []V1PolicyRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *V1ClusterRole) GetRulesOk() (*[]V1PolicyRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *V1ClusterRole) SetRules(v []V1PolicyRule)`

SetRules sets Rules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


