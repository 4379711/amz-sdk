package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the BudgetRecommendationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetRecommendationError{}

// BudgetRecommendationError struct for BudgetRecommendationError
type BudgetRecommendationError struct {
	// encrypted campaignId
	CampaignId string `json:"campaignId"`
	// Correlate the recommendation to the campaign index in the request. Zero-based
	Index int32                          `json:"index"`
	Error SPTORBudgetRecommendationError `json:"Error"`
}

type _BudgetRecommendationError BudgetRecommendationError

// NewBudgetRecommendationError instantiates a new BudgetRecommendationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetRecommendationError(campaignId string, index int32, error_ SPTORBudgetRecommendationError) *BudgetRecommendationError {
	this := BudgetRecommendationError{}
	this.CampaignId = campaignId
	this.Index = index
	this.Error = error_
	return &this
}

// NewBudgetRecommendationErrorWithDefaults instantiates a new BudgetRecommendationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetRecommendationErrorWithDefaults() *BudgetRecommendationError {
	this := BudgetRecommendationError{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *BudgetRecommendationError) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *BudgetRecommendationError) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *BudgetRecommendationError) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetIndex returns the Index field value
func (o *BudgetRecommendationError) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *BudgetRecommendationError) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *BudgetRecommendationError) SetIndex(v int32) {
	o.Index = v
}

// GetError returns the Error field value
func (o *BudgetRecommendationError) GetError() SPTORBudgetRecommendationError {
	if o == nil {
		var ret SPTORBudgetRecommendationError
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *BudgetRecommendationError) GetErrorOk() (*SPTORBudgetRecommendationError, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *BudgetRecommendationError) SetError(v SPTORBudgetRecommendationError) {
	o.Error = v
}

func (o BudgetRecommendationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["index"] = o.Index
	toSerialize["Error"] = o.Error
	return toSerialize, nil
}

type NullableBudgetRecommendationError struct {
	value *BudgetRecommendationError
	isSet bool
}

func (v NullableBudgetRecommendationError) Get() *BudgetRecommendationError {
	return v.value
}

func (v *NullableBudgetRecommendationError) Set(val *BudgetRecommendationError) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetRecommendationError) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetRecommendationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetRecommendationError(val *BudgetRecommendationError) *NullableBudgetRecommendationError {
	return &NullableBudgetRecommendationError{value: val, isSet: true}
}

func (v NullableBudgetRecommendationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBudgetRecommendationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
