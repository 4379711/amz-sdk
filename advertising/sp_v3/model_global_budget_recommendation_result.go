package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalBudgetRecommendationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalBudgetRecommendationResult{}

// GlobalBudgetRecommendationResult struct for GlobalBudgetRecommendationResult
type GlobalBudgetRecommendationResult struct {
	ResultStatus ResultStatus `json:"resultStatus"`
	// The budget recommendation for each country.
	CountryBudgetRecommendations map[string]CountryBudgetRecommendation `json:"countryBudgetRecommendations"`
	// The campaign Id.
	CampaignId string `json:"campaignId"`
	// Correlate the recommendations to the campaign index in the request. Zero-based.
	Index  int32                             `json:"index"`
	Errors []GlobalBudgetRecommendationError `json:"errors,omitempty"`
}

type _GlobalBudgetRecommendationResult GlobalBudgetRecommendationResult

// NewGlobalBudgetRecommendationResult instantiates a new GlobalBudgetRecommendationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalBudgetRecommendationResult(resultStatus ResultStatus, countryBudgetRecommendations map[string]CountryBudgetRecommendation, campaignId string, index int32) *GlobalBudgetRecommendationResult {
	this := GlobalBudgetRecommendationResult{}
	this.ResultStatus = resultStatus
	this.CountryBudgetRecommendations = countryBudgetRecommendations
	this.CampaignId = campaignId
	this.Index = index
	return &this
}

// NewGlobalBudgetRecommendationResultWithDefaults instantiates a new GlobalBudgetRecommendationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalBudgetRecommendationResultWithDefaults() *GlobalBudgetRecommendationResult {
	this := GlobalBudgetRecommendationResult{}
	var resultStatus ResultStatus = RESULTSTATUS_SUCCESS
	this.ResultStatus = resultStatus
	return &this
}

// GetResultStatus returns the ResultStatus field value
func (o *GlobalBudgetRecommendationResult) GetResultStatus() ResultStatus {
	if o == nil {
		var ret ResultStatus
		return ret
	}

	return o.ResultStatus
}

// GetResultStatusOk returns a tuple with the ResultStatus field value
// and a boolean to check if the value has been set.
func (o *GlobalBudgetRecommendationResult) GetResultStatusOk() (*ResultStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultStatus, true
}

// SetResultStatus sets field value
func (o *GlobalBudgetRecommendationResult) SetResultStatus(v ResultStatus) {
	o.ResultStatus = v
}

// GetCountryBudgetRecommendations returns the CountryBudgetRecommendations field value
func (o *GlobalBudgetRecommendationResult) GetCountryBudgetRecommendations() map[string]CountryBudgetRecommendation {
	if o == nil {
		var ret map[string]CountryBudgetRecommendation
		return ret
	}

	return o.CountryBudgetRecommendations
}

// GetCountryBudgetRecommendationsOk returns a tuple with the CountryBudgetRecommendations field value
// and a boolean to check if the value has been set.
func (o *GlobalBudgetRecommendationResult) GetCountryBudgetRecommendationsOk() (*map[string]CountryBudgetRecommendation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryBudgetRecommendations, true
}

// SetCountryBudgetRecommendations sets field value
func (o *GlobalBudgetRecommendationResult) SetCountryBudgetRecommendations(v map[string]CountryBudgetRecommendation) {
	o.CountryBudgetRecommendations = v
}

// GetCampaignId returns the CampaignId field value
func (o *GlobalBudgetRecommendationResult) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *GlobalBudgetRecommendationResult) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *GlobalBudgetRecommendationResult) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetIndex returns the Index field value
func (o *GlobalBudgetRecommendationResult) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *GlobalBudgetRecommendationResult) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *GlobalBudgetRecommendationResult) SetIndex(v int32) {
	o.Index = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GlobalBudgetRecommendationResult) GetErrors() []GlobalBudgetRecommendationError {
	if o == nil || IsNil(o.Errors) {
		var ret []GlobalBudgetRecommendationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalBudgetRecommendationResult) GetErrorsOk() ([]GlobalBudgetRecommendationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GlobalBudgetRecommendationResult) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []GlobalBudgetRecommendationError and assigns it to the Errors field.
func (o *GlobalBudgetRecommendationResult) SetErrors(v []GlobalBudgetRecommendationError) {
	o.Errors = v
}

func (o GlobalBudgetRecommendationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resultStatus"] = o.ResultStatus
	toSerialize["countryBudgetRecommendations"] = o.CountryBudgetRecommendations
	toSerialize["campaignId"] = o.CampaignId
	toSerialize["index"] = o.Index
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGlobalBudgetRecommendationResult struct {
	value *GlobalBudgetRecommendationResult
	isSet bool
}

func (v NullableGlobalBudgetRecommendationResult) Get() *GlobalBudgetRecommendationResult {
	return v.value
}

func (v *NullableGlobalBudgetRecommendationResult) Set(val *GlobalBudgetRecommendationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalBudgetRecommendationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalBudgetRecommendationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalBudgetRecommendationResult(val *GlobalBudgetRecommendationResult) *NullableGlobalBudgetRecommendationResult {
	return &NullableGlobalBudgetRecommendationResult{value: val, isSet: true}
}

func (v NullableGlobalBudgetRecommendationResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalBudgetRecommendationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
