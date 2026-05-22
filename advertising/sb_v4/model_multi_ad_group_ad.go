package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the MultiAdGroupAd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultiAdGroupAd{}

// MultiAdGroupAd struct for MultiAdGroupAd
type MultiAdGroupAd struct {
	// The ad identifier.
	AdId string `json:"adId"`
	// The campaign identifier.
	CampaignId  string       `json:"campaignId"`
	LandingPage *LandingPage `json:"landingPage,omitempty"`
	// The name of the ad.
	Name  string      `json:"name"`
	State EntityState `json:"state"`
	// The adGroup identifier.
	AdGroupId    string          `json:"adGroupId"`
	Creative     *Creative       `json:"creative,omitempty"`
	ExtendedData *AdExtendedData `json:"extendedData,omitempty"`
}

type _MultiAdGroupAd MultiAdGroupAd

// NewMultiAdGroupAd instantiates a new MultiAdGroupAd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiAdGroupAd(adId string, campaignId string, name string, state EntityState, adGroupId string) *MultiAdGroupAd {
	this := MultiAdGroupAd{}
	this.AdId = adId
	this.CampaignId = campaignId
	this.Name = name
	this.State = state
	this.AdGroupId = adGroupId
	return &this
}

// NewMultiAdGroupAdWithDefaults instantiates a new MultiAdGroupAd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiAdGroupAdWithDefaults() *MultiAdGroupAd {
	this := MultiAdGroupAd{}
	return &this
}

// GetAdId returns the AdId field value
func (o *MultiAdGroupAd) GetAdId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value
// and a boolean to check if the value has been set.
func (o *MultiAdGroupAd) GetAdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdId, true
}

// SetAdId sets field value
func (o *MultiAdGroupAd) SetAdId(v string) {
	o.AdId = v
}

// GetCampaignId returns the CampaignId field value
func (o *MultiAdGroupAd) GetCampaignId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *MultiAdGroupAd) GetCampaignIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *MultiAdGroupAd) SetCampaignId(v string) {
	o.CampaignId = v
}

// GetLandingPage returns the LandingPage field value if set, zero value otherwise.
func (o *MultiAdGroupAd) GetLandingPage() LandingPage {
	if o == nil || IsNil(o.LandingPage) {
		var ret LandingPage
		return ret
	}
	return *o.LandingPage
}

// GetLandingPageOk returns a tuple with the LandingPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAdGroupAd) GetLandingPageOk() (*LandingPage, bool) {
	if o == nil || IsNil(o.LandingPage) {
		return nil, false
	}
	return o.LandingPage, true
}

// HasLandingPage returns a boolean if a field has been set.
func (o *MultiAdGroupAd) HasLandingPage() bool {
	if o != nil && !IsNil(o.LandingPage) {
		return true
	}

	return false
}

// SetLandingPage gets a reference to the given LandingPage and assigns it to the LandingPage field.
func (o *MultiAdGroupAd) SetLandingPage(v LandingPage) {
	o.LandingPage = &v
}

// GetName returns the Name field value
func (o *MultiAdGroupAd) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *MultiAdGroupAd) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *MultiAdGroupAd) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *MultiAdGroupAd) GetState() EntityState {
	if o == nil {
		var ret EntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *MultiAdGroupAd) GetStateOk() (*EntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *MultiAdGroupAd) SetState(v EntityState) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *MultiAdGroupAd) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *MultiAdGroupAd) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *MultiAdGroupAd) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetCreative returns the Creative field value if set, zero value otherwise.
func (o *MultiAdGroupAd) GetCreative() Creative {
	if o == nil || IsNil(o.Creative) {
		var ret Creative
		return ret
	}
	return *o.Creative
}

// GetCreativeOk returns a tuple with the Creative field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAdGroupAd) GetCreativeOk() (*Creative, bool) {
	if o == nil || IsNil(o.Creative) {
		return nil, false
	}
	return o.Creative, true
}

// HasCreative returns a boolean if a field has been set.
func (o *MultiAdGroupAd) HasCreative() bool {
	if o != nil && !IsNil(o.Creative) {
		return true
	}

	return false
}

// SetCreative gets a reference to the given Creative and assigns it to the Creative field.
func (o *MultiAdGroupAd) SetCreative(v Creative) {
	o.Creative = &v
}

// GetExtendedData returns the ExtendedData field value if set, zero value otherwise.
func (o *MultiAdGroupAd) GetExtendedData() AdExtendedData {
	if o == nil || IsNil(o.ExtendedData) {
		var ret AdExtendedData
		return ret
	}
	return *o.ExtendedData
}

// GetExtendedDataOk returns a tuple with the ExtendedData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultiAdGroupAd) GetExtendedDataOk() (*AdExtendedData, bool) {
	if o == nil || IsNil(o.ExtendedData) {
		return nil, false
	}
	return o.ExtendedData, true
}

// HasExtendedData returns a boolean if a field has been set.
func (o *MultiAdGroupAd) HasExtendedData() bool {
	if o != nil && !IsNil(o.ExtendedData) {
		return true
	}

	return false
}

// SetExtendedData gets a reference to the given AdExtendedData and assigns it to the ExtendedData field.
func (o *MultiAdGroupAd) SetExtendedData(v AdExtendedData) {
	o.ExtendedData = &v
}

func (o MultiAdGroupAd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adId"] = o.AdId
	toSerialize["campaignId"] = o.CampaignId
	if !IsNil(o.LandingPage) {
		toSerialize["landingPage"] = o.LandingPage
	}
	toSerialize["name"] = o.Name
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

type NullableMultiAdGroupAd struct {
	value *MultiAdGroupAd
	isSet bool
}

func (v NullableMultiAdGroupAd) Get() *MultiAdGroupAd {
	return v.value
}

func (v *NullableMultiAdGroupAd) Set(val *MultiAdGroupAd) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiAdGroupAd) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiAdGroupAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiAdGroupAd(val *MultiAdGroupAd) *NullableMultiAdGroupAd {
	return &NullableMultiAdGroupAd{value: val, isSet: true}
}

func (v NullableMultiAdGroupAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableMultiAdGroupAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
