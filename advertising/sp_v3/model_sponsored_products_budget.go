package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsBudget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsBudget{}

// SponsoredProductsBudget struct for SponsoredProductsBudget
type SponsoredProductsBudget struct {
	BudgetType SponsoredProductsBudgetType `json:"budgetType"`
	// Monetary value
	Budget float64 `json:"budget"`
	// Monetary value
	EffectiveBudget *float64 `json:"effectiveBudget,omitempty"`
}

type _SponsoredProductsBudget SponsoredProductsBudget

// NewSponsoredProductsBudget instantiates a new SponsoredProductsBudget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsBudget(budgetType SponsoredProductsBudgetType, budget float64) *SponsoredProductsBudget {
	this := SponsoredProductsBudget{}
	this.BudgetType = budgetType
	this.Budget = budget
	return &this
}

// NewSponsoredProductsBudgetWithDefaults instantiates a new SponsoredProductsBudget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsBudgetWithDefaults() *SponsoredProductsBudget {
	this := SponsoredProductsBudget{}
	return &this
}

// GetBudgetType returns the BudgetType field value
func (o *SponsoredProductsBudget) GetBudgetType() SponsoredProductsBudgetType {
	if o == nil {
		var ret SponsoredProductsBudgetType
		return ret
	}

	return o.BudgetType
}

// GetBudgetTypeOk returns a tuple with the BudgetType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBudget) GetBudgetTypeOk() (*SponsoredProductsBudgetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BudgetType, true
}

// SetBudgetType sets field value
func (o *SponsoredProductsBudget) SetBudgetType(v SponsoredProductsBudgetType) {
	o.BudgetType = v
}

// GetBudget returns the Budget field value
func (o *SponsoredProductsBudget) GetBudget() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBudget) GetBudgetOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Budget, true
}

// SetBudget sets field value
func (o *SponsoredProductsBudget) SetBudget(v float64) {
	o.Budget = v
}

// GetEffectiveBudget returns the EffectiveBudget field value if set, zero value otherwise.
func (o *SponsoredProductsBudget) GetEffectiveBudget() float64 {
	if o == nil || IsNil(o.EffectiveBudget) {
		var ret float64
		return ret
	}
	return *o.EffectiveBudget
}

// GetEffectiveBudgetOk returns a tuple with the EffectiveBudget field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsBudget) GetEffectiveBudgetOk() (*float64, bool) {
	if o == nil || IsNil(o.EffectiveBudget) {
		return nil, false
	}
	return o.EffectiveBudget, true
}

// HasEffectiveBudget returns a boolean if a field has been set.
func (o *SponsoredProductsBudget) HasEffectiveBudget() bool {
	if o != nil && !IsNil(o.EffectiveBudget) {
		return true
	}

	return false
}

// SetEffectiveBudget gets a reference to the given float64 and assigns it to the EffectiveBudget field.
func (o *SponsoredProductsBudget) SetEffectiveBudget(v float64) {
	o.EffectiveBudget = &v
}

func (o SponsoredProductsBudget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["budgetType"] = o.BudgetType
	toSerialize["budget"] = o.Budget
	if !IsNil(o.EffectiveBudget) {
		toSerialize["effectiveBudget"] = o.EffectiveBudget
	}
	return toSerialize, nil
}

type NullableSponsoredProductsBudget struct {
	value *SponsoredProductsBudget
	isSet bool
}

func (v NullableSponsoredProductsBudget) Get() *SponsoredProductsBudget {
	return v.value
}

func (v *NullableSponsoredProductsBudget) Set(val *SponsoredProductsBudget) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsBudget) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsBudget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsBudget(val *SponsoredProductsBudget) *NullableSponsoredProductsBudget {
	return &NullableSponsoredProductsBudget{value: val, isSet: true}
}

func (v NullableSponsoredProductsBudget) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsBudget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
