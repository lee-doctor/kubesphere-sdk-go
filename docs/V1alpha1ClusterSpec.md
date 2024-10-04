# V1alpha1ClusterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to **string** |  | [optional] 
**Connection** | Pointer to [**V1alpha1Connection**](V1alpha1Connection.md) |  | [optional] 
**Enable** | Pointer to **bool** |  | [optional] 
**ExternalKubeAPIEnabled** | Pointer to **bool** |  | [optional] 
**JoinFederation** | Pointer to **bool** |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 

## Methods

### NewV1alpha1ClusterSpec

`func NewV1alpha1ClusterSpec() *V1alpha1ClusterSpec`

NewV1alpha1ClusterSpec instantiates a new V1alpha1ClusterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha1ClusterSpecWithDefaults

`func NewV1alpha1ClusterSpecWithDefaults() *V1alpha1ClusterSpec`

NewV1alpha1ClusterSpecWithDefaults instantiates a new V1alpha1ClusterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *V1alpha1ClusterSpec) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V1alpha1ClusterSpec) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V1alpha1ClusterSpec) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *V1alpha1ClusterSpec) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetConnection

`func (o *V1alpha1ClusterSpec) GetConnection() V1alpha1Connection`

GetConnection returns the Connection field if non-nil, zero value otherwise.

### GetConnectionOk

`func (o *V1alpha1ClusterSpec) GetConnectionOk() (*V1alpha1Connection, bool)`

GetConnectionOk returns a tuple with the Connection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnection

`func (o *V1alpha1ClusterSpec) SetConnection(v V1alpha1Connection)`

SetConnection sets Connection field to given value.

### HasConnection

`func (o *V1alpha1ClusterSpec) HasConnection() bool`

HasConnection returns a boolean if a field has been set.

### GetEnable

`func (o *V1alpha1ClusterSpec) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *V1alpha1ClusterSpec) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *V1alpha1ClusterSpec) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *V1alpha1ClusterSpec) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetExternalKubeAPIEnabled

`func (o *V1alpha1ClusterSpec) GetExternalKubeAPIEnabled() bool`

GetExternalKubeAPIEnabled returns the ExternalKubeAPIEnabled field if non-nil, zero value otherwise.

### GetExternalKubeAPIEnabledOk

`func (o *V1alpha1ClusterSpec) GetExternalKubeAPIEnabledOk() (*bool, bool)`

GetExternalKubeAPIEnabledOk returns a tuple with the ExternalKubeAPIEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalKubeAPIEnabled

`func (o *V1alpha1ClusterSpec) SetExternalKubeAPIEnabled(v bool)`

SetExternalKubeAPIEnabled sets ExternalKubeAPIEnabled field to given value.

### HasExternalKubeAPIEnabled

`func (o *V1alpha1ClusterSpec) HasExternalKubeAPIEnabled() bool`

HasExternalKubeAPIEnabled returns a boolean if a field has been set.

### GetJoinFederation

`func (o *V1alpha1ClusterSpec) GetJoinFederation() bool`

GetJoinFederation returns the JoinFederation field if non-nil, zero value otherwise.

### GetJoinFederationOk

`func (o *V1alpha1ClusterSpec) GetJoinFederationOk() (*bool, bool)`

GetJoinFederationOk returns a tuple with the JoinFederation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinFederation

`func (o *V1alpha1ClusterSpec) SetJoinFederation(v bool)`

SetJoinFederation sets JoinFederation field to given value.

### HasJoinFederation

`func (o *V1alpha1ClusterSpec) HasJoinFederation() bool`

HasJoinFederation returns a boolean if a field has been set.

### GetProvider

`func (o *V1alpha1ClusterSpec) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *V1alpha1ClusterSpec) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *V1alpha1ClusterSpec) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *V1alpha1ClusterSpec) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


