package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the KeywordBidInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeywordBidInfo{}

// KeywordBidInfo struct for KeywordBidInfo
type KeywordBidInfo struct {
	SuggestedBid *BidSuggestion `json:"suggestedBid,omitempty"`
	// Keyword match type. The default value will be BROAD.
	MatchType *string `json:"matchType,omitempty"`
	// The keyword target rank
	Rank *float32 `json:"rank,omitempty"`
	// The bid value for the keyword, in minor currency units (example: cents). The default value will be the suggested bid.
	Bid *float32 `json:"bid,omitempty"`
}

// NewKeywordBidInfo instantiates a new KeywordBidInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeywordBidInfo() *KeywordBidInfo {
	this := KeywordBidInfo{}
	return &this
}

// NewKeywordBidInfoWithDefaults instantiates a new KeywordBidInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeywordBidInfoWithDefaults() *KeywordBidInfo {
	this := KeywordBidInfo{}
	return &this
}

// GetSuggestedBid returns the SuggestedBid field value if set, zero value otherwise.
func (o *KeywordBidInfo) GetSuggestedBid() BidSuggestion {
	if o == nil || IsNil(o.SuggestedBid) {
		var ret BidSuggestion
		return ret
	}
	return *o.SuggestedBid
}

// GetSuggestedBidOk returns a tuple with the SuggestedBid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordBidInfo) GetSuggestedBidOk() (*BidSuggestion, bool) {
	if o == nil || IsNil(o.SuggestedBid) {
		return nil, false
	}
	return o.SuggestedBid, true
}

// HasSuggestedBid returns a boolean if a field has been set.
func (o *KeywordBidInfo) HasSuggestedBid() bool {
	if o != nil && !IsNil(o.SuggestedBid) {
		return true
	}

	return false
}

// SetSuggestedBid gets a reference to the given BidSuggestion and assigns it to the SuggestedBid field.
func (o *KeywordBidInfo) SetSuggestedBid(v BidSuggestion) {
	o.SuggestedBid = &v
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *KeywordBidInfo) GetMatchType() string {
	if o == nil || IsNil(o.MatchType) {
		var ret string
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordBidInfo) GetMatchTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MatchType) {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *KeywordBidInfo) HasMatchType() bool {
	if o != nil && !IsNil(o.MatchType) {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given string and assigns it to the MatchType field.
func (o *KeywordBidInfo) SetMatchType(v string) {
	o.MatchType = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *KeywordBidInfo) GetRank() float32 {
	if o == nil || IsNil(o.Rank) {
		var ret float32
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordBidInfo) GetRankOk() (*float32, bool) {
	if o == nil || IsNil(o.Rank) {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *KeywordBidInfo) HasRank() bool {
	if o != nil && !IsNil(o.Rank) {
		return true
	}

	return false
}

// SetRank gets a reference to the given float32 and assigns it to the Rank field.
func (o *KeywordBidInfo) SetRank(v float32) {
	o.Rank = &v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *KeywordBidInfo) GetBid() float32 {
	if o == nil || IsNil(o.Bid) {
		var ret float32
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordBidInfo) GetBidOk() (*float32, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *KeywordBidInfo) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given float32 and assigns it to the Bid field.
func (o *KeywordBidInfo) SetBid(v float32) {
	o.Bid = &v
}

func (o KeywordBidInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SuggestedBid) {
		toSerialize["suggestedBid"] = o.SuggestedBid
	}
	if !IsNil(o.MatchType) {
		toSerialize["matchType"] = o.MatchType
	}
	if !IsNil(o.Rank) {
		toSerialize["rank"] = o.Rank
	}
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	return toSerialize, nil
}

type NullableKeywordBidInfo struct {
	value *KeywordBidInfo
	isSet bool
}

func (v NullableKeywordBidInfo) Get() *KeywordBidInfo {
	return v.value
}

func (v *NullableKeywordBidInfo) Set(val *KeywordBidInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableKeywordBidInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableKeywordBidInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeywordBidInfo(val *KeywordBidInfo) *NullableKeywordBidInfo {
	return &NullableKeywordBidInfo{value: val, isSet: true}
}

func (v NullableKeywordBidInfo) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableKeywordBidInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
