package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateOrUpdateBudget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateOrUpdateBudget{}

// SponsoredProductsCreateOrUpdateBudget struct for SponsoredProductsCreateOrUpdateBudget
type SponsoredProductsCreateOrUpdateBudget struct {
	BudgetType SponsoredProductsCreateOrUpdateBudgetType `json:"budgetType"`
	// Monetary value
	Budget float64 `json:"budget"`
}

type _SponsoredProductsCreateOrUpdateBudget SponsoredProductsCreateOrUpdateBudget

// NewSponsoredProductsCreateOrUpdateBudget instantiates a new SponsoredProductsCreateOrUpdateBudget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateOrUpdateBudget(budgetType SponsoredProductsCreateOrUpdateBudgetType, budget float64) *SponsoredProductsCreateOrUpdateBudget {
	this := SponsoredProductsCreateOrUpdateBudget{}
	this.BudgetType = budgetType
	this.Budget = budget
	return &this
}

// NewSponsoredProductsCreateOrUpdateBudgetWithDefaults instantiates a new SponsoredProductsCreateOrUpdateBudget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateOrUpdateBudgetWithDefaults() *SponsoredProductsCreateOrUpdateBudget {
	this := SponsoredProductsCreateOrUpdateBudget{}
	return &this
}

// GetBudgetType returns the BudgetType field value
func (o *SponsoredProductsCreateOrUpdateBudget) GetBudgetType() SponsoredProductsCreateOrUpdateBudgetType {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateBudgetType
		return ret
	}

	return o.BudgetType
}

// GetBudgetTypeOk returns a tuple with the BudgetType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateOrUpdateBudget) GetBudgetTypeOk() (*SponsoredProductsCreateOrUpdateBudgetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BudgetType, true
}

// SetBudgetType sets field value
func (o *SponsoredProductsCreateOrUpdateBudget) SetBudgetType(v SponsoredProductsCreateOrUpdateBudgetType) {
	o.BudgetType = v
}

// GetBudget returns the Budget field value
func (o *SponsoredProductsCreateOrUpdateBudget) GetBudget() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateOrUpdateBudget) GetBudgetOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Budget, true
}

// SetBudget sets field value
func (o *SponsoredProductsCreateOrUpdateBudget) SetBudget(v float64) {
	o.Budget = v
}

func (o SponsoredProductsCreateOrUpdateBudget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["budgetType"] = o.BudgetType
	toSerialize["budget"] = o.Budget
	return toSerialize, nil
}

type NullableSponsoredProductsCreateOrUpdateBudget struct {
	value *SponsoredProductsCreateOrUpdateBudget
	isSet bool
}

func (v NullableSponsoredProductsCreateOrUpdateBudget) Get() *SponsoredProductsCreateOrUpdateBudget {
	return v.value
}

func (v *NullableSponsoredProductsCreateOrUpdateBudget) Set(val *SponsoredProductsCreateOrUpdateBudget) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateOrUpdateBudget) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateOrUpdateBudget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateOrUpdateBudget(val *SponsoredProductsCreateOrUpdateBudget) *NullableSponsoredProductsCreateOrUpdateBudget {
	return &NullableSponsoredProductsCreateOrUpdateBudget{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateOrUpdateBudget) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateOrUpdateBudget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
