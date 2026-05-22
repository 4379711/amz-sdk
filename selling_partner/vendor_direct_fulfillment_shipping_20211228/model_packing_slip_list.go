package vendor_direct_fulfillment_shipping_20211228

import (
	"github.com/bytedance/sonic"
)

// checks if the PackingSlipList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PackingSlipList{}

// PackingSlipList A list of packing slips.
type PackingSlipList struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	// An array of packing slip objects.
	PackingSlips []PackingSlip `json:"packingSlips,omitempty"`
}

// NewPackingSlipList instantiates a new PackingSlipList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackingSlipList() *PackingSlipList {
	this := PackingSlipList{}
	return &this
}

// NewPackingSlipListWithDefaults instantiates a new PackingSlipList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackingSlipListWithDefaults() *PackingSlipList {
	this := PackingSlipList{}
	return &this
}

// GetPagination returns the Pagination field value if set, zero value otherwise.
func (o *PackingSlipList) GetPagination() Pagination {
	if o == nil || IsNil(o.Pagination) {
		var ret Pagination
		return ret
	}
	return *o.Pagination
}

// GetPaginationOk returns a tuple with the Pagination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackingSlipList) GetPaginationOk() (*Pagination, bool) {
	if o == nil || IsNil(o.Pagination) {
		return nil, false
	}
	return o.Pagination, true
}

// HasPagination returns a boolean if a field has been set.
func (o *PackingSlipList) HasPagination() bool {
	if o != nil && !IsNil(o.Pagination) {
		return true
	}

	return false
}

// SetPagination gets a reference to the given Pagination and assigns it to the Pagination field.
func (o *PackingSlipList) SetPagination(v Pagination) {
	o.Pagination = &v
}

// GetPackingSlips returns the PackingSlips field value if set, zero value otherwise.
func (o *PackingSlipList) GetPackingSlips() []PackingSlip {
	if o == nil || IsNil(o.PackingSlips) {
		var ret []PackingSlip
		return ret
	}
	return o.PackingSlips
}

// GetPackingSlipsOk returns a tuple with the PackingSlips field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PackingSlipList) GetPackingSlipsOk() ([]PackingSlip, bool) {
	if o == nil || IsNil(o.PackingSlips) {
		return nil, false
	}
	return o.PackingSlips, true
}

// HasPackingSlips returns a boolean if a field has been set.
func (o *PackingSlipList) HasPackingSlips() bool {
	if o != nil && !IsNil(o.PackingSlips) {
		return true
	}

	return false
}

// SetPackingSlips gets a reference to the given []PackingSlip and assigns it to the PackingSlips field.
func (o *PackingSlipList) SetPackingSlips(v []PackingSlip) {
	o.PackingSlips = v
}

func (o PackingSlipList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pagination) {
		toSerialize["pagination"] = o.Pagination
	}
	if !IsNil(o.PackingSlips) {
		toSerialize["packingSlips"] = o.PackingSlips
	}
	return toSerialize, nil
}

type NullablePackingSlipList struct {
	value *PackingSlipList
	isSet bool
}

func (v NullablePackingSlipList) Get() *PackingSlipList {
	return v.value
}

func (v *NullablePackingSlipList) Set(val *PackingSlipList) {
	v.value = val
	v.isSet = true
}

func (v NullablePackingSlipList) IsSet() bool {
	return v.isSet
}

func (v *NullablePackingSlipList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackingSlipList(val *PackingSlipList) *NullablePackingSlipList {
	return &NullablePackingSlipList{value: val, isSet: true}
}

func (v NullablePackingSlipList) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullablePackingSlipList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
