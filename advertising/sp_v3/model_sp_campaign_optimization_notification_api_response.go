package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SPCampaignOptimizationNotificationAPIResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SPCampaignOptimizationNotificationAPIResponse{}

// SPCampaignOptimizationNotificationAPIResponse struct for SPCampaignOptimizationNotificationAPIResponse
type SPCampaignOptimizationNotificationAPIResponse struct {
	// List of errors that occured when generating campaign optimization notifications.
	CampaignOptimizationRecommendationsError []RuleNotificationError `json:"CampaignOptimizationRecommendationsError,omitempty"`
	// List of successful campaign optimization notifications for campaigns.
	CampaignOptimizationNotifications []RuleNotification `json:"CampaignOptimizationNotifications,omitempty"`
}

// NewSPCampaignOptimizationNotificationAPIResponse instantiates a new SPCampaignOptimizationNotificationAPIResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSPCampaignOptimizationNotificationAPIResponse() *SPCampaignOptimizationNotificationAPIResponse {
	this := SPCampaignOptimizationNotificationAPIResponse{}
	return &this
}

// NewSPCampaignOptimizationNotificationAPIResponseWithDefaults instantiates a new SPCampaignOptimizationNotificationAPIResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSPCampaignOptimizationNotificationAPIResponseWithDefaults() *SPCampaignOptimizationNotificationAPIResponse {
	this := SPCampaignOptimizationNotificationAPIResponse{}
	return &this
}

// GetCampaignOptimizationRecommendationsError returns the CampaignOptimizationRecommendationsError field value if set, zero value otherwise.
func (o *SPCampaignOptimizationNotificationAPIResponse) GetCampaignOptimizationRecommendationsError() []RuleNotificationError {
	if o == nil || IsNil(o.CampaignOptimizationRecommendationsError) {
		var ret []RuleNotificationError
		return ret
	}
	return o.CampaignOptimizationRecommendationsError
}

// GetCampaignOptimizationRecommendationsErrorOk returns a tuple with the CampaignOptimizationRecommendationsError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPCampaignOptimizationNotificationAPIResponse) GetCampaignOptimizationRecommendationsErrorOk() ([]RuleNotificationError, bool) {
	if o == nil || IsNil(o.CampaignOptimizationRecommendationsError) {
		return nil, false
	}
	return o.CampaignOptimizationRecommendationsError, true
}

// HasCampaignOptimizationRecommendationsError returns a boolean if a field has been set.
func (o *SPCampaignOptimizationNotificationAPIResponse) HasCampaignOptimizationRecommendationsError() bool {
	if o != nil && !IsNil(o.CampaignOptimizationRecommendationsError) {
		return true
	}

	return false
}

// SetCampaignOptimizationRecommendationsError gets a reference to the given []RuleNotificationError and assigns it to the CampaignOptimizationRecommendationsError field.
func (o *SPCampaignOptimizationNotificationAPIResponse) SetCampaignOptimizationRecommendationsError(v []RuleNotificationError) {
	o.CampaignOptimizationRecommendationsError = v
}

// GetCampaignOptimizationNotifications returns the CampaignOptimizationNotifications field value if set, zero value otherwise.
func (o *SPCampaignOptimizationNotificationAPIResponse) GetCampaignOptimizationNotifications() []RuleNotification {
	if o == nil || IsNil(o.CampaignOptimizationNotifications) {
		var ret []RuleNotification
		return ret
	}
	return o.CampaignOptimizationNotifications
}

// GetCampaignOptimizationNotificationsOk returns a tuple with the CampaignOptimizationNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPCampaignOptimizationNotificationAPIResponse) GetCampaignOptimizationNotificationsOk() ([]RuleNotification, bool) {
	if o == nil || IsNil(o.CampaignOptimizationNotifications) {
		return nil, false
	}
	return o.CampaignOptimizationNotifications, true
}

// HasCampaignOptimizationNotifications returns a boolean if a field has been set.
func (o *SPCampaignOptimizationNotificationAPIResponse) HasCampaignOptimizationNotifications() bool {
	if o != nil && !IsNil(o.CampaignOptimizationNotifications) {
		return true
	}

	return false
}

// SetCampaignOptimizationNotifications gets a reference to the given []RuleNotification and assigns it to the CampaignOptimizationNotifications field.
func (o *SPCampaignOptimizationNotificationAPIResponse) SetCampaignOptimizationNotifications(v []RuleNotification) {
	o.CampaignOptimizationNotifications = v
}

func (o SPCampaignOptimizationNotificationAPIResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignOptimizationRecommendationsError) {
		toSerialize["CampaignOptimizationRecommendationsError"] = o.CampaignOptimizationRecommendationsError
	}
	if !IsNil(o.CampaignOptimizationNotifications) {
		toSerialize["CampaignOptimizationNotifications"] = o.CampaignOptimizationNotifications
	}
	return toSerialize, nil
}

type NullableSPCampaignOptimizationNotificationAPIResponse struct {
	value *SPCampaignOptimizationNotificationAPIResponse
	isSet bool
}

func (v NullableSPCampaignOptimizationNotificationAPIResponse) Get() *SPCampaignOptimizationNotificationAPIResponse {
	return v.value
}

func (v *NullableSPCampaignOptimizationNotificationAPIResponse) Set(val *SPCampaignOptimizationNotificationAPIResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSPCampaignOptimizationNotificationAPIResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSPCampaignOptimizationNotificationAPIResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPCampaignOptimizationNotificationAPIResponse(val *SPCampaignOptimizationNotificationAPIResponse) *NullableSPCampaignOptimizationNotificationAPIResponse {
	return &NullableSPCampaignOptimizationNotificationAPIResponse{value: val, isSet: true}
}

func (v NullableSPCampaignOptimizationNotificationAPIResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPCampaignOptimizationNotificationAPIResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
