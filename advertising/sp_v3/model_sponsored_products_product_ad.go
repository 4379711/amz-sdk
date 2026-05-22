package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsProductAd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsProductAd{}

// SponsoredProductsProductAd struct for SponsoredProductsProductAd
type SponsoredProductsProductAd struct {
	GlobalStoreSetting *SponsoredProductsGlobalStoreSetting `json:"globalStoreSetting,omitempty"`
	// The product ad identifier.
	AdId string `json:"adId"`
	// The campaign identifier.
	CampaignId string `json:"campaignId"`
	// The custom text that is associated with this ad. Defined for custom text ads only.
	CustomText *string `json:"customText,omitempty"`
	// The ASIN associated with the product. Defined for vendors only.
	Asin  *string                      `json:"asin,omitempty"`
	State SponsoredProductsEntityState `json:"state"`
	// The SKU associated with the product. Defined for seller accounts only.
	Sku *string `json:"sku,omitempty"`
	// The ad group identifier.
	AdGroupId    string                                  `json:"adGroupId"`
	ExtendedData *SponsoredProductsProductAdExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsProductAd SponsoredProductsProductAd

// NewSponsoredProductsProductAd instantiates a new SponsoredProductsProductAd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsProductAd(adId string, campaignId string, state SponsoredProductsEntityState, adGroupId string) *SponsoredProductsProductAd {
	this := SponsoredProductsProductAd{}
	this.AdId = adId
	this.CampaignId = campaignId
	this.State = state
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsProductAdWithDefaults instantiates a new SponsoredProductsProductAd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsProductAdWithDefaults() *SponsoredProductsProductAd {
	this := SponsoredProductsProductAd{}
	return &this
}

// GetGlobalStoreSetting returns the GlobalStoreSetting field value if set, zero value otherwise.
func (o *SponsoredProductsProductAd) GetGlobalStoreSetting() SponsoredProductsGlobalStoreSetting {
	if o == nil || IsNil(o.GlobalStoreSetting) {
		var ret SponsoredProductsGlobalStoreSetting
		return ret
	}
	return *o.GlobalStoreSetting
}

// GetGlobalStoreSettingOk returns a tuple with the GlobalStoreSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAd) GetGlobalStoreSettingOk() (*SponsoredProductsGlobalStoreSetting, bool) {
	if o == nil || IsNil(o.GlobalStoreSetting) {
		return nil, false
	}
	return o.GlobalStoreSetting, true
}

// HasGlobalStoreSetting returns a boolean if a field has been set.
func (o *SponsoredProductsProductAd) HasGlobalStoreSetting() bool {
	if o != nil && !IsNil(o.GlobalStoreSetting) {
		return true
	}

	return false
}

// SetGlobalStoreSetting gets a reference to the given SponsoredProductsGlobalStoreSetting and assigns it to the GlobalStoreSetting field.
func (o *SponsoredProductsProductAd) SetGlobalStoreSetting(v SponsoredProductsGlobalStoreSetting) {
	o.GlobalStoreSetting = &v
}

// GetAdId returns the AdId field value
func (o *SponsoredProductsProductAd) GetAdId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAd) GetAdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdId, true
}

// SetAdId sets field value
func (o *SponsoredProductsProductAd) SetAdId(v string) {
	o.AdId = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsProductAd) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAd) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsProductAd) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetCustomText returns the CustomText field value if set, zero value otherwise.
func (o *SponsoredProductsProductAd) GetCustomText() string {
	if o == nil || IsNil(o.CustomText) {
		var ret string
		return ret
	}
	return *o.CustomText
}

// GetCustomTextOk returns a tuple with the CustomText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAd) GetCustomTextOk() (*string, bool) {
	if o == nil || IsNil(o.CustomText) {
		return nil, false
	}
	return o.CustomText, true
}

// HasCustomText returns a boolean if a field has been set.
func (o *SponsoredProductsProductAd) HasCustomText() bool {
	if o != nil && !IsNil(o.CustomText) {
		return true
	}

	return false
}

// SetCustomText gets a reference to the given string and assigns it to the CustomText field.
func (o *SponsoredProductsProductAd) SetCustomText(v string) {
	o.CustomText = &v
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *SponsoredProductsProductAd) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAd) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *SponsoredProductsProductAd) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *SponsoredProductsProductAd) SetAsin(v string) {
	o.Asin = &v
}

// GetState returns the State field value
func (o *SponsoredProductsProductAd) GetState() SponsoredProductsEntityState {
	if o == nil {
		var ret SponsoredProductsEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAd) GetStateOk() (*SponsoredProductsEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsProductAd) SetState(v SponsoredProductsEntityState) {
	o.State = v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *SponsoredProductsProductAd) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAd) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *SponsoredProductsProductAd) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *SponsoredProductsProductAd) SetSku(v string) {
	o.Sku = &v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsProductAd) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAd) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsProductAd) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsProductAd) GetExtendedData() SponsoredProductsProductAdExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsProductAdExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsProductAd) GetExtendedDataOk() (*SponsoredProductsProductAdExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsProductAd) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsProductAdExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsProductAd) SetExtendedData(v SponsoredProductsProductAdExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsProductAd) ToMap() (map[string]interface{}, error) {
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
	toSerialize["state"] = o.State
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	toSerialize["adGroupId"] = o.AdGroupId
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableSponsoredProductsProductAd struct {
	value *SponsoredProductsProductAd
	isSet bool
}

func (v NullableSponsoredProductsProductAd) Get() *SponsoredProductsProductAd {
	return v.value
}

func (v *NullableSponsoredProductsProductAd) Set(val *SponsoredProductsProductAd) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsProductAd) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsProductAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsProductAd(val *SponsoredProductsProductAd) *NullableSponsoredProductsProductAd {
	return &NullableSponsoredProductsProductAd{value: val, isSet: true}
}

func (v NullableSponsoredProductsProductAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsProductAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
