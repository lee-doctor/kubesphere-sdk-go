# V1PodResourceClaim

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name uniquely identifies this resource claim inside the pod. This must be a DNS_LABEL. | 
**Source** | Pointer to [**V1ClaimSource**](V1ClaimSource.md) |  | [optional] 

## Methods

### NewV1PodResourceClaim

`func NewV1PodResourceClaim(name string, ) *V1PodResourceClaim`

NewV1PodResourceClaim instantiates a new V1PodResourceClaim object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1PodResourceClaimWithDefaults

`func NewV1PodResourceClaimWithDefaults() *V1PodResourceClaim`

NewV1PodResourceClaimWithDefaults instantiates a new V1PodResourceClaim object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *V1PodResourceClaim) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1PodResourceClaim) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1PodResourceClaim) SetName(v string)`

SetName sets Name field to given value.


### GetSource

`func (o *V1PodResourceClaim) GetSource() V1ClaimSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V1PodResourceClaim) GetSourceOk() (*V1ClaimSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V1PodResourceClaim) SetSource(v V1ClaimSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *V1PodResourceClaim) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


