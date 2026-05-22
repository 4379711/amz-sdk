package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsDraftProductAd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsDraftProductAd{}

// SponsoredProductsDraftProductAd struct for SponsoredProductsDraftProductAd
type SponsoredProductsDraftProductAd struct {
	GlobalStoreSetting *SponsoredProductsGlobalStoreSetting `json:"globalStoreSetting,omitempty"`
	// The product ad identifier.
	AdId string `json:"adId"`
	// The campaign identifier.
	CampaignId string `json:"campaignId"`
	// The custom text that is associated with this ad. Defined for custom text ads only.
	CustomText *string `json:"customText,omitempty"`
	// The ASIN associated with the product. Defined for vendors only.
	Asin *string `json:"asin,omitempty"`
	// The SKU associated with the product. Defined for seller accounts only.
	Sku *string `json:"sku,omitempty"`
	// The ad group identifier.
	AdGroupId    string                                       `json:"adGroupId"`
	ExtendedData *SponsoredProductsDraftProductAdExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsDraftProductAd SponsoredProductsDraftProductAd

// NewSponsoredProductsDraftProductAd instantiates a new SponsoredProductsDraftProductAd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsDraftProductAd(adId string, campaignId string, adGroupId string) *SponsoredProductsDraftProductAd {
	this := SponsoredProductsDraftProductAd{}
	this.AdId = adId
	this.CampaignId = campaignId
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsDraftProductAdWithDefaults instantiates a new SponsoredProductsDraftProductAd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsDraftProductAdWithDefaults() *SponsoredProductsDraftProductAd {
	this := SponsoredProductsDraftProductAd{}
	return &this
}

// GetGlobalStoreSetting returns the GlobalStoreSetting field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAd) GetGlobalStoreSetting() SponsoredProductsGlobalStoreSetting {
	if o == nil || IsNil(o.GlobalStoreSetting) {
		var ret SponsoredProductsGlobalStoreSetting
		return ret
	}
	return *o.GlobalStoreSetting
}

// GetGlobalStoreSettingOk returns a tuple with the GlobalStoreSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAd) GetGlobalStoreSettingOk() (*SponsoredProductsGlobalStoreSetting, bool) {
	if o == nil || IsNil(o.GlobalStoreSetting) {
		return nil, false
	}
	return o.GlobalStoreSetting, true
}

// HasGlobalStoreSetting returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAd) HasGlobalStoreSetting() bool {
	if o != nil && !IsNil(o.GlobalStoreSetting) {
		return true
	}

	return false
}

// SetGlobalStoreSetting gets a reference to the given SponsoredProductsGlobalStoreSetting and assigns it to the GlobalStoreSetting field.
func (o *SponsoredProductsDraftProductAd) SetGlobalStoreSetting(v SponsoredProductsGlobalStoreSetting) {
	o.GlobalStoreSetting = &v
}

// GetAdId returns the AdId field value
func (o *SponsoredProductsDraftProductAd) GetAdId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAd) GetAdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdId, true
}

// SetAdId sets field value
func (o *SponsoredProductsDraftProductAd) SetAdId(v string) {
	o.AdId = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsDraftProductAd) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAd) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsDraftProductAd) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetCustomText returns the CustomText field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAd) GetCustomText() string {
	if o == nil || IsNil(o.CustomText) {
		var ret string
		return ret
	}
	return *o.CustomText
}

// GetCustomTextOk returns a tuple with the CustomText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAd) GetCustomTextOk() (*string, bool) {
	if o == nil || IsNil(o.CustomText) {
		return nil, false
	}
	return o.CustomText, true
}

// HasCustomText returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAd) HasCustomText() bool {
	if o != nil && !IsNil(o.CustomText) {
		return true
	}

	return false
}

// SetCustomText gets a reference to the given string and assigns it to the CustomText field.
func (o *SponsoredProductsDraftProductAd) SetCustomText(v string) {
	o.CustomText = &v
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAd) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAd) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAd) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *SponsoredProductsDraftProductAd) SetAsin(v string) {
	o.Asin = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAd) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAd) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAd) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *SponsoredProductsDraftProductAd) SetSku(v string) {
	o.Sku = &v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsDraftProductAd) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAd) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsDraftProductAd) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsDraftProductAd) GetExtendedData() SponsoredProductsDraftProductAdExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsDraftProductAdExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsDraftProductAd) GetExtendedDataOk() (*SponsoredProductsDraftProductAdExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsDraftProductAd) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsDraftProductAdExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsDraftProductAd) SetExtendedData(v SponsoredProductsDraftProductAdExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsDraftProductAd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GlobalStoreSetting) {
		toSerialize["globalStoreSetting"] = o.GlobalStoreSetting
	}
	toSerialize["adId"] = o.AdId
	toSerialize["campaignId"] = o.CampaignId
	if !IsNil(o.CustomText) {
		toSerialize["customText"] = o.CustomText
	}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	toSerialize["adGroupId"] = o.AdGroupId
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsDraftProductAd struct {
	value *SponsoredProductsDraftProductAd
	isSet bool
}

func (v NullableSponsoredProductsDraftProductAd) Get() *SponsoredProductsDraftProductAd {
	return v.value
}

func (v *NullableSponsoredProductsDraftProductAd) Set(val *SponsoredProductsDraftProductAd) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsDraftProductAd) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsDraftProductAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsDraftProductAd(val *SponsoredProductsDraftProductAd) *NullableSponsoredProductsDraftProductAd {
	return &NullableSponsoredProductsDraftProductAd{value: val, isSet: true}
}

func (v NullableSponsoredProductsDraftProductAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsDraftProductAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
