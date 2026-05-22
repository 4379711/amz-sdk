package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the SpdTrackingDetailInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpdTrackingDetailInput{}

// SpdTrackingDetailInput Contains input information to update Small Parcel Delivery (SPD) tracking information.
type SpdTrackingDetailInput struct {
	// List of Small Parcel Delivery (SPD) tracking items input.
	SpdTrackingItems []SpdTrackingItemInput `json:"spdTrackingItems"`
}

type _SpdTrackingDetailInput SpdTrackingDetailInput

// NewSpdTrackingDetailInput instantiates a new SpdTrackingDetailInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpdTrackingDetailInput(spdTrackingItems []SpdTrackingItemInput) *SpdTrackingDetailInput {
	this := SpdTrackingDetailInput{}
	this.SpdTrackingItems = spdTrackingItems
	return &this
}

// NewSpdTrackingDetailInputWithDefaults instantiates a new SpdTrackingDetailInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpdTrackingDetailInputWithDefaults() *SpdTrackingDetailInput {
	this := SpdTrackingDetailInput{}
	return &this
}

// GetSpdTrackingItems returns the SpdTrackingItems field value
func (o *SpdTrackingDetailInput) GetSpdTrackingItems() []SpdTrackingItemInput {
	if o == nil {
		var ret []SpdTrackingItemInput
		return ret
	}

	return o.SpdTrackingItems
}

// GetSpdTrackingItemsOk returns a tuple with the SpdTrackingItems field value
// and a boolean to check if the value has been set.
func (o *SpdTrackingDetailInput) GetSpdTrackingItemsOk() ([]SpdTrackingItemInput, bool) {
	if o == nil {
		return nil, false
	}
	return o.SpdTrackingItems, true
}

// SetSpdTrackingItems sets field value
func (o *SpdTrackingDetailInput) SetSpdTrackingItems(v []SpdTrackingItemInput) {
	o.SpdTrackingItems = v
}

func (o SpdTrackingDetailInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["spdTrackingItems"] = o.SpdTrackingItems
	return toSerialize, nil
}

type NullableSpdTrackingDetailInput struct {
	value *SpdTrackingDetailInput
	isSet bool
}

func (v NullableSpdTrackingDetailInput) Get() *SpdTrackingDetailInput {
	return v.value
}

func (v *NullableSpdTrackingDetailInput) Set(val *SpdTrackingDetailInput) {
	v.value = val
	v.isSet = true
}

func (v NullableSpdTrackingDetailInput) IsSet() bool {
	return v.isSet
}

func (v *NullableSpdTrackingDetailInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpdTrackingDetailInput(val *SpdTrackingDetailInput) *NullableSpdTrackingDetailInput {
	return &NullableSpdTrackingDetailInput{value: val, isSet: true}
}

func (v NullableSpdTrackingDetailInput) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSpdTrackingDetailInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
