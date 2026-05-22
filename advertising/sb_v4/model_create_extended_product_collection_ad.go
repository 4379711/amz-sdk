package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateExtendedProductCollectionAd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateExtendedProductCollectionAd{}

// CreateExtendedProductCollectionAd struct for CreateExtendedProductCollectionAd
type CreateExtendedProductCollectionAd struct {
	LandingPage LandingPage `json:"landingPage"`
	// The name of the ad.
	Name  string                    `json:"name"`
	State CreateOrUpdateEntityState `json:"state"`
	// The adGroup identifier.
	AdGroupId string                                  `json:"adGroupId"`
	Creative  CreateExtendedProductCollectionCreative `json:"creative"`
}

type _CreateExtendedProductCollectionAd CreateExtendedProductCollectionAd

// NewCreateExtendedProductCollectionAd instantiates a new CreateExtendedProductCollectionAd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateExtendedProductCollectionAd(landingPage LandingPage, name string, state CreateOrUpdateEntityState, adGroupId string, creative CreateExtendedProductCollectionCreative) *CreateExtendedProductCollectionAd {
	this := CreateExtendedProductCollectionAd{}
	this.LandingPage = landingPage
	this.Name = name
	this.State = state
	this.AdGroupId = adGroupId
	this.Creative = creative
	return &this
}

// NewCreateExtendedProductCollectionAdWithDefaults instantiates a new CreateExtendedProductCollectionAd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateExtendedProductCollectionAdWithDefaults() *CreateExtendedProductCollectionAd {
	this := CreateExtendedProductCollectionAd{}
	return &this
}

// GetLandingPage returns the LandingPage field value
func (o *CreateExtendedProductCollectionAd) GetLandingPage() LandingPage {
	if o == nil {
		var ret LandingPage
		return ret
	}

	return o.LandingPage
}

// GetLandingPageOk returns a tuple with the LandingPage field value
// and a boolean to check if the value has been set.
func (o *CreateExtendedProductCollectionAd) GetLandingPageOk() (*LandingPage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LandingPage, true
}

// SetLandingPage sets field value
func (o *CreateExtendedProductCollectionAd) SetLandingPage(v LandingPage) {
	o.LandingPage = v
}

// GetName returns the Name field value
func (o *CreateExtendedProductCollectionAd) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateExtendedProductCollectionAd) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateExtendedProductCollectionAd) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *CreateExtendedProductCollectionAd) GetState() CreateOrUpdateEntityState {
	if o == nil {
		var ret CreateOrUpdateEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CreateExtendedProductCollectionAd) GetStateOk() (*CreateOrUpdateEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CreateExtendedProductCollectionAd) SetState(v CreateOrUpdateEntityState) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *CreateExtendedProductCollectionAd) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *CreateExtendedProductCollectionAd) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *CreateExtendedProductCollectionAd) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetCreative returns the Creative field value
func (o *CreateExtendedProductCollectionAd) GetCreative() CreateExtendedProductCollectionCreative {
	if o == nil {
		var ret CreateExtendedProductCollectionCreative
		return ret
	}

	return o.Creative
}

// GetCreativeOk returns a tuple with the Creative field value
// and a boolean to check if the value has been set.
func (o *CreateExtendedProductCollectionAd) GetCreativeOk() (*CreateExtendedProductCollectionCreative, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creative, true
}

// SetCreative sets field value
func (o *CreateExtendedProductCollectionAd) SetCreative(v CreateExtendedProductCollectionCreative) {
	o.Creative = v
}

func (o CreateExtendedProductCollectionAd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["landingPage"] = o.LandingPage
	toSerialize["name"] = o.Name
	toSerialize["state"] = o.State
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["creative"] = o.Creative
	return toSerialize, nil
}

type NullableCreateExtendedProductCollectionAd struct {
	value *CreateExtendedProductCollectionAd
	isSet bool
}

func (v NullableCreateExtendedProductCollectionAd) Get() *CreateExtendedProductCollectionAd {
	return v.value
}

func (v *NullableCreateExtendedProductCollectionAd) Set(val *CreateExtendedProductCollectionAd) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateExtendedProductCollectionAd) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateExtendedProductCollectionAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateExtendedProductCollectionAd(val *CreateExtendedProductCollectionAd) *NullableCreateExtendedProductCollectionAd {
	return &NullableCreateExtendedProductCollectionAd{value: val, isSet: true}
}

func (v NullableCreateExtendedProductCollectionAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateExtendedProductCollectionAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
