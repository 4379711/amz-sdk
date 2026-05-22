package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the ThemedBid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ThemedBid{}

// ThemedBid struct for ThemedBid
type ThemedBid struct {
	SuggestedBid *BidValues `json:"suggestedBid,omitempty"`
	// Keyword match type. The default value will be BROAD.
	MatchType *string `json:"matchType,omitempty"`
	// The keyword target rank.
	Rank *float32 `json:"rank,omitempty"`
	// The theme of the bid recommendation. The default theme is CONVERSION_OPPORTUNITIES.
	Theme *string `json:"theme,omitempty"`
	// The bid value for the keyword, in minor currency units (example: cents). The default value will be the suggested bid.
	Bid *float32 `json:"bid,omitempty"`
}

// NewThemedBid instantiates a new ThemedBid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewThemedBid() *ThemedBid {
	this := ThemedBid{}
	return &this
}

// NewThemedBidWithDefaults instantiates a new ThemedBid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewThemedBidWithDefaults() *ThemedBid {
	this := ThemedBid{}
	return &this
}

// GetSuggestedBid returns the SuggestedBid field value if set, zero value otherwise.
func (o *ThemedBid) GetSuggestedBid() BidValues {
	if o == nil || IsNil(o.SuggestedBid) {
		var ret BidValues
		return ret
	}
	return *o.SuggestedBid
}

// GetSuggestedBidOk returns a tuple with the SuggestedBid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemedBid) GetSuggestedBidOk() (*BidValues, bool) {
	if o == nil || IsNil(o.SuggestedBid) {
		return nil, false
	}
	return o.SuggestedBid, true
}

// HasSuggestedBid returns a boolean if a field has been set.
func (o *ThemedBid) HasSuggestedBid() bool {
	if o != nil && !IsNil(o.SuggestedBid) {
		return true
	}

	return false
}

// SetSuggestedBid gets a reference to the given BidValues and assigns it to the SuggestedBid field.
func (o *ThemedBid) SetSuggestedBid(v BidValues) {
	o.SuggestedBid = &v
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *ThemedBid) GetMatchType() string {
	if o == nil || IsNil(o.MatchType) {
		var ret string
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemedBid) GetMatchTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MatchType) {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *ThemedBid) HasMatchType() bool {
	if o != nil && !IsNil(o.MatchType) {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given string and assigns it to the MatchType field.
func (o *ThemedBid) SetMatchType(v string) {
	o.MatchType = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *ThemedBid) GetRank() float32 {
	if o == nil || IsNil(o.Rank) {
		var ret float32
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemedBid) GetRankOk() (*float32, bool) {
	if o == nil || IsNil(o.Rank) {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *ThemedBid) HasRank() bool {
	if o != nil && !IsNil(o.Rank) {
		return true
	}

	return false
}

// SetRank gets a reference to the given float32 and assigns it to the Rank field.
func (o *ThemedBid) SetRank(v float32) {
	o.Rank = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *ThemedBid) GetTheme() string {
	if o == nil || IsNil(o.Theme) {
		var ret string
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemedBid) GetThemeOk() (*string, bool) {
	if o == nil || IsNil(o.Theme) {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *ThemedBid) HasTheme() bool {
	if o != nil && !IsNil(o.Theme) {
		return true
	}

	return false
}

// SetTheme gets a reference to the given string and assigns it to the Theme field.
func (o *ThemedBid) SetTheme(v string) {
	o.Theme = &v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *ThemedBid) GetBid() float32 {
	if o == nil || IsNil(o.Bid) {
		var ret float32
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ThemedBid) GetBidOk() (*float32, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *ThemedBid) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given float32 and assigns it to the Bid field.
func (o *ThemedBid) SetBid(v float32) {
	o.Bid = &v
}

func (o ThemedBid) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Theme) {
		toSerialize["theme"] = o.Theme
	}
	if !IsNil(o.Bid) {
		toSerialize["bid"] = o.Bid
	}
	return toSerialize, nil
}

type NullableThemedBid struct {
	value *ThemedBid
	isSet bool
}

func (v NullableThemedBid) Get() *ThemedBid {
	return v.value
}

func (v *NullableThemedBid) Set(val *ThemedBid) {
	v.value = val
	v.isSet = true
}

func (v NullableThemedBid) IsSet() bool {
	return v.isSet
}

func (v *NullableThemedBid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableThemedBid(val *ThemedBid) *NullableThemedBid {
	return &NullableThemedBid{value: val, isSet: true}
}

func (v NullableThemedBid) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableThemedBid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
