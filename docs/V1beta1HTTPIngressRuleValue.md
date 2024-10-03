# V1beta1HTTPIngressRuleValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Paths** | [**[]V1beta1HTTPIngressPath**](V1beta1HTTPIngressPath.md) | A collection of paths that map requests to backends. | 

## Methods

### NewV1beta1HTTPIngressRuleValue

`func NewV1beta1HTTPIngressRuleValue(paths []V1beta1HTTPIngressPath, ) *V1beta1HTTPIngressRuleValue`

NewV1beta1HTTPIngressRuleValue instantiates a new V1beta1HTTPIngressRuleValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1beta1HTTPIngressRuleValueWithDefaults

`func NewV1beta1HTTPIngressRuleValueWithDefaults() *V1beta1HTTPIngressRuleValue`

NewV1beta1HTTPIngressRuleValueWithDefaults instantiates a new V1beta1HTTPIngressRuleValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaths

`func (o *V1beta1HTTPIngressRuleValue) GetPaths() []V1beta1HTTPIngressPath`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *V1beta1HTTPIngressRuleValue) GetPathsOk() (*[]V1beta1HTTPIngressPath, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *V1beta1HTTPIngressRuleValue) SetPaths(v []V1beta1HTTPIngressPath)`

SetPaths sets Paths field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


