package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateStoreSpotlightAd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateStoreSpotlightAd{}

// CreateStoreSpotlightAd struct for CreateStoreSpotlightAd
type CreateStoreSpotlightAd struct {
	LandingPage LandingPage `json:"landingPage"`
	// The name of the ad.
	Name  string                    `json:"name"`
	State CreateOrUpdateEntityState `json:"state"`
	// The adGroup identifier.
	AdGroupId string                       `json:"adGroupId"`
	Creative  CreateStoreSpotlightCreative `json:"creative"`
}

type _CreateStoreSpotlightAd CreateStoreSpotlightAd

// NewCreateStoreSpotlightAd instantiates a new CreateStoreSpotlightAd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateStoreSpotlightAd(landingPage LandingPage, name string, state CreateOrUpdateEntityState, adGroupId string, creative CreateStoreSpotlightCreative) *CreateStoreSpotlightAd {
	this := CreateStoreSpotlightAd{}
	this.LandingPage = landingPage
	this.Name = name
	this.State = state
	this.AdGroupId = adGroupId
	this.Creative = creative
	return &this
}

// NewCreateStoreSpotlightAdWithDefaults instantiates a new CreateStoreSpotlightAd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateStoreSpotlightAdWithDefaults() *CreateStoreSpotlightAd {
	this := CreateStoreSpotlightAd{}
	return &this
}

// GetLandingPage returns the LandingPage field value
func (o *CreateStoreSpotlightAd) GetLandingPage() LandingPage {
	if o == nil {
		var ret LandingPage
		return ret
	}

	return o.LandingPage
}

// GetLandingPageOk returns a tuple with the LandingPage field value
// and a boolean to check if the value has been set.
func (o *CreateStoreSpotlightAd) GetLandingPageOk() (*LandingPage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LandingPage, true
}

// SetLandingPage sets field value
func (o *CreateStoreSpotlightAd) SetLandingPage(v LandingPage) {
	o.LandingPage = v
}

// GetName returns the Name field value
func (o *CreateStoreSpotlightAd) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateStoreSpotlightAd) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateStoreSpotlightAd) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *CreateStoreSpotlightAd) GetState() CreateOrUpdateEntityState {
	if o == nil {
		var ret CreateOrUpdateEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CreateStoreSpotlightAd) GetStateOk() (*CreateOrUpdateEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CreateStoreSpotlightAd) SetState(v CreateOrUpdateEntityState) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *CreateStoreSpotlightAd) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *CreateStoreSpotlightAd) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *CreateStoreSpotlightAd) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetCreative returns the Creative field value
func (o *CreateStoreSpotlightAd) GetCreative() CreateStoreSpotlightCreative {
	if o == nil {
		var ret CreateStoreSpotlightCreative
		return ret
	}

	return o.Creative
}

// GetCreativeOk returns a tuple with the Creative field value
// and a boolean to check if the value has been set.
func (o *CreateStoreSpotlightAd) GetCreativeOk() (*CreateStoreSpotlightCreative, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creative, true
}

// SetCreative sets field value
func (o *CreateStoreSpotlightAd) SetCreative(v CreateStoreSpotlightCreative) {
	o.Creative = v
}

func (o CreateStoreSpotlightAd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["landingPage"] = o.LandingPage
	toSerialize["name"] = o.Name
	toSerialize["state"] = o.State
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["creative"] = o.Creative
	return toSerialize, nil
}

type NullableCreateStoreSpotlightAd struct {
	value *CreateStoreSpotlightAd
	isSet bool
}

func (v NullableCreateStoreSpotlightAd) Get() *CreateStoreSpotlightAd {
	return v.value
}

func (v *NullableCreateStoreSpotlightAd) Set(val *CreateStoreSpotlightAd) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateStoreSpotlightAd) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateStoreSpotlightAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateStoreSpotlightAd(val *CreateStoreSpotlightAd) *NullableCreateStoreSpotlightAd {
	return &NullableCreateStoreSpotlightAd{value: val, isSet: true}
}

func (v NullableCreateStoreSpotlightAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateStoreSpotlightAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
