package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the ProductAd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductAd{}

// ProductAd struct for ProductAd
type ProductAd struct {
	// The state of the campaign associated with the product ad.
	State *string `json:"state,omitempty"`
	// The identifier of the product ad.
	AdId *int64 `json:"adId,omitempty"`
	// The identifier of the ad group.
	AdGroupId *int64 `json:"adGroupId,omitempty"`
	// The identifier of the campaign.
	CampaignId *int64 `json:"campaignId,omitempty"`
	// The URL where customers will land after clicking on its link. Must be provided if a LandingPageType is set. Please note that if a single product ad sets the landing page url, only one product ad can be added to the ad group. This field is not supported when using ASIN or SKU fields. ||Specifications| |------------------|------------------| |LandingPageType| Description| |STORE| The url should be in the format of https://www.amazon.com/stores/_* (using a correct Amazon url based on the marketplace)| |OFF_AMAZON_LINK| The url should be in the format of https://www.****.com. Note that this LandingPageType is not supported when using ASIN or SKU fields. A custom creative of headline, logo, image are require for this LandingPageType. | |MOMENT| Not yet supported. The url should be in the format of https://www.amazon.com/moments/promotion/{campaignId} (using a correct Amazon url based on the marketplace)|
	LandingPageURL  *string          `json:"landingPageURL,omitempty"`
	LandingPageType *LandingPageType `json:"landingPageType,omitempty"`
	// The name of the ad. Note that this field is not supported when using ASIN or SKU fields.
	AdName *string `json:"adName,omitempty"`
	// The Amazon ASIN of the product advertised by the product ad. This parameter is included in the response for sellers and vendors.
	Asin *string `json:"asin,omitempty"`
	// The Amazon SKU of the product advertised by the product ad. This parameter is included in the response for sellers.
	Sku *string `json:"sku,omitempty"`
}

// NewProductAd instantiates a new ProductAd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductAd() *ProductAd {
	this := ProductAd{}
	return &this
}

// NewProductAdWithDefaults instantiates a new ProductAd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductAdWithDefaults() *ProductAd {
	this := ProductAd{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ProductAd) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAd) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ProductAd) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ProductAd) SetState(v string) {
	o.State = &v
}

// GetAdId returns the AdId field value if set, zero value otherwise.
func (o *ProductAd) GetAdId() int64 {
	if o == nil || IsNil(o.AdId) {
		var ret int64
		return ret
	}
	return *o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAd) GetAdIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AdId) {
		return nil, false
	}
	return o.AdId, true
}

// HasAdId returns a boolean if a field has been set.
func (o *ProductAd) HasAdId() bool {
	if o != nil && !IsNil(o.AdId) {
		return true
	}

	return false
}

// SetAdId gets a reference to the given int64 and assigns it to the AdId field.
func (o *ProductAd) SetAdId(v int64) {
	o.AdId = &v
}

// GetAdGroupId returns the AdGroupId field value if set, zero value otherwise.
func (o *ProductAd) GetAdGroupId() int64 {
	if o == nil || IsNil(o.AdGroupId) {
		var ret int64
		return ret
	}
	return *o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAd) GetAdGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AdGroupId) {
		return nil, false
	}
	return o.AdGroupId, true
}

// HasAdGroupId returns a boolean if a field has been set.
func (o *ProductAd) HasAdGroupId() bool {
	if o != nil && !IsNil(o.AdGroupId) {
		return true
	}

	return false
}

// SetAdGroupId gets a reference to the given int64 and assigns it to the AdGroupId field.
func (o *ProductAd) SetAdGroupId(v int64) {
	o.AdGroupId = &v
}

// GetCampaignId returns the CampaignId field value if set, zero value otherwise.
func (o *ProductAd) GetCampaignId() int64 {
	if o == nil || IsNil(o.CampaignId) {
		var ret int64
		return ret
	}
	return *o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAd) GetCampaignIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CampaignId) {
		return nil, false
	}
	return o.CampaignId, true
}

// HasCampaignId returns a boolean if a field has been set.
func (o *ProductAd) HasCampaignId() bool {
	if o != nil && !IsNil(o.CampaignId) {
		return true
	}

	return false
}

// SetCampaignId gets a reference to the given int64 and assigns it to the CampaignId field.
func (o *ProductAd) SetCampaignId(v int64) {
	o.CampaignId = &v
}

// GetLandingPageURL returns the LandingPageURL field value if set, zero value otherwise.
func (o *ProductAd) GetLandingPageURL() string {
	if o == nil || IsNil(o.LandingPageURL) {
		var ret string
		return ret
	}
	return *o.LandingPageURL
}

// GetLandingPageURLOk returns a tuple with the LandingPageURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAd) GetLandingPageURLOk() (*string, bool) {
	if o == nil || IsNil(o.LandingPageURL) {
		return nil, false
	}
	return o.LandingPageURL, true
}

// HasLandingPageURL returns a boolean if a field has been set.
func (o *ProductAd) HasLandingPageURL() bool {
	if o != nil && !IsNil(o.LandingPageURL) {
		return true
	}

	return false
}

// SetLandingPageURL gets a reference to the given string and assigns it to the LandingPageURL field.
func (o *ProductAd) SetLandingPageURL(v string) {
	o.LandingPageURL = &v
}

// GetLandingPageType returns the LandingPageType field value if set, zero value otherwise.
func (o *ProductAd) GetLandingPageType() LandingPageType {
	if o == nil || IsNil(o.LandingPageType) {
		var ret LandingPageType
		return ret
	}
	return *o.LandingPageType
}

// GetLandingPageTypeOk returns a tuple with the LandingPageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAd) GetLandingPageTypeOk() (*LandingPageType, bool) {
	if o == nil || IsNil(o.LandingPageType) {
		return nil, false
	}
	return o.LandingPageType, true
}

// HasLandingPageType returns a boolean if a field has been set.
func (o *ProductAd) HasLandingPageType() bool {
	if o != nil && !IsNil(o.LandingPageType) {
		return true
	}

	return false
}

// SetLandingPageType gets a reference to the given LandingPageType and assigns it to the LandingPageType field.
func (o *ProductAd) SetLandingPageType(v LandingPageType) {
	o.LandingPageType = &v
}

// GetAdName returns the AdName field value if set, zero value otherwise.
func (o *ProductAd) GetAdName() string {
	if o == nil || IsNil(o.AdName) {
		var ret string
		return ret
	}
	return *o.AdName
}

// GetAdNameOk returns a tuple with the AdName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAd) GetAdNameOk() (*string, bool) {
	if o == nil || IsNil(o.AdName) {
		return nil, false
	}
	return o.AdName, true
}

// HasAdName returns a boolean if a field has been set.
func (o *ProductAd) HasAdName() bool {
	if o != nil && !IsNil(o.AdName) {
		return true
	}

	return false
}

// SetAdName gets a reference to the given string and assigns it to the AdName field.
func (o *ProductAd) SetAdName(v string) {
	o.AdName = &v
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *ProductAd) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAd) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *ProductAd) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *ProductAd) SetAsin(v string) {
	o.Asin = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *ProductAd) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductAd) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *ProductAd) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *ProductAd) SetSku(v string) {
	o.Sku = &v
}

func (o ProductAd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !IsNil(o.AdId) {
		toSerialize["adId"] = o.AdId
	}
	if !IsNil(o.AdGroupId) {
		toSerialize["adGroupId"] = o.AdGroupId
	}
	if !IsNil(o.CampaignId) {
		toSerialize["campaignId"] = o.CampaignId
	}
	if !IsNil(o.LandingPageURL) {
		toSerialize["landingPageURL"] = o.LandingPageURL
	}
	if !IsNil(o.LandingPageType) {
		toSerialize["landingPageType"] = o.LandingPageType
	}
	if !IsNil(o.AdName) {
		toSerialize["adName"] = o.AdName
	}
	if !IsNil(o.Asin) {
		toSerialize["asin"] = o.Asin
	}
	if !IsNil(o.Sku) {
		toSerialize["sku"] = o.Sku
	}
	return toSerialize, nil
}

type NullableProductAd struct {
	value *ProductAd
	isSet bool
}

func (v NullableProductAd) Get() *ProductAd {
	return v.value
}

func (v *NullableProductAd) Set(val *ProductAd) {
	v.value = val
	v.isSet = true
}

func (v NullableProductAd) IsSet() bool {
	return v.isSet
}

func (v *NullableProductAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductAd(val *ProductAd) *NullableProductAd {
	return &NullableProductAd{value: val, isSet: true}
}

func (v NullableProductAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableProductAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
