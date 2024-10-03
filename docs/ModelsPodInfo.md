# ModelsPodInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | **string** | container name | 
**Namespace** | **string** | namespace | 
**Pod** | **string** | pod name | 

## Methods

### NewModelsPodInfo

`func NewModelsPodInfo(container string, namespace string, pod string, ) *ModelsPodInfo`

NewModelsPodInfo instantiates a new ModelsPodInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsPodInfoWithDefaults

`func NewModelsPodInfoWithDefaults() *ModelsPodInfo`

NewModelsPodInfoWithDefaults instantiates a new ModelsPodInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *ModelsPodInfo) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ModelsPodInfo) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ModelsPodInfo) SetContainer(v string)`

SetContainer sets Container field to given value.


### GetNamespace

`func (o *ModelsPodInfo) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ModelsPodInfo) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ModelsPodInfo) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetPod

`func (o *ModelsPodInfo) GetPod() string`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *ModelsPodInfo) GetPodOk() (*string, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *ModelsPodInfo) SetPod(v string)`

SetPod sets Pod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


