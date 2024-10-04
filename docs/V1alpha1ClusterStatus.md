# V1alpha1ClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]V1alpha1ClusterCondition**](V1alpha1ClusterCondition.md) |  | [optional] 
**Configz** | Pointer to **map[string]bool** |  | [optional] 
**KubeSphereVersion** | Pointer to **string** |  | [optional] 
**KubernetesVersion** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int32** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** |  | [optional] 
**Zones** | Pointer to **[]string** |  | [optional] 

## Methods

### NewV1alpha1ClusterStatus

`func NewV1alpha1ClusterStatus() *V1alpha1ClusterStatus`

NewV1alpha1ClusterStatus instantiates a new V1alpha1ClusterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ClusterStatusWithDefaults

`func NewV1alpha1ClusterStatusWithDefaults() *V1alpha1ClusterStatus`

NewV1alpha1ClusterStatusWithDefaults instantiates a new V1alpha1ClusterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *V1alpha1ClusterStatus) GetConditions() []V1alpha1ClusterCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *V1alpha1ClusterStatus) GetConditionsOk() (*[]V1alpha1ClusterCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *V1alpha1ClusterStatus) SetConditions(v []V1alpha1ClusterCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *V1alpha1ClusterStatus) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetConfigz

`func (o *V1alpha1ClusterStatus) GetConfigz() map[string]bool`

GetConfigz returns the Configz field if non-nil, zero value otherwise.

### GetConfigzOk

`func (o *V1alpha1ClusterStatus) GetConfigzOk() (*map[string]bool, bool)`

GetConfigzOk returns a tuple with the Configz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigz

`func (o *V1alpha1ClusterStatus) SetConfigz(v map[string]bool)`

SetConfigz sets Configz field to given value.

### HasConfigz

`func (o *V1alpha1ClusterStatus) HasConfigz() bool`

HasConfigz returns a boolean if a field has been set.

### GetKubeSphereVersion

`func (o *V1alpha1ClusterStatus) GetKubeSphereVersion() string`

GetKubeSphereVersion returns the KubeSphereVersion field if non-nil, zero value otherwise.

### GetKubeSphereVersionOk

`func (o *V1alpha1ClusterStatus) GetKubeSphereVersionOk() (*string, bool)`

GetKubeSphereVersionOk returns a tuple with the KubeSphereVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubeSphereVersion

`func (o *V1alpha1ClusterStatus) SetKubeSphereVersion(v string)`

SetKubeSphereVersion sets KubeSphereVersion field to given value.

### HasKubeSphereVersion

`func (o *V1alpha1ClusterStatus) HasKubeSphereVersion() bool`

HasKubeSphereVersion returns a boolean if a field has been set.

### GetKubernetesVersion

`func (o *V1alpha1ClusterStatus) GetKubernetesVersion() string`

GetKubernetesVersion returns the KubernetesVersion field if non-nil, zero value otherwise.

### GetKubernetesVersionOk

`func (o *V1alpha1ClusterStatus) GetKubernetesVersionOk() (*string, bool)`

GetKubernetesVersionOk returns a tuple with the KubernetesVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesVersion

`func (o *V1alpha1ClusterStatus) SetKubernetesVersion(v string)`

SetKubernetesVersion sets KubernetesVersion field to given value.

### HasKubernetesVersion

`func (o *V1alpha1ClusterStatus) HasKubernetesVersion() bool`

HasKubernetesVersion returns a boolean if a field has been set.

### GetNodeCount

`func (o *V1alpha1ClusterStatus) GetNodeCount() int32`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *V1alpha1ClusterStatus) GetNodeCountOk() (*int32, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *V1alpha1ClusterStatus) SetNodeCount(v int32)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *V1alpha1ClusterStatus) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetRegion

`func (o *V1alpha1ClusterStatus) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *V1alpha1ClusterStatus) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *V1alpha1ClusterStatus) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *V1alpha1ClusterStatus) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetUid

`func (o *V1alpha1ClusterStatus) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *V1alpha1ClusterStatus) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *V1alpha1ClusterStatus) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *V1alpha1ClusterStatus) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetZones

`func (o *V1alpha1ClusterStatus) GetZones() []string`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *V1alpha1ClusterStatus) GetZonesOk() (*[]string, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *V1alpha1ClusterStatus) SetZones(v []string)`

SetZones sets Zones field to given value.

### HasZones

`func (o *V1alpha1ClusterStatus) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


