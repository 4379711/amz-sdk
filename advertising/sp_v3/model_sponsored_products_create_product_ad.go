package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateProductAd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateProductAd{}

// SponsoredProductsCreateProductAd struct for SponsoredProductsCreateProductAd
type SponsoredProductsCreateProductAd struct {
	GlobalStoreSetting *SponsoredProductsGlobalStoreSetting `json:"globalStoreSetting,omitempty"`
	// The campaign identifier.
	CampaignId string `json:"campaignId"`
	// The custom text to use for creating a custom text ad for the associated ASIN. Defined only for KDP Authors and Book Vendors in US marketplace.
	CustomText *string `json:"customText,omitempty" validate:"regexp=^[^a-z<>^][^<>^]+$"`
	// The ASIN associated with the product. Defined for vendors only.
	Asin  *string                                    `json:"asin,omitempty"`
	State SponsoredProductsCreateOrUpdateEntityState `json:"state"`
	// The SKU associated with the product. Defined for seller accounts only.
	Sku *string `json:"sku,omitempty"`
	// The ad group identifier.
	AdGroupId string `json:"adGroupId"`
}

type _SponsoredProductsCreateProductAd SponsoredProductsCreateProductAd

// NewSponsoredProductsCreateProductAd instantiates a new SponsoredProductsCreateProductAd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateProductAd(campaignId string, state SponsoredProductsCreateOrUpdateEntityState, adGroupId string) *SponsoredProductsCreateProductAd {
	this := SponsoredProductsCreateProductAd{}
	this.CampaignId = campaignId
	this.State = state
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsCreateProductAdWithDefaults instantiates a new SponsoredProductsCreateProductAd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateProductAdWithDefaults() *SponsoredProductsCreateProductAd {
	this := SponsoredProductsCreateProductAd{}
	return &this
}

// GetGlobalStoreSetting returns the GlobalStoreSetting field value if set, zero value otherwise.
func (o *SponsoredProductsCreateProductAd) GetGlobalStoreSetting() SponsoredProductsGlobalStoreSetting {
	if o == nil || IsNil(o.GlobalStoreSetting) {
		var ret SponsoredProductsGlobalStoreSetting
		return ret
	}
	return *o.GlobalStoreSetting
}

// GetGlobalStoreSettingOk returns a tuple with the GlobalStoreSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateProductAd) GetGlobalStoreSettingOk() (*SponsoredProductsGlobalStoreSetting, bool) {
	if o == nil || IsNil(o.GlobalStoreSetting) {
		return nil, false
	}
	return o.GlobalStoreSetting, true
}

// HasGlobalStoreSetting returns a boolean if a field has been set.
func (o *SponsoredProductsCreateProductAd) HasGlobalStoreSetting() bool {
	if o != nil && !IsNil(o.GlobalStoreSetting) {
		return true
	}

	return false
}

// SetGlobalStoreSetting gets a reference to the given SponsoredProductsGlobalStoreSetting and assigns it to the GlobalStoreSetting field.
func (o *SponsoredProductsCreateProductAd) SetGlobalStoreSetting(v SponsoredProductsGlobalStoreSetting) {
	o.GlobalStoreSetting = &v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateProductAd) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateProductAd) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateProductAd) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetCustomText returns the CustomText field value if set, zero value otherwise.
func (o *SponsoredProductsCreateProductAd) GetCustomText() string {
	if o == nil || IsNil(o.CustomText) {
		var ret string
		return ret
	}
	return *o.CustomText
}

// GetCustomTextOk returns a tuple with the CustomText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateProductAd) GetCustomTextOk() (*string, bool) {
	if o == nil || IsNil(o.CustomText) {
		return nil, false
	}
	return o.CustomText, true
}

// HasCustomText returns a boolean if a field has been set.
func (o *SponsoredProductsCreateProductAd) HasCustomText() bool {
	if o != nil && !IsNil(o.CustomText) {
		return true
	}

	return false
}

// SetCustomText gets a reference to the given string and assigns it to the CustomText field.
func (o *SponsoredProductsCreateProductAd) SetCustomText(v string) {
	o.CustomText = &v
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *SponsoredProductsCreateProductAd) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateProductAd) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *SponsoredProductsCreateProductAd) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *SponsoredProductsCreateProductAd) SetAsin(v string) {
	o.Asin = &v
}

// GetState returns the State field value
func (o *SponsoredProductsCreateProductAd) GetState() SponsoredProductsCreateOrUpdateEntityState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateProductAd) GetStateOk() (*SponsoredProductsCreateOrUpdateEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCreateProductAd) SetState(v SponsoredProductsCreateOrUpdateEntityState) {
	o.State = v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *SponsoredProductsCreateProductAd) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateProductAd) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *SponsoredProductsCreateProductAd) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *SponsoredProductsCreateProductAd) SetSku(v string) {
	o.Sku = &v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsCreateProductAd) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateProductAd) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsCreateProductAd) SetAdGroupId(v string) {
	o.AdGroupId = v
}

func (o SponsoredProductsCreateProductAd) ToMap() (map[string]interface{}, error) {
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
	toSerialize["state"] = o.State
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	toSerialize["adGroupId"] = o.AdGroupId
	return toSerialize, nil
}

type NullableSponsoredProductsCreateProductAd struct {
	value *SponsoredProductsCreateProductAd
	isSet bool
}

func (v NullableSponsoredProductsCreateProductAd) Get() *SponsoredProductsCreateProductAd {
	return v.value
}

func (v *NullableSponsoredProductsCreateProductAd) Set(val *SponsoredProductsCreateProductAd) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateProductAd) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateProductAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateProductAd(val *SponsoredProductsCreateProductAd) *NullableSponsoredProductsCreateProductAd {
	return &NullableSponsoredProductsCreateProductAd{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateProductAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateProductAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
