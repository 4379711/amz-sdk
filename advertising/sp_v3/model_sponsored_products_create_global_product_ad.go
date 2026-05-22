package sp_v3

import (
	"github.com/bytedance/sonic"
)

// checks if the SponsoredProductsCreateGlobalProductAd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsCreateGlobalProductAd{}

// SponsoredProductsCreateGlobalProductAd struct for SponsoredProductsCreateGlobalProductAd
type SponsoredProductsCreateGlobalProductAd struct {
	// The campaign identifier.
	CampaignId string `json:"campaignId"`
	// The custom text to use for creating a custom text ad for the associated ASIN. Defined for only KDP Authors and Book Vendors
	CustomText *string `json:"customText,omitempty" validate:"regexp=^[A-Z][^<>\\\\^]+$"`
	// Name for the product Ad
	Name  *string                                          `json:"name,omitempty"`
	Asin  *SponsoredProductsGlobalProductIdentifiers       `json:"asin,omitempty"`
	State SponsoredProductsCreateOrUpdateGlobalEntityState `json:"state"`
	Sku   *SponsoredProductsGlobalProductIdentifiers       `json:"sku,omitempty"`
	// The ad group identifier.
	AdGroupId string `json:"adGroupId"`
}

type _SponsoredProductsCreateGlobalProductAd SponsoredProductsCreateGlobalProductAd

// NewSponsoredProductsCreateGlobalProductAd instantiates a new SponsoredProductsCreateGlobalProductAd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsCreateGlobalProductAd(campaignId string, state SponsoredProductsCreateOrUpdateGlobalEntityState, adGroupId string) *SponsoredProductsCreateGlobalProductAd {
	this := SponsoredProductsCreateGlobalProductAd{}
	this.CampaignId = campaignId
	this.State = state
	this.AdGroupId = adGroupId
	return &this
}

// NewSponsoredProductsCreateGlobalProductAdWithDefaults instantiates a new SponsoredProductsCreateGlobalProductAd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsCreateGlobalProductAdWithDefaults() *SponsoredProductsCreateGlobalProductAd {
	this := SponsoredProductsCreateGlobalProductAd{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *SponsoredProductsCreateGlobalProductAd) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalProductAd) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *SponsoredProductsCreateGlobalProductAd) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetCustomText returns the CustomText field value if set, zero value otherwise.
func (o *SponsoredProductsCreateGlobalProductAd) GetCustomText() string {
	if o == nil || IsNil(o.CustomText) {
		var ret string
		return ret
	}
	return *o.CustomText
}

// GetCustomTextOk returns a tuple with the CustomText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalProductAd) GetCustomTextOk() (*string, bool) {
	if o == nil || IsNil(o.CustomText) {
		return nil, false
	}
	return o.CustomText, true
}

// HasCustomText returns a boolean if a field has been set.
func (o *SponsoredProductsCreateGlobalProductAd) HasCustomText() bool {
	if o != nil && !IsNil(o.CustomText) {
		return true
	}

	return false
}

// SetCustomText gets a reference to the given string and assigns it to the CustomText field.
func (o *SponsoredProductsCreateGlobalProductAd) SetCustomText(v string) {
	o.CustomText = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SponsoredProductsCreateGlobalProductAd) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalProductAd) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SponsoredProductsCreateGlobalProductAd) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SponsoredProductsCreateGlobalProductAd) SetName(v string) {
	o.Name = &v
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *SponsoredProductsCreateGlobalProductAd) GetAsin() SponsoredProductsGlobalProductIdentifiers {
	if o == nil || IsNil(o.Asin) {
		var ret SponsoredProductsGlobalProductIdentifiers
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalProductAd) GetAsinOk() (*SponsoredProductsGlobalProductIdentifiers, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *SponsoredProductsCreateGlobalProductAd) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given SponsoredProductsGlobalProductIdentifiers and assigns it to the Asin field.
func (o *SponsoredProductsCreateGlobalProductAd) SetAsin(v SponsoredProductsGlobalProductIdentifiers) {
	o.Asin = &v
}

// GetState returns the State field value
func (o *SponsoredProductsCreateGlobalProductAd) GetState() SponsoredProductsCreateOrUpdateGlobalEntityState {
	if o == nil {
		var ret SponsoredProductsCreateOrUpdateGlobalEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalProductAd) GetStateOk() (*SponsoredProductsCreateOrUpdateGlobalEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SponsoredProductsCreateGlobalProductAd) SetState(v SponsoredProductsCreateOrUpdateGlobalEntityState) {
	o.State = v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *SponsoredProductsCreateGlobalProductAd) GetSku() SponsoredProductsGlobalProductIdentifiers {
	if o == nil || IsNil(o.Sku) {
		var ret SponsoredProductsGlobalProductIdentifiers
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalProductAd) GetSkuOk() (*SponsoredProductsGlobalProductIdentifiers, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *SponsoredProductsCreateGlobalProductAd) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given SponsoredProductsGlobalProductIdentifiers and assigns it to the Sku field.
func (o *SponsoredProductsCreateGlobalProductAd) SetSku(v SponsoredProductsGlobalProductIdentifiers) {
	o.Sku = &v
}

// GetAdGroupId returns the AdGroupId field value
func (o *SponsoredProductsCreateGlobalProductAd) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsCreateGlobalProductAd) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *SponsoredProductsCreateGlobalProductAd) SetAdGroupId(v string) {
	o.AdGroupId = v
}

func (o SponsoredProductsCreateGlobalProductAd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	return toSerialize, nil
}

type NullableSponsoredProductsCreateGlobalProductAd struct {
	value *SponsoredProductsCreateGlobalProductAd
	isSet bool
}

func (v NullableSponsoredProductsCreateGlobalProductAd) Get() *SponsoredProductsCreateGlobalProductAd {
	return v.value
}

func (v *NullableSponsoredProductsCreateGlobalProductAd) Set(val *SponsoredProductsCreateGlobalProductAd) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsCreateGlobalProductAd) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsCreateGlobalProductAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsCreateGlobalProductAd(val *SponsoredProductsCreateGlobalProductAd) *NullableSponsoredProductsCreateGlobalProductAd {
	return &NullableSponsoredProductsCreateGlobalProductAd{value: val, isSet: true}
}

func (v NullableSponsoredProductsCreateGlobalProductAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsCreateGlobalProductAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
