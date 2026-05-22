package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDBudgetRecommendationError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDBudgetRecommendationError{}

// SDBudgetRecommendationError struct for SDBudgetRecommendationError
type SDBudgetRecommendationError struct {
	// Correlate the recommendation to the campaign index in the request. Zero-based.
	Index int32 `json:"index"`
	// Campaign id.
	CampaignId string `json:"campaignId"`
	// The HTTP status code of the response.
	Code string `json:"code"`
	// A human-readable description of the response.
	Details string `json:"details"`
}

type _SDBudgetRecommendationError SDBudgetRecommendationError

// NewSDBudgetRecommendationError instantiates a new SDBudgetRecommendationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDBudgetRecommendationError(index int32, campaignId string, code string, details string) *SDBudgetRecommendationError {
	this := SDBudgetRecommendationError{}
	this.Index = index
	this.CampaignId = campaignId
	this.Code = code
	this.Details = details
	return &this
}

// NewSDBudgetRecommendationErrorWithDefaults instantiates a new SDBudgetRecommendationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDBudgetRecommendationErrorWithDefaults() *SDBudgetRecommendationError {
	this := SDBudgetRecommendationError{}
	return &this
}

// GetIndex returns the Index field value
func (o *SDBudgetRecommendationError) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SDBudgetRecommendationError) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SDBudgetRecommendationError) SetIndex(v int32) {
	o.Index = v
}

// GetCampaignId returns the CampaignId field value
func (o *SDBudgetRecommendationError) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SDBudgetRecommendationError) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SDBudgetRecommendationError) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetCode returns the Code field value
func (o *SDBudgetRecommendationError) GetCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *SDBudgetRecommendationError) GetCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *SDBudgetRecommendationError) SetCode(v string) {
	o.Code = v
}

// GetDetails returns the Details field value
func (o *SDBudgetRecommendationError) GetDetails() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *SDBudgetRecommendationError) GetDetailsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *SDBudgetRecommendationError) SetDetails(v string) {
	o.Details = v
}

func (o SDBudgetRecommendationError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["code"] = o.Code
	toSerialize["details"] = o.Details
	return toSerialize, nil
}

type NullableSDBudgetRecommendationError struct {
	value *SDBudgetRecommendationError
	isSet bool
}

func (v NullableSDBudgetRecommendationError) Get() *SDBudgetRecommendationError {
	return v.value
}

func (v *NullableSDBudgetRecommendationError) Set(val *SDBudgetRecommendationError) {
	v.value = val
	v.isSet = true
}

func (v NullableSDBudgetRecommendationError) IsSet() bool {
	return v.isSet
}

func (v *NullableSDBudgetRecommendationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDBudgetRecommendationError(val *SDBudgetRecommendationError) *NullableSDBudgetRecommendationError {
	return &NullableSDBudgetRecommendationError{value: val, isSet: true}
}

func (v NullableSDBudgetRecommendationError) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDBudgetRecommendationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
