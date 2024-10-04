# V1ClaimSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceClaimName** | Pointer to **string** | ResourceClaimName is the name of a ResourceClaim object in the same namespace as this pod. | [optional] 
**ResourceClaimTemplateName** | Pointer to **string** | ResourceClaimTemplateName is the name of a ResourceClaimTemplate object in the same namespace as this pod.  The template will be used to create a new ResourceClaim, which will be bound to this pod. When this pod is deleted, the ResourceClaim will also be deleted. The pod name and resource name, along with a generated component, will be used to form a unique name for the ResourceClaim, which will be recorded in pod.status.resourceClaimStatuses.  This field is immutable and no changes will be made to the corresponding ResourceClaim by the control plane after creating the ResourceClaim. | [optional] 

## Methods

### NewV1ClaimSource

`func NewV1ClaimSource() *V1ClaimSource`

NewV1ClaimSource instantiates a new V1ClaimSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ClaimSourceWithDefaults

`func NewV1ClaimSourceWithDefaults() *V1ClaimSource`

NewV1ClaimSourceWithDefaults instantiates a new V1ClaimSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceClaimName

`func (o *V1ClaimSource) GetResourceClaimName() string`

GetResourceClaimName returns the ResourceClaimName field if non-nil, zero value otherwise.

### GetResourceClaimNameOk

`func (o *V1ClaimSource) GetResourceClaimNameOk() (*string, bool)`

GetResourceClaimNameOk returns a tuple with the ResourceClaimName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceClaimName

`func (o *V1ClaimSource) SetResourceClaimName(v string)`

SetResourceClaimName sets ResourceClaimName field to given value.

### HasResourceClaimName

`func (o *V1ClaimSource) HasResourceClaimName() bool`

HasResourceClaimName returns a boolean if a field has been set.

### GetResourceClaimTemplateName

`func (o *V1ClaimSource) GetResourceClaimTemplateName() string`

GetResourceClaimTemplateName returns the ResourceClaimTemplateName field if non-nil, zero value otherwise.

### GetResourceClaimTemplateNameOk

`func (o *V1ClaimSource) GetResourceClaimTemplateNameOk() (*string, bool)`

GetResourceClaimTemplateNameOk returns a tuple with the ResourceClaimTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceClaimTemplateName

`func (o *V1ClaimSource) SetResourceClaimTemplateName(v string)`

SetResourceClaimTemplateName sets ResourceClaimTemplateName field to given value.

### HasResourceClaimTemplateName

`func (o *V1ClaimSource) HasResourceClaimTemplateName() bool`

HasResourceClaimTemplateName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


