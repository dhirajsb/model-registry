/*
Model Registry REST API

REST API for Model Registry to create and manage ML model metadata

API version: v1alpha3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CatalogModelVersionList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogModelVersionList{}

// CatalogModelVersionList List of CatalogModelVersion entities.
type CatalogModelVersionList struct {
	// Token to use to retrieve next page of results.
	NextPageToken string `json:"nextPageToken"`
	// Maximum number of resources to return in the result.
	PageSize int32 `json:"pageSize"`
	// Number of items in result list.
	Size int32 `json:"size"`
	// Array of `CatalogModelVersion` entities.
	Items []CatalogModelVersion `json:"items,omitempty"`
}

// NewCatalogModelVersionList instantiates a new CatalogModelVersionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogModelVersionList(nextPageToken string, pageSize int32, size int32) *CatalogModelVersionList {
	this := CatalogModelVersionList{}
	this.NextPageToken = nextPageToken
	this.PageSize = pageSize
	this.Size = size
	return &this
}

// NewCatalogModelVersionListWithDefaults instantiates a new CatalogModelVersionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogModelVersionListWithDefaults() *CatalogModelVersionList {
	this := CatalogModelVersionList{}
	return &this
}

// GetNextPageToken returns the NextPageToken field value
func (o *CatalogModelVersionList) GetNextPageToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
func (o *CatalogModelVersionList) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPageToken, true
}

// SetNextPageToken sets field value
func (o *CatalogModelVersionList) SetNextPageToken(v string) {
	o.NextPageToken = v
}

// GetPageSize returns the PageSize field value
func (o *CatalogModelVersionList) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *CatalogModelVersionList) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *CatalogModelVersionList) SetPageSize(v int32) {
	o.PageSize = v
}

// GetSize returns the Size field value
func (o *CatalogModelVersionList) GetSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *CatalogModelVersionList) GetSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *CatalogModelVersionList) SetSize(v int32) {
	o.Size = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *CatalogModelVersionList) GetItems() []CatalogModelVersion {
	if o == nil || IsNil(o.Items) {
		var ret []CatalogModelVersion
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogModelVersionList) GetItemsOk() ([]CatalogModelVersion, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *CatalogModelVersionList) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []CatalogModelVersion and assigns it to the Items field.
func (o *CatalogModelVersionList) SetItems(v []CatalogModelVersion) {
	o.Items = v
}

func (o CatalogModelVersionList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogModelVersionList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nextPageToken"] = o.NextPageToken
	toSerialize["pageSize"] = o.PageSize
	toSerialize["size"] = o.Size
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableCatalogModelVersionList struct {
	value *CatalogModelVersionList
	isSet bool
}

func (v NullableCatalogModelVersionList) Get() *CatalogModelVersionList {
	return v.value
}

func (v *NullableCatalogModelVersionList) Set(val *CatalogModelVersionList) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogModelVersionList) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogModelVersionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogModelVersionList(val *CatalogModelVersionList) *NullableCatalogModelVersionList {
	return &NullableCatalogModelVersionList{value: val, isSet: true}
}

func (v NullableCatalogModelVersionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogModelVersionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
