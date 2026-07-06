package orders_20260101

import (
	"github.com/bytedance/sonic"
)

// checks if the ItemProceeds type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ItemProceeds{}

// ItemProceeds The money that the seller receives from the sale of this specific item.
type ItemProceeds struct {
	ProceedsTotal *Money `json:"proceedsTotal,omitempty"`
	// The breakdown of proceeds.
	Breakdowns []ItemProceedsBreakdown `json:"breakdowns,omitempty"`
}

// NewItemProceeds instantiates a new ItemProceeds object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemProceeds() *ItemProceeds {
	this := ItemProceeds{}
	return &this
}

// NewItemProceedsWithDefaults instantiates a new ItemProceeds object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemProceedsWithDefaults() *ItemProceeds {
	this := ItemProceeds{}
	return &this
}

// GetProceedsTotal returns the ProceedsTotal field value if set, zero value otherwise.
func (o *ItemProceeds) GetProceedsTotal() Money {
	if o == nil || IsNil(o.ProceedsTotal) {
		var ret Money
		return ret
	}
	return *o.ProceedsTotal
}

// GetProceedsTotalOk returns a tuple with the ProceedsTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemProceeds) GetProceedsTotalOk() (*Money, bool) {
	if o == nil || IsNil(o.ProceedsTotal) {
		return nil, false
	}
	return o.ProceedsTotal, true
}

// HasProceedsTotal returns a boolean if a field has been set.
func (o *ItemProceeds) HasProceedsTotal() bool {
	if o != nil && !IsNil(o.ProceedsTotal) {
		return true
	}

	return false
}

// SetProceedsTotal gets a reference to the given Money and assigns it to the ProceedsTotal field.
func (o *ItemProceeds) SetProceedsTotal(v Money) {
	o.ProceedsTotal = &v
}

// GetBreakdowns returns the Breakdowns field value if set, zero value otherwise.
func (o *ItemProceeds) GetBreakdowns() []ItemProceedsBreakdown {
	if o == nil || IsNil(o.Breakdowns) {
		var ret []ItemProceedsBreakdown
		return ret
	}
	return o.Breakdowns
}

// GetBreakdownsOk returns a tuple with the Breakdowns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ItemProceeds) GetBreakdownsOk() ([]ItemProceedsBreakdown, bool) {
	if o == nil || IsNil(o.Breakdowns) {
		return nil, false
	}
	return o.Breakdowns, true
}

// HasBreakdowns returns a boolean if a field has been set.
func (o *ItemProceeds) HasBreakdowns() bool {
	if o != nil && !IsNil(o.Breakdowns) {
		return true
	}

	return false
}

// SetBreakdowns gets a reference to the given []ItemProceedsBreakdown and assigns it to the Breakdowns field.
func (o *ItemProceeds) SetBreakdowns(v []ItemProceedsBreakdown) {
	o.Breakdowns = v
}

func (o ItemProceeds) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProceedsTotal) {
		toSerialize["proceedsTotal"] = o.ProceedsTotal
	}
	if !IsNil(o.Breakdowns) {
		toSerialize["breakdowns"] = o.Breakdowns
	}
	return toSerialize, nil
}

type NullableItemProceeds struct {
	value *ItemProceeds
	isSet bool
}

func (v NullableItemProceeds) Get() *ItemProceeds {
	return v.value
}

func (v *NullableItemProceeds) Set(val *ItemProceeds) {
	v.value = val
	v.isSet = true
}

func (v NullableItemProceeds) IsSet() bool {
	return v.isSet
}

func (v *NullableItemProceeds) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemProceeds(val *ItemProceeds) *NullableItemProceeds {
	return &NullableItemProceeds{value: val, isSet: true}
}

func (v NullableItemProceeds) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableItemProceeds) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
