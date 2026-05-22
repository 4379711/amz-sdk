package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the RecKeywordTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RecKeywordTarget{}

// RecKeywordTarget struct for RecKeywordTarget
type RecKeywordTarget struct {
	// The account-level ad-attributed impression share for the search-term/keyword. Provides percentage share of all ad impressions the advertiser has for the keyword in the last 30 days. This metric helps advertisers identify potential opportunities based on their share on relevant keywords. This information is only available for keywords the advertiser targeted with ad impressions.
	SearchTermImpressionShare *float32 `json:"searchTermImpressionShare,omitempty"`
	// The translation of keyword if a locale is passed in
	Translation *string `json:"translation,omitempty"`
	// A list of keyword bid info
	BidInfo []KeywordBidInfo `json:"bidInfo,omitempty"`
	// The account-level ad-attributed impression rank for the search-term/keyword. Provides [1:N] place the advertiser ranks among all advertisers for the keyword by ad impressions in a marketplace in the last 30 days. It tells an advertiser how many advertisers had higher share of ad impressions. This information is only available for keywords the advertiser targeted with ad impressions.
	SearchTermImpressionRank *float32 `json:"searchTermImpressionRank,omitempty"`
	// The keyword value
	Keyword *string `json:"keyword,omitempty"`
	// Flag that tells if keyword was selected by the user or was recommended by KRS
	UserSelectedKeyword *bool `json:"userSelectedKeyword,omitempty"`
	// The recommended keyword target id
	RecId *string `json:"recId,omitempty"`
}

// NewRecKeywordTarget instantiates a new RecKeywordTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecKeywordTarget() *RecKeywordTarget {
	this := RecKeywordTarget{}
	return &this
}

// NewRecKeywordTargetWithDefaults instantiates a new RecKeywordTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecKeywordTargetWithDefaults() *RecKeywordTarget {
	this := RecKeywordTarget{}
	return &this
}

// GetSearchTermImpressionShare returns the SearchTermImpressionShare field value if set, zero value otherwise.
func (o *RecKeywordTarget) GetSearchTermImpressionShare() float32 {
	if o == nil || IsNil(o.SearchTermImpressionShare) {
		var ret float32
		return ret
	}
	return *o.SearchTermImpressionShare
}

// GetSearchTermImpressionShareOk returns a tuple with the SearchTermImpressionShare field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecKeywordTarget) GetSearchTermImpressionShareOk() (*float32, bool) {
	if o == nil || IsNil(o.SearchTermImpressionShare) {
		return nil, false
	}
	return o.SearchTermImpressionShare, true
}

// HasSearchTermImpressionShare returns a boolean if a field has been set.
func (o *RecKeywordTarget) HasSearchTermImpressionShare() bool {
	if o != nil && !IsNil(o.SearchTermImpressionShare) {
		return true
	}

	return false
}

// SetSearchTermImpressionShare gets a reference to the given float32 and assigns it to the SearchTermImpressionShare field.
func (o *RecKeywordTarget) SetSearchTermImpressionShare(v float32) {
	o.SearchTermImpressionShare = &v
}

// GetTranslation returns the Translation field value if set, zero value otherwise.
func (o *RecKeywordTarget) GetTranslation() string {
	if o == nil || IsNil(o.Translation) {
		var ret string
		return ret
	}
	return *o.Translation
}

// GetTranslationOk returns a tuple with the Translation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecKeywordTarget) GetTranslationOk() (*string, bool) {
	if o == nil || IsNil(o.Translation) {
		return nil, false
	}
	return o.Translation, true
}

// HasTranslation returns a boolean if a field has been set.
func (o *RecKeywordTarget) HasTranslation() bool {
	if o != nil && !IsNil(o.Translation) {
		return true
	}

	return false
}

// SetTranslation gets a reference to the given string and assigns it to the Translation field.
func (o *RecKeywordTarget) SetTranslation(v string) {
	o.Translation = &v
}

// GetBidInfo returns the BidInfo field value if set, zero value otherwise.
func (o *RecKeywordTarget) GetBidInfo() []KeywordBidInfo {
	if o == nil || IsNil(o.BidInfo) {
		var ret []KeywordBidInfo
		return ret
	}
	return o.BidInfo
}

// GetBidInfoOk returns a tuple with the BidInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecKeywordTarget) GetBidInfoOk() ([]KeywordBidInfo, bool) {
	if o == nil || IsNil(o.BidInfo) {
		return nil, false
	}
	return o.BidInfo, true
}

// HasBidInfo returns a boolean if a field has been set.
func (o *RecKeywordTarget) HasBidInfo() bool {
	if o != nil && !IsNil(o.BidInfo) {
		return true
	}

	return false
}

// SetBidInfo gets a reference to the given []KeywordBidInfo and assigns it to the BidInfo field.
func (o *RecKeywordTarget) SetBidInfo(v []KeywordBidInfo) {
	o.BidInfo = v
}

// GetSearchTermImpressionRank returns the SearchTermImpressionRank field value if set, zero value otherwise.
func (o *RecKeywordTarget) GetSearchTermImpressionRank() float32 {
	if o == nil || IsNil(o.SearchTermImpressionRank) {
		var ret float32
		return ret
	}
	return *o.SearchTermImpressionRank
}

// GetSearchTermImpressionRankOk returns a tuple with the SearchTermImpressionRank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecKeywordTarget) GetSearchTermImpressionRankOk() (*float32, bool) {
	if o == nil || IsNil(o.SearchTermImpressionRank) {
		return nil, false
	}
	return o.SearchTermImpressionRank, true
}

// HasSearchTermImpressionRank returns a boolean if a field has been set.
func (o *RecKeywordTarget) HasSearchTermImpressionRank() bool {
	if o != nil && !IsNil(o.SearchTermImpressionRank) {
		return true
	}

	return false
}

// SetSearchTermImpressionRank gets a reference to the given float32 and assigns it to the SearchTermImpressionRank field.
func (o *RecKeywordTarget) SetSearchTermImpressionRank(v float32) {
	o.SearchTermImpressionRank = &v
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *RecKeywordTarget) GetKeyword() string {
	if o == nil || IsNil(o.Keyword) {
		var ret string
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecKeywordTarget) GetKeywordOk() (*string, bool) {
	if o == nil || IsNil(o.Keyword) {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *RecKeywordTarget) HasKeyword() bool {
	if o != nil && !IsNil(o.Keyword) {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given string and assigns it to the Keyword field.
func (o *RecKeywordTarget) SetKeyword(v string) {
	o.Keyword = &v
}

// GetUserSelectedKeyword returns the UserSelectedKeyword field value if set, zero value otherwise.
func (o *RecKeywordTarget) GetUserSelectedKeyword() bool {
	if o == nil || IsNil(o.UserSelectedKeyword) {
		var ret bool
		return ret
	}
	return *o.UserSelectedKeyword
}

// GetUserSelectedKeywordOk returns a tuple with the UserSelectedKeyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecKeywordTarget) GetUserSelectedKeywordOk() (*bool, bool) {
	if o == nil || IsNil(o.UserSelectedKeyword) {
		return nil, false
	}
	return o.UserSelectedKeyword, true
}

// HasUserSelectedKeyword returns a boolean if a field has been set.
func (o *RecKeywordTarget) HasUserSelectedKeyword() bool {
	if o != nil && !IsNil(o.UserSelectedKeyword) {
		return true
	}

	return false
}

// SetUserSelectedKeyword gets a reference to the given bool and assigns it to the UserSelectedKeyword field.
func (o *RecKeywordTarget) SetUserSelectedKeyword(v bool) {
	o.UserSelectedKeyword = &v
}

// GetRecId returns the RecId field value if set, zero value otherwise.
func (o *RecKeywordTarget) GetRecId() string {
	if o == nil || IsNil(o.RecId) {
		var ret string
		return ret
	}
	return *o.RecId
}

// GetRecIdOk returns a tuple with the RecId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RecKeywordTarget) GetRecIdOk() (*string, bool) {
	if o == nil || IsNil(o.RecId) {
		return nil, false
	}
	return o.RecId, true
}

// HasRecId returns a boolean if a field has been set.
func (o *RecKeywordTarget) HasRecId() bool {
	if o != nil && !IsNil(o.RecId) {
		return true
	}

	return false
}

// SetRecId gets a reference to the given string and assigns it to the RecId field.
func (o *RecKeywordTarget) SetRecId(v string) {
	o.RecId = &v
}

func (o RecKeywordTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SearchTermImpressionShare) {
		toSerialize["searchTermImpressionShare"] = o.SearchTermImpressionShare
	}
	if !IsNil(o.Translation) {
		toSerialize["translation"] = o.Translation
	}
	if !IsNil(o.BidInfo) {
		toSerialize["bidInfo"] = o.BidInfo
	}
	if !IsNil(o.SearchTermImpressionRank) {
		toSerialize["searchTermImpressionRank"] = o.SearchTermImpressionRank
	}
	if !IsNil(o.Keyword) {
		toSerialize["keyword"] = o.Keyword
	}
	if !IsNil(o.UserSelectedKeyword) {
		toSerialize["userSelectedKeyword"] = o.UserSelectedKeyword
	}
	if !IsNil(o.RecId) {
		toSerialize["recId"] = o.RecId
	}
	return toSerialize, nil
}

type NullableRecKeywordTarget struct {
	value *RecKeywordTarget
	isSet bool
}

func (v NullableRecKeywordTarget) Get() *RecKeywordTarget {
	return v.value
}

func (v *NullableRecKeywordTarget) Set(val *RecKeywordTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableRecKeywordTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableRecKeywordTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecKeywordTarget(val *RecKeywordTarget) *NullableRecKeywordTarget {
	return &NullableRecKeywordTarget{value: val, isSet: true}
}

func (v NullableRecKeywordTarget) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableRecKeywordTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
