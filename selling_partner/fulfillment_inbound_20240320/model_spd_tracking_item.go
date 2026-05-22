package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the SpdTrackingItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpdTrackingItem{}

// SpdTrackingItem Contains information used to track and identify a Small Parcel Delivery (SPD) item.
type SpdTrackingItem struct {
	// The ID provided by Amazon that identifies a given box. This ID is comprised of the external shipment ID (which is generated after transportation has been confirmed) and the index of the box.
	BoxId *string `json:"boxId,omitempty"`
	// The tracking ID associated with each box in a non-Amazon partnered Small Parcel Delivery (SPD) shipment.
	TrackingId *string `json:"trackingId,omitempty"`
	// Whether or not Amazon has validated the tracking number. If more than 24 hours have passed and the status is not yet 'VALIDATED', please verify the number and update if necessary. Possible values: `VALIDATED`, `NOT_VALIDATED`.
	TrackingNumberValidationStatus *string `json:"trackingNumberValidationStatus,omitempty"`
}

// NewSpdTrackingItem instantiates a new SpdTrackingItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpdTrackingItem() *SpdTrackingItem {
	this := SpdTrackingItem{}
	return &this
}

// NewSpdTrackingItemWithDefaults instantiates a new SpdTrackingItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpdTrackingItemWithDefaults() *SpdTrackingItem {
	this := SpdTrackingItem{}
	return &this
}

// GetBoxId returns the BoxId field value if set, zero value otherwise.
func (o *SpdTrackingItem) GetBoxId() string {
	if o == nil || IsNil(o.BoxId) {
		var ret string
		return ret
	}
	return *o.BoxId
}

// GetBoxIdOk returns a tuple with the BoxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpdTrackingItem) GetBoxIdOk() (*string, bool) {
	if o == nil || IsNil(o.BoxId) {
		return nil, false
	}
	return o.BoxId, true
}

// HasBoxId returns a boolean if a field has been set.
func (o *SpdTrackingItem) HasBoxId() bool {
	if o != nil && !IsNil(o.BoxId) {
		return true
	}

	return false
}

// SetBoxId gets a reference to the given string and assigns it to the BoxId field.
func (o *SpdTrackingItem) SetBoxId(v string) {
	o.BoxId = &v
}

// GetTrackingId returns the TrackingId field value if set, zero value otherwise.
func (o *SpdTrackingItem) GetTrackingId() string {
	if o == nil || IsNil(o.TrackingId) {
		var ret string
		return ret
	}
	return *o.TrackingId
}

// GetTrackingIdOk returns a tuple with the TrackingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpdTrackingItem) GetTrackingIdOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingId) {
		return nil, false
	}
	return o.TrackingId, true
}

// HasTrackingId returns a boolean if a field has been set.
func (o *SpdTrackingItem) HasTrackingId() bool {
	if o != nil && !IsNil(o.TrackingId) {
		return true
	}

	return false
}

// SetTrackingId gets a reference to the given string and assigns it to the TrackingId field.
func (o *SpdTrackingItem) SetTrackingId(v string) {
	o.TrackingId = &v
}

// GetTrackingNumberValidationStatus returns the TrackingNumberValidationStatus field value if set, zero value otherwise.
func (o *SpdTrackingItem) GetTrackingNumberValidationStatus() string {
	if o == nil || IsNil(o.TrackingNumberValidationStatus) {
		var ret string
		return ret
	}
	return *o.TrackingNumberValidationStatus
}

// GetTrackingNumberValidationStatusOk returns a tuple with the TrackingNumberValidationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpdTrackingItem) GetTrackingNumberValidationStatusOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingNumberValidationStatus) {
		return nil, false
	}
	return o.TrackingNumberValidationStatus, true
}

// HasTrackingNumberValidationStatus returns a boolean if a field has been set.
func (o *SpdTrackingItem) HasTrackingNumberValidationStatus() bool {
	if o != nil && !IsNil(o.TrackingNumberValidationStatus) {
		return true
	}

	return false
}

// SetTrackingNumberValidationStatus gets a reference to the given string and assigns it to the TrackingNumberValidationStatus field.
func (o *SpdTrackingItem) SetTrackingNumberValidationStatus(v string) {
	o.TrackingNumberValidationStatus = &v
}

func (o SpdTrackingItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BoxId) {
		toSerialize["boxId"] = o.BoxId
	}
	if !IsNil(o.TrackingId) {
		toSerialize["trackingId"] = o.TrackingId
	}
	if !IsNil(o.TrackingNumberValidationStatus) {
		toSerialize["trackingNumberValidationStatus"] = o.TrackingNumberValidationStatus
	}
	return toSerialize, nil
}

type NullableSpdTrackingItem struct {
	value *SpdTrackingItem
	isSet bool
}

func (v NullableSpdTrackingItem) Get() *SpdTrackingItem {
	return v.value
}

func (v *NullableSpdTrackingItem) Set(val *SpdTrackingItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSpdTrackingItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSpdTrackingItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpdTrackingItem(val *SpdTrackingItem) *NullableSpdTrackingItem {
	return &NullableSpdTrackingItem{value: val, isSet: true}
}

func (v NullableSpdTrackingItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSpdTrackingItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
