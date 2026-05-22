package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the Ad type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ad{}

// Ad struct for Ad
type Ad struct {
	// The ad identifier. Note: Ads created using version 3/non-multi ad group campaigns do not have an associated adId. [Learn more](https://advertising.amazon.com/API/docs/en-us/sponsored-brands/campaigns/managing-multi-ad-group-campaigns#ads).
	AdId *string `json:"adId,omitempty"`
	// The campaign identifier.
	CampaignId  string       `json:"campaignId"`
	LandingPage *LandingPage `json:"landingPage,omitempty"`
	// The name of the ad. Note: Ads created using version 3/non-multi ad group campaigns do not have an associated name. [Learn more](https://advertising.amazon.com/API/docs/en-us/sponsored-brands/campaigns/managing-multi-ad-group-campaigns#ads).
	Name  *string     `json:"name,omitempty"`
	State EntityState `json:"state"`
	// The adGroup identifier.
	AdGroupId    string          `json:"adGroupId"`
	Creative     *Creative       `json:"creative,omitempty"`
	ExtendedData *AdExtendedData `json:"extendedData,omitempty"`
}

type _Ad Ad

// NewAd instantiates a new Ad object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAd(campaignId string, state EntityState, adGroupId string) *Ad {
	this := Ad{}
	this.CampaignId = campaignId
	this.State = state
	this.AdGroupId = adGroupId
	return &this
}

// NewAdWithDefaults instantiates a new Ad object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdWithDefaults() *Ad {
	this := Ad{}
	return &this
}

// GetAdId returns the AdId field value if set, zero value otherwise.
func (o *Ad) GetAdId() string {
	if o == nil || IsNil(o.AdId) {
		var ret string
		return ret
	}
	return *o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ad) GetAdIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdId) {
		return nil, false
	}
	return o.AdId, true
}

// HasAdId returns a boolean if a field has been set.
func (o *Ad) HasAdId() bool {
	if o != nil && !IsNil(o.AdId) {
		return true
	}

	return false
}

// SetAdId gets a reference to the given string and assigns it to the AdId field.
func (o *Ad) SetAdId(v string) {
	o.AdId = &v
}

// GetCampaignId returns the CampaignId field value
func (o *Ad) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *Ad) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *Ad) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetLandingPage returns the LandingPage field value if set, zero value otherwise.
func (o *Ad) GetLandingPage() LandingPage {
	if o == nil || IsNil(o.LandingPage) {
		var ret LandingPage
		return ret
	}
	return *o.LandingPage
}

// GetLandingPageOk returns a tuple with the LandingPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ad) GetLandingPageOk() (*LandingPage, bool) {
	if o == nil || IsNil(o.LandingPage) {
		return nil, false
	}
	return o.LandingPage, true
}

// HasLandingPage returns a boolean if a field has been set.
func (o *Ad) HasLandingPage() bool {
	if o != nil && !IsNil(o.LandingPage) {
		return true
	}

	return false
}

// SetLandingPage gets a reference to the given LandingPage and assigns it to the LandingPage field.
func (o *Ad) SetLandingPage(v LandingPage) {
	o.LandingPage = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Ad) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ad) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Ad) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Ad) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value
func (o *Ad) GetState() EntityState {
	if o == nil {
		var ret EntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Ad) GetStateOk() (*EntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Ad) SetState(v EntityState) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *Ad) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *Ad) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *Ad) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetCreative returns the Creative field value if set, zero value otherwise.
func (o *Ad) GetCreative() Creative {
	if o == nil || IsNil(o.Creative) {
		var ret Creative
		return ret
	}
	return *o.Creative
}

// GetCreativeOk returns a tuple with the Creative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ad) GetCreativeOk() (*Creative, bool) {
	if o == nil || IsNil(o.Creative) {
		return nil, false
	}
	return o.Creative, true
}

// HasCreative returns a boolean if a field has been set.
func (o *Ad) HasCreative() bool {
	if o != nil && !IsNil(o.Creative) {
		return true
	}

	return false
}

// SetCreative gets a reference to the given Creative and assigns it to the Creative field.
func (o *Ad) SetCreative(v Creative) {
	o.Creative = &v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *Ad) GetExtendedData() AdExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret AdExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Ad) GetExtendedDataOk() (*AdExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *Ad) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given AdExtendedData and assigns it to the ExtendedData field.
func (o *Ad) SetExtendedData(v AdExtendedData) {
	o.ExtendedData = &v
}

func (o Ad) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdId) {
		toSerialize["adId"] = o.AdId
	}
	toSerialize["campaignId"] = o.CampaignId
	if !IsNil(o.LandingPage) {
		toSerialize["landingPage"] = o.LandingPage
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["state"] = o.State
	toSerialize["adGroupId"] = o.AdGroupId
	if !IsNil(o.Creative) {
		toSerialize["creative"] = o.Creative
	}
	if !IsNil(o.ExtendedData) {
		toSerialize["extendedData"] = o.ExtendedData
	}
	return toSerialize, nil
}

type NullableAd struct {
	value *Ad
	isSet bool
}

func (v NullableAd) Get() *Ad {
	return v.value
}

func (v *NullableAd) Set(val *Ad) {
	v.value = val
	v.isSet = true
}

func (v NullableAd) IsSet() bool {
	return v.isSet
}

func (v *NullableAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAd(val *Ad) *NullableAd {
	return &NullableAd{value: val, isSet: true}
}

func (v NullableAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
