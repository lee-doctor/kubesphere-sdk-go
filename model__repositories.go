/*
KubeSphere

KubeSphere OpenAPI

API version: v3.1.0
Contact: info@kubesphere.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the Repositories type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Repositories{}

// Repositories struct for Repositories
type Repositories struct {
	// It’s a fully qualified name and is an identifier of the producer of this resource's capability.
	Class *string `json:"_class,omitempty"`
	Links *Links `json:"_links,omitempty"`
	Items []RepositoriesItems `json:"items,omitempty"`
	LastPage map[string]interface{} `json:"lastPage,omitempty"`
	NextPage map[string]interface{} `json:"nextPage,omitempty"`
	// page size
	PageSize *int32 `json:"pageSize,omitempty"`
}

// NewRepositories instantiates a new Repositories object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRepositories() *Repositories {
	this := Repositories{}
	return &this
}

// NewRepositoriesWithDefaults instantiates a new Repositories object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRepositoriesWithDefaults() *Repositories {
	this := Repositories{}
	return &this
}

// GetClass returns the Class field value if set, zero value otherwise.
func (o *Repositories) GetClass() string {
	if o == nil || IsNil(o.Class) {
		var ret string
		return ret
	}
	return *o.Class
}

// GetClassOk returns a tuple with the Class field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repositories) GetClassOk() (*string, bool) {
	if o == nil || IsNil(o.Class) {
		return nil, false
	}
	return o.Class, true
}

// HasClass returns a boolean if a field has been set.
func (o *Repositories) HasClass() bool {
	if o != nil && !IsNil(o.Class) {
		return true
	}

	return false
}

// SetClass gets a reference to the given string and assigns it to the Class field.
func (o *Repositories) SetClass(v string) {
	o.Class = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Repositories) GetLinks() Links {
	if o == nil || IsNil(o.Links) {
		var ret Links
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repositories) GetLinksOk() (*Links, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Repositories) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given Links and assigns it to the Links field.
func (o *Repositories) SetLinks(v Links) {
	o.Links = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Repositories) GetItems() []RepositoriesItems {
	if o == nil || IsNil(o.Items) {
		var ret []RepositoriesItems
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repositories) GetItemsOk() ([]RepositoriesItems, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Repositories) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []RepositoriesItems and assigns it to the Items field.
func (o *Repositories) SetItems(v []RepositoriesItems) {
	o.Items = v
}

// GetLastPage returns the LastPage field value if set, zero value otherwise.
func (o *Repositories) GetLastPage() map[string]interface{} {
	if o == nil || IsNil(o.LastPage) {
		var ret map[string]interface{}
		return ret
	}
	return o.LastPage
}

// GetLastPageOk returns a tuple with the LastPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repositories) GetLastPageOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.LastPage) {
		return map[string]interface{}{}, false
	}
	return o.LastPage, true
}

// HasLastPage returns a boolean if a field has been set.
func (o *Repositories) HasLastPage() bool {
	if o != nil && !IsNil(o.LastPage) {
		return true
	}

	return false
}

// SetLastPage gets a reference to the given map[string]interface{} and assigns it to the LastPage field.
func (o *Repositories) SetLastPage(v map[string]interface{}) {
	o.LastPage = v
}

// GetNextPage returns the NextPage field value if set, zero value otherwise.
func (o *Repositories) GetNextPage() map[string]interface{} {
	if o == nil || IsNil(o.NextPage) {
		var ret map[string]interface{}
		return ret
	}
	return o.NextPage
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repositories) GetNextPageOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.NextPage) {
		return map[string]interface{}{}, false
	}
	return o.NextPage, true
}

// HasNextPage returns a boolean if a field has been set.
func (o *Repositories) HasNextPage() bool {
	if o != nil && !IsNil(o.NextPage) {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given map[string]interface{} and assigns it to the NextPage field.
func (o *Repositories) SetNextPage(v map[string]interface{}) {
	o.NextPage = v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *Repositories) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Repositories) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *Repositories) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *Repositories) SetPageSize(v int32) {
	o.PageSize = &v
}

func (o Repositories) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Repositories) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Class) {
		toSerialize["_class"] = o.Class
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.LastPage) {
		toSerialize["lastPage"] = o.LastPage
	}
	if !IsNil(o.NextPage) {
		toSerialize["nextPage"] = o.NextPage
	}
	if !IsNil(o.PageSize) {
		toSerialize["pageSize"] = o.PageSize
	}
	return toSerialize, nil
}

type NullableRepositories struct {
	value *Repositories
	isSet bool
}

func (v NullableRepositories) Get() *Repositories {
	return v.value
}

func (v *NullableRepositories) Set(val *Repositories) {
	v.value = val
	v.isSet = true
}

func (v NullableRepositories) IsSet() bool {
	return v.isSet
}

func (v *NullableRepositories) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRepositories(val *Repositories) *NullableRepositories {
	return &NullableRepositories{value: val, isSet: true}
}

func (v NullableRepositories) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRepositories) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

