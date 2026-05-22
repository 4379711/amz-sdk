package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the MultiCountryAsinsThemeBasedBidRecommendationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiCountryAsinsThemeBasedBidRecommendationRequest{}

// MultiCountryAsinsThemeBasedBidRecommendationRequest struct for MultiCountryAsinsThemeBasedBidRecommendationRequest
type MultiCountryAsinsThemeBasedBidRecommendationRequest struct {
	// The list of multi country asins
	Asins []map[string]string `json:"asins"`
	// The list of targeting expressions. Maximum of 100 per request per country, use pagination for more if needed.
	TargetingExpressions []MultiCountryTargetingExpression `json:"targetingExpressions"`
	// A list of country codes. Supported country codes: | Country Code |  Country            | |-------------|----------------------| | US          | United States        | | CA          | Canada               | | MX          | Mexico               | | BR          | Brazil               | | UK          | United Kingdom       | | DE          | Germany              | | FR          | France               | | ES          | Spain                | | IT          | Italy                | | IN          | India                | | AE          | United Arab Emirates | | SA          | Saudi Arabia         | | NL          | Netherlands          | | PL          | Poland               | | BE          | Belgium              | | SE          | Sweden               | | TR          | Turkey               | | EG          | Egypt                | | JP          | Japan                | | AU          | Australia            | | SG          | Singapore            |
	CountryCodes []string                                                   `json:"countryCodes"`
	Bidding      MultiCountryAsinsThemeBasedBidRecommendationRequestBidding `json:"bidding"`
	// The bid recommendation type.
	RecommendationType string `json:"recommendationType"`
}

type _MultiCountryAsinsThemeBasedBidRecommendationRequest MultiCountryAsinsThemeBasedBidRecommendationRequest

// NewMultiCountryAsinsThemeBasedBidRecommendationRequest instantiates a new MultiCountryAsinsThemeBasedBidRecommendationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiCountryAsinsThemeBasedBidRecommendationRequest(asins []map[string]string, targetingExpressions []MultiCountryTargetingExpression, countryCodes []string, bidding MultiCountryAsinsThemeBasedBidRecommendationRequestBidding, recommendationType string) *MultiCountryAsinsThemeBasedBidRecommendationRequest {
	this := MultiCountryAsinsThemeBasedBidRecommendationRequest{}
	this.Asins = asins
	this.TargetingExpressions = targetingExpressions
	this.CountryCodes = countryCodes
	this.Bidding = bidding
	this.RecommendationType = recommendationType
	return &this
}

// NewMultiCountryAsinsThemeBasedBidRecommendationRequestWithDefaults instantiates a new MultiCountryAsinsThemeBasedBidRecommendationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiCountryAsinsThemeBasedBidRecommendationRequestWithDefaults() *MultiCountryAsinsThemeBasedBidRecommendationRequest {
	this := MultiCountryAsinsThemeBasedBidRecommendationRequest{}
	return &this
}

// GetAsins returns the Asins field value
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequest) GetAsins() []map[string]string {
	if o == nil {
		var ret []map[string]string
		return ret
	}

	return o.Asins
}

// GetAsinsOk returns a tuple with the Asins field value
// and a boolean to check if the value has been set.
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequest) GetAsinsOk() ([]map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Asins, true
}

// SetAsins sets field value
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequest) SetAsins(v []map[string]string) {
	o.Asins = v
}

// GetTargetingExpressions returns the TargetingExpressions field value
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequest) GetTargetingExpressions() []MultiCountryTargetingExpression {
	if o == nil {
		var ret []MultiCountryTargetingExpression
		return ret
	}

	return o.TargetingExpressions
}

// GetTargetingExpressionsOk returns a tuple with the TargetingExpressions field value
// and a boolean to check if the value has been set.
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequest) GetTargetingExpressionsOk() ([]MultiCountryTargetingExpression, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetingExpressions, true
}

// SetTargetingExpressions sets field value
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequest) SetTargetingExpressions(v []MultiCountryTargetingExpression) {
	o.TargetingExpressions = v
}

// GetCountryCodes returns the CountryCodes field value
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequest) GetCountryCodes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value
// and a boolean to check if the value has been set.
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequest) GetCountryCodesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CountryCodes, true
}

// SetCountryCodes sets field value
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequest) SetCountryCodes(v []string) {
	o.CountryCodes = v
}

// GetBidding returns the Bidding field value
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequest) GetBidding() MultiCountryAsinsThemeBasedBidRecommendationRequestBidding {
	if o == nil {
		var ret MultiCountryAsinsThemeBasedBidRecommendationRequestBidding
		return ret
	}

	return o.Bidding
}

// GetBiddingOk returns a tuple with the Bidding field value
// and a boolean to check if the value has been set.
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequest) GetBiddingOk() (*MultiCountryAsinsThemeBasedBidRecommendationRequestBidding, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bidding, true
}

// SetBidding sets field value
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequest) SetBidding(v MultiCountryAsinsThemeBasedBidRecommendationRequestBidding) {
	o.Bidding = v
}

// GetRecommendationType returns the RecommendationType field value
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequest) GetRecommendationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecommendationType
}

// GetRecommendationTypeOk returns a tuple with the RecommendationType field value
// and a boolean to check if the value has been set.
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequest) GetRecommendationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendationType, true
}

// SetRecommendationType sets field value
func (o *MultiCountryAsinsThemeBasedBidRecommendationRequest) SetRecommendationType(v string) {
	o.RecommendationType = v
}

func (o MultiCountryAsinsThemeBasedBidRecommendationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asins"] = o.Asins
	toSerialize["targetingExpressions"] = o.TargetingExpressions
	toSerialize["countryCodes"] = o.CountryCodes
	toSerialize["bidding"] = o.Bidding
	toSerialize["recommendationType"] = o.RecommendationType
	return toSerialize, nil
}

type NullableMultiCountryAsinsThemeBasedBidRecommendationRequest struct {
	value *MultiCountryAsinsThemeBasedBidRecommendationRequest
	isSet bool
}

func (v NullableMultiCountryAsinsThemeBasedBidRecommendationRequest) Get() *MultiCountryAsinsThemeBasedBidRecommendationRequest {
	return v.value
}

func (v *NullableMultiCountryAsinsThemeBasedBidRecommendationRequest) Set(val *MultiCountryAsinsThemeBasedBidRecommendationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiCountryAsinsThemeBasedBidRecommendationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiCountryAsinsThemeBasedBidRecommendationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiCountryAsinsThemeBasedBidRecommendationRequest(val *MultiCountryAsinsThemeBasedBidRecommendationRequest) *NullableMultiCountryAsinsThemeBasedBidRecommendationRequest {
	return &NullableMultiCountryAsinsThemeBasedBidRecommendationRequest{value: val, isSet: true}
}

func (v NullableMultiCountryAsinsThemeBasedBidRecommendationRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMultiCountryAsinsThemeBasedBidRecommendationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
