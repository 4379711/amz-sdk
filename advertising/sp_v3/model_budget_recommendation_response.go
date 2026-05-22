package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BudgetRecommendationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetRecommendationResponse{}

// BudgetRecommendationResponse struct for BudgetRecommendationResponse
type BudgetRecommendationResponse struct {
	// List of successful budget recomendation for campagins.
	BudgetRecommendationsSuccessResults []BudgetRecommendationForExistingCampaign `json:"budgetRecommendationsSuccessResults"`
	// List of errors that occured when generating bduget recommendation.
	BudgetRecommendationsErrorResults []BudgetRecommendationError `json:"budgetRecommendationsErrorResults"`
}

type _BudgetRecommendationResponse BudgetRecommendationResponse

// NewBudgetRecommendationResponse instantiates a new BudgetRecommendationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetRecommendationResponse(budgetRecommendationsSuccessResults []BudgetRecommendationForExistingCampaign, budgetRecommendationsErrorResults []BudgetRecommendationError) *BudgetRecommendationResponse {
	this := BudgetRecommendationResponse{}
	this.BudgetRecommendationsSuccessResults = budgetRecommendationsSuccessResults
	this.BudgetRecommendationsErrorResults = budgetRecommendationsErrorResults
	return &this
}

// NewBudgetRecommendationResponseWithDefaults instantiates a new BudgetRecommendationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetRecommendationResponseWithDefaults() *BudgetRecommendationResponse {
	this := BudgetRecommendationResponse{}
	return &this
}

// GetBudgetRecommendationsSuccessResults returns the BudgetRecommendationsSuccessResults field value
func (o *BudgetRecommendationResponse) GetBudgetRecommendationsSuccessResults() []BudgetRecommendationForExistingCampaign {
	if o == nil {
		var ret []BudgetRecommendationForExistingCampaign
		return ret
	}

	return o.BudgetRecommendationsSuccessResults
}

// GetBudgetRecommendationsSuccessResultsOk returns a tuple with the BudgetRecommendationsSuccessResults field value
// and a boolean to check if the value has been set.
func (o *BudgetRecommendationResponse) GetBudgetRecommendationsSuccessResultsOk() ([]BudgetRecommendationForExistingCampaign, bool) {
	if o == nil {
		return nil, false
	}
	return o.BudgetRecommendationsSuccessResults, true
}

// SetBudgetRecommendationsSuccessResults sets field value
func (o *BudgetRecommendationResponse) SetBudgetRecommendationsSuccessResults(v []BudgetRecommendationForExistingCampaign) {
	o.BudgetRecommendationsSuccessResults = v
}

// GetBudgetRecommendationsErrorResults returns the BudgetRecommendationsErrorResults field value
func (o *BudgetRecommendationResponse) GetBudgetRecommendationsErrorResults() []BudgetRecommendationError {
	if o == nil {
		var ret []BudgetRecommendationError
		return ret
	}

	return o.BudgetRecommendationsErrorResults
}

// GetBudgetRecommendationsErrorResultsOk returns a tuple with the BudgetRecommendationsErrorResults field value
// and a boolean to check if the value has been set.
func (o *BudgetRecommendationResponse) GetBudgetRecommendationsErrorResultsOk() ([]BudgetRecommendationError, bool) {
	if o == nil {
		return nil, false
	}
	return o.BudgetRecommendationsErrorResults, true
}

// SetBudgetRecommendationsErrorResults sets field value
func (o *BudgetRecommendationResponse) SetBudgetRecommendationsErrorResults(v []BudgetRecommendationError) {
	o.BudgetRecommendationsErrorResults = v
}

func (o BudgetRecommendationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["budgetRecommendationsSuccessResults"] = o.BudgetRecommendationsSuccessResults
	toSerialize["budgetRecommendationsErrorResults"] = o.BudgetRecommendationsErrorResults
	return toSerialize, nil
}

type NullableBudgetRecommendationResponse struct {
	value *BudgetRecommendationResponse
	isSet bool
}

func (v NullableBudgetRecommendationResponse) Get() *BudgetRecommendationResponse {
	return v.value
}

func (v *NullableBudgetRecommendationResponse) Set(val *BudgetRecommendationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetRecommendationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetRecommendationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetRecommendationResponse(val *BudgetRecommendationResponse) *NullableBudgetRecommendationResponse {
	return &NullableBudgetRecommendationResponse{value: val, isSet: true}
}

func (v NullableBudgetRecommendationResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetRecommendationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
