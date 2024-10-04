# V1StatefulSetOrdinals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | **int32** | start is the number representing the first replica&#39;s index. It may be used to number replicas from an alternate index (eg: 1-indexed) over the default 0-indexed names, or to orchestrate progressive movement of replicas from one StatefulSet to another. If set, replica indices will be in the range:   [.spec.ordinals.start, .spec.ordinals.start + .spec.replicas). If unset, defaults to 0. Replica indices will be in the range:   [0, .spec.replicas). | 

## Methods

### NewV1StatefulSetOrdinals

`func NewV1StatefulSetOrdinals(start int32, ) *V1StatefulSetOrdinals`

NewV1StatefulSetOrdinals instantiates a new V1StatefulSetOrdinals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1StatefulSetOrdinalsWithDefaults

`func NewV1StatefulSetOrdinalsWithDefaults() *V1StatefulSetOrdinals`

NewV1StatefulSetOrdinalsWithDefaults instantiates a new V1StatefulSetOrdinals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *V1StatefulSetOrdinals) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *V1StatefulSetOrdinals) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *V1StatefulSetOrdinals) SetStart(v int32)`

SetStart sets Start field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


