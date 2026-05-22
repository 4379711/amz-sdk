package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsKeywordSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsKeywordSuccessResponseItem{}

// SponsoredProductsKeywordSuccessResponseItem struct for SponsoredProductsKeywordSuccessResponseItem
type SponsoredProductsKeywordSuccessResponseItem struct {
	// the keyword ID
	KeywordId *string `json:"keywordId,omitempty"`
	// the index of the keyword in the array from the request body
	Index   int32                     `json:"index"`
	Keyword *SponsoredProductsKeyword `json:"keyword,omitempty"`
}

type _SponsoredProductsKeywordSuccessResponseItem SponsoredProductsKeywordSuccessResponseItem

// NewSponsoredProductsKeywordSuccessResponseItem instantiates a new SponsoredProductsKeywordSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsKeywordSuccessResponseItem(index int32) *SponsoredProductsKeywordSuccessResponseItem {
	this := SponsoredProductsKeywordSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsKeywordSuccessResponseItemWithDefaults instantiates a new SponsoredProductsKeywordSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsKeywordSuccessResponseItemWithDefaults() *SponsoredProductsKeywordSuccessResponseItem {
	this := SponsoredProductsKeywordSuccessResponseItem{}
	return &this
}

// GetKeywordId returns the KeywordId field value if set, zero value otherwise.
func (o *SponsoredProductsKeywordSuccessResponseItem) GetKeywordId() string {
	if o == nil || IsNil(o.KeywordId) {
		var ret string
		return ret
	}
	return *o.KeywordId
}

// GetKeywordIdOk returns a tuple with the KeywordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordSuccessResponseItem) GetKeywordIdOk() (*string, bool) {
	if o == nil || IsNil(o.KeywordId) {
		return nil, false
	}
	return o.KeywordId, true
}

// HasKeywordId returns a boolean if a field has been set.
func (o *SponsoredProductsKeywordSuccessResponseItem) HasKeywordId() bool {
	if o != nil && !IsNil(o.KeywordId) {
		return true
	}

	return false
}

// SetKeywordId gets a reference to the given string and assigns it to the KeywordId field.
func (o *SponsoredProductsKeywordSuccessResponseItem) SetKeywordId(v string) {
	o.KeywordId = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsKeywordSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsKeywordSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *SponsoredProductsKeywordSuccessResponseItem) GetKeyword() SponsoredProductsKeyword {
	if o == nil || IsNil(o.Keyword) {
		var ret SponsoredProductsKeyword
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsKeywordSuccessResponseItem) GetKeywordOk() (*SponsoredProductsKeyword, bool) {
	if o == nil || IsNil(o.Keyword) {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *SponsoredProductsKeywordSuccessResponseItem) HasKeyword() bool {
	if o != nil && !IsNil(o.Keyword) {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given SponsoredProductsKeyword and assigns it to the Keyword field.
func (o *SponsoredProductsKeywordSuccessResponseItem) SetKeyword(v SponsoredProductsKeyword) {
	o.Keyword = &v
}

func (o SponsoredProductsKeywordSuccessResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.KeywordId) {
		toSerialize["keywordId"] = o.KeywordId
	}
	toSerialize["index"] = o.Index
	if !IsNil(o.Keyword) {
		toSerialize["keyword"] = o.Keyword
	}
	return toSerialize, nil
}

type NullableSponsoredProductsKeywordSuccessResponseItem struct {
	value *SponsoredProductsKeywordSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsKeywordSuccessResponseItem) Get() *SponsoredProductsKeywordSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsKeywordSuccessResponseItem) Set(val *SponsoredProductsKeywordSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsKeywordSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsKeywordSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsKeywordSuccessResponseItem(val *SponsoredProductsKeywordSuccessResponseItem) *NullableSponsoredProductsKeywordSuccessResponseItem {
	return &NullableSponsoredProductsKeywordSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsKeywordSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsKeywordSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
