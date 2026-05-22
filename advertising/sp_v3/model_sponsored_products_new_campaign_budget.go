package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsNewCampaignBudget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNewCampaignBudget{}

// SponsoredProductsNewCampaignBudget The budget for the campaigns in the target promotion group.
type SponsoredProductsNewCampaignBudget struct {
	// DAILY.
	BudgetType string `json:"budgetType"`
	// The value of the budget.
	Budget float64 `json:"budget"`
}

type _SponsoredProductsNewCampaignBudget SponsoredProductsNewCampaignBudget

// NewSponsoredProductsNewCampaignBudget instantiates a new SponsoredProductsNewCampaignBudget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNewCampaignBudget(budgetType string, budget float64) *SponsoredProductsNewCampaignBudget {
	this := SponsoredProductsNewCampaignBudget{}
	this.BudgetType = budgetType
	this.Budget = budget
	return &this
}

// NewSponsoredProductsNewCampaignBudgetWithDefaults instantiates a new SponsoredProductsNewCampaignBudget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNewCampaignBudgetWithDefaults() *SponsoredProductsNewCampaignBudget {
	this := SponsoredProductsNewCampaignBudget{}
	return &this
}

// GetBudgetType returns the BudgetType field value
func (o *SponsoredProductsNewCampaignBudget) GetBudgetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BudgetType
}

// GetBudgetTypeOk returns a tuple with the BudgetType field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNewCampaignBudget) GetBudgetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BudgetType, true
}

// SetBudgetType sets field value
func (o *SponsoredProductsNewCampaignBudget) SetBudgetType(v string) {
	o.BudgetType = v
}

// GetBudget returns the Budget field value
func (o *SponsoredProductsNewCampaignBudget) GetBudget() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Budget
}

// GetBudgetOk returns a tuple with the Budget field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNewCampaignBudget) GetBudgetOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Budget, true
}

// SetBudget sets field value
func (o *SponsoredProductsNewCampaignBudget) SetBudget(v float64) {
	o.Budget = v
}

func (o SponsoredProductsNewCampaignBudget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["budgetType"] = o.BudgetType
	toSerialize["budget"] = o.Budget
	return toSerialize, nil
}

type NullableSponsoredProductsNewCampaignBudget struct {
	value *SponsoredProductsNewCampaignBudget
	isSet bool
}

func (v NullableSponsoredProductsNewCampaignBudget) Get() *SponsoredProductsNewCampaignBudget {
	return v.value
}

func (v *NullableSponsoredProductsNewCampaignBudget) Set(val *SponsoredProductsNewCampaignBudget) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNewCampaignBudget) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNewCampaignBudget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNewCampaignBudget(val *SponsoredProductsNewCampaignBudget) *NullableSponsoredProductsNewCampaignBudget {
	return &NullableSponsoredProductsNewCampaignBudget{value: val, isSet: true}
}

func (v NullableSponsoredProductsNewCampaignBudget) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNewCampaignBudget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
