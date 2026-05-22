package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateSPCampaignOptimizationRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSPCampaignOptimizationRuleResponse{}

// UpdateSPCampaignOptimizationRuleResponse struct for UpdateSPCampaignOptimizationRuleResponse
type UpdateSPCampaignOptimizationRuleResponse struct {
	// The persistent rule identifier.
	CampaignOptimizationId *string `json:"campaignOptimizationId,omitempty"`
	// An enumerated success or error code for machine use.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the error, if unsuccessful
	Details *string `json:"details,omitempty"`
}

// NewUpdateSPCampaignOptimizationRuleResponse instantiates a new UpdateSPCampaignOptimizationRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSPCampaignOptimizationRuleResponse() *UpdateSPCampaignOptimizationRuleResponse {
	this := UpdateSPCampaignOptimizationRuleResponse{}
	return &this
}

// NewUpdateSPCampaignOptimizationRuleResponseWithDefaults instantiates a new UpdateSPCampaignOptimizationRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSPCampaignOptimizationRuleResponseWithDefaults() *UpdateSPCampaignOptimizationRuleResponse {
	this := UpdateSPCampaignOptimizationRuleResponse{}
	return &this
}

// GetCampaignOptimizationId returns the CampaignOptimizationId field value if set, zero value otherwise.
func (o *UpdateSPCampaignOptimizationRuleResponse) GetCampaignOptimizationId() string {
	if o == nil || IsNil(o.CampaignOptimizationId) {
		var ret string
		return ret
	}
	return *o.CampaignOptimizationId
}

// GetCampaignOptimizationIdOk returns a tuple with the CampaignOptimizationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSPCampaignOptimizationRuleResponse) GetCampaignOptimizationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CampaignOptimizationId) {
		return nil, false
	}
	return o.CampaignOptimizationId, true
}

// HasCampaignOptimizationId returns a boolean if a field has been set.
func (o *UpdateSPCampaignOptimizationRuleResponse) HasCampaignOptimizationId() bool {
	if o != nil && !IsNil(o.CampaignOptimizationId) {
		return true
	}

	return false
}

// SetCampaignOptimizationId gets a reference to the given string and assigns it to the CampaignOptimizationId field.
func (o *UpdateSPCampaignOptimizationRuleResponse) SetCampaignOptimizationId(v string) {
	o.CampaignOptimizationId = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UpdateSPCampaignOptimizationRuleResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSPCampaignOptimizationRuleResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UpdateSPCampaignOptimizationRuleResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UpdateSPCampaignOptimizationRuleResponse) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *UpdateSPCampaignOptimizationRuleResponse) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSPCampaignOptimizationRuleResponse) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *UpdateSPCampaignOptimizationRuleResponse) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *UpdateSPCampaignOptimizationRuleResponse) SetDetails(v string) {
	o.Details = &v
}

func (o UpdateSPCampaignOptimizationRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CampaignOptimizationId) {
		toSerialize["campaignOptimizationId"] = o.CampaignOptimizationId
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableUpdateSPCampaignOptimizationRuleResponse struct {
	value *UpdateSPCampaignOptimizationRuleResponse
	isSet bool
}

func (v NullableUpdateSPCampaignOptimizationRuleResponse) Get() *UpdateSPCampaignOptimizationRuleResponse {
	return v.value
}

func (v *NullableUpdateSPCampaignOptimizationRuleResponse) Set(val *UpdateSPCampaignOptimizationRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSPCampaignOptimizationRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSPCampaignOptimizationRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSPCampaignOptimizationRuleResponse(val *UpdateSPCampaignOptimizationRuleResponse) *NullableUpdateSPCampaignOptimizationRuleResponse {
	return &NullableUpdateSPCampaignOptimizationRuleResponse{value: val, isSet: true}
}

func (v NullableUpdateSPCampaignOptimizationRuleResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateSPCampaignOptimizationRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
