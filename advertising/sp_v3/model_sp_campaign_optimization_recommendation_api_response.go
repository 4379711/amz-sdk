package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SPCampaignOptimizationRecommendationAPIResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SPCampaignOptimizationRecommendationAPIResponse{}

// SPCampaignOptimizationRecommendationAPIResponse struct for SPCampaignOptimizationRecommendationAPIResponse
type SPCampaignOptimizationRecommendationAPIResponse struct {
	// List of campaigns eligible for optimization rule.
	CampaignOptimizationRecommendations []RuleRecommendation `json:"CampaignOptimizationRecommendations,omitempty"`
	// List of campaigns not eligible for optimization rule.
	CampaignOptimizationRecommendationsError []RuleRecommendationError `json:"CampaignOptimizationRecommendationsError,omitempty"`
}

// NewSPCampaignOptimizationRecommendationAPIResponse instantiates a new SPCampaignOptimizationRecommendationAPIResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSPCampaignOptimizationRecommendationAPIResponse() *SPCampaignOptimizationRecommendationAPIResponse {
	this := SPCampaignOptimizationRecommendationAPIResponse{}
	return &this
}

// NewSPCampaignOptimizationRecommendationAPIResponseWithDefaults instantiates a new SPCampaignOptimizationRecommendationAPIResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSPCampaignOptimizationRecommendationAPIResponseWithDefaults() *SPCampaignOptimizationRecommendationAPIResponse {
	this := SPCampaignOptimizationRecommendationAPIResponse{}
	return &this
}

// GetCampaignOptimizationRecommendations returns the CampaignOptimizationRecommendations field value if set, zero value otherwise.
func (o *SPCampaignOptimizationRecommendationAPIResponse) GetCampaignOptimizationRecommendations() []RuleRecommendation {
	if o == nil || IsNil(o.CampaignOptimizationRecommendations) {
		var ret []RuleRecommendation
		return ret
	}
	return o.CampaignOptimizationRecommendations
}

// GetCampaignOptimizationRecommendationsOk returns a tuple with the CampaignOptimizationRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPCampaignOptimizationRecommendationAPIResponse) GetCampaignOptimizationRecommendationsOk() ([]RuleRecommendation, bool) {
	if o == nil || IsNil(o.CampaignOptimizationRecommendations) {
		return nil, false
	}
	return o.CampaignOptimizationRecommendations, true
}

// HasCampaignOptimizationRecommendations returns a boolean if a field has been set.
func (o *SPCampaignOptimizationRecommendationAPIResponse) HasCampaignOptimizationRecommendations() bool {
	if o != nil && !IsNil(o.CampaignOptimizationRecommendations) {
		return true
	}

	return false
}

// SetCampaignOptimizationRecommendations gets a reference to the given []RuleRecommendation and assigns it to the CampaignOptimizationRecommendations field.
func (o *SPCampaignOptimizationRecommendationAPIResponse) SetCampaignOptimizationRecommendations(v []RuleRecommendation) {
	o.CampaignOptimizationRecommendations = v
}

// GetCampaignOptimizationRecommendationsError returns the CampaignOptimizationRecommendationsError field value if set, zero value otherwise.
func (o *SPCampaignOptimizationRecommendationAPIResponse) GetCampaignOptimizationRecommendationsError() []RuleRecommendationError {
	if o == nil || IsNil(o.CampaignOptimizationRecommendationsError) {
		var ret []RuleRecommendationError
		return ret
	}
	return o.CampaignOptimizationRecommendationsError
}

// GetCampaignOptimizationRecommendationsErrorOk returns a tuple with the CampaignOptimizationRecommendationsError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SPCampaignOptimizationRecommendationAPIResponse) GetCampaignOptimizationRecommendationsErrorOk() ([]RuleRecommendationError, bool) {
	if o == nil || IsNil(o.CampaignOptimizationRecommendationsError) {
		return nil, false
	}
	return o.CampaignOptimizationRecommendationsError, true
}

// HasCampaignOptimizationRecommendationsError returns a boolean if a field has been set.
func (o *SPCampaignOptimizationRecommendationAPIResponse) HasCampaignOptimizationRecommendationsError() bool {
	if o != nil && !IsNil(o.CampaignOptimizationRecommendationsError) {
		return true
	}

	return false
}

// SetCampaignOptimizationRecommendationsError gets a reference to the given []RuleRecommendationError and assigns it to the CampaignOptimizationRecommendationsError field.
func (o *SPCampaignOptimizationRecommendationAPIResponse) SetCampaignOptimizationRecommendationsError(v []RuleRecommendationError) {
	o.CampaignOptimizationRecommendationsError = v
}

func (o SPCampaignOptimizationRecommendationAPIResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignOptimizationRecommendations) {
		toSerialize["CampaignOptimizationRecommendations"] = o.CampaignOptimizationRecommendations
	}
	if !IsNil(o.CampaignOptimizationRecommendationsError) {
		toSerialize["CampaignOptimizationRecommendationsError"] = o.CampaignOptimizationRecommendationsError
	}
	return toSerialize, nil
}

type NullableSPCampaignOptimizationRecommendationAPIResponse struct {
	value *SPCampaignOptimizationRecommendationAPIResponse
	isSet bool
}

func (v NullableSPCampaignOptimizationRecommendationAPIResponse) Get() *SPCampaignOptimizationRecommendationAPIResponse {
	return v.value
}

func (v *NullableSPCampaignOptimizationRecommendationAPIResponse) Set(val *SPCampaignOptimizationRecommendationAPIResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSPCampaignOptimizationRecommendationAPIResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSPCampaignOptimizationRecommendationAPIResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSPCampaignOptimizationRecommendationAPIResponse(val *SPCampaignOptimizationRecommendationAPIResponse) *NullableSPCampaignOptimizationRecommendationAPIResponse {
	return &NullableSPCampaignOptimizationRecommendationAPIResponse{value: val, isSet: true}
}

func (v NullableSPCampaignOptimizationRecommendationAPIResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSPCampaignOptimizationRecommendationAPIResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
