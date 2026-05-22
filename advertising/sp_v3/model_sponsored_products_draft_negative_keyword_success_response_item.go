package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftNegativeKeywordSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftNegativeKeywordSuccessResponseItem{}

// SponsoredProductsDraftNegativeKeywordSuccessResponseItem struct for SponsoredProductsDraftNegativeKeywordSuccessResponseItem
type SponsoredProductsDraftNegativeKeywordSuccessResponseItem struct {
	// the index of the negativeKeyword in the array from the request body
	Index           int32                                  `json:"index"`
	NegativeKeyword *SponsoredProductsDraftNegativeKeyword `json:"negativeKeyword,omitempty"`
	// the negativeKeyword ID
	NegativeKeywordId *string `json:"negativeKeywordId,omitempty"`
}

type _SponsoredProductsDraftNegativeKeywordSuccessResponseItem SponsoredProductsDraftNegativeKeywordSuccessResponseItem

// NewSponsoredProductsDraftNegativeKeywordSuccessResponseItem instantiates a new SponsoredProductsDraftNegativeKeywordSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftNegativeKeywordSuccessResponseItem(index int32) *SponsoredProductsDraftNegativeKeywordSuccessResponseItem {
	this := SponsoredProductsDraftNegativeKeywordSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsDraftNegativeKeywordSuccessResponseItemWithDefaults instantiates a new SponsoredProductsDraftNegativeKeywordSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftNegativeKeywordSuccessResponseItemWithDefaults() *SponsoredProductsDraftNegativeKeywordSuccessResponseItem {
	this := SponsoredProductsDraftNegativeKeywordSuccessResponseItem{}
	return &this
}

// GetIndex returns the Index field value
func (o *SponsoredProductsDraftNegativeKeywordSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeKeywordSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsDraftNegativeKeywordSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetNegativeKeyword returns the NegativeKeyword field value if set, zero value otherwise.
func (o *SponsoredProductsDraftNegativeKeywordSuccessResponseItem) GetNegativeKeyword() SponsoredProductsDraftNegativeKeyword {
	if o == nil || IsNil(o.NegativeKeyword) {
		var ret SponsoredProductsDraftNegativeKeyword
		return ret
	}
	return *o.NegativeKeyword
}

// GetNegativeKeywordOk returns a tuple with the NegativeKeyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeKeywordSuccessResponseItem) GetNegativeKeywordOk() (*SponsoredProductsDraftNegativeKeyword, bool) {
	if o == nil || IsNil(o.NegativeKeyword) {
		return nil, false
	}
	return o.NegativeKeyword, true
}

// HasNegativeKeyword returns a boolean if a field has been set.
func (o *SponsoredProductsDraftNegativeKeywordSuccessResponseItem) HasNegativeKeyword() bool {
	if o != nil && !IsNil(o.NegativeKeyword) {
		return true
	}

	return false
}

// SetNegativeKeyword gets a reference to the given SponsoredProductsDraftNegativeKeyword and assigns it to the NegativeKeyword field.
func (o *SponsoredProductsDraftNegativeKeywordSuccessResponseItem) SetNegativeKeyword(v SponsoredProductsDraftNegativeKeyword) {
	o.NegativeKeyword = &v
}

// GetNegativeKeywordId returns the NegativeKeywordId field value if set, zero value otherwise.
func (o *SponsoredProductsDraftNegativeKeywordSuccessResponseItem) GetNegativeKeywordId() string {
	if o == nil || IsNil(o.NegativeKeywordId) {
		var ret string
		return ret
	}
	return *o.NegativeKeywordId
}

// GetNegativeKeywordIdOk returns a tuple with the NegativeKeywordId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftNegativeKeywordSuccessResponseItem) GetNegativeKeywordIdOk() (*string, bool) {
	if o == nil || IsNil(o.NegativeKeywordId) {
		return nil, false
	}
	return o.NegativeKeywordId, true
}

// HasNegativeKeywordId returns a boolean if a field has been set.
func (o *SponsoredProductsDraftNegativeKeywordSuccessResponseItem) HasNegativeKeywordId() bool {
	if o != nil && !IsNil(o.NegativeKeywordId) {
		return true
	}

	return false
}

// SetNegativeKeywordId gets a reference to the given string and assigns it to the NegativeKeywordId field.
func (o *SponsoredProductsDraftNegativeKeywordSuccessResponseItem) SetNegativeKeywordId(v string) {
	o.NegativeKeywordId = &v
}

func (o SponsoredProductsDraftNegativeKeywordSuccessResponseItem) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsDraftNegativeKeywordSuccessResponseItem struct {
	value *SponsoredProductsDraftNegativeKeywordSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsDraftNegativeKeywordSuccessResponseItem) Get() *SponsoredProductsDraftNegativeKeywordSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsDraftNegativeKeywordSuccessResponseItem) Set(val *SponsoredProductsDraftNegativeKeywordSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftNegativeKeywordSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftNegativeKeywordSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftNegativeKeywordSuccessResponseItem(val *SponsoredProductsDraftNegativeKeywordSuccessResponseItem) *NullableSponsoredProductsDraftNegativeKeywordSuccessResponseItem {
	return &NullableSponsoredProductsDraftNegativeKeywordSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftNegativeKeywordSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftNegativeKeywordSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
