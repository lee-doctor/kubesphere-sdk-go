# Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]DataErrors**](DataErrors.md) |  | [optional] 
**Jenkinsfile** | Pointer to **string** | jenkinsfile | [optional] 
**Result** | Pointer to **string** | result e.g. success | [optional] 

## Methods

### NewData

`func NewData() *Data`

NewData instantiates a new Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataWithDefaults

`func NewDataWithDefaults() *Data`

NewDataWithDefaults instantiates a new Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *Data) GetErrors() []DataErrors`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Data) GetErrorsOk() (*[]DataErrors, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Data) SetErrors(v []DataErrors)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Data) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetJenkinsfile

`func (o *Data) GetJenkinsfile() string`

GetJenkinsfile returns the Jenkinsfile field if non-nil, zero value otherwise.

### GetJenkinsfileOk

`func (o *Data) GetJenkinsfileOk() (*string, bool)`

GetJenkinsfileOk returns a tuple with the Jenkinsfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJenkinsfile

`func (o *Data) SetJenkinsfile(v string)`

SetJenkinsfile sets Jenkinsfile field to given value.

### HasJenkinsfile

`func (o *Data) HasJenkinsfile() bool`

HasJenkinsfile returns a boolean if a field has been set.

### GetResult

`func (o *Data) GetResult() string`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *Data) GetResultOk() (*string, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *Data) SetResult(v string)`

SetResult sets Result field to given value.

### HasResult

`func (o *Data) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


