package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the GlobalRankedKeywordTargetsForAsinsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalRankedKeywordTargetsForAsinsRequest{}

// GlobalRankedKeywordTargetsForAsinsRequest This request type is used to retrieve recommended keyword targets for ASINs. Set the recommendationType to KEYWORDS_FOR_ASINS to use this request type.
type GlobalRankedKeywordTargetsForAsinsRequest struct {
	// A list of countries with targets that need to be ranked or ASINs to generate recommended keywords. The 2-letter country code is the key, and an object with ASINs and targets is the value.
	CountryCodes map[string]CountryWithTargetsAndAsins `json:"countryCodes"`
	// The bid recommendations returned will depend on the bidding strategy. <br> LEGACY_FOR_SALES - Dynamic Bids (Down only) <br> AUTO_FOR_SALES - Dynamic Bids (Up or down) <br> MANUAL - Fixed Bids
	BiddingStrategy *string `json:"biddingStrategy,omitempty"`
	// The recommendationType to retrieve recommended keyword targets for a list of ASINs.
	RecommendationType string `json:"recommendationType"`
	// Set this parameter to false if you do not want to retrieve bid suggestions for your keyword targets. Defaults to true.
	BidsEnabled *bool `json:"bidsEnabled,omitempty"`
	// The max size of recommended target. Set it to 0 if you only want to rank user-defined keywords.
	MaxRecommendations *float32 `json:"maxRecommendations,omitempty"`
	// The ranking metric value. Supported values: CLICKS, CONVERSIONS, DEFAULT. DEFAULT will be applied if no value passed in.
	SortDimension *string `json:"sortDimension,omitempty"`
	// Translations are for readability and do not affect the targeting of ads. Supported marketplace to locale mappings can be found at the <a href='https://advertising.amazon.com/API/docs/en-us/localization/#/Keyword%20Localization'>POST /keywords/localize</a> endpoint. Note: Translations will be null if locale is unsupported.
	Locale *string `json:"locale,omitempty"`
}

type _GlobalRankedKeywordTargetsForAsinsRequest GlobalRankedKeywordTargetsForAsinsRequest

// NewGlobalRankedKeywordTargetsForAsinsRequest instantiates a new GlobalRankedKeywordTargetsForAsinsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalRankedKeywordTargetsForAsinsRequest(countryCodes map[string]CountryWithTargetsAndAsins, recommendationType string) *GlobalRankedKeywordTargetsForAsinsRequest {
	this := GlobalRankedKeywordTargetsForAsinsRequest{}
	var biddingStrategy string = "LEGACY_FOR_SALES"
	this.BiddingStrategy = &biddingStrategy
	this.RecommendationType = recommendationType
	var bidsEnabled bool = true
	this.BidsEnabled = &bidsEnabled
	return &this
}

// NewGlobalRankedKeywordTargetsForAsinsRequestWithDefaults instantiates a new GlobalRankedKeywordTargetsForAsinsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalRankedKeywordTargetsForAsinsRequestWithDefaults() *GlobalRankedKeywordTargetsForAsinsRequest {
	this := GlobalRankedKeywordTargetsForAsinsRequest{}
	var biddingStrategy string = "LEGACY_FOR_SALES"
	this.BiddingStrategy = &biddingStrategy
	var bidsEnabled bool = true
	this.BidsEnabled = &bidsEnabled
	return &this
}

// GetCountryCodes returns the CountryCodes field value
func (o *GlobalRankedKeywordTargetsForAsinsRequest) GetCountryCodes() map[string]CountryWithTargetsAndAsins {
	if o == nil {
		var ret map[string]CountryWithTargetsAndAsins
		return ret
	}

	return o.CountryCodes
}

// GetCountryCodesOk returns a tuple with the CountryCodes field value
// and a boolean to check if the value has been set.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) GetCountryCodesOk() (*map[string]CountryWithTargetsAndAsins, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCodes, true
}

// SetCountryCodes sets field value
func (o *GlobalRankedKeywordTargetsForAsinsRequest) SetCountryCodes(v map[string]CountryWithTargetsAndAsins) {
	o.CountryCodes = v
}

// GetBiddingStrategy returns the BiddingStrategy field value if set, zero value otherwise.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) GetBiddingStrategy() string {
	if o == nil || IsNil(o.BiddingStrategy) {
		var ret string
		return ret
	}
	return *o.BiddingStrategy
}

// GetBiddingStrategyOk returns a tuple with the BiddingStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) GetBiddingStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.BiddingStrategy) {
		return nil, false
	}
	return o.BiddingStrategy, true
}

// HasBiddingStrategy returns a boolean if a field has been set.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) HasBiddingStrategy() bool {
	if o != nil && !IsNil(o.BiddingStrategy) {
		return true
	}

	return false
}

// SetBiddingStrategy gets a reference to the given string and assigns it to the BiddingStrategy field.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) SetBiddingStrategy(v string) {
	o.BiddingStrategy = &v
}

// GetRecommendationType returns the RecommendationType field value
func (o *GlobalRankedKeywordTargetsForAsinsRequest) GetRecommendationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecommendationType
}

// GetRecommendationTypeOk returns a tuple with the RecommendationType field value
// and a boolean to check if the value has been set.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) GetRecommendationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecommendationType, true
}

// SetRecommendationType sets field value
func (o *GlobalRankedKeywordTargetsForAsinsRequest) SetRecommendationType(v string) {
	o.RecommendationType = v
}

// GetBidsEnabled returns the BidsEnabled field value if set, zero value otherwise.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) GetBidsEnabled() bool {
	if o == nil || IsNil(o.BidsEnabled) {
		var ret bool
		return ret
	}
	return *o.BidsEnabled
}

// GetBidsEnabledOk returns a tuple with the BidsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) GetBidsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BidsEnabled) {
		return nil, false
	}
	return o.BidsEnabled, true
}

// HasBidsEnabled returns a boolean if a field has been set.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) HasBidsEnabled() bool {
	if o != nil && !IsNil(o.BidsEnabled) {
		return true
	}

	return false
}

// SetBidsEnabled gets a reference to the given bool and assigns it to the BidsEnabled field.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) SetBidsEnabled(v bool) {
	o.BidsEnabled = &v
}

// GetMaxRecommendations returns the MaxRecommendations field value if set, zero value otherwise.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) GetMaxRecommendations() float32 {
	if o == nil || IsNil(o.MaxRecommendations) {
		var ret float32
		return ret
	}
	return *o.MaxRecommendations
}

// GetMaxRecommendationsOk returns a tuple with the MaxRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) GetMaxRecommendationsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxRecommendations) {
		return nil, false
	}
	return o.MaxRecommendations, true
}

// HasMaxRecommendations returns a boolean if a field has been set.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) HasMaxRecommendations() bool {
	if o != nil && !IsNil(o.MaxRecommendations) {
		return true
	}

	return false
}

// SetMaxRecommendations gets a reference to the given float32 and assigns it to the MaxRecommendations field.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) SetMaxRecommendations(v float32) {
	o.MaxRecommendations = &v
}

// GetSortDimension returns the SortDimension field value if set, zero value otherwise.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) GetSortDimension() string {
	if o == nil || IsNil(o.SortDimension) {
		var ret string
		return ret
	}
	return *o.SortDimension
}

// GetSortDimensionOk returns a tuple with the SortDimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) GetSortDimensionOk() (*string, bool) {
	if o == nil || IsNil(o.SortDimension) {
		return nil, false
	}
	return o.SortDimension, true
}

// HasSortDimension returns a boolean if a field has been set.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) HasSortDimension() bool {
	if o != nil && !IsNil(o.SortDimension) {
		return true
	}

	return false
}

// SetSortDimension gets a reference to the given string and assigns it to the SortDimension field.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) SetSortDimension(v string) {
	o.SortDimension = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *GlobalRankedKeywordTargetsForAsinsRequest) SetLocale(v string) {
	o.Locale = &v
}

func (o GlobalRankedKeywordTargetsForAsinsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["countryCodes"] = o.CountryCodes
	if !IsNil(o.BiddingStrategy) {
		toSerialize["biddingStrategy"] = o.BiddingStrategy
	}
	toSerialize["recommendationType"] = o.RecommendationType
	if !IsNil(o.BidsEnabled) {
		toSerialize["bidsEnabled"] = o.BidsEnabled
	}
	if !IsNil(o.MaxRecommendations) {
		toSerialize["maxRecommendations"] = o.MaxRecommendations
	}
	if !IsNil(o.SortDimension) {
		toSerialize["sortDimension"] = o.SortDimension
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	return toSerialize, nil
}

type NullableGlobalRankedKeywordTargetsForAsinsRequest struct {
	value *GlobalRankedKeywordTargetsForAsinsRequest
	isSet bool
}

func (v NullableGlobalRankedKeywordTargetsForAsinsRequest) Get() *GlobalRankedKeywordTargetsForAsinsRequest {
	return v.value
}

func (v *NullableGlobalRankedKeywordTargetsForAsinsRequest) Set(val *GlobalRankedKeywordTargetsForAsinsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalRankedKeywordTargetsForAsinsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalRankedKeywordTargetsForAsinsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalRankedKeywordTargetsForAsinsRequest(val *GlobalRankedKeywordTargetsForAsinsRequest) *NullableGlobalRankedKeywordTargetsForAsinsRequest {
	return &NullableGlobalRankedKeywordTargetsForAsinsRequest{value: val, isSet: true}
}

func (v NullableGlobalRankedKeywordTargetsForAsinsRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGlobalRankedKeywordTargetsForAsinsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
