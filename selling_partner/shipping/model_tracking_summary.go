package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the TrackingSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackingSummary{}

// TrackingSummary The tracking summary.
type TrackingSummary struct {
	// The derived status based on the events in the eventHistory.
	Status *string `json:"status,omitempty"`
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
func (o *TrackingSummary) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrackingSummary) GetStatusOk() (*string, bool) {
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

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TrackingSummary) SetStatus(v string) {
	o.Status = &v
}

func (o TrackingSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
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
