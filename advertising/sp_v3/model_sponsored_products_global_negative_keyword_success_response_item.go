package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalNegativeKeywordSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalNegativeKeywordSuccessResponseItem{}

// SponsoredProductsGlobalNegativeKeywordSuccessResponseItem struct for SponsoredProductsGlobalNegativeKeywordSuccessResponseItem
type SponsoredProductsGlobalNegativeKeywordSuccessResponseItem struct {
	// the index of the negativeKeyword in the array from the request body
	Index           int32                                   `json:"index"`
	NegativeKeyword *SponsoredProductsGlobalNegativeKeyword `json:"negativeKeyword,omitempty"`
	// the negativeKeyword ID
	NegativeKeywordId *string `json:"negativeKeywordId,omitempty"`
}

type _SponsoredProductsGlobalNegativeKeywordSuccessResponseItem SponsoredProductsGlobalNegativeKeywordSuccessResponseItem

// NewSponsoredProductsGlobalNegativeKeywordSuccessResponseItem instantiates a new SponsoredProductsGlobalNegativeKeywordSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalNegativeKeywordSuccessResponseItem(index int32) *SponsoredProductsGlobalNegativeKeywordSuccessResponseItem {
	this := SponsoredProductsGlobalNegativeKeywordSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsGlobalNegativeKeywordSuccessResponseItemWithDefaults instantiates a new SponsoredProductsGlobalNegativeKeywordSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalNegativeKeywordSuccessResponseItemWithDefaults() *SponsoredProductsGlobalNegativeKeywordSuccessResponseItem {
	this := SponsoredProductsGlobalNegativeKeywordSuccessResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsGlobalNegativeKeywordSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeKeywordSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsGlobalNegativeKeywordSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetNegativeKeyword returns the NegativeKeyword field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalNegativeKeywordSuccessResponseItem) GetNegativeKeyword() SponsoredProductsGlobalNegativeKeyword {
	if o == nil || IsNil(o.NegativeKeyword) {
		var ret SponsoredProductsGlobalNegativeKeyword
		return ret
	}
	return *o.NegativeKeyword
}

// GetNegativeKeywordOk returns a tuple with the NegativeKeyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeKeywordSuccessResponseItem) GetNegativeKeywordOk() (*SponsoredProductsGlobalNegativeKeyword, bool) {
	if o == nil || IsNil(o.NegativeKeyword) {
		return nil, false
	}
	return o.NegativeKeyword, true
}

// HasNegativeKeyword returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalNegativeKeywordSuccessResponseItem) HasNegativeKeyword() bool {
	if o != nil && !IsNil(o.NegativeKeyword) {
		return true
	}

	return false
}

// SetNegativeKeyword gets a reference to the given SponsoredProductsGlobalNegativeKeyword and assigns it to the NegativeKeyword field.
func (o *SponsoredProductsGlobalNegativeKeywordSuccessResponseItem) SetNegativeKeyword(v SponsoredProductsGlobalNegativeKeyword) {
	o.NegativeKeyword = &v
}

// GetNegativeKeywordId returns the NegativeKeywordId field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalNegativeKeywordSuccessResponseItem) GetNegativeKeywordId() string {
	if o == nil || IsNil(o.NegativeKeywordId) {
		var ret string
		return ret
	}
	return *o.NegativeKeywordId
}

// GetNegativeKeywordIdOk returns a tuple with the NegativeKeywordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalNegativeKeywordSuccessResponseItem) GetNegativeKeywordIdOk() (*string, bool) {
	if o == nil || IsNil(o.NegativeKeywordId) {
		return nil, false
	}
	return o.NegativeKeywordId, true
}

// HasNegativeKeywordId returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalNegativeKeywordSuccessResponseItem) HasNegativeKeywordId() bool {
	if o != nil && !IsNil(o.NegativeKeywordId) {
		return true
	}

	return false
}

// SetNegativeKeywordId gets a reference to the given string and assigns it to the NegativeKeywordId field.
func (o *SponsoredProductsGlobalNegativeKeywordSuccessResponseItem) SetNegativeKeywordId(v string) {
	o.NegativeKeywordId = &v
}

func (o SponsoredProductsGlobalNegativeKeywordSuccessResponseItem) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsGlobalNegativeKeywordSuccessResponseItem struct {
	value *SponsoredProductsGlobalNegativeKeywordSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsGlobalNegativeKeywordSuccessResponseItem) Get() *SponsoredProductsGlobalNegativeKeywordSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsGlobalNegativeKeywordSuccessResponseItem) Set(val *SponsoredProductsGlobalNegativeKeywordSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalNegativeKeywordSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalNegativeKeywordSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalNegativeKeywordSuccessResponseItem(val *SponsoredProductsGlobalNegativeKeywordSuccessResponseItem) *NullableSponsoredProductsGlobalNegativeKeywordSuccessResponseItem {
	return &NullableSponsoredProductsGlobalNegativeKeywordSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalNegativeKeywordSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalNegativeKeywordSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
