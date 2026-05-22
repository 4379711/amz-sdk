package easy_ship_20220323

import (
	"github.com/bytedance/sonic"
)

// checks if the TrackingDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackingDetails{}

// TrackingDetails Representation of tracking metadata.
type TrackingDetails struct {
	// A string of up to 255 characters.
	TrackingId *string `json:"trackingId,omitempty"`
}

// NewTrackingDetails instantiates a new TrackingDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackingDetails() *TrackingDetails {
	this := TrackingDetails{}
	return &this
}

// NewTrackingDetailsWithDefaults instantiates a new TrackingDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingDetailsWithDefaults() *TrackingDetails {
	this := TrackingDetails{}
	return &this
}

// GetTrackingId returns the TrackingId field value if set, zero value otherwise.
func (o *TrackingDetails) GetTrackingId() string {
	if o == nil || IsNil(o.TrackingId) {
		var ret string
		return ret
	}
	return *o.TrackingId
}

// GetTrackingIdOk returns a tuple with the TrackingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingDetails) GetTrackingIdOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingId) {
		return nil, false
	}
	return o.TrackingId, true
}

// HasTrackingId returns a boolean if a field has been set.
func (o *TrackingDetails) HasTrackingId() bool {
	if o != nil && !IsNil(o.TrackingId) {
		return true
	}

	return false
}

// SetTrackingId gets a reference to the given string and assigns it to the TrackingId field.
func (o *TrackingDetails) SetTrackingId(v string) {
	o.TrackingId = &v
}

func (o TrackingDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrackingId) {
		toSerialize["trackingId"] = o.TrackingId
	}
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
