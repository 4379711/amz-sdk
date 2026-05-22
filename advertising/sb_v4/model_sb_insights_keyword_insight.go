package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the SBInsightsKeywordInsight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SBInsightsKeywordInsight{}

// SBInsightsKeywordInsight Insights for keywords selected for targeting.
type SBInsightsKeywordInsight struct {
	Alerts []SBInsightsKeywordAlertType `json:"alerts,omitempty"`
	// The account-level ad-attributed impression share for the search-term / keyword. Provides percentage share of all ad impressions the advertiser has for the keyword in the last 7 days. This metric helps advertisers identify potential opportunities based on their share of relevant keywords. This information is only available for keywords the advertiser targeted with ad impressions. Only available in the following marketplaces: US, CA, MX, UK, DE, FR, ES, IT, IN, JP.
	SearchTermImpressionShare *float64             `json:"searchTermImpressionShare,omitempty"`
	MatchType                 *SBInsightsMatchType `json:"matchType,omitempty"`
	// Correlates the ad group to the ad group array index specified in the request. Zero-based.
	AdGroupIndex *int32 `json:"adGroupIndex,omitempty"`
	// The account-level ad-attributed impression rank for the search-term / keyword. Provides the [1:N] place the advertiser ranks among all advertisers for the keyword by ad impressions in a marketplace in the last 7 days. It tells an advertiser how many advertisers had higher share of ad impressions. This information is only available for keywords the advertiser targeted with ad impressions. Only available in the following marketplaces: US, CA, MX, UK, DE, FR, ES, IT, IN, JP.
	SearchTermImpressionRank *int32 `json:"searchTermImpressionRank,omitempty"`
	// The associated bid. Note that this value must be less than the budget associated with the Advertiser account. For more information, see [supported features](https://advertising.amazon.com/API/docs/v2/guides/supported_features).
	Bid *float64 `json:"bid,omitempty"`
	// Correlates the keyword to the keyword array index specified in the request. Zero-based.
	KeywordIndex *int32 `json:"keywordIndex,omitempty"`
	// The keyword text. Maximum of 10 words.
	KeywordText *string `json:"keywordText,omitempty"`
}

// NewSBInsightsKeywordInsight instantiates a new SBInsightsKeywordInsight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSBInsightsKeywordInsight() *SBInsightsKeywordInsight {
	this := SBInsightsKeywordInsight{}
	return &this
}

// NewSBInsightsKeywordInsightWithDefaults instantiates a new SBInsightsKeywordInsight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSBInsightsKeywordInsightWithDefaults() *SBInsightsKeywordInsight {
	this := SBInsightsKeywordInsight{}
	return &this
}

// GetAlerts returns the Alerts field value if set, zero value otherwise.
func (o *SBInsightsKeywordInsight) GetAlerts() []SBInsightsKeywordAlertType {
	if o == nil || IsNil(o.Alerts) {
		var ret []SBInsightsKeywordAlertType
		return ret
	}
	return o.Alerts
}

// GetAlertsOk returns a tuple with the Alerts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBInsightsKeywordInsight) GetAlertsOk() ([]SBInsightsKeywordAlertType, bool) {
	if o == nil || IsNil(o.Alerts) {
		return nil, false
	}
	return o.Alerts, true
}

// HasAlerts returns a boolean if a field has been set.
func (o *SBInsightsKeywordInsight) HasAlerts() bool {
	if o != nil && !IsNil(o.Alerts) {
		return true
	}

	return false
}

// SetAlerts gets a reference to the given []SBInsightsKeywordAlertType and assigns it to the Alerts field.
func (o *SBInsightsKeywordInsight) SetAlerts(v []SBInsightsKeywordAlertType) {
	o.Alerts = v
}

// GetSearchTermImpressionShare returns the SearchTermImpressionShare field value if set, zero value otherwise.
func (o *SBInsightsKeywordInsight) GetSearchTermImpressionShare() float64 {
	if o == nil || IsNil(o.SearchTermImpressionShare) {
		var ret float64
		return ret
	}
	return *o.SearchTermImpressionShare
}

// GetSearchTermImpressionShareOk returns a tuple with the SearchTermImpressionShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBInsightsKeywordInsight) GetSearchTermImpressionShareOk() (*float64, bool) {
	if o == nil || IsNil(o.SearchTermImpressionShare) {
		return nil, false
	}
	return o.SearchTermImpressionShare, true
}

// HasSearchTermImpressionShare returns a boolean if a field has been set.
func (o *SBInsightsKeywordInsight) HasSearchTermImpressionShare() bool {
	if o != nil && !IsNil(o.SearchTermImpressionShare) {
		return true
	}

	return false
}

// SetSearchTermImpressionShare gets a reference to the given float64 and assigns it to the SearchTermImpressionShare field.
func (o *SBInsightsKeywordInsight) SetSearchTermImpressionShare(v float64) {
	o.SearchTermImpressionShare = &v
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *SBInsightsKeywordInsight) GetMatchType() SBInsightsMatchType {
	if o == nil || IsNil(o.MatchType) {
		var ret SBInsightsMatchType
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBInsightsKeywordInsight) GetMatchTypeOk() (*SBInsightsMatchType, bool) {
	if o == nil || IsNil(o.MatchType) {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *SBInsightsKeywordInsight) HasMatchType() bool {
	if o != nil && !IsNil(o.MatchType) {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given SBInsightsMatchType and assigns it to the MatchType field.
func (o *SBInsightsKeywordInsight) SetMatchType(v SBInsightsMatchType) {
	o.MatchType = &v
}

// GetAdGroupIndex returns the AdGroupIndex field value if set, zero value otherwise.
func (o *SBInsightsKeywordInsight) GetAdGroupIndex() int32 {
	if o == nil || IsNil(o.AdGroupIndex) {
		var ret int32
		return ret
	}
	return *o.AdGroupIndex
}

// GetAdGroupIndexOk returns a tuple with the AdGroupIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBInsightsKeywordInsight) GetAdGroupIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.AdGroupIndex) {
		return nil, false
	}
	return o.AdGroupIndex, true
}

// HasAdGroupIndex returns a boolean if a field has been set.
func (o *SBInsightsKeywordInsight) HasAdGroupIndex() bool {
	if o != nil && !IsNil(o.AdGroupIndex) {
		return true
	}

	return false
}

// SetAdGroupIndex gets a reference to the given int32 and assigns it to the AdGroupIndex field.
func (o *SBInsightsKeywordInsight) SetAdGroupIndex(v int32) {
	o.AdGroupIndex = &v
}

// GetSearchTermImpressionRank returns the SearchTermImpressionRank field value if set, zero value otherwise.
func (o *SBInsightsKeywordInsight) GetSearchTermImpressionRank() int32 {
	if o == nil || IsNil(o.SearchTermImpressionRank) {
		var ret int32
		return ret
	}
	return *o.SearchTermImpressionRank
}

// GetSearchTermImpressionRankOk returns a tuple with the SearchTermImpressionRank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBInsightsKeywordInsight) GetSearchTermImpressionRankOk() (*int32, bool) {
	if o == nil || IsNil(o.SearchTermImpressionRank) {
		return nil, false
	}
	return o.SearchTermImpressionRank, true
}

// HasSearchTermImpressionRank returns a boolean if a field has been set.
func (o *SBInsightsKeywordInsight) HasSearchTermImpressionRank() bool {
	if o != nil && !IsNil(o.SearchTermImpressionRank) {
		return true
	}

	return false
}

// SetSearchTermImpressionRank gets a reference to the given int32 and assigns it to the SearchTermImpressionRank field.
func (o *SBInsightsKeywordInsight) SetSearchTermImpressionRank(v int32) {
	o.SearchTermImpressionRank = &v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *SBInsightsKeywordInsight) GetBid() float64 {
	if o == nil || IsNil(o.Bid) {
		var ret float64
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBInsightsKeywordInsight) GetBidOk() (*float64, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *SBInsightsKeywordInsight) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given float64 and assigns it to the Bid field.
func (o *SBInsightsKeywordInsight) SetBid(v float64) {
	o.Bid = &v
}

// GetKeywordIndex returns the KeywordIndex field value if set, zero value otherwise.
func (o *SBInsightsKeywordInsight) GetKeywordIndex() int32 {
	if o == nil || IsNil(o.KeywordIndex) {
		var ret int32
		return ret
	}
	return *o.KeywordIndex
}

// GetKeywordIndexOk returns a tuple with the KeywordIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBInsightsKeywordInsight) GetKeywordIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.KeywordIndex) {
		return nil, false
	}
	return o.KeywordIndex, true
}

// HasKeywordIndex returns a boolean if a field has been set.
func (o *SBInsightsKeywordInsight) HasKeywordIndex() bool {
	if o != nil && !IsNil(o.KeywordIndex) {
		return true
	}

	return false
}

// SetKeywordIndex gets a reference to the given int32 and assigns it to the KeywordIndex field.
func (o *SBInsightsKeywordInsight) SetKeywordIndex(v int32) {
	o.KeywordIndex = &v
}

// GetKeywordText returns the KeywordText field value if set, zero value otherwise.
func (o *SBInsightsKeywordInsight) GetKeywordText() string {
	if o == nil || IsNil(o.KeywordText) {
		var ret string
		return ret
	}
	return *o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SBInsightsKeywordInsight) GetKeywordTextOk() (*string, bool) {
	if o == nil || IsNil(o.KeywordText) {
		return nil, false
	}
	return o.KeywordText, true
}

// HasKeywordText returns a boolean if a field has been set.
func (o *SBInsightsKeywordInsight) HasKeywordText() bool {
	if o != nil && !IsNil(o.KeywordText) {
		return true
	}

	return false
}

// SetKeywordText gets a reference to the given string and assigns it to the KeywordText field.
func (o *SBInsightsKeywordInsight) SetKeywordText(v string) {
	o.KeywordText = &v
}

func (o SBInsightsKeywordInsight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alerts) {
		toSerialize["alerts"] = o.Alerts
	}
	if !IsNil(o.SearchTermImpressionShare) {
		toSerialize["searchTermImpressionShare"] = o.SearchTermImpressionShare
	}
	if !IsNil(o.MatchType) {
		toSerialize["matchType"] = o.MatchType
	}
	if !IsNil(o.AdGroupIndex) {
		toSerialize["adGroupIndex"] = o.AdGroupIndex
	}
	if !IsNil(o.SearchTermImpressionRank) {
		toSerialize["searchTermImpressionRank"] = o.SearchTermImpressionRank
	}
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	if !IsNil(o.KeywordIndex) {
		toSerialize["keywordIndex"] = o.KeywordIndex
	}
	if !IsNil(o.KeywordText) {
		toSerialize["keywordText"] = o.KeywordText
	}
	return toSerialize, nil
}

type NullableSBInsightsKeywordInsight struct {
	value *SBInsightsKeywordInsight
	isSet bool
}

func (v NullableSBInsightsKeywordInsight) Get() *SBInsightsKeywordInsight {
	return v.value
}

func (v *NullableSBInsightsKeywordInsight) Set(val *SBInsightsKeywordInsight) {
	v.value = val
	v.isSet = true
}

func (v NullableSBInsightsKeywordInsight) IsSet() bool {
	return v.isSet
}

func (v *NullableSBInsightsKeywordInsight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSBInsightsKeywordInsight(val *SBInsightsKeywordInsight) *NullableSBInsightsKeywordInsight {
	return &NullableSBInsightsKeywordInsight{value: val, isSet: true}
}

func (v NullableSBInsightsKeywordInsight) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSBInsightsKeywordInsight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
