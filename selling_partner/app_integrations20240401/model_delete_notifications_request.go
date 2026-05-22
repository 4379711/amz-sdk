package app_integrations20240401

import (
	"github.com/bytedance/sonic"
)

// checks if the DeleteNotificationsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteNotificationsRequest{}

// DeleteNotificationsRequest The request for the `deleteNotifications` operation.
type DeleteNotificationsRequest struct {
	// The unique identifier of the notification template you used to onboard your application.
	TemplateId string `json:"templateId"`
	// The unique identifier that maps each notification status to a reason code.
	DeletionReason string `json:"deletionReason"`
}

type _DeleteNotificationsRequest DeleteNotificationsRequest

// NewDeleteNotificationsRequest instantiates a new DeleteNotificationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteNotificationsRequest(templateId string, deletionReason string) *DeleteNotificationsRequest {
	this := DeleteNotificationsRequest{}
	this.TemplateId = templateId
	this.DeletionReason = deletionReason
	return &this
}

// NewDeleteNotificationsRequestWithDefaults instantiates a new DeleteNotificationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteNotificationsRequestWithDefaults() *DeleteNotificationsRequest {
	this := DeleteNotificationsRequest{}
	return &this
}

// GetTemplateId returns the TemplateId field value
func (o *DeleteNotificationsRequest) GetTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *DeleteNotificationsRequest) GetTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *DeleteNotificationsRequest) SetTemplateId(v string) {
	o.TemplateId = v
}

// GetDeletionReason returns the DeletionReason field value
func (o *DeleteNotificationsRequest) GetDeletionReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeletionReason
}

// GetDeletionReasonOk returns a tuple with the DeletionReason field value
// and a boolean to check if the value has been set.
func (o *DeleteNotificationsRequest) GetDeletionReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletionReason, true
}

// SetDeletionReason sets field value
func (o *DeleteNotificationsRequest) SetDeletionReason(v string) {
	o.DeletionReason = v
}

func (o DeleteNotificationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["templateId"] = o.TemplateId
	toSerialize["deletionReason"] = o.DeletionReason
	return toSerialize, nil
}

type NullableDeleteNotificationsRequest struct {
	value *DeleteNotificationsRequest
	isSet bool
}

func (v NullableDeleteNotificationsRequest) Get() *DeleteNotificationsRequest {
	return v.value
}

func (v *NullableDeleteNotificationsRequest) Set(val *DeleteNotificationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteNotificationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteNotificationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteNotificationsRequest(val *DeleteNotificationsRequest) *NullableDeleteNotificationsRequest {
	return &NullableDeleteNotificationsRequest{value: val, isSet: true}
}

func (v NullableDeleteNotificationsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeleteNotificationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
