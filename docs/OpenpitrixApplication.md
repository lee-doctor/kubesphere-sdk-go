# OpenpitrixApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**OpenpitrixApp**](OpenpitrixApp.md) |  | [optional] 
**Cluster** | Pointer to [**OpenpitrixCluster**](OpenpitrixCluster.md) |  | [optional] 
**Ingresses** | Pointer to [**[]V1beta1Ingress**](V1beta1Ingress.md) | application ingresses | [optional] 
**Name** | **string** | application name | 
**Services** | Pointer to [**[]V1Service**](V1Service.md) | application services | [optional] 
**Version** | Pointer to [**OpenpitrixAppVersion**](OpenpitrixAppVersion.md) |  | [optional] 
**Workloads** | Pointer to [**OpenpitrixWorkLoads**](OpenpitrixWorkLoads.md) |  | [optional] 

## Methods

### NewOpenpitrixApplication

`func NewOpenpitrixApplication(name string, ) *OpenpitrixApplication`

NewOpenpitrixApplication instantiates a new OpenpitrixApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenpitrixApplicationWithDefaults

`func NewOpenpitrixApplicationWithDefaults() *OpenpitrixApplication`

NewOpenpitrixApplicationWithDefaults instantiates a new OpenpitrixApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *OpenpitrixApplication) GetApp() OpenpitrixApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *OpenpitrixApplication) GetAppOk() (*OpenpitrixApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *OpenpitrixApplication) SetApp(v OpenpitrixApp)`

SetApp sets App field to given value.

### HasApp

`func (o *OpenpitrixApplication) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetCluster

`func (o *OpenpitrixApplication) GetCluster() OpenpitrixCluster`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *OpenpitrixApplication) GetClusterOk() (*OpenpitrixCluster, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *OpenpitrixApplication) SetCluster(v OpenpitrixCluster)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *OpenpitrixApplication) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetIngresses

`func (o *OpenpitrixApplication) GetIngresses() []V1beta1Ingress`

GetIngresses returns the Ingresses field if non-nil, zero value otherwise.

### GetIngressesOk

`func (o *OpenpitrixApplication) GetIngressesOk() (*[]V1beta1Ingress, bool)`

GetIngressesOk returns a tuple with the Ingresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngresses

`func (o *OpenpitrixApplication) SetIngresses(v []V1beta1Ingress)`

SetIngresses sets Ingresses field to given value.

### HasIngresses

`func (o *OpenpitrixApplication) HasIngresses() bool`

HasIngresses returns a boolean if a field has been set.

### GetName

`func (o *OpenpitrixApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OpenpitrixApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OpenpitrixApplication) SetName(v string)`

SetName sets Name field to given value.


### GetServices

`func (o *OpenpitrixApplication) GetServices() []V1Service`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *OpenpitrixApplication) GetServicesOk() (*[]V1Service, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *OpenpitrixApplication) SetServices(v []V1Service)`

SetServices sets Services field to given value.

### HasServices

`func (o *OpenpitrixApplication) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetVersion

`func (o *OpenpitrixApplication) GetVersion() OpenpitrixAppVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *OpenpitrixApplication) GetVersionOk() (*OpenpitrixAppVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *OpenpitrixApplication) SetVersion(v OpenpitrixAppVersion)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *OpenpitrixApplication) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWorkloads

`func (o *OpenpitrixApplication) GetWorkloads() OpenpitrixWorkLoads`

GetWorkloads returns the Workloads field if non-nil, zero value otherwise.

### GetWorkloadsOk

`func (o *OpenpitrixApplication) GetWorkloadsOk() (*OpenpitrixWorkLoads, bool)`

GetWorkloadsOk returns a tuple with the Workloads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloads

`func (o *OpenpitrixApplication) SetWorkloads(v OpenpitrixWorkLoads)`

SetWorkloads sets Workloads field to given value.

### HasWorkloads

`func (o *OpenpitrixApplication) HasWorkloads() bool`

HasWorkloads returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


