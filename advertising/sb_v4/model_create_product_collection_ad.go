package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateProductCollectionAd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateProductCollectionAd{}

// CreateProductCollectionAd struct for CreateProductCollectionAd
type CreateProductCollectionAd struct {
	LandingPage LandingPage `json:"landingPage"`
	// The name of the ad.
	Name  string                    `json:"name"`
	State CreateOrUpdateEntityState `json:"state"`
	// The adGroup identifier.
	AdGroupId string                          `json:"adGroupId"`
	Creative  CreateProductCollectionCreative `json:"creative"`
}

type _CreateProductCollectionAd CreateProductCollectionAd

// NewCreateProductCollectionAd instantiates a new CreateProductCollectionAd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProductCollectionAd(landingPage LandingPage, name string, state CreateOrUpdateEntityState, adGroupId string, creative CreateProductCollectionCreative) *CreateProductCollectionAd {
	this := CreateProductCollectionAd{}
	this.LandingPage = landingPage
	this.Name = name
	this.State = state
	this.AdGroupId = adGroupId
	this.Creative = creative
	return &this
}

// NewCreateProductCollectionAdWithDefaults instantiates a new CreateProductCollectionAd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProductCollectionAdWithDefaults() *CreateProductCollectionAd {
	this := CreateProductCollectionAd{}
	return &this
}

// GetLandingPage returns the LandingPage field value
func (o *CreateProductCollectionAd) GetLandingPage() LandingPage {
	if o == nil {
		var ret LandingPage
		return ret
	}

	return o.LandingPage
}

// GetLandingPageOk returns a tuple with the LandingPage field value
// and a boolean to check if the value has been set.
func (o *CreateProductCollectionAd) GetLandingPageOk() (*LandingPage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LandingPage, true
}

// SetLandingPage sets field value
func (o *CreateProductCollectionAd) SetLandingPage(v LandingPage) {
	o.LandingPage = v
}

// GetName returns the Name field value
func (o *CreateProductCollectionAd) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateProductCollectionAd) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateProductCollectionAd) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *CreateProductCollectionAd) GetState() CreateOrUpdateEntityState {
	if o == nil {
		var ret CreateOrUpdateEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CreateProductCollectionAd) GetStateOk() (*CreateOrUpdateEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CreateProductCollectionAd) SetState(v CreateOrUpdateEntityState) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *CreateProductCollectionAd) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *CreateProductCollectionAd) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *CreateProductCollectionAd) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetCreative returns the Creative field value
func (o *CreateProductCollectionAd) GetCreative() CreateProductCollectionCreative {
	if o == nil {
		var ret CreateProductCollectionCreative
		return ret
	}

	return o.Creative
}

// GetCreativeOk returns a tuple with the Creative field value
// and a boolean to check if the value has been set.
func (o *CreateProductCollectionAd) GetCreativeOk() (*CreateProductCollectionCreative, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creative, true
}

// SetCreative sets field value
func (o *CreateProductCollectionAd) SetCreative(v CreateProductCollectionCreative) {
	o.Creative = v
}

func (o CreateProductCollectionAd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["landingPage"] = o.LandingPage
	toSerialize["name"] = o.Name
	toSerialize["state"] = o.State
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["creative"] = o.Creative
	return toSerialize, nil
}

type NullableCreateProductCollectionAd struct {
	value *CreateProductCollectionAd
	isSet bool
}

func (v NullableCreateProductCollectionAd) Get() *CreateProductCollectionAd {
	return v.value
}

func (v *NullableCreateProductCollectionAd) Set(val *CreateProductCollectionAd) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProductCollectionAd) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProductCollectionAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProductCollectionAd(val *CreateProductCollectionAd) *NullableCreateProductCollectionAd {
	return &NullableCreateProductCollectionAd{value: val, isSet: true}
}

func (v NullableCreateProductCollectionAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateProductCollectionAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
