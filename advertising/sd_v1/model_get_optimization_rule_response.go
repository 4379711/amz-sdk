package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the GetOptimizationRuleResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetOptimizationRuleResponse{}

// GetOptimizationRuleResponse struct for GetOptimizationRuleResponse
type GetOptimizationRuleResponse struct {
	OptimizationRule *OptimizationRule `json:"optimizationRule,omitempty"`
	// A list of adGroup identifiers that the optimization rule associates with.
	AdGroupIds []int64 `json:"adGroupIds,omitempty"`
}

// NewGetOptimizationRuleResponse instantiates a new GetOptimizationRuleResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetOptimizationRuleResponse() *GetOptimizationRuleResponse {
	this := GetOptimizationRuleResponse{}
	return &this
}

// NewGetOptimizationRuleResponseWithDefaults instantiates a new GetOptimizationRuleResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetOptimizationRuleResponseWithDefaults() *GetOptimizationRuleResponse {
	this := GetOptimizationRuleResponse{}
	return &this
}

// GetOptimizationRule returns the OptimizationRule field value if set, zero value otherwise.
func (o *GetOptimizationRuleResponse) GetOptimizationRule() OptimizationRule {
	if o == nil || IsNil(o.OptimizationRule) {
		var ret OptimizationRule
		return ret
	}
	return *o.OptimizationRule
}

// GetOptimizationRuleOk returns a tuple with the OptimizationRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOptimizationRuleResponse) GetOptimizationRuleOk() (*OptimizationRule, bool) {
	if o == nil || IsNil(o.OptimizationRule) {
		return nil, false
	}
	return o.OptimizationRule, true
}

// HasOptimizationRule returns a boolean if a field has been set.
func (o *GetOptimizationRuleResponse) HasOptimizationRule() bool {
	if o != nil && !IsNil(o.OptimizationRule) {
		return true
	}

	return false
}

// SetOptimizationRule gets a reference to the given OptimizationRule and assigns it to the OptimizationRule field.
func (o *GetOptimizationRuleResponse) SetOptimizationRule(v OptimizationRule) {
	o.OptimizationRule = &v
}

// GetAdGroupIds returns the AdGroupIds field value if set, zero value otherwise.
func (o *GetOptimizationRuleResponse) GetAdGroupIds() []int64 {
	if o == nil || IsNil(o.AdGroupIds) {
		var ret []int64
		return ret
	}
	return o.AdGroupIds
}

// GetAdGroupIdsOk returns a tuple with the AdGroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetOptimizationRuleResponse) GetAdGroupIdsOk() ([]int64, bool) {
	if o == nil || IsNil(o.AdGroupIds) {
		return nil, false
	}
	return o.AdGroupIds, true
}

// HasAdGroupIds returns a boolean if a field has been set.
func (o *GetOptimizationRuleResponse) HasAdGroupIds() bool {
	if o != nil && !IsNil(o.AdGroupIds) {
		return true
	}

	return false
}

// SetAdGroupIds gets a reference to the given []int64 and assigns it to the AdGroupIds field.
func (o *GetOptimizationRuleResponse) SetAdGroupIds(v []int64) {
	o.AdGroupIds = v
}

func (o GetOptimizationRuleResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptimizationRule) {
		toSerialize["optimizationRule"] = o.OptimizationRule
	}
	if !IsNil(o.AdGroupIds) {
		toSerialize["adGroupIds"] = o.AdGroupIds
	}
	return toSerialize, nil
}

type NullableGetOptimizationRuleResponse struct {
	value *GetOptimizationRuleResponse
	isSet bool
}

func (v NullableGetOptimizationRuleResponse) Get() *GetOptimizationRuleResponse {
	return v.value
}

func (v *NullableGetOptimizationRuleResponse) Set(val *GetOptimizationRuleResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetOptimizationRuleResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetOptimizationRuleResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetOptimizationRuleResponse(val *GetOptimizationRuleResponse) *NullableGetOptimizationRuleResponse {
	return &NullableGetOptimizationRuleResponse{value: val, isSet: true}
}

func (v NullableGetOptimizationRuleResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGetOptimizationRuleResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
