package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the BrandSafetyGetResponsePagination type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandSafetyGetResponsePagination{}

// BrandSafetyGetResponsePagination Response pagination info for Brand Safety Deny List GET requests
type BrandSafetyGetResponsePagination struct {
	// The total number of deny list domains created by the advertiser
	Total *int32 `json:"total,omitempty"`
	// The maximum number of deny list domains returned from GET request
	Limit *int32 `json:"limit,omitempty"`
	// The number of deny list domains skipped
	Offset *int32 `json:"offset,omitempty"`
}

// NewBrandSafetyGetResponsePagination instantiates a new BrandSafetyGetResponsePagination object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandSafetyGetResponsePagination() *BrandSafetyGetResponsePagination {
	this := BrandSafetyGetResponsePagination{}
	return &this
}

// NewBrandSafetyGetResponsePaginationWithDefaults instantiates a new BrandSafetyGetResponsePagination object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandSafetyGetResponsePaginationWithDefaults() *BrandSafetyGetResponsePagination {
	this := BrandSafetyGetResponsePagination{}
	return &this
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *BrandSafetyGetResponsePagination) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyGetResponsePagination) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *BrandSafetyGetResponsePagination) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *BrandSafetyGetResponsePagination) SetTotal(v int32) {
	o.Total = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *BrandSafetyGetResponsePagination) GetLimit() int32 {
	if o == nil || IsNil(o.Limit) {
		var ret int32
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyGetResponsePagination) GetLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.Limit) {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *BrandSafetyGetResponsePagination) HasLimit() bool {
	if o != nil && !IsNil(o.Limit) {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int32 and assigns it to the Limit field.
func (o *BrandSafetyGetResponsePagination) SetLimit(v int32) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *BrandSafetyGetResponsePagination) GetOffset() int32 {
	if o == nil || IsNil(o.Offset) {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandSafetyGetResponsePagination) GetOffsetOk() (*int32, bool) {
	if o == nil || IsNil(o.Offset) {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *BrandSafetyGetResponsePagination) HasOffset() bool {
	if o != nil && !IsNil(o.Offset) {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *BrandSafetyGetResponsePagination) SetOffset(v int32) {
	o.Offset = &v
}

func (o BrandSafetyGetResponsePagination) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Limit) {
		toSerialize["limit"] = o.Limit
	}
	if !IsNil(o.Offset) {
		toSerialize["offset"] = o.Offset
	}
	return toSerialize, nil
}

type NullableBrandSafetyGetResponsePagination struct {
	value *BrandSafetyGetResponsePagination
	isSet bool
}

func (v NullableBrandSafetyGetResponsePagination) Get() *BrandSafetyGetResponsePagination {
	return v.value
}

func (v *NullableBrandSafetyGetResponsePagination) Set(val *BrandSafetyGetResponsePagination) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandSafetyGetResponsePagination) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandSafetyGetResponsePagination) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandSafetyGetResponsePagination(val *BrandSafetyGetResponsePagination) *NullableBrandSafetyGetResponsePagination {
	return &NullableBrandSafetyGetResponsePagination{value: val, isSet: true}
}

func (v NullableBrandSafetyGetResponsePagination) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBrandSafetyGetResponsePagination) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
