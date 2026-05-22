package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the SpdTrackingItemInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpdTrackingItemInput{}

// SpdTrackingItemInput Small Parcel Delivery (SPD) tracking items input information.
type SpdTrackingItemInput struct {
	// The ID provided by Amazon that identifies a given box. This ID is comprised of the external shipment ID (which is generated after transportation has been confirmed) and the index of the box.
	BoxId string `json:"boxId"`
	// The tracking Id associated with each box in a non-Amazon partnered Small Parcel Delivery (SPD) shipment. The seller must provide this information.
	TrackingId string `json:"trackingId"`
}

type _SpdTrackingItemInput SpdTrackingItemInput

// NewSpdTrackingItemInput instantiates a new SpdTrackingItemInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpdTrackingItemInput(boxId string, trackingId string) *SpdTrackingItemInput {
	this := SpdTrackingItemInput{}
	this.BoxId = boxId
	this.TrackingId = trackingId
	return &this
}

// NewSpdTrackingItemInputWithDefaults instantiates a new SpdTrackingItemInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpdTrackingItemInputWithDefaults() *SpdTrackingItemInput {
	this := SpdTrackingItemInput{}
	return &this
}

// GetBoxId returns the BoxId field value
func (o *SpdTrackingItemInput) GetBoxId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BoxId
}

// GetBoxIdOk returns a tuple with the BoxId field value
// and a boolean to check if the value has been set.
func (o *SpdTrackingItemInput) GetBoxIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BoxId, true
}

// SetBoxId sets field value
func (o *SpdTrackingItemInput) SetBoxId(v string) {
	o.BoxId = v
}

// GetTrackingId returns the TrackingId field value
func (o *SpdTrackingItemInput) GetTrackingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrackingId
}

// GetTrackingIdOk returns a tuple with the TrackingId field value
// and a boolean to check if the value has been set.
func (o *SpdTrackingItemInput) GetTrackingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrackingId, true
}

// SetTrackingId sets field value
func (o *SpdTrackingItemInput) SetTrackingId(v string) {
	o.TrackingId = v
}

func (o SpdTrackingItemInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["boxId"] = o.BoxId
	toSerialize["trackingId"] = o.TrackingId
	return toSerialize, nil
}

type NullableSpdTrackingItemInput struct {
	value *SpdTrackingItemInput
	isSet bool
}

func (v NullableSpdTrackingItemInput) Get() *SpdTrackingItemInput {
	return v.value
}

func (v *NullableSpdTrackingItemInput) Set(val *SpdTrackingItemInput) {
	v.value = val
	v.isSet = true
}

func (v NullableSpdTrackingItemInput) IsSet() bool {
	return v.isSet
}

func (v *NullableSpdTrackingItemInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpdTrackingItemInput(val *SpdTrackingItemInput) *NullableSpdTrackingItemInput {
	return &NullableSpdTrackingItemInput{value: val, isSet: true}
}

func (v NullableSpdTrackingItemInput) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSpdTrackingItemInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
