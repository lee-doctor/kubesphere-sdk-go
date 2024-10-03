# V1alpha3NoScmPipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | description of pipeline | [optional] 
**DisableConcurrent** | Pointer to **bool** | Whether to prohibit the pipeline from running in parallel | [optional] 
**Discarder** | Pointer to [**V1alpha3DiscarderProperty**](V1alpha3DiscarderProperty.md) |  | [optional] 
**Jenkinsfile** | Pointer to **string** | Jenkinsfile&#39;s content&#39; | [optional] 
**Name** | **string** | name of pipeline | 
**Parameters** | Pointer to [**[]V1alpha3Parameter**](V1alpha3Parameter.md) | Parameters define of pipeline,user could pass param when run pipeline | [optional] 
**RemoteTrigger** | Pointer to [**V1alpha3RemoteTrigger**](V1alpha3RemoteTrigger.md) |  | [optional] 
**TimerTrigger** | Pointer to [**V1alpha3TimerTrigger**](V1alpha3TimerTrigger.md) |  | [optional] 

## Methods

### NewV1alpha3NoScmPipeline

`func NewV1alpha3NoScmPipeline(name string, ) *V1alpha3NoScmPipeline`

NewV1alpha3NoScmPipeline instantiates a new V1alpha3NoScmPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1alpha3NoScmPipelineWithDefaults

`func NewV1alpha3NoScmPipelineWithDefaults() *V1alpha3NoScmPipeline`

NewV1alpha3NoScmPipelineWithDefaults instantiates a new V1alpha3NoScmPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *V1alpha3NoScmPipeline) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *V1alpha3NoScmPipeline) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *V1alpha3NoScmPipeline) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *V1alpha3NoScmPipeline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisableConcurrent

`func (o *V1alpha3NoScmPipeline) GetDisableConcurrent() bool`

GetDisableConcurrent returns the DisableConcurrent field if non-nil, zero value otherwise.

### GetDisableConcurrentOk

`func (o *V1alpha3NoScmPipeline) GetDisableConcurrentOk() (*bool, bool)`

GetDisableConcurrentOk returns a tuple with the DisableConcurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableConcurrent

`func (o *V1alpha3NoScmPipeline) SetDisableConcurrent(v bool)`

SetDisableConcurrent sets DisableConcurrent field to given value.

### HasDisableConcurrent

`func (o *V1alpha3NoScmPipeline) HasDisableConcurrent() bool`

HasDisableConcurrent returns a boolean if a field has been set.

### GetDiscarder

`func (o *V1alpha3NoScmPipeline) GetDiscarder() V1alpha3DiscarderProperty`

GetDiscarder returns the Discarder field if non-nil, zero value otherwise.

### GetDiscarderOk

`func (o *V1alpha3NoScmPipeline) GetDiscarderOk() (*V1alpha3DiscarderProperty, bool)`

GetDiscarderOk returns a tuple with the Discarder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscarder

`func (o *V1alpha3NoScmPipeline) SetDiscarder(v V1alpha3DiscarderProperty)`

SetDiscarder sets Discarder field to given value.

### HasDiscarder

`func (o *V1alpha3NoScmPipeline) HasDiscarder() bool`

HasDiscarder returns a boolean if a field has been set.

### GetJenkinsfile

`func (o *V1alpha3NoScmPipeline) GetJenkinsfile() string`

GetJenkinsfile returns the Jenkinsfile field if non-nil, zero value otherwise.

### GetJenkinsfileOk

`func (o *V1alpha3NoScmPipeline) GetJenkinsfileOk() (*string, bool)`

GetJenkinsfileOk returns a tuple with the Jenkinsfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJenkinsfile

`func (o *V1alpha3NoScmPipeline) SetJenkinsfile(v string)`

SetJenkinsfile sets Jenkinsfile field to given value.

### HasJenkinsfile

`func (o *V1alpha3NoScmPipeline) HasJenkinsfile() bool`

HasJenkinsfile returns a boolean if a field has been set.

### GetName

`func (o *V1alpha3NoScmPipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1alpha3NoScmPipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1alpha3NoScmPipeline) SetName(v string)`

SetName sets Name field to given value.


### GetParameters

`func (o *V1alpha3NoScmPipeline) GetParameters() []V1alpha3Parameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *V1alpha3NoScmPipeline) GetParametersOk() (*[]V1alpha3Parameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *V1alpha3NoScmPipeline) SetParameters(v []V1alpha3Parameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *V1alpha3NoScmPipeline) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetRemoteTrigger

`func (o *V1alpha3NoScmPipeline) GetRemoteTrigger() V1alpha3RemoteTrigger`

GetRemoteTrigger returns the RemoteTrigger field if non-nil, zero value otherwise.

### GetRemoteTriggerOk

`func (o *V1alpha3NoScmPipeline) GetRemoteTriggerOk() (*V1alpha3RemoteTrigger, bool)`

GetRemoteTriggerOk returns a tuple with the RemoteTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteTrigger

`func (o *V1alpha3NoScmPipeline) SetRemoteTrigger(v V1alpha3RemoteTrigger)`

SetRemoteTrigger sets RemoteTrigger field to given value.

### HasRemoteTrigger

`func (o *V1alpha3NoScmPipeline) HasRemoteTrigger() bool`

HasRemoteTrigger returns a boolean if a field has been set.

### GetTimerTrigger

`func (o *V1alpha3NoScmPipeline) GetTimerTrigger() V1alpha3TimerTrigger`

GetTimerTrigger returns the TimerTrigger field if non-nil, zero value otherwise.

### GetTimerTriggerOk

`func (o *V1alpha3NoScmPipeline) GetTimerTriggerOk() (*V1alpha3TimerTrigger, bool)`

GetTimerTriggerOk returns a tuple with the TimerTrigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimerTrigger

`func (o *V1alpha3NoScmPipeline) SetTimerTrigger(v V1alpha3TimerTrigger)`

SetTimerTrigger sets TimerTrigger field to given value.

### HasTimerTrigger

`func (o *V1alpha3NoScmPipeline) HasTimerTrigger() bool`

HasTimerTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


