package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the OptimizationRulesAPIDisassociateOptimizationRulesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIDisassociateOptimizationRulesRequest{}

// OptimizationRulesAPIDisassociateOptimizationRulesRequest Request body for disassociating optimization rules from advertising entities, in bulk.
type OptimizationRulesAPIDisassociateOptimizationRulesRequest struct {
	// Pairs of a rule with an advertising entity to disassociate from each other.
	Associations []OptimizationRulesAPIRuleAssociationDefinition `json:"associations"`
}

type _OptimizationRulesAPIDisassociateOptimizationRulesRequest OptimizationRulesAPIDisassociateOptimizationRulesRequest

// NewOptimizationRulesAPIDisassociateOptimizationRulesRequest instantiates a new OptimizationRulesAPIDisassociateOptimizationRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIDisassociateOptimizationRulesRequest(associations []OptimizationRulesAPIRuleAssociationDefinition) *OptimizationRulesAPIDisassociateOptimizationRulesRequest {
	this := OptimizationRulesAPIDisassociateOptimizationRulesRequest{}
	this.Associations = associations
	return &this
}

// NewOptimizationRulesAPIDisassociateOptimizationRulesRequestWithDefaults instantiates a new OptimizationRulesAPIDisassociateOptimizationRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIDisassociateOptimizationRulesRequestWithDefaults() *OptimizationRulesAPIDisassociateOptimizationRulesRequest {
	this := OptimizationRulesAPIDisassociateOptimizationRulesRequest{}
	return &this
}

// GetAssociations returns the Associations field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesRequest) GetAssociations() []OptimizationRulesAPIRuleAssociationDefinition {
	if o == nil {
		var ret []OptimizationRulesAPIRuleAssociationDefinition
		return ret
	}

	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesRequest) GetAssociationsOk() ([]OptimizationRulesAPIRuleAssociationDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Associations, true
}

// SetAssociations sets field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesRequest) SetAssociations(v []OptimizationRulesAPIRuleAssociationDefinition) {
	o.Associations = v
}

func (o OptimizationRulesAPIDisassociateOptimizationRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["associations"] = o.Associations
	return toSerialize, nil
}

type NullableOptimizationRulesAPIDisassociateOptimizationRulesRequest struct {
	value *OptimizationRulesAPIDisassociateOptimizationRulesRequest
	isSet bool
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesRequest) Get() *OptimizationRulesAPIDisassociateOptimizationRulesRequest {
	return v.value
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesRequest) Set(val *OptimizationRulesAPIDisassociateOptimizationRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIDisassociateOptimizationRulesRequest(val *OptimizationRulesAPIDisassociateOptimizationRulesRequest) *NullableOptimizationRulesAPIDisassociateOptimizationRulesRequest {
	return &NullableOptimizationRulesAPIDisassociateOptimizationRulesRequest{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
