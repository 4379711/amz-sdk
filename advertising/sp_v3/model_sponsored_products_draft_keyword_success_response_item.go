package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsDraftKeywordSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftKeywordSuccessResponseItem{}

// SponsoredProductsDraftKeywordSuccessResponseItem struct for SponsoredProductsDraftKeywordSuccessResponseItem
type SponsoredProductsDraftKeywordSuccessResponseItem struct {
	// the draft keyword ID
	KeywordId *string `json:"keywordId,omitempty"`
	// the index of the draft keyword in the array from the request body
	Index   int32                          `json:"index"`
	Keyword *SponsoredProductsDraftKeyword `json:"keyword,omitempty"`
}

type _SponsoredProductsDraftKeywordSuccessResponseItem SponsoredProductsDraftKeywordSuccessResponseItem

// NewSponsoredProductsDraftKeywordSuccessResponseItem instantiates a new SponsoredProductsDraftKeywordSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftKeywordSuccessResponseItem(index int32) *SponsoredProductsDraftKeywordSuccessResponseItem {
	this := SponsoredProductsDraftKeywordSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsDraftKeywordSuccessResponseItemWithDefaults instantiates a new SponsoredProductsDraftKeywordSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftKeywordSuccessResponseItemWithDefaults() *SponsoredProductsDraftKeywordSuccessResponseItem {
	this := SponsoredProductsDraftKeywordSuccessResponseItem{}
	return &this
}

// GetKeywordId returns the KeywordId field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeywordSuccessResponseItem) GetKeywordId() string {
	if o == nil || IsNil(o.KeywordId) {
		var ret string
		return ret
	}
	return *o.KeywordId
}

// GetKeywordIdOk returns a tuple with the KeywordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordSuccessResponseItem) GetKeywordIdOk() (*string, bool) {
	if o == nil || IsNil(o.KeywordId) {
		return nil, false
	}
	return o.KeywordId, true
}

// HasKeywordId returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeywordSuccessResponseItem) HasKeywordId() bool {
	if o != nil && !IsNil(o.KeywordId) {
		return true
	}

	return false
}

// SetKeywordId gets a reference to the given string and assigns it to the KeywordId field.
func (o *SponsoredProductsDraftKeywordSuccessResponseItem) SetKeywordId(v string) {
	o.KeywordId = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsDraftKeywordSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsDraftKeywordSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *SponsoredProductsDraftKeywordSuccessResponseItem) GetKeyword() SponsoredProductsDraftKeyword {
	if o == nil || IsNil(o.Keyword) {
		var ret SponsoredProductsDraftKeyword
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftKeywordSuccessResponseItem) GetKeywordOk() (*SponsoredProductsDraftKeyword, bool) {
	if o == nil || IsNil(o.Keyword) {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *SponsoredProductsDraftKeywordSuccessResponseItem) HasKeyword() bool {
	if o != nil && !IsNil(o.Keyword) {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given SponsoredProductsDraftKeyword and assigns it to the Keyword field.
func (o *SponsoredProductsDraftKeywordSuccessResponseItem) SetKeyword(v SponsoredProductsDraftKeyword) {
	o.Keyword = &v
}

func (o SponsoredProductsDraftKeywordSuccessResponseItem) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsDraftKeywordSuccessResponseItem struct {
	value *SponsoredProductsDraftKeywordSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsDraftKeywordSuccessResponseItem) Get() *SponsoredProductsDraftKeywordSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsDraftKeywordSuccessResponseItem) Set(val *SponsoredProductsDraftKeywordSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftKeywordSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftKeywordSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftKeywordSuccessResponseItem(val *SponsoredProductsDraftKeywordSuccessResponseItem) *NullableSponsoredProductsDraftKeywordSuccessResponseItem {
	return &NullableSponsoredProductsDraftKeywordSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftKeywordSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftKeywordSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
