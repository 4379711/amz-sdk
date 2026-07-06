package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemInternationalShipping type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemInternationalShipping{}

// ItemInternationalShipping Additional requirements needed for cross-border shipping of an order item.
type ItemInternationalShipping struct {
	// Import One-Stop Shop registration number required for EU VAT compliance when shipping from outside the European Union. Sellers shipping to the EU from outside the EU must provide this IOSS number to their carrier when Amazon has collected the VAT on the sale.
	IossNumber *string `json:"iossNumber,omitempty"`
}

// NewItemInternationalShipping instantiates a new ItemInternationalShipping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemInternationalShipping() *ItemInternationalShipping {
	this := ItemInternationalShipping{}
	return &this
}

// NewItemInternationalShippingWithDefaults instantiates a new ItemInternationalShipping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemInternationalShippingWithDefaults() *ItemInternationalShipping {
	this := ItemInternationalShipping{}
	return &this
}

// GetIossNumber returns the IossNumber field value if set, zero value otherwise.
func (o *ItemInternationalShipping) GetIossNumber() string {
	if o == nil || IsNil(o.IossNumber) {
		var ret string
		return ret
	}
	return *o.IossNumber
}

// GetIossNumberOk returns a tuple with the IossNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemInternationalShipping) GetIossNumberOk() (*string, bool) {
	if o == nil || IsNil(o.IossNumber) {
		return nil, false
	}
	return o.IossNumber, true
}

// HasIossNumber returns a boolean if a field has been set.
func (o *ItemInternationalShipping) HasIossNumber() bool {
	if o != nil && !IsNil(o.IossNumber) {
		return true
	}

	return false
}

// SetIossNumber gets a reference to the given string and assigns it to the IossNumber field.
func (o *ItemInternationalShipping) SetIossNumber(v string) {
	o.IossNumber = &v
}

func (o ItemInternationalShipping) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IossNumber) {
		toSerialize["iossNumber"] = o.IossNumber
	}
	return toSerialize, nil
}

type NullableItemInternationalShipping struct {
	value *ItemInternationalShipping
	isSet bool
}

func (v NullableItemInternationalShipping) Get() *ItemInternationalShipping {
	return v.value
}

func (v *NullableItemInternationalShipping) Set(val *ItemInternationalShipping) {
	v.value = val
	v.isSet = true
}

func (v NullableItemInternationalShipping) IsSet() bool {
	return v.isSet
}

func (v *NullableItemInternationalShipping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemInternationalShipping(val *ItemInternationalShipping) *NullableItemInternationalShipping {
	return &NullableItemInternationalShipping{value: val, isSet: true}
}

func (v NullableItemInternationalShipping) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemInternationalShipping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
