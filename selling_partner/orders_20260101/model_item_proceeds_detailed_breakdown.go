package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemProceedsDetailedBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemProceedsDetailedBreakdown{}

// ItemProceedsDetailedBreakdown Further granular breakdown of the subtotal of the proceeds breakdown, only available for TAX and DISCOUNT proceeds types.
type ItemProceedsDetailedBreakdown struct {
	// Specific classification of the further granular breakdown.   **Possible values**: `ITEM`, `SHIPPING`, `GIFT_WRAP`, `COD_FEE`, `OTHER`, `DISCOUNT`
	Subtype *string `json:"subtype,omitempty"`
	Value   *Money  `json:"value,omitempty"`
}

// NewItemProceedsDetailedBreakdown instantiates a new ItemProceedsDetailedBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemProceedsDetailedBreakdown() *ItemProceedsDetailedBreakdown {
	this := ItemProceedsDetailedBreakdown{}
	return &this
}

// NewItemProceedsDetailedBreakdownWithDefaults instantiates a new ItemProceedsDetailedBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemProceedsDetailedBreakdownWithDefaults() *ItemProceedsDetailedBreakdown {
	this := ItemProceedsDetailedBreakdown{}
	return &this
}

// GetSubtype returns the Subtype field value if set, zero value otherwise.
func (o *ItemProceedsDetailedBreakdown) GetSubtype() string {
	if o == nil || IsNil(o.Subtype) {
		var ret string
		return ret
	}
	return *o.Subtype
}

// GetSubtypeOk returns a tuple with the Subtype field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemProceedsDetailedBreakdown) GetSubtypeOk() (*string, bool) {
	if o == nil || IsNil(o.Subtype) {
		return nil, false
	}
	return o.Subtype, true
}

// HasSubtype returns a boolean if a field has been set.
func (o *ItemProceedsDetailedBreakdown) HasSubtype() bool {
	if o != nil && !IsNil(o.Subtype) {
		return true
	}

	return false
}

// SetSubtype gets a reference to the given string and assigns it to the Subtype field.
func (o *ItemProceedsDetailedBreakdown) SetSubtype(v string) {
	o.Subtype = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ItemProceedsDetailedBreakdown) GetValue() Money {
	if o == nil || IsNil(o.Value) {
		var ret Money
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemProceedsDetailedBreakdown) GetValueOk() (*Money, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ItemProceedsDetailedBreakdown) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given Money and assigns it to the Value field.
func (o *ItemProceedsDetailedBreakdown) SetValue(v Money) {
	o.Value = &v
}

func (o ItemProceedsDetailedBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subtype) {
		toSerialize["subtype"] = o.Subtype
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableItemProceedsDetailedBreakdown struct {
	value *ItemProceedsDetailedBreakdown
	isSet bool
}

func (v NullableItemProceedsDetailedBreakdown) Get() *ItemProceedsDetailedBreakdown {
	return v.value
}

func (v *NullableItemProceedsDetailedBreakdown) Set(val *ItemProceedsDetailedBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableItemProceedsDetailedBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableItemProceedsDetailedBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemProceedsDetailedBreakdown(val *ItemProceedsDetailedBreakdown) *NullableItemProceedsDetailedBreakdown {
	return &NullableItemProceedsDetailedBreakdown{value: val, isSet: true}
}

func (v NullableItemProceedsDetailedBreakdown) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemProceedsDetailedBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
