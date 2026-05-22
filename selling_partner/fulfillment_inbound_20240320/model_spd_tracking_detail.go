package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the SpdTrackingDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpdTrackingDetail{}

// SpdTrackingDetail Contains information related to Small Parcel Delivery (SPD) shipment tracking.
type SpdTrackingDetail struct {
	// List of Small Parcel Delivery (SPD) tracking items.
	SpdTrackingItems []SpdTrackingItem `json:"spdTrackingItems,omitempty"`
}

// NewSpdTrackingDetail instantiates a new SpdTrackingDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpdTrackingDetail() *SpdTrackingDetail {
	this := SpdTrackingDetail{}
	return &this
}

// NewSpdTrackingDetailWithDefaults instantiates a new SpdTrackingDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpdTrackingDetailWithDefaults() *SpdTrackingDetail {
	this := SpdTrackingDetail{}
	return &this
}

// GetSpdTrackingItems returns the SpdTrackingItems field value if set, zero value otherwise.
func (o *SpdTrackingDetail) GetSpdTrackingItems() []SpdTrackingItem {
	if o == nil || IsNil(o.SpdTrackingItems) {
		var ret []SpdTrackingItem
		return ret
	}
	return o.SpdTrackingItems
}

// GetSpdTrackingItemsOk returns a tuple with the SpdTrackingItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpdTrackingDetail) GetSpdTrackingItemsOk() ([]SpdTrackingItem, bool) {
	if o == nil || IsNil(o.SpdTrackingItems) {
		return nil, false
	}
	return o.SpdTrackingItems, true
}

// HasSpdTrackingItems returns a boolean if a field has been set.
func (o *SpdTrackingDetail) HasSpdTrackingItems() bool {
	if o != nil && !IsNil(o.SpdTrackingItems) {
		return true
	}

	return false
}

// SetSpdTrackingItems gets a reference to the given []SpdTrackingItem and assigns it to the SpdTrackingItems field.
func (o *SpdTrackingDetail) SetSpdTrackingItems(v []SpdTrackingItem) {
	o.SpdTrackingItems = v
}

func (o SpdTrackingDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SpdTrackingItems) {
		toSerialize["spdTrackingItems"] = o.SpdTrackingItems
	}
	return toSerialize, nil
}

type NullableSpdTrackingDetail struct {
	value *SpdTrackingDetail
	isSet bool
}

func (v NullableSpdTrackingDetail) Get() *SpdTrackingDetail {
	return v.value
}

func (v *NullableSpdTrackingDetail) Set(val *SpdTrackingDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableSpdTrackingDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableSpdTrackingDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpdTrackingDetail(val *SpdTrackingDetail) *NullableSpdTrackingDetail {
	return &NullableSpdTrackingDetail{value: val, isSet: true}
}

func (v NullableSpdTrackingDetail) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSpdTrackingDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
