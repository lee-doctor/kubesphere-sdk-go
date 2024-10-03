# OpenpitrixWorkLoads

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Daemonsets** | Pointer to [**[]V1DaemonSet**](V1DaemonSet.md) | daemonset list | [optional] 
**Deployments** | Pointer to [**[]V1Deployment**](V1Deployment.md) | deployment list | [optional] 
**Statefulsets** | Pointer to [**[]V1StatefulSet**](V1StatefulSet.md) | statefulset list | [optional] 

## Methods

### NewOpenpitrixWorkLoads

`func NewOpenpitrixWorkLoads() *OpenpitrixWorkLoads`

NewOpenpitrixWorkLoads instantiates a new OpenpitrixWorkLoads object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenpitrixWorkLoadsWithDefaults

`func NewOpenpitrixWorkLoadsWithDefaults() *OpenpitrixWorkLoads`

NewOpenpitrixWorkLoadsWithDefaults instantiates a new OpenpitrixWorkLoads object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDaemonsets

`func (o *OpenpitrixWorkLoads) GetDaemonsets() []V1DaemonSet`

GetDaemonsets returns the Daemonsets field if non-nil, zero value otherwise.

### GetDaemonsetsOk

`func (o *OpenpitrixWorkLoads) GetDaemonsetsOk() (*[]V1DaemonSet, bool)`

GetDaemonsetsOk returns a tuple with the Daemonsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonsets

`func (o *OpenpitrixWorkLoads) SetDaemonsets(v []V1DaemonSet)`

SetDaemonsets sets Daemonsets field to given value.

### HasDaemonsets

`func (o *OpenpitrixWorkLoads) HasDaemonsets() bool`

HasDaemonsets returns a boolean if a field has been set.

### GetDeployments

`func (o *OpenpitrixWorkLoads) GetDeployments() []V1Deployment`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *OpenpitrixWorkLoads) GetDeploymentsOk() (*[]V1Deployment, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *OpenpitrixWorkLoads) SetDeployments(v []V1Deployment)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *OpenpitrixWorkLoads) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetStatefulsets

`func (o *OpenpitrixWorkLoads) GetStatefulsets() []V1StatefulSet`

GetStatefulsets returns the Statefulsets field if non-nil, zero value otherwise.

### GetStatefulsetsOk

`func (o *OpenpitrixWorkLoads) GetStatefulsetsOk() (*[]V1StatefulSet, bool)`

GetStatefulsetsOk returns a tuple with the Statefulsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatefulsets

`func (o *OpenpitrixWorkLoads) SetStatefulsets(v []V1StatefulSet)`

SetStatefulsets sets Statefulsets field to given value.

### HasStatefulsets

`func (o *OpenpitrixWorkLoads) HasStatefulsets() bool`

HasStatefulsets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


