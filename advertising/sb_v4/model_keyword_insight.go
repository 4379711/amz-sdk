package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the KeywordInsight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KeywordInsight{}

// KeywordInsight struct for KeywordInsight
type KeywordInsight struct {
	KeywordInsight SBInsightsKeywordInsight `json:"keywordInsight"`
}

type _KeywordInsight KeywordInsight

// NewKeywordInsight instantiates a new KeywordInsight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKeywordInsight(keywordInsight SBInsightsKeywordInsight) *KeywordInsight {
	this := KeywordInsight{}
	this.KeywordInsight = keywordInsight
	return &this
}

// NewKeywordInsightWithDefaults instantiates a new KeywordInsight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeywordInsightWithDefaults() *KeywordInsight {
	this := KeywordInsight{}
	return &this
}

// GetKeywordInsight returns the KeywordInsight field value
func (o *KeywordInsight) GetKeywordInsight() SBInsightsKeywordInsight {
	if o == nil {
		var ret SBInsightsKeywordInsight
		return ret
	}

	return o.KeywordInsight
}

// GetKeywordInsightOk returns a tuple with the KeywordInsight field value
// and a boolean to check if the value has been set.
func (o *KeywordInsight) GetKeywordInsightOk() (*SBInsightsKeywordInsight, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeywordInsight, true
}

// SetKeywordInsight sets field value
func (o *KeywordInsight) SetKeywordInsight(v SBInsightsKeywordInsight) {
	o.KeywordInsight = v
}

func (o KeywordInsight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["keywordInsight"] = o.KeywordInsight
	return toSerialize, nil
}

type NullableKeywordInsight struct {
	value *KeywordInsight
	isSet bool
}

func (v NullableKeywordInsight) Get() *KeywordInsight {
	return v.value
}

func (v *NullableKeywordInsight) Set(val *KeywordInsight) {
	v.value = val
	v.isSet = true
}

func (v NullableKeywordInsight) IsSet() bool {
	return v.isSet
}

func (v *NullableKeywordInsight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeywordInsight(val *KeywordInsight) *NullableKeywordInsight {
	return &NullableKeywordInsight{value: val, isSet: true}
}

func (v NullableKeywordInsight) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableKeywordInsight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
