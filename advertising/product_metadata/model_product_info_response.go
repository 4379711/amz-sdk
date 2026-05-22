package product_metadata

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductInfoResponse{}

// ProductInfoResponse struct for ProductInfoResponse
type ProductInfoResponse struct {
	ProductInfoList []ProductInfoModel `json:"ProductInfoList,omitempty"`
	// Pagination token for later requests with specific sort type to use as the page index instead. Empty cursorToken means no further data is present at Server side.
	CursorToken *string `json:"cursorToken,omitempty"`
	// Returned listing under AMAZON, BRAND, KDP or GLOBAL_STORE catalog
	CatalogType *string `json:"catalogType,omitempty"`
	TotalCount  *int32  `json:"totalCount,omitempty"`
}

// NewProductInfoResponse instantiates a new ProductInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductInfoResponse() *ProductInfoResponse {
	this := ProductInfoResponse{}
	return &this
}

// NewProductInfoResponseWithDefaults instantiates a new ProductInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductInfoResponseWithDefaults() *ProductInfoResponse {
	this := ProductInfoResponse{}
	return &this
}

// GetProductInfoList returns the ProductInfoList field value if set, zero value otherwise.
func (o *ProductInfoResponse) GetProductInfoList() []ProductInfoModel {
	if o == nil || IsNil(o.ProductInfoList) {
		var ret []ProductInfoModel
		return ret
	}
	return o.ProductInfoList
}

// GetProductInfoListOk returns a tuple with the ProductInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoResponse) GetProductInfoListOk() ([]ProductInfoModel, bool) {
	if o == nil || IsNil(o.ProductInfoList) {
		return nil, false
	}
	return o.ProductInfoList, true
}

// HasProductInfoList returns a boolean if a field has been set.
func (o *ProductInfoResponse) HasProductInfoList() bool {
	if o != nil && !IsNil(o.ProductInfoList) {
		return true
	}

	return false
}

// SetProductInfoList gets a reference to the given []ProductInfoModel and assigns it to the ProductInfoList field.
func (o *ProductInfoResponse) SetProductInfoList(v []ProductInfoModel) {
	o.ProductInfoList = v
}

// GetCursorToken returns the CursorToken field value if set, zero value otherwise.
func (o *ProductInfoResponse) GetCursorToken() string {
	if o == nil || IsNil(o.CursorToken) {
		var ret string
		return ret
	}
	return *o.CursorToken
}

// GetCursorTokenOk returns a tuple with the CursorToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoResponse) GetCursorTokenOk() (*string, bool) {
	if o == nil || IsNil(o.CursorToken) {
		return nil, false
	}
	return o.CursorToken, true
}

// HasCursorToken returns a boolean if a field has been set.
func (o *ProductInfoResponse) HasCursorToken() bool {
	if o != nil && !IsNil(o.CursorToken) {
		return true
	}

	return false
}

// SetCursorToken gets a reference to the given string and assigns it to the CursorToken field.
func (o *ProductInfoResponse) SetCursorToken(v string) {
	o.CursorToken = &v
}

// GetCatalogType returns the CatalogType field value if set, zero value otherwise.
func (o *ProductInfoResponse) GetCatalogType() string {
	if o == nil || IsNil(o.CatalogType) {
		var ret string
		return ret
	}
	return *o.CatalogType
}

// GetCatalogTypeOk returns a tuple with the CatalogType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoResponse) GetCatalogTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CatalogType) {
		return nil, false
	}
	return o.CatalogType, true
}

// HasCatalogType returns a boolean if a field has been set.
func (o *ProductInfoResponse) HasCatalogType() bool {
	if o != nil && !IsNil(o.CatalogType) {
		return true
	}

	return false
}

// SetCatalogType gets a reference to the given string and assigns it to the CatalogType field.
func (o *ProductInfoResponse) SetCatalogType(v string) {
	o.CatalogType = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ProductInfoResponse) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductInfoResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ProductInfoResponse) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ProductInfoResponse) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o ProductInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProductInfoList) {
		toSerialize["ProductInfoList"] = o.ProductInfoList
	}
	if !IsNil(o.CursorToken) {
		toSerialize["cursorToken"] = o.CursorToken
	}
	if !IsNil(o.CatalogType) {
		toSerialize["catalogType"] = o.CatalogType
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableProductInfoResponse struct {
	value *ProductInfoResponse
	isSet bool
}

func (v NullableProductInfoResponse) Get() *ProductInfoResponse {
	return v.value
}

func (v *NullableProductInfoResponse) Set(val *ProductInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductInfoResponse(val *ProductInfoResponse) *NullableProductInfoResponse {
	return &NullableProductInfoResponse{value: val, isSet: true}
}

func (v NullableProductInfoResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
