package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIAssociateOptimizationRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIAssociateOptimizationRulesRequest{}

// OptimizationRulesAPIAssociateOptimizationRulesRequest Request body for creating optimization rules associations, in bulk.
type OptimizationRulesAPIAssociateOptimizationRulesRequest struct {
	// Pairs of a rule with an advertising entity to associate together.
	Associations []OptimizationRulesAPIRuleAssociationDefinition `json:"associations"`
}

type _OptimizationRulesAPIAssociateOptimizationRulesRequest OptimizationRulesAPIAssociateOptimizationRulesRequest

// NewOptimizationRulesAPIAssociateOptimizationRulesRequest instantiates a new OptimizationRulesAPIAssociateOptimizationRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIAssociateOptimizationRulesRequest(associations []OptimizationRulesAPIRuleAssociationDefinition) *OptimizationRulesAPIAssociateOptimizationRulesRequest {
	this := OptimizationRulesAPIAssociateOptimizationRulesRequest{}
	this.Associations = associations
	return &this
}

// NewOptimizationRulesAPIAssociateOptimizationRulesRequestWithDefaults instantiates a new OptimizationRulesAPIAssociateOptimizationRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIAssociateOptimizationRulesRequestWithDefaults() *OptimizationRulesAPIAssociateOptimizationRulesRequest {
	this := OptimizationRulesAPIAssociateOptimizationRulesRequest{}
	return &this
}

// GetAssociations returns the Associations field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesRequest) GetAssociations() []OptimizationRulesAPIRuleAssociationDefinition {
	if o == nil {
		var ret []OptimizationRulesAPIRuleAssociationDefinition
		return ret
	}

	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIAssociateOptimizationRulesRequest) GetAssociationsOk() ([]OptimizationRulesAPIRuleAssociationDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Associations, true
}

// SetAssociations sets field value
func (o *OptimizationRulesAPIAssociateOptimizationRulesRequest) SetAssociations(v []OptimizationRulesAPIRuleAssociationDefinition) {
	o.Associations = v
}

func (o OptimizationRulesAPIAssociateOptimizationRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["associations"] = o.Associations
	return toSerialize, nil
}

type NullableOptimizationRulesAPIAssociateOptimizationRulesRequest struct {
	value *OptimizationRulesAPIAssociateOptimizationRulesRequest
	isSet bool
}

func (v NullableOptimizationRulesAPIAssociateOptimizationRulesRequest) Get() *OptimizationRulesAPIAssociateOptimizationRulesRequest {
	return v.value
}

func (v *NullableOptimizationRulesAPIAssociateOptimizationRulesRequest) Set(val *OptimizationRulesAPIAssociateOptimizationRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIAssociateOptimizationRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIAssociateOptimizationRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIAssociateOptimizationRulesRequest(val *OptimizationRulesAPIAssociateOptimizationRulesRequest) *NullableOptimizationRulesAPIAssociateOptimizationRulesRequest {
	return &NullableOptimizationRulesAPIAssociateOptimizationRulesRequest{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIAssociateOptimizationRulesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIAssociateOptimizationRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
