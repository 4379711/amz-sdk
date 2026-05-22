package sp_v3

import "github.com/bytedance/sonic"

// checks if the OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem{}

// OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem struct for OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem
type OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem struct {
	// Index of the failed disassociation pair in the request.
	Index  int32                                       `json:"index"`
	Errors []OptimizationRulesAPIOptimizationRuleError `json:"errors"`
}

type _OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem

// NewOptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem instantiates a new OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem(index int32, errors []OptimizationRulesAPIOptimizationRuleError) *OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem {
	this := OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem{}
	this.Index = index
	this.Errors = errors
	return &this
}

// NewOptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItemWithDefaults instantiates a new OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItemWithDefaults() *OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem {
	this := OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem) GetErrors() []OptimizationRulesAPIOptimizationRuleError {
	if o == nil {
		var ret []OptimizationRulesAPIOptimizationRuleError
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem) GetErrorsOk() ([]OptimizationRulesAPIOptimizationRuleError, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem) SetErrors(v []OptimizationRulesAPIOptimizationRuleError) {
	o.Errors = v
}

func (o OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["errors"] = o.Errors
	return toSerialize, nil
}

type NullableOptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem struct {
	value *OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem
	isSet bool
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem) Get() *OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem {
	return v.value
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem) Set(val *OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem) {
	v.value = val
	v.isSet = true
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem) IsSet() bool {
	return v.isSet
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem(val *OptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem) *NullableOptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem {
	return &NullableOptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem{value: val, isSet: true}
}

func (v NullableOptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOptimizationRulesAPIDisassociateOptimizationRulesResponseErrorItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
