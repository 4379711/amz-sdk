package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsGlobalProductAd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsGlobalProductAd{}

// SponsoredProductsGlobalProductAd struct for SponsoredProductsGlobalProductAd
type SponsoredProductsGlobalProductAd struct {
	// The product ad identifier.
	AdId string `json:"adId"`
	// The campaign identifier.
	CampaignId string `json:"campaignId"`
	// The custom text that is associated with this ad. Defined for custom text ads only.
	CustomText *string `json:"customText,omitempty"`
	// Name for the product Ad
	Name  *string                                    `json:"name,omitempty"`
	Asin  *SponsoredProductsGlobalProductIdentifiers `json:"asin,omitempty"`
	State SponsoredProductsGlobalEntityState         `json:"state"`
	Sku   *SponsoredProductsGlobalProductIdentifiers `json:"sku,omitempty"`
	// The ad group identifier.
	AdGroupId    string                                        `json:"adGroupId"`
	ExtendedData *SponsoredProductsGlobalProductAdExtendedData `json:"extendedData,omitempty"`
}

type _SponsoredProductsGlobalProductAd SponsoredProductsGlobalProductAd

// NewSponsoredProductsGlobalProductAd instantiates a new SponsoredProductsGlobalProductAd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsGlobalProductAd(adId string, campaignId string, state SponsoredProductsGlobalEntityState, adGroupId string) *SponsoredProductsGlobalProductAd {
	this := SponsoredProductsGlobalProductAd{}
	this.AdId = adId
	this.CampaignId = campaignId
	this.State = state
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsGlobalProductAdWithDefaults instantiates a new SponsoredProductsGlobalProductAd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsGlobalProductAdWithDefaults() *SponsoredProductsGlobalProductAd {
	this := SponsoredProductsGlobalProductAd{}
	return &this
}

// GetAdId returns the AdId field value
func (o *SponsoredProductsGlobalProductAd) GetAdId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalProductAd) GetAdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdId, true
}

// SetAdId sets field value
func (o *SponsoredProductsGlobalProductAd) SetAdId(v string) {
	o.AdId = v
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsGlobalProductAd) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalProductAd) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsGlobalProductAd) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetCustomText returns the CustomText field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalProductAd) GetCustomText() string {
	if o == nil || IsNil(o.CustomText) {
		var ret string
		return ret
	}
	return *o.CustomText
}

// GetCustomTextOk returns a tuple with the CustomText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalProductAd) GetCustomTextOk() (*string, bool) {
	if o == nil || IsNil(o.CustomText) {
		return nil, false
	}
	return o.CustomText, true
}

// HasCustomText returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalProductAd) HasCustomText() bool {
	if o != nil && !IsNil(o.CustomText) {
		return true
	}

	return false
}

// SetCustomText gets a reference to the given string and assigns it to the CustomText field.
func (o *SponsoredProductsGlobalProductAd) SetCustomText(v string) {
	o.CustomText = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalProductAd) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalProductAd) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalProductAd) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsGlobalProductAd) SetName(v string) {
	o.Name = &v
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalProductAd) GetAsin() SponsoredProductsGlobalProductIdentifiers {
	if o == nil || IsNil(o.Asin) {
		var ret SponsoredProductsGlobalProductIdentifiers
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalProductAd) GetAsinOk() (*SponsoredProductsGlobalProductIdentifiers, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalProductAd) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given SponsoredProductsGlobalProductIdentifiers and assigns it to the Asin field.
func (o *SponsoredProductsGlobalProductAd) SetAsin(v SponsoredProductsGlobalProductIdentifiers) {
	o.Asin = &v
}

// GetState returns the State field value
func (o *SponsoredProductsGlobalProductAd) GetState() SponsoredProductsGlobalEntityState {
	if o == nil {
		var ret SponsoredProductsGlobalEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalProductAd) GetStateOk() (*SponsoredProductsGlobalEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsGlobalProductAd) SetState(v SponsoredProductsGlobalEntityState) {
	o.State = v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalProductAd) GetSku() SponsoredProductsGlobalProductIdentifiers {
	if o == nil || IsNil(o.Sku) {
		var ret SponsoredProductsGlobalProductIdentifiers
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalProductAd) GetSkuOk() (*SponsoredProductsGlobalProductIdentifiers, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalProductAd) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given SponsoredProductsGlobalProductIdentifiers and assigns it to the Sku field.
func (o *SponsoredProductsGlobalProductAd) SetSku(v SponsoredProductsGlobalProductIdentifiers) {
	o.Sku = &v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsGlobalProductAd) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalProductAd) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsGlobalProductAd) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *SponsoredProductsGlobalProductAd) GetExtendedData() SponsoredProductsGlobalProductAdExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret SponsoredProductsGlobalProductAdExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsGlobalProductAd) GetExtendedDataOk() (*SponsoredProductsGlobalProductAdExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *SponsoredProductsGlobalProductAd) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given SponsoredProductsGlobalProductAdExtendedData and assigns it to the ExtendedData field.
func (o *SponsoredProductsGlobalProductAd) SetExtendedData(v SponsoredProductsGlobalProductAdExtendedData) {
	o.ExtendedData = &v
}

func (o SponsoredProductsGlobalProductAd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adId"] = o.AdId
	toSerialize["campaignId"] = o.CampaignId
	if !IsNil(o.CustomText) {
		toSerialize["customText"] = o.CustomText
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
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

type NullableSponsoredProductsGlobalProductAd struct {
	value *SponsoredProductsGlobalProductAd
	isSet bool
}

func (v NullableSponsoredProductsGlobalProductAd) Get() *SponsoredProductsGlobalProductAd {
	return v.value
}

func (v *NullableSponsoredProductsGlobalProductAd) Set(val *SponsoredProductsGlobalProductAd) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsGlobalProductAd) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsGlobalProductAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsGlobalProductAd(val *SponsoredProductsGlobalProductAd) *NullableSponsoredProductsGlobalProductAd {
	return &NullableSponsoredProductsGlobalProductAd{value: val, isSet: true}
}

func (v NullableSponsoredProductsGlobalProductAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsGlobalProductAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
