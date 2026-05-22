package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the AsinsBasedRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AsinsBasedRequest{}

// AsinsBasedRequest struct for AsinsBasedRequest
type AsinsBasedRequest struct {
	// The bid recommendations returned will depend on the bidding strategy. <br> LEGACY_FOR_SALES - Dynamic Bids (Down only) <br> AUTO_FOR_SALES - Dynamic Bids (Up or down) <br> MANUAL - Fixed Bids
	BiddingStrategy *string `json:"biddingStrategy,omitempty"`
	// The recommendationType to retrieve recommended keyword targets for a list of ASINs.
	RecommendationType *string `json:"recommendationType,omitempty"`
	// Set this parameter to false if you do not want to retrieve bid suggestions for your keyword targets. Defaults to true.
	BidsEnabled *bool `json:"bidsEnabled,omitempty"`
	// The max size of recommended target. Set it to 0 if you only want to rank user-defined keywords.
	MaxRecommendations *float32 `json:"maxRecommendations,omitempty"`
	// The ranking metric value. Supported values: CLICKS, CONVERSIONS, DEFAULT. DEFAULT will be applied if no value passed in.
	SortDimension *string `json:"sortDimension,omitempty"`
	// Translations are for readability and do not affect the targeting of ads. Supported marketplace to locale mappings can be found at the <a href='https://advertising.amazon.com/API/docs/en-us/localization/#/Keyword%20Localization'>POST /keywords/localize</a> endpoint. Note: Translations will be null if locale is unsupported.
	Locale *string `json:"locale,omitempty"`
}

// NewAsinsBasedRequest instantiates a new AsinsBasedRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAsinsBasedRequest() *AsinsBasedRequest {
	this := AsinsBasedRequest{}
	return &this
}

// NewAsinsBasedRequestWithDefaults instantiates a new AsinsBasedRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAsinsBasedRequestWithDefaults() *AsinsBasedRequest {
	this := AsinsBasedRequest{}
	var biddingStrategy string = "LEGACY_FOR_SALES"
	this.BiddingStrategy = &biddingStrategy
	var bidsEnabled bool = true
	this.BidsEnabled = &bidsEnabled
	return &this
}

// GetBiddingStrategy returns the BiddingStrategy field value if set, zero value otherwise.
func (o *AsinsBasedRequest) GetBiddingStrategy() string {
	if o == nil || IsNil(o.BiddingStrategy) {
		var ret string
		return ret
	}
	return *o.BiddingStrategy
}

// GetBiddingStrategyOk returns a tuple with the BiddingStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinsBasedRequest) GetBiddingStrategyOk() (*string, bool) {
	if o == nil || IsNil(o.BiddingStrategy) {
		return nil, false
	}
	return o.BiddingStrategy, true
}

// HasBiddingStrategy returns a boolean if a field has been set.
func (o *AsinsBasedRequest) HasBiddingStrategy() bool {
	if o != nil && !IsNil(o.BiddingStrategy) {
		return true
	}

	return false
}

// SetBiddingStrategy gets a reference to the given string and assigns it to the BiddingStrategy field.
func (o *AsinsBasedRequest) SetBiddingStrategy(v string) {
	o.BiddingStrategy = &v
}

// GetRecommendationType returns the RecommendationType field value if set, zero value otherwise.
func (o *AsinsBasedRequest) GetRecommendationType() string {
	if o == nil || IsNil(o.RecommendationType) {
		var ret string
		return ret
	}
	return *o.RecommendationType
}

// GetRecommendationTypeOk returns a tuple with the RecommendationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinsBasedRequest) GetRecommendationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecommendationType) {
		return nil, false
	}
	return o.RecommendationType, true
}

// HasRecommendationType returns a boolean if a field has been set.
func (o *AsinsBasedRequest) HasRecommendationType() bool {
	if o != nil && !IsNil(o.RecommendationType) {
		return true
	}

	return false
}

// SetRecommendationType gets a reference to the given string and assigns it to the RecommendationType field.
func (o *AsinsBasedRequest) SetRecommendationType(v string) {
	o.RecommendationType = &v
}

// GetBidsEnabled returns the BidsEnabled field value if set, zero value otherwise.
func (o *AsinsBasedRequest) GetBidsEnabled() bool {
	if o == nil || IsNil(o.BidsEnabled) {
		var ret bool
		return ret
	}
	return *o.BidsEnabled
}

// GetBidsEnabledOk returns a tuple with the BidsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinsBasedRequest) GetBidsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BidsEnabled) {
		return nil, false
	}
	return o.BidsEnabled, true
}

// HasBidsEnabled returns a boolean if a field has been set.
func (o *AsinsBasedRequest) HasBidsEnabled() bool {
	if o != nil && !IsNil(o.BidsEnabled) {
		return true
	}

	return false
}

// SetBidsEnabled gets a reference to the given bool and assigns it to the BidsEnabled field.
func (o *AsinsBasedRequest) SetBidsEnabled(v bool) {
	o.BidsEnabled = &v
}

// GetMaxRecommendations returns the MaxRecommendations field value if set, zero value otherwise.
func (o *AsinsBasedRequest) GetMaxRecommendations() float32 {
	if o == nil || IsNil(o.MaxRecommendations) {
		var ret float32
		return ret
	}
	return *o.MaxRecommendations
}

// GetMaxRecommendationsOk returns a tuple with the MaxRecommendations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinsBasedRequest) GetMaxRecommendationsOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxRecommendations) {
		return nil, false
	}
	return o.MaxRecommendations, true
}

// HasMaxRecommendations returns a boolean if a field has been set.
func (o *AsinsBasedRequest) HasMaxRecommendations() bool {
	if o != nil && !IsNil(o.MaxRecommendations) {
		return true
	}

	return false
}

// SetMaxRecommendations gets a reference to the given float32 and assigns it to the MaxRecommendations field.
func (o *AsinsBasedRequest) SetMaxRecommendations(v float32) {
	o.MaxRecommendations = &v
}

// GetSortDimension returns the SortDimension field value if set, zero value otherwise.
func (o *AsinsBasedRequest) GetSortDimension() string {
	if o == nil || IsNil(o.SortDimension) {
		var ret string
		return ret
	}
	return *o.SortDimension
}

// GetSortDimensionOk returns a tuple with the SortDimension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinsBasedRequest) GetSortDimensionOk() (*string, bool) {
	if o == nil || IsNil(o.SortDimension) {
		return nil, false
	}
	return o.SortDimension, true
}

// HasSortDimension returns a boolean if a field has been set.
func (o *AsinsBasedRequest) HasSortDimension() bool {
	if o != nil && !IsNil(o.SortDimension) {
		return true
	}

	return false
}

// SetSortDimension gets a reference to the given string and assigns it to the SortDimension field.
func (o *AsinsBasedRequest) SetSortDimension(v string) {
	o.SortDimension = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *AsinsBasedRequest) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AsinsBasedRequest) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *AsinsBasedRequest) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *AsinsBasedRequest) SetLocale(v string) {
	o.Locale = &v
}

func (o AsinsBasedRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BiddingStrategy) {
		toSerialize["biddingStrategy"] = o.BiddingStrategy
	}
	if !IsNil(o.RecommendationType) {
		toSerialize["recommendationType"] = o.RecommendationType
	}
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

type NullableAsinsBasedRequest struct {
	value *AsinsBasedRequest
	isSet bool
}

func (v NullableAsinsBasedRequest) Get() *AsinsBasedRequest {
	return v.value
}

func (v *NullableAsinsBasedRequest) Set(val *AsinsBasedRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAsinsBasedRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAsinsBasedRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAsinsBasedRequest(val *AsinsBasedRequest) *NullableAsinsBasedRequest {
	return &NullableAsinsBasedRequest{value: val, isSet: true}
}

func (v NullableAsinsBasedRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAsinsBasedRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
