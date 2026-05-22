package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateDraftProductAd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateDraftProductAd{}

// SponsoredProductsCreateDraftProductAd struct for SponsoredProductsCreateDraftProductAd
type SponsoredProductsCreateDraftProductAd struct {
	GlobalStoreSetting *SponsoredProductsGlobalStoreSetting `json:"globalStoreSetting,omitempty"`
	// The campaign identifier.
	CampaignId string `json:"campaignId"`
	// The custom text to use for creating a custom text ad for the associated ASIN. Defined only for KDP Authors and Book Vendors in US marketplace.
	CustomText *string `json:"customText,omitempty"`
	// The ASIN associated with the product. Defined for vendors only.
	Asin *string `json:"asin,omitempty"`
	// The SKU associated with the product. Defined for seller accounts only.
	Sku *string `json:"sku,omitempty"`
	// The ad group identifier.
	AdGroupId string `json:"adGroupId"`
}

type _SponsoredProductsCreateDraftProductAd SponsoredProductsCreateDraftProductAd

// NewSponsoredProductsCreateDraftProductAd instantiates a new SponsoredProductsCreateDraftProductAd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateDraftProductAd(campaignId string, adGroupId string) *SponsoredProductsCreateDraftProductAd {
	this := SponsoredProductsCreateDraftProductAd{}
	this.CampaignId = campaignId
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsCreateDraftProductAdWithDefaults instantiates a new SponsoredProductsCreateDraftProductAd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateDraftProductAdWithDefaults() *SponsoredProductsCreateDraftProductAd {
	this := SponsoredProductsCreateDraftProductAd{}
	return &this
}

// GetGlobalStoreSetting returns the GlobalStoreSetting field value if set, zero value otherwise.
func (o *SponsoredProductsCreateDraftProductAd) GetGlobalStoreSetting() SponsoredProductsGlobalStoreSetting {
	if o == nil || IsNil(o.GlobalStoreSetting) {
		var ret SponsoredProductsGlobalStoreSetting
		return ret
	}
	return *o.GlobalStoreSetting
}

// GetGlobalStoreSettingOk returns a tuple with the GlobalStoreSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftProductAd) GetGlobalStoreSettingOk() (*SponsoredProductsGlobalStoreSetting, bool) {
	if o == nil || IsNil(o.GlobalStoreSetting) {
		return nil, false
	}
	return o.GlobalStoreSetting, true
}

// HasGlobalStoreSetting returns a boolean if a field has been set.
func (o *SponsoredProductsCreateDraftProductAd) HasGlobalStoreSetting() bool {
	if o != nil && !IsNil(o.GlobalStoreSetting) {
		return true
	}

	return false
}

// SetGlobalStoreSetting gets a reference to the given SponsoredProductsGlobalStoreSetting and assigns it to the GlobalStoreSetting field.
func (o *SponsoredProductsCreateDraftProductAd) SetGlobalStoreSetting(v SponsoredProductsGlobalStoreSetting) {
	o.GlobalStoreSetting = &v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateDraftProductAd) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftProductAd) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateDraftProductAd) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetCustomText returns the CustomText field value if set, zero value otherwise.
func (o *SponsoredProductsCreateDraftProductAd) GetCustomText() string {
	if o == nil || IsNil(o.CustomText) {
		var ret string
		return ret
	}
	return *o.CustomText
}

// GetCustomTextOk returns a tuple with the CustomText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftProductAd) GetCustomTextOk() (*string, bool) {
	if o == nil || IsNil(o.CustomText) {
		return nil, false
	}
	return o.CustomText, true
}

// HasCustomText returns a boolean if a field has been set.
func (o *SponsoredProductsCreateDraftProductAd) HasCustomText() bool {
	if o != nil && !IsNil(o.CustomText) {
		return true
	}

	return false
}

// SetCustomText gets a reference to the given string and assigns it to the CustomText field.
func (o *SponsoredProductsCreateDraftProductAd) SetCustomText(v string) {
	o.CustomText = &v
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *SponsoredProductsCreateDraftProductAd) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftProductAd) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *SponsoredProductsCreateDraftProductAd) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *SponsoredProductsCreateDraftProductAd) SetAsin(v string) {
	o.Asin = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *SponsoredProductsCreateDraftProductAd) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftProductAd) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *SponsoredProductsCreateDraftProductAd) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *SponsoredProductsCreateDraftProductAd) SetSku(v string) {
	o.Sku = &v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsCreateDraftProductAd) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateDraftProductAd) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsCreateDraftProductAd) SetAdGroupId(v string) {
	o.AdGroupId = v
}

func (o SponsoredProductsCreateDraftProductAd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GlobalStoreSetting) {
		toSerialize["globalStoreSetting"] = o.GlobalStoreSetting
	}
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
	return toSerialize, nil
}

type NullableSponsoredProductsCreateDraftProductAd struct {
	value *SponsoredProductsCreateDraftProductAd
	isSet bool
}

func (v NullableSponsoredProductsCreateDraftProductAd) Get() *SponsoredProductsCreateDraftProductAd {
	return v.value
}

func (v *NullableSponsoredProductsCreateDraftProductAd) Set(val *SponsoredProductsCreateDraftProductAd) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateDraftProductAd) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateDraftProductAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateDraftProductAd(val *SponsoredProductsCreateDraftProductAd) *NullableSponsoredProductsCreateDraftProductAd {
	return &NullableSponsoredProductsCreateDraftProductAd{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateDraftProductAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateDraftProductAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
