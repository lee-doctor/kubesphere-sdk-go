# V1NamespaceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]V1NamespaceCondition**](V1NamespaceCondition.md) | Represents the latest available observations of a namespace&#39;s current state. | [optional] 
**Phase** | Pointer to **string** | Phase is the current lifecycle phase of the namespace. More info: https://kubernetes.io/docs/tasks/administer-cluster/namespaces/ | [optional] 

## Methods

### NewV1NamespaceStatus

`func NewV1NamespaceStatus() *V1NamespaceStatus`

NewV1NamespaceStatus instantiates a new V1NamespaceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1NamespaceStatusWithDefaults

`func NewV1NamespaceStatusWithDefaults() *V1NamespaceStatus`

NewV1NamespaceStatusWithDefaults instantiates a new V1NamespaceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *V1NamespaceStatus) GetConditions() []V1NamespaceCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1NamespaceStatus) GetConditionsOk() (*[]V1NamespaceCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1NamespaceStatus) SetConditions(v []V1NamespaceCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1NamespaceStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetPhase

`func (o *V1NamespaceStatus) GetPhase() string`

GetPhase returns the Phase field if non-nil, zero value otherwise.

### GetPhaseOk

`func (o *V1NamespaceStatus) GetPhaseOk() (*string, bool)`

GetPhaseOk returns a tuple with the Phase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase

`func (o *V1NamespaceStatus) SetPhase(v string)`

SetPhase sets Phase field to given value.

### HasPhase

`func (o *V1NamespaceStatus) HasPhase() bool`

HasPhase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


