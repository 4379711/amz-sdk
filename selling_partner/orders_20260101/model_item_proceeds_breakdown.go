package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemProceedsBreakdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemProceedsBreakdown{}

// ItemProceedsBreakdown Detailed proceeds breakdown for a specific order item.
type ItemProceedsBreakdown struct {
	// Category classification of the proceeds breakdown.   **Possible values**: `ITEM`, `SHIPPING`, `GIFT_WRAP`, `COD_FEE`, `OTHER`, `TAX`, `DISCOUNT`
	Type     string `json:"type"`
	Subtotal Money  `json:"subtotal"`
	// Further granular breakdown of the subtotal.
	DetailedBreakdowns []ItemProceedsDetailedBreakdown `json:"detailedBreakdowns,omitempty"`
}

// NewItemProceedsBreakdown instantiates a new ItemProceedsBreakdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemProceedsBreakdown(type_ string, subtotal Money) *ItemProceedsBreakdown {
	this := ItemProceedsBreakdown{}
	this.Type = type_
	this.Subtotal = subtotal
	return &this
}

// NewItemProceedsBreakdownWithDefaults instantiates a new ItemProceedsBreakdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemProceedsBreakdownWithDefaults() *ItemProceedsBreakdown {
	this := ItemProceedsBreakdown{}
	return &this
}

// GetType returns the Type field value
func (o *ItemProceedsBreakdown) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ItemProceedsBreakdown) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ItemProceedsBreakdown) SetType(v string) {
	o.Type = v
}

// GetSubtotal returns the Subtotal field value
func (o *ItemProceedsBreakdown) GetSubtotal() Money {
	if o == nil {
		var ret Money
		return ret
	}

	return o.Subtotal
}

// GetSubtotalOk returns a tuple with the Subtotal field value
// and a boolean to check if the value has been set.
func (o *ItemProceedsBreakdown) GetSubtotalOk() (*Money, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtotal, true
}

// SetSubtotal sets field value
func (o *ItemProceedsBreakdown) SetSubtotal(v Money) {
	o.Subtotal = v
}

// GetDetailedBreakdowns returns the DetailedBreakdowns field value if set, zero value otherwise.
func (o *ItemProceedsBreakdown) GetDetailedBreakdowns() []ItemProceedsDetailedBreakdown {
	if o == nil || IsNil(o.DetailedBreakdowns) {
		var ret []ItemProceedsDetailedBreakdown
		return ret
	}
	return o.DetailedBreakdowns
}

// GetDetailedBreakdownsOk returns a tuple with the DetailedBreakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemProceedsBreakdown) GetDetailedBreakdownsOk() ([]ItemProceedsDetailedBreakdown, bool) {
	if o == nil || IsNil(o.DetailedBreakdowns) {
		return nil, false
	}
	return o.DetailedBreakdowns, true
}

// HasDetailedBreakdowns returns a boolean if a field has been set.
func (o *ItemProceedsBreakdown) HasDetailedBreakdowns() bool {
	if o != nil && !IsNil(o.DetailedBreakdowns) {
		return true
	}

	return false
}

// SetDetailedBreakdowns gets a reference to the given []ItemProceedsDetailedBreakdown and assigns it to the DetailedBreakdowns field.
func (o *ItemProceedsBreakdown) SetDetailedBreakdowns(v []ItemProceedsDetailedBreakdown) {
	o.DetailedBreakdowns = v
}

func (o ItemProceedsBreakdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["subtotal"] = o.Subtotal
	if !IsNil(o.DetailedBreakdowns) {
		toSerialize["detailedBreakdowns"] = o.DetailedBreakdowns
	}
	return toSerialize, nil
}

type NullableItemProceedsBreakdown struct {
	value *ItemProceedsBreakdown
	isSet bool
}

func (v NullableItemProceedsBreakdown) Get() *ItemProceedsBreakdown {
	return v.value
}

func (v *NullableItemProceedsBreakdown) Set(val *ItemProceedsBreakdown) {
	v.value = val
	v.isSet = true
}

func (v NullableItemProceedsBreakdown) IsSet() bool {
	return v.isSet
}

func (v *NullableItemProceedsBreakdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemProceedsBreakdown(val *ItemProceedsBreakdown) *NullableItemProceedsBreakdown {
	return &NullableItemProceedsBreakdown{value: val, isSet: true}
}

func (v NullableItemProceedsBreakdown) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemProceedsBreakdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
