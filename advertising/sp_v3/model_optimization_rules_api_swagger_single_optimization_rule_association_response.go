package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPISingleOptimizationRuleAssociationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPISingleOptimizationRuleAssociationResponse{}

// OptimizationRulesAPISingleOptimizationRuleAssociationResponse Response object for operations involving associating a single optimization rule.
type OptimizationRulesAPISingleOptimizationRuleAssociationResponse struct {
	// An enumerated success or error code for machine use.
	Code *string `json:"code,omitempty"`
	// A human-readable description of the error, if unsuccessful.
	Details *string `json:"details,omitempty"`
	// The rule identifier.
	OptimizationRuleId *string `json:"optimizationRuleId,omitempty"`
}

// NewOptimizationRulesAPISingleOptimizationRuleAssociationResponse instantiates a new OptimizationRulesAPISingleOptimizationRuleAssociationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPISingleOptimizationRuleAssociationResponse() *OptimizationRulesAPISingleOptimizationRuleAssociationResponse {
	this := OptimizationRulesAPISingleOptimizationRuleAssociationResponse{}
	return &this
}

// NewOptimizationRulesAPISingleOptimizationRuleAssociationResponseWithDefaults instantiates a new OptimizationRulesAPISingleOptimizationRuleAssociationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPISingleOptimizationRuleAssociationResponseWithDefaults() *OptimizationRulesAPISingleOptimizationRuleAssociationResponse {
	this := OptimizationRulesAPISingleOptimizationRuleAssociationResponse{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *OptimizationRulesAPISingleOptimizationRuleAssociationResponse) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISingleOptimizationRuleAssociationResponse) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *OptimizationRulesAPISingleOptimizationRuleAssociationResponse) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *OptimizationRulesAPISingleOptimizationRuleAssociationResponse) SetCode(v string) {
	o.Code = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *OptimizationRulesAPISingleOptimizationRuleAssociationResponse) GetDetails() string {
	if o == nil || IsNil(o.Details) {
		var ret string
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISingleOptimizationRuleAssociationResponse) GetDetailsOk() (*string, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *OptimizationRulesAPISingleOptimizationRuleAssociationResponse) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given string and assigns it to the Details field.
func (o *OptimizationRulesAPISingleOptimizationRuleAssociationResponse) SetDetails(v string) {
	o.Details = &v
}

// GetOptimizationRuleId returns the OptimizationRuleId field value if set, zero value otherwise.
func (o *OptimizationRulesAPISingleOptimizationRuleAssociationResponse) GetOptimizationRuleId() string {
	if o == nil || IsNil(o.OptimizationRuleId) {
		var ret string
		return ret
	}
	return *o.OptimizationRuleId
}

// GetOptimizationRuleIdOk returns a tuple with the OptimizationRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPISingleOptimizationRuleAssociationResponse) GetOptimizationRuleIdOk() (*string, bool) {
	if o == nil || IsNil(o.OptimizationRuleId) {
		return nil, false
	}
	return o.OptimizationRuleId, true
}

// HasOptimizationRuleId returns a boolean if a field has been set.
func (o *OptimizationRulesAPISingleOptimizationRuleAssociationResponse) HasOptimizationRuleId() bool {
	if o != nil && !IsNil(o.OptimizationRuleId) {
		return true
	}

	return false
}

// SetOptimizationRuleId gets a reference to the given string and assigns it to the OptimizationRuleId field.
func (o *OptimizationRulesAPISingleOptimizationRuleAssociationResponse) SetOptimizationRuleId(v string) {
	o.OptimizationRuleId = &v
}

func (o OptimizationRulesAPISingleOptimizationRuleAssociationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.OptimizationRuleId) {
		toSerialize["optimizationRuleId"] = o.OptimizationRuleId
	}
	return toSerialize, nil
}

type NullableOptimizationRulesAPISingleOptimizationRuleAssociationResponse struct {
	value *OptimizationRulesAPISingleOptimizationRuleAssociationResponse
	isSet bool
}

func (v NullableOptimizationRulesAPISingleOptimizationRuleAssociationResponse) Get() *OptimizationRulesAPISingleOptimizationRuleAssociationResponse {
	return v.value
}

func (v *NullableOptimizationRulesAPISingleOptimizationRuleAssociationResponse) Set(val *OptimizationRulesAPISingleOptimizationRuleAssociationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPISingleOptimizationRuleAssociationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPISingleOptimizationRuleAssociationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPISingleOptimizationRuleAssociationResponse(val *OptimizationRulesAPISingleOptimizationRuleAssociationResponse) *NullableOptimizationRulesAPISingleOptimizationRuleAssociationResponse {
	return &NullableOptimizationRulesAPISingleOptimizationRuleAssociationResponse{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPISingleOptimizationRuleAssociationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPISingleOptimizationRuleAssociationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
