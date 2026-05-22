package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the KeywordTargetResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeywordTargetResponse{}

// KeywordTargetResponse struct for KeywordTargetResponse
type KeywordTargetResponse struct {
	SuggestedBid *BidSuggestion `json:"suggestedBid,omitempty"`
	// The translation of keyword if a locale is passed in
	Translation *string `json:"translation,omitempty"`
	// The keyword target rank
	Rank *float32 `json:"rank,omitempty"`
	// Keyword match type. The default value will be BROAD.
	MatchType *string `json:"matchType,omitempty"`
	// The keyword value
	Keyword *string `json:"keyword,omitempty"`
	// The bid value for the keyword, in minor currency units (example: cents). The default value will be the suggested bid.
	Bid *float64 `json:"bid,omitempty"`
	// Flag that tells if keyword was selected by the user or was recommended by KRS
	UserSelectedKeyword *bool `json:"userSelectedKeyword,omitempty"`
}

// NewKeywordTargetResponse instantiates a new KeywordTargetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeywordTargetResponse() *KeywordTargetResponse {
	this := KeywordTargetResponse{}
	return &this
}

// NewKeywordTargetResponseWithDefaults instantiates a new KeywordTargetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeywordTargetResponseWithDefaults() *KeywordTargetResponse {
	this := KeywordTargetResponse{}
	return &this
}

// GetSuggestedBid returns the SuggestedBid field value if set, zero value otherwise.
func (o *KeywordTargetResponse) GetSuggestedBid() BidSuggestion {
	if o == nil || IsNil(o.SuggestedBid) {
		var ret BidSuggestion
		return ret
	}
	return *o.SuggestedBid
}

// GetSuggestedBidOk returns a tuple with the SuggestedBid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTargetResponse) GetSuggestedBidOk() (*BidSuggestion, bool) {
	if o == nil || IsNil(o.SuggestedBid) {
		return nil, false
	}
	return o.SuggestedBid, true
}

// HasSuggestedBid returns a boolean if a field has been set.
func (o *KeywordTargetResponse) HasSuggestedBid() bool {
	if o != nil && !IsNil(o.SuggestedBid) {
		return true
	}

	return false
}

// SetSuggestedBid gets a reference to the given BidSuggestion and assigns it to the SuggestedBid field.
func (o *KeywordTargetResponse) SetSuggestedBid(v BidSuggestion) {
	o.SuggestedBid = &v
}

// GetTranslation returns the Translation field value if set, zero value otherwise.
func (o *KeywordTargetResponse) GetTranslation() string {
	if o == nil || IsNil(o.Translation) {
		var ret string
		return ret
	}
	return *o.Translation
}

// GetTranslationOk returns a tuple with the Translation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTargetResponse) GetTranslationOk() (*string, bool) {
	if o == nil || IsNil(o.Translation) {
		return nil, false
	}
	return o.Translation, true
}

// HasTranslation returns a boolean if a field has been set.
func (o *KeywordTargetResponse) HasTranslation() bool {
	if o != nil && !IsNil(o.Translation) {
		return true
	}

	return false
}

// SetTranslation gets a reference to the given string and assigns it to the Translation field.
func (o *KeywordTargetResponse) SetTranslation(v string) {
	o.Translation = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *KeywordTargetResponse) GetRank() float32 {
	if o == nil || IsNil(o.Rank) {
		var ret float32
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTargetResponse) GetRankOk() (*float32, bool) {
	if o == nil || IsNil(o.Rank) {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *KeywordTargetResponse) HasRank() bool {
	if o != nil && !IsNil(o.Rank) {
		return true
	}

	return false
}

// SetRank gets a reference to the given float32 and assigns it to the Rank field.
func (o *KeywordTargetResponse) SetRank(v float32) {
	o.Rank = &v
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *KeywordTargetResponse) GetMatchType() string {
	if o == nil || IsNil(o.MatchType) {
		var ret string
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTargetResponse) GetMatchTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MatchType) {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *KeywordTargetResponse) HasMatchType() bool {
	if o != nil && !IsNil(o.MatchType) {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given string and assigns it to the MatchType field.
func (o *KeywordTargetResponse) SetMatchType(v string) {
	o.MatchType = &v
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *KeywordTargetResponse) GetKeyword() string {
	if o == nil || IsNil(o.Keyword) {
		var ret string
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTargetResponse) GetKeywordOk() (*string, bool) {
	if o == nil || IsNil(o.Keyword) {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *KeywordTargetResponse) HasKeyword() bool {
	if o != nil && !IsNil(o.Keyword) {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given string and assigns it to the Keyword field.
func (o *KeywordTargetResponse) SetKeyword(v string) {
	o.Keyword = &v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *KeywordTargetResponse) GetBid() float64 {
	if o == nil || IsNil(o.Bid) {
		var ret float64
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTargetResponse) GetBidOk() (*float64, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *KeywordTargetResponse) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given float64 and assigns it to the Bid field.
func (o *KeywordTargetResponse) SetBid(v float64) {
	o.Bid = &v
}

// GetUserSelectedKeyword returns the UserSelectedKeyword field value if set, zero value otherwise.
func (o *KeywordTargetResponse) GetUserSelectedKeyword() bool {
	if o == nil || IsNil(o.UserSelectedKeyword) {
		var ret bool
		return ret
	}
	return *o.UserSelectedKeyword
}

// GetUserSelectedKeywordOk returns a tuple with the UserSelectedKeyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTargetResponse) GetUserSelectedKeywordOk() (*bool, bool) {
	if o == nil || IsNil(o.UserSelectedKeyword) {
		return nil, false
	}
	return o.UserSelectedKeyword, true
}

// HasUserSelectedKeyword returns a boolean if a field has been set.
func (o *KeywordTargetResponse) HasUserSelectedKeyword() bool {
	if o != nil && !IsNil(o.UserSelectedKeyword) {
		return true
	}

	return false
}

// SetUserSelectedKeyword gets a reference to the given bool and assigns it to the UserSelectedKeyword field.
func (o *KeywordTargetResponse) SetUserSelectedKeyword(v bool) {
	o.UserSelectedKeyword = &v
}

func (o KeywordTargetResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SuggestedBid) {
		toSerialize["suggestedBid"] = o.SuggestedBid
	}
	if !IsNil(o.Translation) {
		toSerialize["translation"] = o.Translation
	}
	if !IsNil(o.Rank) {
		toSerialize["rank"] = o.Rank
	}
	if !IsNil(o.MatchType) {
		toSerialize["matchType"] = o.MatchType
	}
	if !IsNil(o.Keyword) {
		toSerialize["keyword"] = o.Keyword
	}
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	if !IsNil(o.UserSelectedKeyword) {
		toSerialize["userSelectedKeyword"] = o.UserSelectedKeyword
	}
	return toSerialize, nil
}

type NullableKeywordTargetResponse struct {
	value *KeywordTargetResponse
	isSet bool
}

func (v NullableKeywordTargetResponse) Get() *KeywordTargetResponse {
	return v.value
}

func (v *NullableKeywordTargetResponse) Set(val *KeywordTargetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableKeywordTargetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableKeywordTargetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeywordTargetResponse(val *KeywordTargetResponse) *NullableKeywordTargetResponse {
	return &NullableKeywordTargetResponse{value: val, isSet: true}
}

func (v NullableKeywordTargetResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableKeywordTargetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
