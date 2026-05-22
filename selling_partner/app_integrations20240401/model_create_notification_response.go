package app_integrations20240401

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateNotificationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNotificationResponse{}

// CreateNotificationResponse The response for the `createNotification` operation.
type CreateNotificationResponse struct {
	// The unique identifier assigned to each notification.
	NotificationId *string `json:"notificationId,omitempty"`
}

// NewCreateNotificationResponse instantiates a new CreateNotificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNotificationResponse() *CreateNotificationResponse {
	this := CreateNotificationResponse{}
	return &this
}

// NewCreateNotificationResponseWithDefaults instantiates a new CreateNotificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNotificationResponseWithDefaults() *CreateNotificationResponse {
	this := CreateNotificationResponse{}
	return &this
}

// GetNotificationId returns the NotificationId field value if set, zero value otherwise.
func (o *CreateNotificationResponse) GetNotificationId() string {
	if o == nil || IsNil(o.NotificationId) {
		var ret string
		return ret
	}
	return *o.NotificationId
}

// GetNotificationIdOk returns a tuple with the NotificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotificationResponse) GetNotificationIdOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationId) {
		return nil, false
	}
	return o.NotificationId, true
}

// HasNotificationId returns a boolean if a field has been set.
func (o *CreateNotificationResponse) HasNotificationId() bool {
	if o != nil && !IsNil(o.NotificationId) {
		return true
	}

	return false
}

// SetNotificationId gets a reference to the given string and assigns it to the NotificationId field.
func (o *CreateNotificationResponse) SetNotificationId(v string) {
	o.NotificationId = &v
}

func (o CreateNotificationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotificationId) {
		toSerialize["notificationId"] = o.NotificationId
	}
	return toSerialize, nil
}

type NullableCreateNotificationResponse struct {
	value *CreateNotificationResponse
	isSet bool
}

func (v NullableCreateNotificationResponse) Get() *CreateNotificationResponse {
	return v.value
}

func (v *NullableCreateNotificationResponse) Set(val *CreateNotificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNotificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNotificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNotificationResponse(val *CreateNotificationResponse) *NullableCreateNotificationResponse {
	return &NullableCreateNotificationResponse{value: val, isSet: true}
}

func (v NullableCreateNotificationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateNotificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
