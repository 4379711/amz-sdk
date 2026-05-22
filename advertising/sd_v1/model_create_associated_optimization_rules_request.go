package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateAssociatedOptimizationRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAssociatedOptimizationRulesRequest{}

// CreateAssociatedOptimizationRulesRequest struct for CreateAssociatedOptimizationRulesRequest
type CreateAssociatedOptimizationRulesRequest struct {
	// A list of optimization rule identifiers.
	OptimizationRuleIds []string `json:"optimizationRuleIds,omitempty"`
}

// NewCreateAssociatedOptimizationRulesRequest instantiates a new CreateAssociatedOptimizationRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAssociatedOptimizationRulesRequest() *CreateAssociatedOptimizationRulesRequest {
	this := CreateAssociatedOptimizationRulesRequest{}
	return &this
}

// NewCreateAssociatedOptimizationRulesRequestWithDefaults instantiates a new CreateAssociatedOptimizationRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAssociatedOptimizationRulesRequestWithDefaults() *CreateAssociatedOptimizationRulesRequest {
	this := CreateAssociatedOptimizationRulesRequest{}
	return &this
}

// GetOptimizationRuleIds returns the OptimizationRuleIds field value if set, zero value otherwise.
func (o *CreateAssociatedOptimizationRulesRequest) GetOptimizationRuleIds() []string {
	if o == nil || IsNil(o.OptimizationRuleIds) {
		var ret []string
		return ret
	}
	return o.OptimizationRuleIds
}

// GetOptimizationRuleIdsOk returns a tuple with the OptimizationRuleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAssociatedOptimizationRulesRequest) GetOptimizationRuleIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.OptimizationRuleIds) {
		return nil, false
	}
	return o.OptimizationRuleIds, true
}

// HasOptimizationRuleIds returns a boolean if a field has been set.
func (o *CreateAssociatedOptimizationRulesRequest) HasOptimizationRuleIds() bool {
	if o != nil && !IsNil(o.OptimizationRuleIds) {
		return true
	}

	return false
}

// SetOptimizationRuleIds gets a reference to the given []string and assigns it to the OptimizationRuleIds field.
func (o *CreateAssociatedOptimizationRulesRequest) SetOptimizationRuleIds(v []string) {
	o.OptimizationRuleIds = v
}

func (o CreateAssociatedOptimizationRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OptimizationRuleIds) {
		toSerialize["optimizationRuleIds"] = o.OptimizationRuleIds
	}
	return toSerialize, nil
}

type NullableCreateAssociatedOptimizationRulesRequest struct {
	value *CreateAssociatedOptimizationRulesRequest
	isSet bool
}

func (v NullableCreateAssociatedOptimizationRulesRequest) Get() *CreateAssociatedOptimizationRulesRequest {
	return v.value
}

func (v *NullableCreateAssociatedOptimizationRulesRequest) Set(val *CreateAssociatedOptimizationRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAssociatedOptimizationRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAssociatedOptimizationRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAssociatedOptimizationRulesRequest(val *CreateAssociatedOptimizationRulesRequest) *NullableCreateAssociatedOptimizationRulesRequest {
	return &NullableCreateAssociatedOptimizationRulesRequest{value: val, isSet: true}
}

func (v NullableCreateAssociatedOptimizationRulesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateAssociatedOptimizationRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
