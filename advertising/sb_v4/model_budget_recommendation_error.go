package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the BudgetRecommendationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BudgetRecommendationError{}

// BudgetRecommendationError Error that occurred when generating budget recommendations.
type BudgetRecommendationError struct {
	// A human-readable description of the enumerated response code in the `code` field.
	Code string `json:"code"`
	// The identifier of a campaign.
	CampaignId string `json:"campaignId"`
	// Correlate the recommendation to the campaign index in the request. Zero-based.
	Index float32 `json:"index"`
	// An enumerated response code.
	Details string `json:"details"`
}

type _BudgetRecommendationError BudgetRecommendationError

// NewBudgetRecommendationError instantiates a new BudgetRecommendationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBudgetRecommendationError(code string, campaignId string, index float32, details string) *BudgetRecommendationError {
	this := BudgetRecommendationError{}
	this.Code = code
	this.CampaignId = campaignId
	this.Index = index
	this.Details = details
	return &this
}

// NewBudgetRecommendationErrorWithDefaults instantiates a new BudgetRecommendationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBudgetRecommendationErrorWithDefaults() *BudgetRecommendationError {
	this := BudgetRecommendationError{}
	return &this
}

// GetCode returns the Code field value
func (o *BudgetRecommendationError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *BudgetRecommendationError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *BudgetRecommendationError) SetCode(v string) {
	o.Code = v
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
func (o *BudgetRecommendationError) GetIndex() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *BudgetRecommendationError) GetIndexOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *BudgetRecommendationError) SetIndex(v float32) {
	o.Index = v
}

// GetDetails returns the Details field value
func (o *BudgetRecommendationError) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *BudgetRecommendationError) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *BudgetRecommendationError) SetDetails(v string) {
	o.Details = v
}

func (o BudgetRecommendationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["index"] = o.Index
	toSerialize["details"] = o.Details
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
