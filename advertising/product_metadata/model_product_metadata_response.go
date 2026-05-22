package product_metadata

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductMetadataResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductMetadataResponse{}

// ProductMetadataResponse struct for ProductMetadataResponse
type ProductMetadataResponse struct {
	// Pagination token for later requests with specific sort type to use as the page index instead. Empty cursorToken means no further data is present at Server side.
	CursorToken         *string                `json:"cursorToken,omitempty"`
	ProductMetadataList []ProductMetadataModel `json:"ProductMetadataList,omitempty"`
}

// NewProductMetadataResponse instantiates a new ProductMetadataResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductMetadataResponse() *ProductMetadataResponse {
	this := ProductMetadataResponse{}
	return &this
}

// NewProductMetadataResponseWithDefaults instantiates a new ProductMetadataResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductMetadataResponseWithDefaults() *ProductMetadataResponse {
	this := ProductMetadataResponse{}
	return &this
}

// GetCursorToken returns the CursorToken field value if set, zero value otherwise.
func (o *ProductMetadataResponse) GetCursorToken() string {
	if o == nil || IsNil(o.CursorToken) {
		var ret string
		return ret
	}
	return *o.CursorToken
}

// GetCursorTokenOk returns a tuple with the CursorToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataResponse) GetCursorTokenOk() (*string, bool) {
	if o == nil || IsNil(o.CursorToken) {
		return nil, false
	}
	return o.CursorToken, true
}

// HasCursorToken returns a boolean if a field has been set.
func (o *ProductMetadataResponse) HasCursorToken() bool {
	if o != nil && !IsNil(o.CursorToken) {
		return true
	}

	return false
}

// SetCursorToken gets a reference to the given string and assigns it to the CursorToken field.
func (o *ProductMetadataResponse) SetCursorToken(v string) {
	o.CursorToken = &v
}

// GetProductMetadataList returns the ProductMetadataList field value if set, zero value otherwise.
func (o *ProductMetadataResponse) GetProductMetadataList() []ProductMetadataModel {
	if o == nil || IsNil(o.ProductMetadataList) {
		var ret []ProductMetadataModel
		return ret
	}
	return o.ProductMetadataList
}

// GetProductMetadataListOk returns a tuple with the ProductMetadataList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductMetadataResponse) GetProductMetadataListOk() ([]ProductMetadataModel, bool) {
	if o == nil || IsNil(o.ProductMetadataList) {
		return nil, false
	}
	return o.ProductMetadataList, true
}

// HasProductMetadataList returns a boolean if a field has been set.
func (o *ProductMetadataResponse) HasProductMetadataList() bool {
	if o != nil && !IsNil(o.ProductMetadataList) {
		return true
	}

	return false
}

// SetProductMetadataList gets a reference to the given []ProductMetadataModel and assigns it to the ProductMetadataList field.
func (o *ProductMetadataResponse) SetProductMetadataList(v []ProductMetadataModel) {
	o.ProductMetadataList = v
}

func (o ProductMetadataResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CursorToken) {
		toSerialize["cursorToken"] = o.CursorToken
	}
	if !IsNil(o.ProductMetadataList) {
		toSerialize["ProductMetadataList"] = o.ProductMetadataList
	}
	return toSerialize, nil
}

type NullableProductMetadataResponse struct {
	value *ProductMetadataResponse
	isSet bool
}

func (v NullableProductMetadataResponse) Get() *ProductMetadataResponse {
	return v.value
}

func (v *NullableProductMetadataResponse) Set(val *ProductMetadataResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductMetadataResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductMetadataResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductMetadataResponse(val *ProductMetadataResponse) *NullableProductMetadataResponse {
	return &NullableProductMetadataResponse{value: val, isSet: true}
}

func (v NullableProductMetadataResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductMetadataResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
