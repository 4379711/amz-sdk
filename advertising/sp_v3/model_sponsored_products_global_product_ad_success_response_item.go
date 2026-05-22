package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalProductAdSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalProductAdSuccessResponseItem{}

// SponsoredProductsGlobalProductAdSuccessResponseItem struct for SponsoredProductsGlobalProductAdSuccessResponseItem
type SponsoredProductsGlobalProductAdSuccessResponseItem struct {
	// the ProductAd ID
	AdId *string `json:"adId,omitempty"`
	// The index in the original list from the request.
	Index     int32                             `json:"index"`
	ProductAd *SponsoredProductsGlobalProductAd `json:"productAd,omitempty"`
}

type _SponsoredProductsGlobalProductAdSuccessResponseItem SponsoredProductsGlobalProductAdSuccessResponseItem

// NewSponsoredProductsGlobalProductAdSuccessResponseItem instantiates a new SponsoredProductsGlobalProductAdSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalProductAdSuccessResponseItem(index int32) *SponsoredProductsGlobalProductAdSuccessResponseItem {
	this := SponsoredProductsGlobalProductAdSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsGlobalProductAdSuccessResponseItemWithDefaults instantiates a new SponsoredProductsGlobalProductAdSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalProductAdSuccessResponseItemWithDefaults() *SponsoredProductsGlobalProductAdSuccessResponseItem {
	this := SponsoredProductsGlobalProductAdSuccessResponseItem{}
	return &this
}

// GetAdId returns the AdId field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalProductAdSuccessResponseItem) GetAdId() string {
	if o == nil || IsNil(o.AdId) {
		var ret string
		return ret
	}
	return *o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalProductAdSuccessResponseItem) GetAdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdId) {
		return nil, false
	}
	return o.AdId, true
}

// HasAdId returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalProductAdSuccessResponseItem) HasAdId() bool {
	if o != nil && !IsNil(o.AdId) {
		return true
	}

	return false
}

// SetAdId gets a reference to the given string and assigns it to the AdId field.
func (o *SponsoredProductsGlobalProductAdSuccessResponseItem) SetAdId(v string) {
	o.AdId = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsGlobalProductAdSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalProductAdSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsGlobalProductAdSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetProductAd returns the ProductAd field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalProductAdSuccessResponseItem) GetProductAd() SponsoredProductsGlobalProductAd {
	if o == nil || IsNil(o.ProductAd) {
		var ret SponsoredProductsGlobalProductAd
		return ret
	}
	return *o.ProductAd
}

// GetProductAdOk returns a tuple with the ProductAd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalProductAdSuccessResponseItem) GetProductAdOk() (*SponsoredProductsGlobalProductAd, bool) {
	if o == nil || IsNil(o.ProductAd) {
		return nil, false
	}
	return o.ProductAd, true
}

// HasProductAd returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalProductAdSuccessResponseItem) HasProductAd() bool {
	if o != nil && !IsNil(o.ProductAd) {
		return true
	}

	return false
}

// SetProductAd gets a reference to the given SponsoredProductsGlobalProductAd and assigns it to the ProductAd field.
func (o *SponsoredProductsGlobalProductAdSuccessResponseItem) SetProductAd(v SponsoredProductsGlobalProductAd) {
	o.ProductAd = &v
}

func (o SponsoredProductsGlobalProductAdSuccessResponseItem) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsGlobalProductAdSuccessResponseItem struct {
	value *SponsoredProductsGlobalProductAdSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsGlobalProductAdSuccessResponseItem) Get() *SponsoredProductsGlobalProductAdSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsGlobalProductAdSuccessResponseItem) Set(val *SponsoredProductsGlobalProductAdSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalProductAdSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalProductAdSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalProductAdSuccessResponseItem(val *SponsoredProductsGlobalProductAdSuccessResponseItem) *NullableSponsoredProductsGlobalProductAdSuccessResponseItem {
	return &NullableSponsoredProductsGlobalProductAdSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalProductAdSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalProductAdSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
