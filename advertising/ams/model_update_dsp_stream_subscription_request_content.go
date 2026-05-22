package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateDspStreamSubscriptionRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDspStreamSubscriptionRequestContent{}

// UpdateDspStreamSubscriptionRequestContent struct for UpdateDspStreamSubscriptionRequestContent
type UpdateDspStreamSubscriptionRequestContent struct {
	// Additional details associated with the subscription
	Notes  *string             `json:"notes,omitempty"`
	Status *UpdateEntityStatus `json:"status,omitempty"`
}

// NewUpdateDspStreamSubscriptionRequestContent instantiates a new UpdateDspStreamSubscriptionRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDspStreamSubscriptionRequestContent() *UpdateDspStreamSubscriptionRequestContent {
	this := UpdateDspStreamSubscriptionRequestContent{}
	return &this
}

// NewUpdateDspStreamSubscriptionRequestContentWithDefaults instantiates a new UpdateDspStreamSubscriptionRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDspStreamSubscriptionRequestContentWithDefaults() *UpdateDspStreamSubscriptionRequestContent {
	this := UpdateDspStreamSubscriptionRequestContent{}
	return &this
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *UpdateDspStreamSubscriptionRequestContent) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDspStreamSubscriptionRequestContent) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *UpdateDspStreamSubscriptionRequestContent) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *UpdateDspStreamSubscriptionRequestContent) SetNotes(v string) {
	o.Notes = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateDspStreamSubscriptionRequestContent) GetStatus() UpdateEntityStatus {
	if o == nil || IsNil(o.Status) {
		var ret UpdateEntityStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDspStreamSubscriptionRequestContent) GetStatusOk() (*UpdateEntityStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateDspStreamSubscriptionRequestContent) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given UpdateEntityStatus and assigns it to the Status field.
func (o *UpdateDspStreamSubscriptionRequestContent) SetStatus(v UpdateEntityStatus) {
	o.Status = &v
}

func (o UpdateDspStreamSubscriptionRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableUpdateDspStreamSubscriptionRequestContent struct {
	value *UpdateDspStreamSubscriptionRequestContent
	isSet bool
}

func (v NullableUpdateDspStreamSubscriptionRequestContent) Get() *UpdateDspStreamSubscriptionRequestContent {
	return v.value
}

func (v *NullableUpdateDspStreamSubscriptionRequestContent) Set(val *UpdateDspStreamSubscriptionRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDspStreamSubscriptionRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDspStreamSubscriptionRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDspStreamSubscriptionRequestContent(val *UpdateDspStreamSubscriptionRequestContent) *NullableUpdateDspStreamSubscriptionRequestContent {
	return &NullableUpdateDspStreamSubscriptionRequestContent{value: val, isSet: true}
}

func (v NullableUpdateDspStreamSubscriptionRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateDspStreamSubscriptionRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
