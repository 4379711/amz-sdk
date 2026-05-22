package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the ListGlobalTargetableCategoriesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListGlobalTargetableCategoriesResponse{}

// ListGlobalTargetableCategoriesResponse struct for ListGlobalTargetableCategoriesResponse
type ListGlobalTargetableCategoriesResponse struct {
	// Map containing category tree for countries passed in the request Ex: { \"US\" : \"jsonString\"}.
	CountryCategoryTree *map[string]string `json:"countryCategoryTree,omitempty"`
}

// NewListGlobalTargetableCategoriesResponse instantiates a new ListGlobalTargetableCategoriesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListGlobalTargetableCategoriesResponse() *ListGlobalTargetableCategoriesResponse {
	this := ListGlobalTargetableCategoriesResponse{}
	return &this
}

// NewListGlobalTargetableCategoriesResponseWithDefaults instantiates a new ListGlobalTargetableCategoriesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListGlobalTargetableCategoriesResponseWithDefaults() *ListGlobalTargetableCategoriesResponse {
	this := ListGlobalTargetableCategoriesResponse{}
	return &this
}

// GetCountryCategoryTree returns the CountryCategoryTree field value if set, zero value otherwise.
func (o *ListGlobalTargetableCategoriesResponse) GetCountryCategoryTree() map[string]string {
	if o == nil || IsNil(o.CountryCategoryTree) {
		var ret map[string]string
		return ret
	}
	return *o.CountryCategoryTree
}

// GetCountryCategoryTreeOk returns a tuple with the CountryCategoryTree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGlobalTargetableCategoriesResponse) GetCountryCategoryTreeOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.CountryCategoryTree) {
		return nil, false
	}
	return o.CountryCategoryTree, true
}

// HasCountryCategoryTree returns a boolean if a field has been set.
func (o *ListGlobalTargetableCategoriesResponse) HasCountryCategoryTree() bool {
	if o != nil && !IsNil(o.CountryCategoryTree) {
		return true
	}

	return false
}

// SetCountryCategoryTree gets a reference to the given map[string]string and assigns it to the CountryCategoryTree field.
func (o *ListGlobalTargetableCategoriesResponse) SetCountryCategoryTree(v map[string]string) {
	o.CountryCategoryTree = &v
}

func (o ListGlobalTargetableCategoriesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryCategoryTree) {
		toSerialize["countryCategoryTree"] = o.CountryCategoryTree
	}
	return toSerialize, nil
}

type NullableListGlobalTargetableCategoriesResponse struct {
	value *ListGlobalTargetableCategoriesResponse
	isSet bool
}

func (v NullableListGlobalTargetableCategoriesResponse) Get() *ListGlobalTargetableCategoriesResponse {
	return v.value
}

func (v *NullableListGlobalTargetableCategoriesResponse) Set(val *ListGlobalTargetableCategoriesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListGlobalTargetableCategoriesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListGlobalTargetableCategoriesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListGlobalTargetableCategoriesResponse(val *ListGlobalTargetableCategoriesResponse) *NullableListGlobalTargetableCategoriesResponse {
	return &NullableListGlobalTargetableCategoriesResponse{value: val, isSet: true}
}

func (v NullableListGlobalTargetableCategoriesResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableListGlobalTargetableCategoriesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
