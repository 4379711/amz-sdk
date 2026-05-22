package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateBrandVideoAd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBrandVideoAd{}

// CreateBrandVideoAd struct for CreateBrandVideoAd
type CreateBrandVideoAd struct {
	LandingPage LandingPage `json:"landingPage"`
	// The name of the ad.
	Name  string                    `json:"name"`
	State CreateOrUpdateEntityState `json:"state"`
	// The adGroup identifier.
	AdGroupId string                   `json:"adGroupId"`
	Creative  CreateBrandVideoCreative `json:"creative"`
}

type _CreateBrandVideoAd CreateBrandVideoAd

// NewCreateBrandVideoAd instantiates a new CreateBrandVideoAd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBrandVideoAd(landingPage LandingPage, name string, state CreateOrUpdateEntityState, adGroupId string, creative CreateBrandVideoCreative) *CreateBrandVideoAd {
	this := CreateBrandVideoAd{}
	this.LandingPage = landingPage
	this.Name = name
	this.State = state
	this.AdGroupId = adGroupId
	this.Creative = creative
	return &this
}

// NewCreateBrandVideoAdWithDefaults instantiates a new CreateBrandVideoAd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBrandVideoAdWithDefaults() *CreateBrandVideoAd {
	this := CreateBrandVideoAd{}
	return &this
}

// GetLandingPage returns the LandingPage field value
func (o *CreateBrandVideoAd) GetLandingPage() LandingPage {
	if o == nil {
		var ret LandingPage
		return ret
	}

	return o.LandingPage
}

// GetLandingPageOk returns a tuple with the LandingPage field value
// and a boolean to check if the value has been set.
func (o *CreateBrandVideoAd) GetLandingPageOk() (*LandingPage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LandingPage, true
}

// SetLandingPage sets field value
func (o *CreateBrandVideoAd) SetLandingPage(v LandingPage) {
	o.LandingPage = v
}

// GetName returns the Name field value
func (o *CreateBrandVideoAd) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateBrandVideoAd) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateBrandVideoAd) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *CreateBrandVideoAd) GetState() CreateOrUpdateEntityState {
	if o == nil {
		var ret CreateOrUpdateEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CreateBrandVideoAd) GetStateOk() (*CreateOrUpdateEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CreateBrandVideoAd) SetState(v CreateOrUpdateEntityState) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *CreateBrandVideoAd) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *CreateBrandVideoAd) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *CreateBrandVideoAd) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetCreative returns the Creative field value
func (o *CreateBrandVideoAd) GetCreative() CreateBrandVideoCreative {
	if o == nil {
		var ret CreateBrandVideoCreative
		return ret
	}

	return o.Creative
}

// GetCreativeOk returns a tuple with the Creative field value
// and a boolean to check if the value has been set.
func (o *CreateBrandVideoAd) GetCreativeOk() (*CreateBrandVideoCreative, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creative, true
}

// SetCreative sets field value
func (o *CreateBrandVideoAd) SetCreative(v CreateBrandVideoCreative) {
	o.Creative = v
}

func (o CreateBrandVideoAd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["landingPage"] = o.LandingPage
	toSerialize["name"] = o.Name
	toSerialize["state"] = o.State
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["creative"] = o.Creative
	return toSerialize, nil
}

type NullableCreateBrandVideoAd struct {
	value *CreateBrandVideoAd
	isSet bool
}

func (v NullableCreateBrandVideoAd) Get() *CreateBrandVideoAd {
	return v.value
}

func (v *NullableCreateBrandVideoAd) Set(val *CreateBrandVideoAd) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBrandVideoAd) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBrandVideoAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBrandVideoAd(val *CreateBrandVideoAd) *NullableCreateBrandVideoAd {
	return &NullableCreateBrandVideoAd{value: val, isSet: true}
}

func (v NullableCreateBrandVideoAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateBrandVideoAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
