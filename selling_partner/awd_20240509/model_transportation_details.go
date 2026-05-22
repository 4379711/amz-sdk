package awd_20240509

import (
	"github.com/bytedance/sonic"
)

// checks if the TransportationDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransportationDetails{}

// TransportationDetails Transportation details for the shipment.
type TransportationDetails struct {
	// Tracking details for the shipment. If using SPD transportation, this can be for each case. If not using SPD transportation, this is a single tracking entry for the entire shipment.
	TrackingDetails []TrackingDetails `json:"trackingDetails"`
}

type _TransportationDetails TransportationDetails

// NewTransportationDetails instantiates a new TransportationDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportationDetails(trackingDetails []TrackingDetails) *TransportationDetails {
	this := TransportationDetails{}
	this.TrackingDetails = trackingDetails
	return &this
}

// NewTransportationDetailsWithDefaults instantiates a new TransportationDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportationDetailsWithDefaults() *TransportationDetails {
	this := TransportationDetails{}
	return &this
}

// GetTrackingDetails returns the TrackingDetails field value
func (o *TransportationDetails) GetTrackingDetails() []TrackingDetails {
	if o == nil {
		var ret []TrackingDetails
		return ret
	}

	return o.TrackingDetails
}

// GetTrackingDetailsOk returns a tuple with the TrackingDetails field value
// and a boolean to check if the value has been set.
func (o *TransportationDetails) GetTrackingDetailsOk() ([]TrackingDetails, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrackingDetails, true
}

// SetTrackingDetails sets field value
func (o *TransportationDetails) SetTrackingDetails(v []TrackingDetails) {
	o.TrackingDetails = v
}

func (o TransportationDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["trackingDetails"] = o.TrackingDetails
	return toSerialize, nil
}

type NullableTransportationDetails struct {
	value *TransportationDetails
	isSet bool
}

func (v NullableTransportationDetails) Get() *TransportationDetails {
	return v.value
}

func (v *NullableTransportationDetails) Set(val *TransportationDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportationDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportationDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportationDetails(val *TransportationDetails) *NullableTransportationDetails {
	return &NullableTransportationDetails{value: val, isSet: true}
}

func (v NullableTransportationDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTransportationDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
