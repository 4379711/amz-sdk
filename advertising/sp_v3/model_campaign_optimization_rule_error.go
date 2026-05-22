package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the CampaignOptimizationRuleError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CampaignOptimizationRuleError{}

// CampaignOptimizationRuleError The Error Response Object.
type CampaignOptimizationRuleError struct {
	// An enumerated error code for machine use.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the response.
	Details *string `json:"details,omitempty"`
}

// NewCampaignOptimizationRuleError instantiates a new CampaignOptimizationRuleError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignOptimizationRuleError() *CampaignOptimizationRuleError {
	this := CampaignOptimizationRuleError{}
	return &this
}

// NewCampaignOptimizationRuleErrorWithDefaults instantiates a new CampaignOptimizationRuleError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignOptimizationRuleErrorWithDefaults() *CampaignOptimizationRuleError {
	this := CampaignOptimizationRuleError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *CampaignOptimizationRuleError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignOptimizationRuleError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *CampaignOptimizationRuleError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *CampaignOptimizationRuleError) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *CampaignOptimizationRuleError) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignOptimizationRuleError) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *CampaignOptimizationRuleError) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *CampaignOptimizationRuleError) SetDetails(v string) {
	o.Details = &v
}

func (o CampaignOptimizationRuleError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	return toSerialize, nil
}

type NullableCampaignOptimizationRuleError struct {
	value *CampaignOptimizationRuleError
	isSet bool
}

func (v NullableCampaignOptimizationRuleError) Get() *CampaignOptimizationRuleError {
	return v.value
}

func (v *NullableCampaignOptimizationRuleError) Set(val *CampaignOptimizationRuleError) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignOptimizationRuleError) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignOptimizationRuleError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignOptimizationRuleError(val *CampaignOptimizationRuleError) *NullableCampaignOptimizationRuleError {
	return &NullableCampaignOptimizationRuleError{value: val, isSet: true}
}

func (v NullableCampaignOptimizationRuleError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCampaignOptimizationRuleError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
