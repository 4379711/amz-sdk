package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateOrUpdateDraftCampaignBudget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateOrUpdateDraftCampaignBudget{}

// SponsoredProductsCreateOrUpdateDraftCampaignBudget struct for SponsoredProductsCreateOrUpdateDraftCampaignBudget
type SponsoredProductsCreateOrUpdateDraftCampaignBudget struct {
	BudgetType SponsoredProductsCreateOrUpdateDraftCampaignBudgetType `json:"budgetType"`
	// Monetary value
	Budget float64 `json:"budget"`
}

type _SponsoredProductsCreateOrUpdateDraftCampaignBudget SponsoredProductsCreateOrUpdateDraftCampaignBudget

// NewSponsoredProductsCreateOrUpdateDraftCampaignBudget instantiates a new SponsoredProductsCreateOrUpdateDraftCampaignBudget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateOrUpdateDraftCampaignBudget(budgetType SponsoredProductsCreateOrUpdateDraftCampaignBudgetType, budget float64) *SponsoredProductsCreateOrUpdateDraftCampaignBudget {
	this := SponsoredProductsCreateOrUpdateDraftCampaignBudget{}
	this.BudgetType = budgetType
	this.Budget = budget
	return &this
}

// NewSponsoredProductsCreateOrUpdateDraftCampaignBudgetWithDefaults instantiates a new SponsoredProductsCreateOrUpdateDraftCampaignBudget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateOrUpdateDraftCampaignBudgetWithDefaults() *SponsoredProductsCreateOrUpdateDraftCampaignBudget {
	this := SponsoredProductsCreateOrUpdateDraftCampaignBudget{}
	return &this
}

// GetBudgetType returns the BudgetType field value
func (o *SponsoredProductsCreateOrUpdateDraftCampaignBudget) GetBudgetType() SponsoredProductsCreateOrUpdateDraftCampaignBudgetType {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateDraftCampaignBudgetType
		return ret
	}

	return o.BudgetType
}

// GetBudgetTypeOk returns a tuple with the BudgetType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateOrUpdateDraftCampaignBudget) GetBudgetTypeOk() (*SponsoredProductsCreateOrUpdateDraftCampaignBudgetType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BudgetType, true
}

// SetBudgetType sets field value
func (o *SponsoredProductsCreateOrUpdateDraftCampaignBudget) SetBudgetType(v SponsoredProductsCreateOrUpdateDraftCampaignBudgetType) {
	o.BudgetType = v
}

// GetBudget returns the Budget field value
func (o *SponsoredProductsCreateOrUpdateDraftCampaignBudget) GetBudget() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateOrUpdateDraftCampaignBudget) GetBudgetOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Budget, true
}

// SetBudget sets field value
func (o *SponsoredProductsCreateOrUpdateDraftCampaignBudget) SetBudget(v float64) {
	o.Budget = v
}

func (o SponsoredProductsCreateOrUpdateDraftCampaignBudget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["budgetType"] = o.BudgetType
	toSerialize["budget"] = o.Budget
	return toSerialize, nil
}

type NullableSponsoredProductsCreateOrUpdateDraftCampaignBudget struct {
	value *SponsoredProductsCreateOrUpdateDraftCampaignBudget
	isSet bool
}

func (v NullableSponsoredProductsCreateOrUpdateDraftCampaignBudget) Get() *SponsoredProductsCreateOrUpdateDraftCampaignBudget {
	return v.value
}

func (v *NullableSponsoredProductsCreateOrUpdateDraftCampaignBudget) Set(val *SponsoredProductsCreateOrUpdateDraftCampaignBudget) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateOrUpdateDraftCampaignBudget) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateOrUpdateDraftCampaignBudget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateOrUpdateDraftCampaignBudget(val *SponsoredProductsCreateOrUpdateDraftCampaignBudget) *NullableSponsoredProductsCreateOrUpdateDraftCampaignBudget {
	return &NullableSponsoredProductsCreateOrUpdateDraftCampaignBudget{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateOrUpdateDraftCampaignBudget) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateOrUpdateDraftCampaignBudget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
