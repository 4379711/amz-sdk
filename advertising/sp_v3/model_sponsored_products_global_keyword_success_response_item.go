package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalKeywordSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalKeywordSuccessResponseItem{}

// SponsoredProductsGlobalKeywordSuccessResponseItem struct for SponsoredProductsGlobalKeywordSuccessResponseItem
type SponsoredProductsGlobalKeywordSuccessResponseItem struct {
	// the keyword ID
	KeywordId *string `json:"keywordId,omitempty"`
	// the index of the keyword in the array from the request body
	Index   int32                           `json:"index"`
	Keyword *SponsoredProductsGlobalKeyword `json:"keyword,omitempty"`
}

type _SponsoredProductsGlobalKeywordSuccessResponseItem SponsoredProductsGlobalKeywordSuccessResponseItem

// NewSponsoredProductsGlobalKeywordSuccessResponseItem instantiates a new SponsoredProductsGlobalKeywordSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalKeywordSuccessResponseItem(index int32) *SponsoredProductsGlobalKeywordSuccessResponseItem {
	this := SponsoredProductsGlobalKeywordSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsGlobalKeywordSuccessResponseItemWithDefaults instantiates a new SponsoredProductsGlobalKeywordSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalKeywordSuccessResponseItemWithDefaults() *SponsoredProductsGlobalKeywordSuccessResponseItem {
	this := SponsoredProductsGlobalKeywordSuccessResponseItem{}
	return &this
}

// GetKeywordId returns the KeywordId field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalKeywordSuccessResponseItem) GetKeywordId() string {
	if o == nil || IsNil(o.KeywordId) {
		var ret string
		return ret
	}
	return *o.KeywordId
}

// GetKeywordIdOk returns a tuple with the KeywordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeywordSuccessResponseItem) GetKeywordIdOk() (*string, bool) {
	if o == nil || IsNil(o.KeywordId) {
		return nil, false
	}
	return o.KeywordId, true
}

// HasKeywordId returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalKeywordSuccessResponseItem) HasKeywordId() bool {
	if o != nil && !IsNil(o.KeywordId) {
		return true
	}

	return false
}

// SetKeywordId gets a reference to the given string and assigns it to the KeywordId field.
func (o *SponsoredProductsGlobalKeywordSuccessResponseItem) SetKeywordId(v string) {
	o.KeywordId = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsGlobalKeywordSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeywordSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsGlobalKeywordSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalKeywordSuccessResponseItem) GetKeyword() SponsoredProductsGlobalKeyword {
	if o == nil || IsNil(o.Keyword) {
		var ret SponsoredProductsGlobalKeyword
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalKeywordSuccessResponseItem) GetKeywordOk() (*SponsoredProductsGlobalKeyword, bool) {
	if o == nil || IsNil(o.Keyword) {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalKeywordSuccessResponseItem) HasKeyword() bool {
	if o != nil && !IsNil(o.Keyword) {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given SponsoredProductsGlobalKeyword and assigns it to the Keyword field.
func (o *SponsoredProductsGlobalKeywordSuccessResponseItem) SetKeyword(v SponsoredProductsGlobalKeyword) {
	o.Keyword = &v
}

func (o SponsoredProductsGlobalKeywordSuccessResponseItem) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsGlobalKeywordSuccessResponseItem struct {
	value *SponsoredProductsGlobalKeywordSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsGlobalKeywordSuccessResponseItem) Get() *SponsoredProductsGlobalKeywordSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsGlobalKeywordSuccessResponseItem) Set(val *SponsoredProductsGlobalKeywordSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalKeywordSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalKeywordSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalKeywordSuccessResponseItem(val *SponsoredProductsGlobalKeywordSuccessResponseItem) *NullableSponsoredProductsGlobalKeywordSuccessResponseItem {
	return &NullableSponsoredProductsGlobalKeywordSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalKeywordSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalKeywordSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
