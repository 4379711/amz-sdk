package ams

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateMessagesSubscriptionRequestContent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMessagesSubscriptionRequestContent{}

// UpdateMessagesSubscriptionRequestContent struct for UpdateMessagesSubscriptionRequestContent
type UpdateMessagesSubscriptionRequestContent struct {
	// Additional details which can be used to identity the  Subscription. (Max Size of 128 characters)
	Notes *string `json:"notes,omitempty"`
	// version of the  Subscription to apply the update to. This should match the current active version for the selected consumerSubscriptionId.
	Version float32             `json:"version"`
	Status  *UpdateEntityStatus `json:"status,omitempty"`
}

type _UpdateMessagesSubscriptionRequestContent UpdateMessagesSubscriptionRequestContent

// NewUpdateMessagesSubscriptionRequestContent instantiates a new UpdateMessagesSubscriptionRequestContent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMessagesSubscriptionRequestContent(version float32) *UpdateMessagesSubscriptionRequestContent {
	this := UpdateMessagesSubscriptionRequestContent{}
	this.Version = version
	return &this
}

// NewUpdateMessagesSubscriptionRequestContentWithDefaults instantiates a new UpdateMessagesSubscriptionRequestContent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMessagesSubscriptionRequestContentWithDefaults() *UpdateMessagesSubscriptionRequestContent {
	this := UpdateMessagesSubscriptionRequestContent{}
	return &this
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *UpdateMessagesSubscriptionRequestContent) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMessagesSubscriptionRequestContent) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *UpdateMessagesSubscriptionRequestContent) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *UpdateMessagesSubscriptionRequestContent) SetNotes(v string) {
	o.Notes = &v
}

// GetVersion returns the Version field value
func (o *UpdateMessagesSubscriptionRequestContent) GetVersion() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *UpdateMessagesSubscriptionRequestContent) GetVersionOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *UpdateMessagesSubscriptionRequestContent) SetVersion(v float32) {
	o.Version = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateMessagesSubscriptionRequestContent) GetStatus() UpdateEntityStatus {
	if o == nil || IsNil(o.Status) {
		var ret UpdateEntityStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMessagesSubscriptionRequestContent) GetStatusOk() (*UpdateEntityStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateMessagesSubscriptionRequestContent) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given UpdateEntityStatus and assigns it to the Status field.
func (o *UpdateMessagesSubscriptionRequestContent) SetStatus(v UpdateEntityStatus) {
	o.Status = &v
}

func (o UpdateMessagesSubscriptionRequestContent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	toSerialize["version"] = o.Version
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableUpdateMessagesSubscriptionRequestContent struct {
	value *UpdateMessagesSubscriptionRequestContent
	isSet bool
}

func (v NullableUpdateMessagesSubscriptionRequestContent) Get() *UpdateMessagesSubscriptionRequestContent {
	return v.value
}

func (v *NullableUpdateMessagesSubscriptionRequestContent) Set(val *UpdateMessagesSubscriptionRequestContent) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMessagesSubscriptionRequestContent) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMessagesSubscriptionRequestContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMessagesSubscriptionRequestContent(val *UpdateMessagesSubscriptionRequestContent) *NullableUpdateMessagesSubscriptionRequestContent {
	return &NullableUpdateMessagesSubscriptionRequestContent{value: val, isSet: true}
}

func (v NullableUpdateMessagesSubscriptionRequestContent) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateMessagesSubscriptionRequestContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
