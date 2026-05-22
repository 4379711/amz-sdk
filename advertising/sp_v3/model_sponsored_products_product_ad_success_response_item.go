package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsProductAdSuccessResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsProductAdSuccessResponseItem{}

// SponsoredProductsProductAdSuccessResponseItem struct for SponsoredProductsProductAdSuccessResponseItem
type SponsoredProductsProductAdSuccessResponseItem struct {
	// the ProductAd ID
	AdId *string `json:"adId,omitempty"`
	// The index in the original list from the request.
	Index     int32                       `json:"index"`
	ProductAd *SponsoredProductsProductAd `json:"productAd,omitempty"`
}

type _SponsoredProductsProductAdSuccessResponseItem SponsoredProductsProductAdSuccessResponseItem

// NewSponsoredProductsProductAdSuccessResponseItem instantiates a new SponsoredProductsProductAdSuccessResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsProductAdSuccessResponseItem(index int32) *SponsoredProductsProductAdSuccessResponseItem {
	this := SponsoredProductsProductAdSuccessResponseItem{}
	this.Index = index
	return &this
}

// NewSponsoredProductsProductAdSuccessResponseItemWithDefaults instantiates a new SponsoredProductsProductAdSuccessResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsProductAdSuccessResponseItemWithDefaults() *SponsoredProductsProductAdSuccessResponseItem {
	this := SponsoredProductsProductAdSuccessResponseItem{}
	return &this
}

// GetAdId returns the AdId field value if set, zero value otherwise.
func (o *SponsoredProductsProductAdSuccessResponseItem) GetAdId() string {
	if o == nil || IsNil(o.AdId) {
		var ret string
		return ret
	}
	return *o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAdSuccessResponseItem) GetAdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdId) {
		return nil, false
	}
	return o.AdId, true
}

// HasAdId returns a boolean if a field has been set.
func (o *SponsoredProductsProductAdSuccessResponseItem) HasAdId() bool {
	if o != nil && !IsNil(o.AdId) {
		return true
	}

	return false
}

// SetAdId gets a reference to the given string and assigns it to the AdId field.
func (o *SponsoredProductsProductAdSuccessResponseItem) SetAdId(v string) {
	o.AdId = &v
}

// GetIndex returns the Index field value
func (o *SponsoredProductsProductAdSuccessResponseItem) GetIndex() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAdSuccessResponseItem) GetIndexOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *SponsoredProductsProductAdSuccessResponseItem) SetIndex(v int32) {
	o.Index = v
}

// GetProductAd returns the ProductAd field value if set, zero value otherwise.
func (o *SponsoredProductsProductAdSuccessResponseItem) GetProductAd() SponsoredProductsProductAd {
	if o == nil || IsNil(o.ProductAd) {
		var ret SponsoredProductsProductAd
		return ret
	}
	return *o.ProductAd
}

// GetProductAdOk returns a tuple with the ProductAd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAdSuccessResponseItem) GetProductAdOk() (*SponsoredProductsProductAd, bool) {
	if o == nil || IsNil(o.ProductAd) {
		return nil, false
	}
	return o.ProductAd, true
}

// HasProductAd returns a boolean if a field has been set.
func (o *SponsoredProductsProductAdSuccessResponseItem) HasProductAd() bool {
	if o != nil && !IsNil(o.ProductAd) {
		return true
	}

	return false
}

// SetProductAd gets a reference to the given SponsoredProductsProductAd and assigns it to the ProductAd field.
func (o *SponsoredProductsProductAdSuccessResponseItem) SetProductAd(v SponsoredProductsProductAd) {
	o.ProductAd = &v
}

func (o SponsoredProductsProductAdSuccessResponseItem) ToMap() (map[string]interface{}, error) {
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

type NullableSponsoredProductsProductAdSuccessResponseItem struct {
	value *SponsoredProductsProductAdSuccessResponseItem
	isSet bool
}

func (v NullableSponsoredProductsProductAdSuccessResponseItem) Get() *SponsoredProductsProductAdSuccessResponseItem {
	return v.value
}

func (v *NullableSponsoredProductsProductAdSuccessResponseItem) Set(val *SponsoredProductsProductAdSuccessResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsProductAdSuccessResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsProductAdSuccessResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsProductAdSuccessResponseItem(val *SponsoredProductsProductAdSuccessResponseItem) *NullableSponsoredProductsProductAdSuccessResponseItem {
	return &NullableSponsoredProductsProductAdSuccessResponseItem{value: val, isSet: true}
}

func (v NullableSponsoredProductsProductAdSuccessResponseItem) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsProductAdSuccessResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
