# V1alpha2ComponentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HealthyBackends** | **int32** | the number of healthy backend components | 
**Label** | **map[string]interface{}** |  | 
**Name** | **string** | component name | 
**Namespace** | **string** | the name of the namespace | 
**SelfLink** | **string** | self link | 
**StartedAt** | **time.Time** | started time | 
**TotalBackends** | **int32** | the total replicas of each backend system component | 

## Methods

### NewV1alpha2ComponentStatus

`func NewV1alpha2ComponentStatus(healthyBackends int32, label map[string]interface{}, name string, namespace string, selfLink string, startedAt time.Time, totalBackends int32, ) *V1alpha2ComponentStatus`

NewV1alpha2ComponentStatus instantiates a new V1alpha2ComponentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha2ComponentStatusWithDefaults

`func NewV1alpha2ComponentStatusWithDefaults() *V1alpha2ComponentStatus`

NewV1alpha2ComponentStatusWithDefaults instantiates a new V1alpha2ComponentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealthyBackends

`func (o *V1alpha2ComponentStatus) GetHealthyBackends() int32`

GetHealthyBackends returns the HealthyBackends field if non-nil, zero value otherwise.

### GetHealthyBackendsOk

`func (o *V1alpha2ComponentStatus) GetHealthyBackendsOk() (*int32, bool)`

GetHealthyBackendsOk returns a tuple with the HealthyBackends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthyBackends

`func (o *V1alpha2ComponentStatus) SetHealthyBackends(v int32)`

SetHealthyBackends sets HealthyBackends field to given value.


### GetLabel

`func (o *V1alpha2ComponentStatus) GetLabel() map[string]interface{}`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *V1alpha2ComponentStatus) GetLabelOk() (*map[string]interface{}, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *V1alpha2ComponentStatus) SetLabel(v map[string]interface{})`

SetLabel sets Label field to given value.


### GetName

`func (o *V1alpha2ComponentStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1alpha2ComponentStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1alpha2ComponentStatus) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *V1alpha2ComponentStatus) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *V1alpha2ComponentStatus) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *V1alpha2ComponentStatus) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetSelfLink

`func (o *V1alpha2ComponentStatus) GetSelfLink() string`

GetSelfLink returns the SelfLink field if non-nil, zero value otherwise.

### GetSelfLinkOk

`func (o *V1alpha2ComponentStatus) GetSelfLinkOk() (*string, bool)`

GetSelfLinkOk returns a tuple with the SelfLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfLink

`func (o *V1alpha2ComponentStatus) SetSelfLink(v string)`

SetSelfLink sets SelfLink field to given value.


### GetStartedAt

`func (o *V1alpha2ComponentStatus) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *V1alpha2ComponentStatus) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *V1alpha2ComponentStatus) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetTotalBackends

`func (o *V1alpha2ComponentStatus) GetTotalBackends() int32`

GetTotalBackends returns the TotalBackends field if non-nil, zero value otherwise.

### GetTotalBackendsOk

`func (o *V1alpha2ComponentStatus) GetTotalBackendsOk() (*int32, bool)`

GetTotalBackendsOk returns a tuple with the TotalBackends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBackends

`func (o *V1alpha2ComponentStatus) SetTotalBackends(v int32)`

SetTotalBackends sets TotalBackends field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


