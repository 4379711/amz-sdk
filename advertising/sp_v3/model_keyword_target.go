package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the KeywordTarget type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeywordTarget{}

// KeywordTarget struct for KeywordTarget
type KeywordTarget struct {
	// Keyword match type. The default value will be BROAD.
	MatchType *string `json:"matchType,omitempty"`
	// The keyword value
	Keyword *string `json:"keyword,omitempty"`
	// The bid value for the keyword, in minor currency units (example: cents). The default value will be the suggested bid.
	Bid *float64 `json:"bid,omitempty"`
	// Flag that tells if keyword was selected by the user or was recommended by KRS
	UserSelectedKeyword *bool `json:"userSelectedKeyword,omitempty"`
}

// NewKeywordTarget instantiates a new KeywordTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeywordTarget() *KeywordTarget {
	this := KeywordTarget{}
	return &this
}

// NewKeywordTargetWithDefaults instantiates a new KeywordTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeywordTargetWithDefaults() *KeywordTarget {
	this := KeywordTarget{}
	return &this
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *KeywordTarget) GetMatchType() string {
	if o == nil || IsNil(o.MatchType) {
		var ret string
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTarget) GetMatchTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MatchType) {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *KeywordTarget) HasMatchType() bool {
	if o != nil && !IsNil(o.MatchType) {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given string and assigns it to the MatchType field.
func (o *KeywordTarget) SetMatchType(v string) {
	o.MatchType = &v
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *KeywordTarget) GetKeyword() string {
	if o == nil || IsNil(o.Keyword) {
		var ret string
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTarget) GetKeywordOk() (*string, bool) {
	if o == nil || IsNil(o.Keyword) {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *KeywordTarget) HasKeyword() bool {
	if o != nil && !IsNil(o.Keyword) {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given string and assigns it to the Keyword field.
func (o *KeywordTarget) SetKeyword(v string) {
	o.Keyword = &v
}

// GetBid returns the Bid field value if set, zero value otherwise.
func (o *KeywordTarget) GetBid() float64 {
	if o == nil || IsNil(o.Bid) {
		var ret float64
		return ret
	}
	return *o.Bid
}

// GetBidOk returns a tuple with the Bid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTarget) GetBidOk() (*float64, bool) {
	if o == nil || IsNil(o.Bid) {
		return nil, false
	}
	return o.Bid, true
}

// HasBid returns a boolean if a field has been set.
func (o *KeywordTarget) HasBid() bool {
	if o != nil && !IsNil(o.Bid) {
		return true
	}

	return false
}

// SetBid gets a reference to the given float64 and assigns it to the Bid field.
func (o *KeywordTarget) SetBid(v float64) {
	o.Bid = &v
}

// GetUserSelectedKeyword returns the UserSelectedKeyword field value if set, zero value otherwise.
func (o *KeywordTarget) GetUserSelectedKeyword() bool {
	if o == nil || IsNil(o.UserSelectedKeyword) {
		var ret bool
		return ret
	}
	return *o.UserSelectedKeyword
}

// GetUserSelectedKeywordOk returns a tuple with the UserSelectedKeyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KeywordTarget) GetUserSelectedKeywordOk() (*bool, bool) {
	if o == nil || IsNil(o.UserSelectedKeyword) {
		return nil, false
	}
	return o.UserSelectedKeyword, true
}

// HasUserSelectedKeyword returns a boolean if a field has been set.
func (o *KeywordTarget) HasUserSelectedKeyword() bool {
	if o != nil && !IsNil(o.UserSelectedKeyword) {
		return true
	}

	return false
}

// SetUserSelectedKeyword gets a reference to the given bool and assigns it to the UserSelectedKeyword field.
func (o *KeywordTarget) SetUserSelectedKeyword(v bool) {
	o.UserSelectedKeyword = &v
}

func (o KeywordTarget) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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

type NullableKeywordTarget struct {
	value *KeywordTarget
	isSet bool
}

func (v NullableKeywordTarget) Get() *KeywordTarget {
	return v.value
}

func (v *NullableKeywordTarget) Set(val *KeywordTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableKeywordTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableKeywordTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeywordTarget(val *KeywordTarget) *NullableKeywordTarget {
	return &NullableKeywordTarget{value: val, isSet: true}
}

func (v NullableKeywordTarget) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableKeywordTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
