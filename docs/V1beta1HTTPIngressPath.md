# V1beta1HTTPIngressPath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backend** | [**V1beta1IngressBackend**](V1beta1IngressBackend.md) |  | 
**Path** | Pointer to **string** | Path is matched against the path of an incoming request. Currently it can contain characters disallowed from the conventional \&quot;path\&quot; part of a URL as defined by RFC 3986. Paths must begin with a &#39;/&#39;. When unspecified, all paths from incoming requests are matched. | [optional] 
**PathType** | Pointer to **string** | PathType determines the interpretation of the Path matching. PathType can be one of the following values: * Exact: Matches the URL path exactly. * Prefix: Matches based on a URL path prefix split by &#39;/&#39;. Matching is   done on a path element by element basis. A path element refers is the   list of labels in the path split by the &#39;/&#39; separator. A request is a   match for path p if every p is an element-wise prefix of p of the   request path. Note that if the last element of the path is a substring   of the last element in request path, it is not a match (e.g. /foo/bar   matches /foo/bar/baz, but does not match /foo/barbaz). * ImplementationSpecific: Interpretation of the Path matching is up to   the IngressClass. Implementations can treat this as a separate PathType   or treat it identically to Prefix or Exact path types. Implementations are required to support all path types. Defaults to ImplementationSpecific. | [optional] 

## Methods

### NewV1beta1HTTPIngressPath

`func NewV1beta1HTTPIngressPath(backend V1beta1IngressBackend, ) *V1beta1HTTPIngressPath`

NewV1beta1HTTPIngressPath instantiates a new V1beta1HTTPIngressPath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1beta1HTTPIngressPathWithDefaults

`func NewV1beta1HTTPIngressPathWithDefaults() *V1beta1HTTPIngressPath`

NewV1beta1HTTPIngressPathWithDefaults instantiates a new V1beta1HTTPIngressPath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackend

`func (o *V1beta1HTTPIngressPath) GetBackend() V1beta1IngressBackend`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *V1beta1HTTPIngressPath) GetBackendOk() (*V1beta1IngressBackend, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *V1beta1HTTPIngressPath) SetBackend(v V1beta1IngressBackend)`

SetBackend sets Backend field to given value.


### GetPath

`func (o *V1beta1HTTPIngressPath) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *V1beta1HTTPIngressPath) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *V1beta1HTTPIngressPath) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *V1beta1HTTPIngressPath) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPathType

`func (o *V1beta1HTTPIngressPath) GetPathType() string`

GetPathType returns the PathType field if non-nil, zero value otherwise.

### GetPathTypeOk

`func (o *V1beta1HTTPIngressPath) GetPathTypeOk() (*string, bool)`

GetPathTypeOk returns a tuple with the PathType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathType

`func (o *V1beta1HTTPIngressPath) SetPathType(v string)`

SetPathType sets PathType field to given value.

### HasPathType

`func (o *V1beta1HTTPIngressPath) HasPathType() bool`

HasPathType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


