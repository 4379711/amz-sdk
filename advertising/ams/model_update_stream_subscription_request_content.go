package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateStreamSubscriptionRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateStreamSubscriptionRequestContent{}

// UpdateStreamSubscriptionRequestContent struct for UpdateStreamSubscriptionRequestContent
type UpdateStreamSubscriptionRequestContent struct {
	// Additional details associated with the subscription
	Notes  *string             `json:"notes,omitempty"`
	Status *UpdateEntityStatus `json:"status,omitempty"`
}

// NewUpdateStreamSubscriptionRequestContent instantiates a new UpdateStreamSubscriptionRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStreamSubscriptionRequestContent() *UpdateStreamSubscriptionRequestContent {
	this := UpdateStreamSubscriptionRequestContent{}
	return &this
}

// NewUpdateStreamSubscriptionRequestContentWithDefaults instantiates a new UpdateStreamSubscriptionRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStreamSubscriptionRequestContentWithDefaults() *UpdateStreamSubscriptionRequestContent {
	this := UpdateStreamSubscriptionRequestContent{}
	return &this
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *UpdateStreamSubscriptionRequestContent) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStreamSubscriptionRequestContent) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *UpdateStreamSubscriptionRequestContent) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *UpdateStreamSubscriptionRequestContent) SetNotes(v string) {
	o.Notes = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateStreamSubscriptionRequestContent) GetStatus() UpdateEntityStatus {
	if o == nil || IsNil(o.Status) {
		var ret UpdateEntityStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStreamSubscriptionRequestContent) GetStatusOk() (*UpdateEntityStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateStreamSubscriptionRequestContent) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given UpdateEntityStatus and assigns it to the Status field.
func (o *UpdateStreamSubscriptionRequestContent) SetStatus(v UpdateEntityStatus) {
	o.Status = &v
}

func (o UpdateStreamSubscriptionRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableUpdateStreamSubscriptionRequestContent struct {
	value *UpdateStreamSubscriptionRequestContent
	isSet bool
}

func (v NullableUpdateStreamSubscriptionRequestContent) Get() *UpdateStreamSubscriptionRequestContent {
	return v.value
}

func (v *NullableUpdateStreamSubscriptionRequestContent) Set(val *UpdateStreamSubscriptionRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStreamSubscriptionRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStreamSubscriptionRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStreamSubscriptionRequestContent(val *UpdateStreamSubscriptionRequestContent) *NullableUpdateStreamSubscriptionRequestContent {
	return &NullableUpdateStreamSubscriptionRequestContent{value: val, isSet: true}
}

func (v NullableUpdateStreamSubscriptionRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateStreamSubscriptionRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
