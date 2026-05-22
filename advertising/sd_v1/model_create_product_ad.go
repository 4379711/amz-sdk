package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateProductAd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProductAd{}

// CreateProductAd struct for CreateProductAd
type CreateProductAd struct {
	// The state of the campaign associated with the product ad.
	State string `json:"state"`
	// The identifier of the ad group.
	AdGroupId int64 `json:"adGroupId"`
	// The identifier of the campaign.
	CampaignId int64 `json:"campaignId"`
	// The URL where customers will land after clicking on its link. Must be provided if a LandingPageType is set. Please note that if a single product ad sets the landing page url, only one product ad can be added to the ad group. This field is not supported when using ASIN or SKU fields. ||Specifications| |------------------|------------------| |LandingPageType| Description| |STORE| The url should be in the format of https://www.amazon.com/stores/_* (using a correct Amazon url based on the marketplace)| |OFF_AMAZON_LINK| The url should be in the format of https://www.****.com. Note that this LandingPageType is not supported when using ASIN or SKU fields. A custom creative of headline, logo, image are require for this LandingPageType. | |MOMENT| Not yet supported. The url should be in the format of https://www.amazon.com/moments/promotion/{campaignId} (using a correct Amazon url based on the marketplace)|
	LandingPageURL  *string          `json:"landingPageURL,omitempty"`
	LandingPageType *LandingPageType `json:"landingPageType,omitempty"`
	// The name of the ad. Note that this field is not supported when using ASIN or SKU fields.
	AdName *string `json:"adName,omitempty"`
	// The ASIN of the product advertised by the product ad. Defined for vendors only.
	Asin *string `json:"asin,omitempty"`
	// The SKU of the product advertised by the product ad. Defined for seller accounts only.
	Sku *string `json:"sku,omitempty"`
}

type _CreateProductAd CreateProductAd

// NewCreateProductAd instantiates a new CreateProductAd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProductAd(state string, adGroupId int64, campaignId int64) *CreateProductAd {
	this := CreateProductAd{}
	this.State = state
	this.AdGroupId = adGroupId
	this.CampaignId = campaignId
	return &this
}

// NewCreateProductAdWithDefaults instantiates a new CreateProductAd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProductAdWithDefaults() *CreateProductAd {
	this := CreateProductAd{}
	return &this
}

// GetState returns the State field value
func (o *CreateProductAd) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CreateProductAd) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CreateProductAd) SetState(v string) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *CreateProductAd) GetAdGroupId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *CreateProductAd) GetAdGroupIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *CreateProductAd) SetAdGroupId(v int64) {
	o.AdGroupId = v
}

// GetCampaignId returns the CampaignId field value
func (o *CreateProductAd) GetCampaignId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *CreateProductAd) GetCampaignIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *CreateProductAd) SetCampaignId(v int64) {
	o.CampaignId = v
}

// GetLandingPageURL returns the LandingPageURL field value if set, zero value otherwise.
func (o *CreateProductAd) GetLandingPageURL() string {
	if o == nil || IsNil(o.LandingPageURL) {
		var ret string
		return ret
	}
	return *o.LandingPageURL
}

// GetLandingPageURLOk returns a tuple with the LandingPageURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductAd) GetLandingPageURLOk() (*string, bool) {
	if o == nil || IsNil(o.LandingPageURL) {
		return nil, false
	}
	return o.LandingPageURL, true
}

// HasLandingPageURL returns a boolean if a field has been set.
func (o *CreateProductAd) HasLandingPageURL() bool {
	if o != nil && !IsNil(o.LandingPageURL) {
		return true
	}

	return false
}

// SetLandingPageURL gets a reference to the given string and assigns it to the LandingPageURL field.
func (o *CreateProductAd) SetLandingPageURL(v string) {
	o.LandingPageURL = &v
}

// GetLandingPageType returns the LandingPageType field value if set, zero value otherwise.
func (o *CreateProductAd) GetLandingPageType() LandingPageType {
	if o == nil || IsNil(o.LandingPageType) {
		var ret LandingPageType
		return ret
	}
	return *o.LandingPageType
}

// GetLandingPageTypeOk returns a tuple with the LandingPageType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductAd) GetLandingPageTypeOk() (*LandingPageType, bool) {
	if o == nil || IsNil(o.LandingPageType) {
		return nil, false
	}
	return o.LandingPageType, true
}

// HasLandingPageType returns a boolean if a field has been set.
func (o *CreateProductAd) HasLandingPageType() bool {
	if o != nil && !IsNil(o.LandingPageType) {
		return true
	}

	return false
}

// SetLandingPageType gets a reference to the given LandingPageType and assigns it to the LandingPageType field.
func (o *CreateProductAd) SetLandingPageType(v LandingPageType) {
	o.LandingPageType = &v
}

// GetAdName returns the AdName field value if set, zero value otherwise.
func (o *CreateProductAd) GetAdName() string {
	if o == nil || IsNil(o.AdName) {
		var ret string
		return ret
	}
	return *o.AdName
}

// GetAdNameOk returns a tuple with the AdName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductAd) GetAdNameOk() (*string, bool) {
	if o == nil || IsNil(o.AdName) {
		return nil, false
	}
	return o.AdName, true
}

// HasAdName returns a boolean if a field has been set.
func (o *CreateProductAd) HasAdName() bool {
	if o != nil && !IsNil(o.AdName) {
		return true
	}

	return false
}

// SetAdName gets a reference to the given string and assigns it to the AdName field.
func (o *CreateProductAd) SetAdName(v string) {
	o.AdName = &v
}

// GetAsin returns the Asin field value if set, zero value otherwise.
func (o *CreateProductAd) GetAsin() string {
	if o == nil || IsNil(o.Asin) {
		var ret string
		return ret
	}
	return *o.Asin
}

// GetAsinOk returns a tuple with the Asin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductAd) GetAsinOk() (*string, bool) {
	if o == nil || IsNil(o.Asin) {
		return nil, false
	}
	return o.Asin, true
}

// HasAsin returns a boolean if a field has been set.
func (o *CreateProductAd) HasAsin() bool {
	if o != nil && !IsNil(o.Asin) {
		return true
	}

	return false
}

// SetAsin gets a reference to the given string and assigns it to the Asin field.
func (o *CreateProductAd) SetAsin(v string) {
	o.Asin = &v
}

// GetSku returns the Sku field value if set, zero value otherwise.
func (o *CreateProductAd) GetSku() string {
	if o == nil || IsNil(o.Sku) {
		var ret string
		return ret
	}
	return *o.Sku
}

// GetSkuOk returns a tuple with the Sku field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProductAd) GetSkuOk() (*string, bool) {
	if o == nil || IsNil(o.Sku) {
		return nil, false
	}
	return o.Sku, true
}

// HasSku returns a boolean if a field has been set.
func (o *CreateProductAd) HasSku() bool {
	if o != nil && !IsNil(o.Sku) {
		return true
	}

	return false
}

// SetSku gets a reference to the given string and assigns it to the Sku field.
func (o *CreateProductAd) SetSku(v string) {
	o.Sku = &v
}

func (o CreateProductAd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["campaignId"] = o.CampaignId
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

type NullableCreateProductAd struct {
	value *CreateProductAd
	isSet bool
}

func (v NullableCreateProductAd) Get() *CreateProductAd {
	return v.value
}

func (v *NullableCreateProductAd) Set(val *CreateProductAd) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProductAd) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProductAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProductAd(val *CreateProductAd) *NullableCreateProductAd {
	return &NullableCreateProductAd{value: val, isSet: true}
}

func (v NullableCreateProductAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateProductAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
