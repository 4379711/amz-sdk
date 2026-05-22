package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the KeywordTargetingRecommendation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeywordTargetingRecommendation{}

// KeywordTargetingRecommendation Contains suggested recommendation for the keyword targeting.
type KeywordTargetingRecommendation struct {
	// The identifier of the keyword targeting.
	KeywordId *string `json:"keywordId,omitempty"`
	// The suggested bid value associated with this keyword targeting clause.
	SuggestedBid *float64 `json:"suggestedBid,omitempty"`
	// Keyword match type. | Value | Description | | --- | --- | | `BROAD` | Use BROAD to broadly match your keyword targeting ads with search queries.| | `EXACT` | Use EXACT to exactly match your keyword targeting ads with search queries.| | `PHRASE` | Use PHRASE to match your keyword targeting ads with search phrases.|
	MatchType *string `json:"matchType,omitempty"`
	// Type of action for the keyword targeting.
	Action *string `json:"action,omitempty"`
	// The ad group identifier.
	AdGroupId *string `json:"adGroupId,omitempty"`
	// The keyword text.
	KeywordText *string `json:"keywordText,omitempty"`
}

// NewKeywordTargetingRecommendation instantiates a new KeywordTargetingRecommendation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeywordTargetingRecommendation() *KeywordTargetingRecommendation {
	this := KeywordTargetingRecommendation{}
	return &this
}

// NewKeywordTargetingRecommendationWithDefaults instantiates a new KeywordTargetingRecommendation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeywordTargetingRecommendationWithDefaults() *KeywordTargetingRecommendation {
	this := KeywordTargetingRecommendation{}
	return &this
}

// GetKeywordId returns the KeywordId field value if set, zero value otherwise.
func (o *KeywordTargetingRecommendation) GetKeywordId() string {
	if o == nil || IsNil(o.KeywordId) {
		var ret string
		return ret
	}
	return *o.KeywordId
}

// GetKeywordIdOk returns a tuple with the KeywordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTargetingRecommendation) GetKeywordIdOk() (*string, bool) {
	if o == nil || IsNil(o.KeywordId) {
		return nil, false
	}
	return o.KeywordId, true
}

// HasKeywordId returns a boolean if a field has been set.
func (o *KeywordTargetingRecommendation) HasKeywordId() bool {
	if o != nil && !IsNil(o.KeywordId) {
		return true
	}

	return false
}

// SetKeywordId gets a reference to the given string and assigns it to the KeywordId field.
func (o *KeywordTargetingRecommendation) SetKeywordId(v string) {
	o.KeywordId = &v
}

// GetSuggestedBid returns the SuggestedBid field value if set, zero value otherwise.
func (o *KeywordTargetingRecommendation) GetSuggestedBid() float64 {
	if o == nil || IsNil(o.SuggestedBid) {
		var ret float64
		return ret
	}
	return *o.SuggestedBid
}

// GetSuggestedBidOk returns a tuple with the SuggestedBid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTargetingRecommendation) GetSuggestedBidOk() (*float64, bool) {
	if o == nil || IsNil(o.SuggestedBid) {
		return nil, false
	}
	return o.SuggestedBid, true
}

// HasSuggestedBid returns a boolean if a field has been set.
func (o *KeywordTargetingRecommendation) HasSuggestedBid() bool {
	if o != nil && !IsNil(o.SuggestedBid) {
		return true
	}

	return false
}

// SetSuggestedBid gets a reference to the given float64 and assigns it to the SuggestedBid field.
func (o *KeywordTargetingRecommendation) SetSuggestedBid(v float64) {
	o.SuggestedBid = &v
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *KeywordTargetingRecommendation) GetMatchType() string {
	if o == nil || IsNil(o.MatchType) {
		var ret string
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTargetingRecommendation) GetMatchTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MatchType) {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *KeywordTargetingRecommendation) HasMatchType() bool {
	if o != nil && !IsNil(o.MatchType) {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given string and assigns it to the MatchType field.
func (o *KeywordTargetingRecommendation) SetMatchType(v string) {
	o.MatchType = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *KeywordTargetingRecommendation) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTargetingRecommendation) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *KeywordTargetingRecommendation) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *KeywordTargetingRecommendation) SetAction(v string) {
	o.Action = &v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *KeywordTargetingRecommendation) GetAdGroupId() string {
	if o == nil || IsNil(o.AdGroupId) {
		var ret string
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTargetingRecommendation) GetAdGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *KeywordTargetingRecommendation) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given string and assigns it to the AdGroupId field.
func (o *KeywordTargetingRecommendation) SetAdGroupId(v string) {
	o.AdGroupId = &v
}

// GetKeywordText returns the KeywordText field value if set, zero value otherwise.
func (o *KeywordTargetingRecommendation) GetKeywordText() string {
	if o == nil || IsNil(o.KeywordText) {
		var ret string
		return ret
	}
	return *o.KeywordText
}

// GetKeywordTextOk returns a tuple with the KeywordText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTargetingRecommendation) GetKeywordTextOk() (*string, bool) {
	if o == nil || IsNil(o.KeywordText) {
		return nil, false
	}
	return o.KeywordText, true
}

// HasKeywordText returns a boolean if a field has been set.
func (o *KeywordTargetingRecommendation) HasKeywordText() bool {
	if o != nil && !IsNil(o.KeywordText) {
		return true
	}

	return false
}

// SetKeywordText gets a reference to the given string and assigns it to the KeywordText field.
func (o *KeywordTargetingRecommendation) SetKeywordText(v string) {
	o.KeywordText = &v
}

func (o KeywordTargetingRecommendation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KeywordId) {
		toSerialize["keywordId"] = o.KeywordId
	}
	if !IsNil(o.SuggestedBid) {
		toSerialize["suggestedBid"] = o.SuggestedBid
	}
	if !IsNil(o.MatchType) {
		toSerialize["matchType"] = o.MatchType
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	if !IsNil(o.KeywordText) {
		toSerialize["keywordText"] = o.KeywordText
	}
	return toSerialize, nil
}

type NullableKeywordTargetingRecommendation struct {
	value *KeywordTargetingRecommendation
	isSet bool
}

func (v NullableKeywordTargetingRecommendation) Get() *KeywordTargetingRecommendation {
	return v.value
}

func (v *NullableKeywordTargetingRecommendation) Set(val *KeywordTargetingRecommendation) {
	v.value = val
	v.isSet = true
}

func (v NullableKeywordTargetingRecommendation) IsSet() bool {
	return v.isSet
}

func (v *NullableKeywordTargetingRecommendation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeywordTargetingRecommendation(val *KeywordTargetingRecommendation) *NullableKeywordTargetingRecommendation {
	return &NullableKeywordTargetingRecommendation{value: val, isSet: true}
}

func (v NullableKeywordTargetingRecommendation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableKeywordTargetingRecommendation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
