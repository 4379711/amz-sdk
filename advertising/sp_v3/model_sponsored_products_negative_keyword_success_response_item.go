package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsNegativeKeywordSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsNegativeKeywordSuccessResponseItem{}

// SponsoredProductsNegativeKeywordSuccessResponseItem struct for SponsoredProductsNegativeKeywordSuccessResponseItem
type SponsoredProductsNegativeKeywordSuccessResponseItem struct {
	// the index of the negativeKeyword in the array from the request body
	Index           int32                             `json:"index"`
	NegativeKeyword *SponsoredProductsNegativeKeyword `json:"negativeKeyword,omitempty"`
	// the negativeKeyword ID
	NegativeKeywordId *string `json:"negativeKeywordId,omitempty"`
}

type _SponsoredProductsNegativeKeywordSuccessResponseItem SponsoredProductsNegativeKeywordSuccessResponseItem

// NewSponsoredProductsNegativeKeywordSuccessResponseItem instantiates a new SponsoredProductsNegativeKeywordSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsNegativeKeywordSuccessResponseItem(index int32) *SponsoredProductsNegativeKeywordSuccessResponseItem {
	this := SponsoredProductsNegativeKeywordSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsNegativeKeywordSuccessResponseItemWithDefaults instantiates a new SponsoredProductsNegativeKeywordSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsNegativeKeywordSuccessResponseItemWithDefaults() *SponsoredProductsNegativeKeywordSuccessResponseItem {
	this := SponsoredProductsNegativeKeywordSuccessResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsNegativeKeywordSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsNegativeKeywordSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetNegativeKeyword returns the NegativeKeyword field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordSuccessResponseItem) GetNegativeKeyword() SponsoredProductsNegativeKeyword {
	if o == nil || IsNil(o.NegativeKeyword) {
		var ret SponsoredProductsNegativeKeyword
		return ret
	}
	return *o.NegativeKeyword
}

// GetNegativeKeywordOk returns a tuple with the NegativeKeyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordSuccessResponseItem) GetNegativeKeywordOk() (*SponsoredProductsNegativeKeyword, bool) {
	if o == nil || IsNil(o.NegativeKeyword) {
		return nil, false
	}
	return o.NegativeKeyword, true
}

// HasNegativeKeyword returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordSuccessResponseItem) HasNegativeKeyword() bool {
	if o != nil && !IsNil(o.NegativeKeyword) {
		return true
	}

	return false
}

// SetNegativeKeyword gets a reference to the given SponsoredProductsNegativeKeyword and assigns it to the NegativeKeyword field.
func (o *SponsoredProductsNegativeKeywordSuccessResponseItem) SetNegativeKeyword(v SponsoredProductsNegativeKeyword) {
	o.NegativeKeyword = &v
}

// GetNegativeKeywordId returns the NegativeKeywordId field value if set, zero value otherwise.
func (o *SponsoredProductsNegativeKeywordSuccessResponseItem) GetNegativeKeywordId() string {
	if o == nil || IsNil(o.NegativeKeywordId) {
		var ret string
		return ret
	}
	return *o.NegativeKeywordId
}

// GetNegativeKeywordIdOk returns a tuple with the NegativeKeywordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsNegativeKeywordSuccessResponseItem) GetNegativeKeywordIdOk() (*string, bool) {
	if o == nil || IsNil(o.NegativeKeywordId) {
		return nil, false
	}
	return o.NegativeKeywordId, true
}

// HasNegativeKeywordId returns a boolean if a field has been set.
func (o *SponsoredProductsNegativeKeywordSuccessResponseItem) HasNegativeKeywordId() bool {
	if o != nil && !IsNil(o.NegativeKeywordId) {
		return true
	}

	return false
}

// SetNegativeKeywordId gets a reference to the given string and assigns it to the NegativeKeywordId field.
func (o *SponsoredProductsNegativeKeywordSuccessResponseItem) SetNegativeKeywordId(v string) {
	o.NegativeKeywordId = &v
}

func (o SponsoredProductsNegativeKeywordSuccessResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["index"] = o.Index
	if !IsNil(o.NegativeKeyword) {
		toSerialize["negativeKeyword"] = o.NegativeKeyword
	}
	if !IsNil(o.NegativeKeywordId) {
		toSerialize["negativeKeywordId"] = o.NegativeKeywordId
	}
	return toSerialize, nil
}

type NullableSponsoredProductsNegativeKeywordSuccessResponseItem struct {
	value *SponsoredProductsNegativeKeywordSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsNegativeKeywordSuccessResponseItem) Get() *SponsoredProductsNegativeKeywordSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsNegativeKeywordSuccessResponseItem) Set(val *SponsoredProductsNegativeKeywordSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsNegativeKeywordSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsNegativeKeywordSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsNegativeKeywordSuccessResponseItem(val *SponsoredProductsNegativeKeywordSuccessResponseItem) *NullableSponsoredProductsNegativeKeywordSuccessResponseItem {
	return &NullableSponsoredProductsNegativeKeywordSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsNegativeKeywordSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsNegativeKeywordSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
