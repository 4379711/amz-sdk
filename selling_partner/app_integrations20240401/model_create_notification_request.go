package app_integrations20240401

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateNotificationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNotificationRequest{}

// CreateNotificationRequest The request for the `createNotification` operation.
type CreateNotificationRequest struct {
	// The unique identifier of the notification template you used to onboard your application.
	TemplateId string `json:"templateId"`
	// The dynamic parameters required by the notification templated specified by `templateId`.
	NotificationParameters map[string]map[string]interface{} `json:"notificationParameters"`
	// An encrypted marketplace identifier for the posted notification.
	MarketplaceId *string `json:"marketplaceId,omitempty"`
}

type _CreateNotificationRequest CreateNotificationRequest

// NewCreateNotificationRequest instantiates a new CreateNotificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNotificationRequest(templateId string, notificationParameters map[string]map[string]interface{}) *CreateNotificationRequest {
	this := CreateNotificationRequest{}
	this.TemplateId = templateId
	this.NotificationParameters = notificationParameters
	return &this
}

// NewCreateNotificationRequestWithDefaults instantiates a new CreateNotificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNotificationRequestWithDefaults() *CreateNotificationRequest {
	this := CreateNotificationRequest{}
	return &this
}

// GetTemplateId returns the TemplateId field value
func (o *CreateNotificationRequest) GetTemplateId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *CreateNotificationRequest) GetTemplateIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *CreateNotificationRequest) SetTemplateId(v string) {
	o.TemplateId = v
}

// GetNotificationParameters returns the NotificationParameters field value
func (o *CreateNotificationRequest) GetNotificationParameters() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.NotificationParameters
}

// GetNotificationParametersOk returns a tuple with the NotificationParameters field value
// and a boolean to check if the value has been set.
func (o *CreateNotificationRequest) GetNotificationParametersOk() (map[string]map[string]interface{}, bool) {
	if o == nil {
		return map[string]map[string]interface{}{}, false
	}
	return o.NotificationParameters, true
}

// SetNotificationParameters sets field value
func (o *CreateNotificationRequest) SetNotificationParameters(v map[string]map[string]interface{}) {
	o.NotificationParameters = v
}

// GetMarketplaceId returns the MarketplaceId field value if set, zero value otherwise.
func (o *CreateNotificationRequest) GetMarketplaceId() string {
	if o == nil || IsNil(o.MarketplaceId) {
		var ret string
		return ret
	}
	return *o.MarketplaceId
}

// GetMarketplaceIdOk returns a tuple with the MarketplaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNotificationRequest) GetMarketplaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.MarketplaceId) {
		return nil, false
	}
	return o.MarketplaceId, true
}

// HasMarketplaceId returns a boolean if a field has been set.
func (o *CreateNotificationRequest) HasMarketplaceId() bool {
	if o != nil && !IsNil(o.MarketplaceId) {
		return true
	}

	return false
}

// SetMarketplaceId gets a reference to the given string and assigns it to the MarketplaceId field.
func (o *CreateNotificationRequest) SetMarketplaceId(v string) {
	o.MarketplaceId = &v
}

func (o CreateNotificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["templateId"] = o.TemplateId
	toSerialize["notificationParameters"] = o.NotificationParameters
	if !IsNil(o.MarketplaceId) {
		toSerialize["marketplaceId"] = o.MarketplaceId
	}
	return toSerialize, nil
}

type NullableCreateNotificationRequest struct {
	value *CreateNotificationRequest
	isSet bool
}

func (v NullableCreateNotificationRequest) Get() *CreateNotificationRequest {
	return v.value
}

func (v *NullableCreateNotificationRequest) Set(val *CreateNotificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNotificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNotificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNotificationRequest(val *CreateNotificationRequest) *NullableCreateNotificationRequest {
	return &NullableCreateNotificationRequest{value: val, isSet: true}
}

func (v NullableCreateNotificationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateNotificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
