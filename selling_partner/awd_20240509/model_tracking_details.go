package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the TrackingDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackingDetails{}

// TrackingDetails Tracking details for the shipment. If using SPD transportation, this can be for each case. If not using SPD transportation, this is a single tracking entry for the entire shipment.
type TrackingDetails struct {
	CarrierCode *CarrierCode `json:"carrierCode,omitempty"`
	// The identifier that is received from transportation to uniquely identify a booking.
	BookingId string `json:"bookingId"`
}

type _TrackingDetails TrackingDetails

// NewTrackingDetails instantiates a new TrackingDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackingDetails(bookingId string) *TrackingDetails {
	this := TrackingDetails{}
	this.BookingId = bookingId
	return &this
}

// NewTrackingDetailsWithDefaults instantiates a new TrackingDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingDetailsWithDefaults() *TrackingDetails {
	this := TrackingDetails{}
	return &this
}

// GetCarrierCode returns the CarrierCode field value if set, zero value otherwise.
func (o *TrackingDetails) GetCarrierCode() CarrierCode {
	if o == nil || IsNil(o.CarrierCode) {
		var ret CarrierCode
		return ret
	}
	return *o.CarrierCode
}

// GetCarrierCodeOk returns a tuple with the CarrierCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingDetails) GetCarrierCodeOk() (*CarrierCode, bool) {
	if o == nil || IsNil(o.CarrierCode) {
		return nil, false
	}
	return o.CarrierCode, true
}

// HasCarrierCode returns a boolean if a field has been set.
func (o *TrackingDetails) HasCarrierCode() bool {
	if o != nil && !IsNil(o.CarrierCode) {
		return true
	}

	return false
}

// SetCarrierCode gets a reference to the given CarrierCode and assigns it to the CarrierCode field.
func (o *TrackingDetails) SetCarrierCode(v CarrierCode) {
	o.CarrierCode = &v
}

// GetBookingId returns the BookingId field value
func (o *TrackingDetails) GetBookingId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BookingId
}

// GetBookingIdOk returns a tuple with the BookingId field value
// and a boolean to check if the value has been set.
func (o *TrackingDetails) GetBookingIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BookingId, true
}

// SetBookingId sets field value
func (o *TrackingDetails) SetBookingId(v string) {
	o.BookingId = v
}

func (o TrackingDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CarrierCode) {
		toSerialize["carrierCode"] = o.CarrierCode
	}
	toSerialize["bookingId"] = o.BookingId
	return toSerialize, nil
}

type NullableTrackingDetails struct {
	value *TrackingDetails
	isSet bool
}

func (v NullableTrackingDetails) Get() *TrackingDetails {
	return v.value
}

func (v *NullableTrackingDetails) Set(val *TrackingDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingDetails(val *TrackingDetails) *NullableTrackingDetails {
	return &NullableTrackingDetails{value: val, isSet: true}
}

func (v NullableTrackingDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTrackingDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
