package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDBudgetRecommendationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDBudgetRecommendationsResponse{}

// SDBudgetRecommendationsResponse struct for SDBudgetRecommendationsResponse
type SDBudgetRecommendationsResponse struct {
	// List of successful budget recommendation for campaigns.
	BudgetRecommendationsSuccessResults []SDBudgetRecommendation `json:"budgetRecommendationsSuccessResults"`
	// List of errors that occurred when generating budget recommendation.
	BudgetRecommendationsErrorResults []SDBudgetRecommendationError `json:"budgetRecommendationsErrorResults"`
}

type _SDBudgetRecommendationsResponse SDBudgetRecommendationsResponse

// NewSDBudgetRecommendationsResponse instantiates a new SDBudgetRecommendationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDBudgetRecommendationsResponse(budgetRecommendationsSuccessResults []SDBudgetRecommendation, budgetRecommendationsErrorResults []SDBudgetRecommendationError) *SDBudgetRecommendationsResponse {
	this := SDBudgetRecommendationsResponse{}
	this.BudgetRecommendationsSuccessResults = budgetRecommendationsSuccessResults
	this.BudgetRecommendationsErrorResults = budgetRecommendationsErrorResults
	return &this
}

// NewSDBudgetRecommendationsResponseWithDefaults instantiates a new SDBudgetRecommendationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDBudgetRecommendationsResponseWithDefaults() *SDBudgetRecommendationsResponse {
	this := SDBudgetRecommendationsResponse{}
	return &this
}

// GetBudgetRecommendationsSuccessResults returns the BudgetRecommendationsSuccessResults field value
func (o *SDBudgetRecommendationsResponse) GetBudgetRecommendationsSuccessResults() []SDBudgetRecommendation {
	if o == nil {
		var ret []SDBudgetRecommendation
		return ret
	}

	return o.BudgetRecommendationsSuccessResults
}

// GetBudgetRecommendationsSuccessResultsOk returns a tuple with the BudgetRecommendationsSuccessResults field value
// and a boolean to check if the value has been set.
func (o *SDBudgetRecommendationsResponse) GetBudgetRecommendationsSuccessResultsOk() ([]SDBudgetRecommendation, bool) {
	if o == nil {
		return nil, false
	}
	return o.BudgetRecommendationsSuccessResults, true
}

// SetBudgetRecommendationsSuccessResults sets field value
func (o *SDBudgetRecommendationsResponse) SetBudgetRecommendationsSuccessResults(v []SDBudgetRecommendation) {
	o.BudgetRecommendationsSuccessResults = v
}

// GetBudgetRecommendationsErrorResults returns the BudgetRecommendationsErrorResults field value
func (o *SDBudgetRecommendationsResponse) GetBudgetRecommendationsErrorResults() []SDBudgetRecommendationError {
	if o == nil {
		var ret []SDBudgetRecommendationError
		return ret
	}

	return o.BudgetRecommendationsErrorResults
}

// GetBudgetRecommendationsErrorResultsOk returns a tuple with the BudgetRecommendationsErrorResults field value
// and a boolean to check if the value has been set.
func (o *SDBudgetRecommendationsResponse) GetBudgetRecommendationsErrorResultsOk() ([]SDBudgetRecommendationError, bool) {
	if o == nil {
		return nil, false
	}
	return o.BudgetRecommendationsErrorResults, true
}

// SetBudgetRecommendationsErrorResults sets field value
func (o *SDBudgetRecommendationsResponse) SetBudgetRecommendationsErrorResults(v []SDBudgetRecommendationError) {
	o.BudgetRecommendationsErrorResults = v
}

func (o SDBudgetRecommendationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["budgetRecommendationsSuccessResults"] = o.BudgetRecommendationsSuccessResults
	toSerialize["budgetRecommendationsErrorResults"] = o.BudgetRecommendationsErrorResults
	return toSerialize, nil
}

type NullableSDBudgetRecommendationsResponse struct {
	value *SDBudgetRecommendationsResponse
	isSet bool
}

func (v NullableSDBudgetRecommendationsResponse) Get() *SDBudgetRecommendationsResponse {
	return v.value
}

func (v *NullableSDBudgetRecommendationsResponse) Set(val *SDBudgetRecommendationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSDBudgetRecommendationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSDBudgetRecommendationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDBudgetRecommendationsResponse(val *SDBudgetRecommendationsResponse) *NullableSDBudgetRecommendationsResponse {
	return &NullableSDBudgetRecommendationsResponse{value: val, isSet: true}
}

func (v NullableSDBudgetRecommendationsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDBudgetRecommendationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
