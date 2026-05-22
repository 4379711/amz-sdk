package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the TrackingSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackingSummary{}

// TrackingSummary A package status summary.
type TrackingSummary struct {
	Status              *Status              `json:"status,omitempty"`
	TrackingDetailCodes *TrackingDetailCodes `json:"trackingDetailCodes,omitempty"`
}

// NewTrackingSummary instantiates a new TrackingSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackingSummary() *TrackingSummary {
	this := TrackingSummary{}
	return &this
}

// NewTrackingSummaryWithDefaults instantiates a new TrackingSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingSummaryWithDefaults() *TrackingSummary {
	this := TrackingSummary{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TrackingSummary) GetStatus() Status {
	if o == nil || IsNil(o.Status) {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingSummary) GetStatusOk() (*Status, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TrackingSummary) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *TrackingSummary) SetStatus(v Status) {
	o.Status = &v
}

// GetTrackingDetailCodes returns the TrackingDetailCodes field value if set, zero value otherwise.
func (o *TrackingSummary) GetTrackingDetailCodes() TrackingDetailCodes {
	if o == nil || IsNil(o.TrackingDetailCodes) {
		var ret TrackingDetailCodes
		return ret
	}
	return *o.TrackingDetailCodes
}

// GetTrackingDetailCodesOk returns a tuple with the TrackingDetailCodes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingSummary) GetTrackingDetailCodesOk() (*TrackingDetailCodes, bool) {
	if o == nil || IsNil(o.TrackingDetailCodes) {
		return nil, false
	}
	return o.TrackingDetailCodes, true
}

// HasTrackingDetailCodes returns a boolean if a field has been set.
func (o *TrackingSummary) HasTrackingDetailCodes() bool {
	if o != nil && !IsNil(o.TrackingDetailCodes) {
		return true
	}

	return false
}

// SetTrackingDetailCodes gets a reference to the given TrackingDetailCodes and assigns it to the TrackingDetailCodes field.
func (o *TrackingSummary) SetTrackingDetailCodes(v TrackingDetailCodes) {
	o.TrackingDetailCodes = &v
}

func (o TrackingSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TrackingDetailCodes) {
		toSerialize["trackingDetailCodes"] = o.TrackingDetailCodes
	}
	return toSerialize, nil
}

type NullableTrackingSummary struct {
	value *TrackingSummary
	isSet bool
}

func (v NullableTrackingSummary) Get() *TrackingSummary {
	return v.value
}

func (v *NullableTrackingSummary) Set(val *TrackingSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingSummary(val *TrackingSummary) *NullableTrackingSummary {
	return &NullableTrackingSummary{value: val, isSet: true}
}

func (v NullableTrackingSummary) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTrackingSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
