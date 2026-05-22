package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftCampaignBudget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftCampaignBudget{}

// SponsoredProductsDraftCampaignBudget struct for SponsoredProductsDraftCampaignBudget
type SponsoredProductsDraftCampaignBudget struct {
	BudgetType SponsoredProductsDraftCampaignBudgetType `json:"budgetType"`
	// Monetary value
	Budget float64 `json:"budget"`
}

type _SponsoredProductsDraftCampaignBudget SponsoredProductsDraftCampaignBudget

// NewSponsoredProductsDraftCampaignBudget instantiates a new SponsoredProductsDraftCampaignBudget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftCampaignBudget(budgetType SponsoredProductsDraftCampaignBudgetType, budget float64) *SponsoredProductsDraftCampaignBudget {
	this := SponsoredProductsDraftCampaignBudget{}
	this.BudgetType = budgetType
	this.Budget = budget
	return &this
}

// NewSponsoredProductsDraftCampaignBudgetWithDefaults instantiates a new SponsoredProductsDraftCampaignBudget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftCampaignBudgetWithDefaults() *SponsoredProductsDraftCampaignBudget {
	this := SponsoredProductsDraftCampaignBudget{}
	return &this
}

// GetBudgetType returns the BudgetType field value
func (o *SponsoredProductsDraftCampaignBudget) GetBudgetType() SponsoredProductsDraftCampaignBudgetType {
	if o == nil {
		var ret SponsoredProductsDraftCampaignBudgetType
		return ret
	}

	return o.BudgetType
}

// GetBudgetTypeOk returns a tuple with the BudgetType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignBudget) GetBudgetTypeOk() (*SponsoredProductsDraftCampaignBudgetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BudgetType, true
}

// SetBudgetType sets field value
func (o *SponsoredProductsDraftCampaignBudget) SetBudgetType(v SponsoredProductsDraftCampaignBudgetType) {
	o.BudgetType = v
}

// GetBudget returns the Budget field value
func (o *SponsoredProductsDraftCampaignBudget) GetBudget() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftCampaignBudget) GetBudgetOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Budget, true
}

// SetBudget sets field value
func (o *SponsoredProductsDraftCampaignBudget) SetBudget(v float64) {
	o.Budget = v
}

func (o SponsoredProductsDraftCampaignBudget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["budgetType"] = o.BudgetType
	toSerialize["budget"] = o.Budget
	return toSerialize, nil
}

type NullableSponsoredProductsDraftCampaignBudget struct {
	value *SponsoredProductsDraftCampaignBudget
	isSet bool
}

func (v NullableSponsoredProductsDraftCampaignBudget) Get() *SponsoredProductsDraftCampaignBudget {
	return v.value
}

func (v *NullableSponsoredProductsDraftCampaignBudget) Set(val *SponsoredProductsDraftCampaignBudget) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftCampaignBudget) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftCampaignBudget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftCampaignBudget(val *SponsoredProductsDraftCampaignBudget) *NullableSponsoredProductsDraftCampaignBudget {
	return &NullableSponsoredProductsDraftCampaignBudget{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftCampaignBudget) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftCampaignBudget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
