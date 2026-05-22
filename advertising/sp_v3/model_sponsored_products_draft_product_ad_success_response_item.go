package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftProductAdSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftProductAdSuccessResponseItem{}

// SponsoredProductsDraftProductAdSuccessResponseItem struct for SponsoredProductsDraftProductAdSuccessResponseItem
type SponsoredProductsDraftProductAdSuccessResponseItem struct {
	// the DraftProductAd ID
	AdId *string `json:"adId,omitempty"`
	// The index in the original list from the request.
	Index     int32                            `json:"index"`
	ProductAd *SponsoredProductsDraftProductAd `json:"productAd,omitempty"`
}

type _SponsoredProductsDraftProductAdSuccessResponseItem SponsoredProductsDraftProductAdSuccessResponseItem

// NewSponsoredProductsDraftProductAdSuccessResponseItem instantiates a new SponsoredProductsDraftProductAdSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftProductAdSuccessResponseItem(index int32) *SponsoredProductsDraftProductAdSuccessResponseItem {
	this := SponsoredProductsDraftProductAdSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsDraftProductAdSuccessResponseItemWithDefaults instantiates a new SponsoredProductsDraftProductAdSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftProductAdSuccessResponseItemWithDefaults() *SponsoredProductsDraftProductAdSuccessResponseItem {
	this := SponsoredProductsDraftProductAdSuccessResponseItem{}
	return &this
}

// GetAdId returns the AdId field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdSuccessResponseItem) GetAdId() string {
	if o == nil || IsNil(o.AdId) {
		var ret string
		return ret
	}
	return *o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdSuccessResponseItem) GetAdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdId) {
		return nil, false
	}
	return o.AdId, true
}

// HasAdId returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdSuccessResponseItem) HasAdId() bool {
	if o != nil && !IsNil(o.AdId) {
		return true
	}

	return false
}

// SetAdId gets a reference to the given string and assigns it to the AdId field.
func (o *SponsoredProductsDraftProductAdSuccessResponseItem) SetAdId(v string) {
	o.AdId = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsDraftProductAdSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsDraftProductAdSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetProductAd returns the ProductAd field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAdSuccessResponseItem) GetProductAd() SponsoredProductsDraftProductAd {
	if o == nil || IsNil(o.ProductAd) {
		var ret SponsoredProductsDraftProductAd
		return ret
	}
	return *o.ProductAd
}

// GetProductAdOk returns a tuple with the ProductAd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAdSuccessResponseItem) GetProductAdOk() (*SponsoredProductsDraftProductAd, bool) {
	if o == nil || IsNil(o.ProductAd) {
		return nil, false
	}
	return o.ProductAd, true
}

// HasProductAd returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAdSuccessResponseItem) HasProductAd() bool {
	if o != nil && !IsNil(o.ProductAd) {
		return true
	}

	return false
}

// SetProductAd gets a reference to the given SponsoredProductsDraftProductAd and assigns it to the ProductAd field.
func (o *SponsoredProductsDraftProductAdSuccessResponseItem) SetProductAd(v SponsoredProductsDraftProductAd) {
	o.ProductAd = &v
}

func (o SponsoredProductsDraftProductAdSuccessResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdId) {
		toSerialize["adId"] = o.AdId
	}
	toSerialize["index"] = o.Index
	if !IsNil(o.ProductAd) {
		toSerialize["productAd"] = o.ProductAd
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftProductAdSuccessResponseItem struct {
	value *SponsoredProductsDraftProductAdSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsDraftProductAdSuccessResponseItem) Get() *SponsoredProductsDraftProductAdSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsDraftProductAdSuccessResponseItem) Set(val *SponsoredProductsDraftProductAdSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftProductAdSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftProductAdSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftProductAdSuccessResponseItem(val *SponsoredProductsDraftProductAdSuccessResponseItem) *NullableSponsoredProductsDraftProductAdSuccessResponseItem {
	return &NullableSponsoredProductsDraftProductAdSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftProductAdSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftProductAdSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
