package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalBudgetRecommendationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalBudgetRecommendationResponse{}

// GlobalBudgetRecommendationResponse struct for GlobalBudgetRecommendationResponse
type GlobalBudgetRecommendationResponse struct {
	// List of global budget recommendation results for campaigns.
	BudgetRecommendationsResults []GlobalBudgetRecommendationResult `json:"budgetRecommendationsResults"`
}

type _GlobalBudgetRecommendationResponse GlobalBudgetRecommendationResponse

// NewGlobalBudgetRecommendationResponse instantiates a new GlobalBudgetRecommendationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalBudgetRecommendationResponse(budgetRecommendationsResults []GlobalBudgetRecommendationResult) *GlobalBudgetRecommendationResponse {
	this := GlobalBudgetRecommendationResponse{}
	this.BudgetRecommendationsResults = budgetRecommendationsResults
	return &this
}

// NewGlobalBudgetRecommendationResponseWithDefaults instantiates a new GlobalBudgetRecommendationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalBudgetRecommendationResponseWithDefaults() *GlobalBudgetRecommendationResponse {
	this := GlobalBudgetRecommendationResponse{}
	return &this
}

// GetBudgetRecommendationsResults returns the BudgetRecommendationsResults field value
func (o *GlobalBudgetRecommendationResponse) GetBudgetRecommendationsResults() []GlobalBudgetRecommendationResult {
	if o == nil {
		var ret []GlobalBudgetRecommendationResult
		return ret
	}

	return o.BudgetRecommendationsResults
}

// GetBudgetRecommendationsResultsOk returns a tuple with the BudgetRecommendationsResults field value
// and a boolean to check if the value has been set.
func (o *GlobalBudgetRecommendationResponse) GetBudgetRecommendationsResultsOk() ([]GlobalBudgetRecommendationResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.BudgetRecommendationsResults, true
}

// SetBudgetRecommendationsResults sets field value
func (o *GlobalBudgetRecommendationResponse) SetBudgetRecommendationsResults(v []GlobalBudgetRecommendationResult) {
	o.BudgetRecommendationsResults = v
}

func (o GlobalBudgetRecommendationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["budgetRecommendationsResults"] = o.BudgetRecommendationsResults
	return toSerialize, nil
}

type NullableGlobalBudgetRecommendationResponse struct {
	value *GlobalBudgetRecommendationResponse
	isSet bool
}

func (v NullableGlobalBudgetRecommendationResponse) Get() *GlobalBudgetRecommendationResponse {
	return v.value
}

func (v *NullableGlobalBudgetRecommendationResponse) Set(val *GlobalBudgetRecommendationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalBudgetRecommendationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalBudgetRecommendationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalBudgetRecommendationResponse(val *GlobalBudgetRecommendationResponse) *NullableGlobalBudgetRecommendationResponse {
	return &NullableGlobalBudgetRecommendationResponse{value: val, isSet: true}
}

func (v NullableGlobalBudgetRecommendationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalBudgetRecommendationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
