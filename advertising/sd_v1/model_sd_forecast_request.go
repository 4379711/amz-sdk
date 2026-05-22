package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the SDForecastRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDForecastRequest{}

// SDForecastRequest Request payload for SD forecasting. Below are required and optional fields. Fields not listed will not impact forecast results. |Field              |Object            |Required| |-------------------|------------------|--------| |startDate          |Campaign          |required| |endDate            |Campaign          |optional| |costType           |Campaign          |optional| |bidOptimization    |AdGroup           |required| |creativeType       |AdGroup           |optional| |defaultBid         |AdGroup           |optional| |asin               |ProductAds        |required for vendors| |sku                |ProductAds        |required for sellers| |bid                |TargetingClauses  |required when defaultBid is not set| |expression         |TargetingClauses  |required| |ruleConditions     |OptimizationRules |optional|
type SDForecastRequest struct {
	Campaign Campaign `json:"campaign"`
	AdGroup  AdGroup  `json:"adGroup"`
	// A list of SD optimization rules. Forecast will be affected by the optimization strategy rules.  Currently, supported rule metrics by forecast are `COST_PER_CLICK`, `COST_PER_THOUSAND_VIEWABLE_IMPRESSIONS` and `COST_PER_ORDER`.
	OptimizationRules []OptimizationRule `json:"optimizationRules,omitempty"`
	ProductAds        []ProductAd        `json:"productAds"`
	// A list of SD targeting clauses.
	TargetingClauses []SDForecastRequestTargetingClause `json:"targetingClauses"`
	// A list of SD negative targeting clauses.
	NegativeTargetingClauses []NegativeTargetingClause `json:"negativeTargetingClauses,omitempty"`
}

type _SDForecastRequest SDForecastRequest

// NewSDForecastRequest instantiates a new SDForecastRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDForecastRequest(campaign Campaign, adGroup AdGroup, productAds []ProductAd, targetingClauses []SDForecastRequestTargetingClause) *SDForecastRequest {
	this := SDForecastRequest{}
	this.Campaign = campaign
	this.AdGroup = adGroup
	this.ProductAds = productAds
	this.TargetingClauses = targetingClauses
	return &this
}

// NewSDForecastRequestWithDefaults instantiates a new SDForecastRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDForecastRequestWithDefaults() *SDForecastRequest {
	this := SDForecastRequest{}
	return &this
}

// GetCampaign returns the Campaign field value
func (o *SDForecastRequest) GetCampaign() Campaign {
	if o == nil {
		var ret Campaign
		return ret
	}

	return o.Campaign
}

// GetCampaignOk returns a tuple with the Campaign field value
// and a boolean to check if the value has been set.
func (o *SDForecastRequest) GetCampaignOk() (*Campaign, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Campaign, true
}

// SetCampaign sets field value
func (o *SDForecastRequest) SetCampaign(v Campaign) {
	o.Campaign = v
}

// GetAdGroup returns the AdGroup field value
func (o *SDForecastRequest) GetAdGroup() AdGroup {
	if o == nil {
		var ret AdGroup
		return ret
	}

	return o.AdGroup
}

// GetAdGroupOk returns a tuple with the AdGroup field value
// and a boolean to check if the value has been set.
func (o *SDForecastRequest) GetAdGroupOk() (*AdGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroup, true
}

// SetAdGroup sets field value
func (o *SDForecastRequest) SetAdGroup(v AdGroup) {
	o.AdGroup = v
}

// GetOptimizationRules returns the OptimizationRules field value if set, zero value otherwise.
func (o *SDForecastRequest) GetOptimizationRules() []OptimizationRule {
	if o == nil || IsNil(o.OptimizationRules) {
		var ret []OptimizationRule
		return ret
	}
	return o.OptimizationRules
}

// GetOptimizationRulesOk returns a tuple with the OptimizationRules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDForecastRequest) GetOptimizationRulesOk() ([]OptimizationRule, bool) {
	if o == nil || IsNil(o.OptimizationRules) {
		return nil, false
	}
	return o.OptimizationRules, true
}

// HasOptimizationRules returns a boolean if a field has been set.
func (o *SDForecastRequest) HasOptimizationRules() bool {
	if o != nil && !IsNil(o.OptimizationRules) {
		return true
	}

	return false
}

// SetOptimizationRules gets a reference to the given []OptimizationRule and assigns it to the OptimizationRules field.
func (o *SDForecastRequest) SetOptimizationRules(v []OptimizationRule) {
	o.OptimizationRules = v
}

// GetProductAds returns the ProductAds field value
func (o *SDForecastRequest) GetProductAds() []ProductAd {
	if o == nil {
		var ret []ProductAd
		return ret
	}

	return o.ProductAds
}

// GetProductAdsOk returns a tuple with the ProductAds field value
// and a boolean to check if the value has been set.
func (o *SDForecastRequest) GetProductAdsOk() ([]ProductAd, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductAds, true
}

// SetProductAds sets field value
func (o *SDForecastRequest) SetProductAds(v []ProductAd) {
	o.ProductAds = v
}

// GetTargetingClauses returns the TargetingClauses field value
func (o *SDForecastRequest) GetTargetingClauses() []SDForecastRequestTargetingClause {
	if o == nil {
		var ret []SDForecastRequestTargetingClause
		return ret
	}

	return o.TargetingClauses
}

// GetTargetingClausesOk returns a tuple with the TargetingClauses field value
// and a boolean to check if the value has been set.
func (o *SDForecastRequest) GetTargetingClausesOk() ([]SDForecastRequestTargetingClause, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetingClauses, true
}

// SetTargetingClauses sets field value
func (o *SDForecastRequest) SetTargetingClauses(v []SDForecastRequestTargetingClause) {
	o.TargetingClauses = v
}

// GetNegativeTargetingClauses returns the NegativeTargetingClauses field value if set, zero value otherwise.
func (o *SDForecastRequest) GetNegativeTargetingClauses() []NegativeTargetingClause {
	if o == nil || IsNil(o.NegativeTargetingClauses) {
		var ret []NegativeTargetingClause
		return ret
	}
	return o.NegativeTargetingClauses
}

// GetNegativeTargetingClausesOk returns a tuple with the NegativeTargetingClauses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDForecastRequest) GetNegativeTargetingClausesOk() ([]NegativeTargetingClause, bool) {
	if o == nil || IsNil(o.NegativeTargetingClauses) {
		return nil, false
	}
	return o.NegativeTargetingClauses, true
}

// HasNegativeTargetingClauses returns a boolean if a field has been set.
func (o *SDForecastRequest) HasNegativeTargetingClauses() bool {
	if o != nil && !IsNil(o.NegativeTargetingClauses) {
		return true
	}

	return false
}

// SetNegativeTargetingClauses gets a reference to the given []NegativeTargetingClause and assigns it to the NegativeTargetingClauses field.
func (o *SDForecastRequest) SetNegativeTargetingClauses(v []NegativeTargetingClause) {
	o.NegativeTargetingClauses = v
}

func (o SDForecastRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["campaign"] = o.Campaign
	toSerialize["adGroup"] = o.AdGroup
	if !IsNil(o.OptimizationRules) {
		toSerialize["optimizationRules"] = o.OptimizationRules
	}
	toSerialize["productAds"] = o.ProductAds
	toSerialize["targetingClauses"] = o.TargetingClauses
	if !IsNil(o.NegativeTargetingClauses) {
		toSerialize["negativeTargetingClauses"] = o.NegativeTargetingClauses
	}
	return toSerialize, nil
}

type NullableSDForecastRequest struct {
	value *SDForecastRequest
	isSet bool
}

func (v NullableSDForecastRequest) Get() *SDForecastRequest {
	return v.value
}

func (v *NullableSDForecastRequest) Set(val *SDForecastRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSDForecastRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSDForecastRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDForecastRequest(val *SDForecastRequest) *NullableSDForecastRequest {
	return &NullableSDForecastRequest{value: val, isSet: true}
}

func (v NullableSDForecastRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSDForecastRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
