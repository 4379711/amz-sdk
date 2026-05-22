package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateVideoAd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVideoAd{}

// CreateVideoAd struct for CreateVideoAd
type CreateVideoAd struct {
	// The name of the ad.
	Name  string                    `json:"name"`
	State CreateOrUpdateEntityState `json:"state"`
	// The adGroup identifier.
	AdGroupId string              `json:"adGroupId"`
	Creative  CreateVideoCreative `json:"creative"`
}

type _CreateVideoAd CreateVideoAd

// NewCreateVideoAd instantiates a new CreateVideoAd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVideoAd(name string, state CreateOrUpdateEntityState, adGroupId string, creative CreateVideoCreative) *CreateVideoAd {
	this := CreateVideoAd{}
	this.Name = name
	this.State = state
	this.AdGroupId = adGroupId
	this.Creative = creative
	return &this
}

// NewCreateVideoAdWithDefaults instantiates a new CreateVideoAd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVideoAdWithDefaults() *CreateVideoAd {
	this := CreateVideoAd{}
	return &this
}

// GetName returns the Name field value
func (o *CreateVideoAd) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateVideoAd) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateVideoAd) SetName(v string) {
	o.Name = v
}

// GetState returns the State field value
func (o *CreateVideoAd) GetState() CreateOrUpdateEntityState {
	if o == nil {
		var ret CreateOrUpdateEntityState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *CreateVideoAd) GetStateOk() (*CreateOrUpdateEntityState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *CreateVideoAd) SetState(v CreateOrUpdateEntityState) {
	o.State = v
}

// GetAdGroupId returns the AdGroupId field value
func (o *CreateVideoAd) GetAdGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdGroupId
}

// GetAdGroupIdOk returns a tuple with the AdGroupId field value
// and a boolean to check if the value has been set.
func (o *CreateVideoAd) GetAdGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdGroupId, true
}

// SetAdGroupId sets field value
func (o *CreateVideoAd) SetAdGroupId(v string) {
	o.AdGroupId = v
}

// GetCreative returns the Creative field value
func (o *CreateVideoAd) GetCreative() CreateVideoCreative {
	if o == nil {
		var ret CreateVideoCreative
		return ret
	}

	return o.Creative
}

// GetCreativeOk returns a tuple with the Creative field value
// and a boolean to check if the value has been set.
func (o *CreateVideoAd) GetCreativeOk() (*CreateVideoCreative, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Creative, true
}

// SetCreative sets field value
func (o *CreateVideoAd) SetCreative(v CreateVideoCreative) {
	o.Creative = v
}

func (o CreateVideoAd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["state"] = o.State
	toSerialize["adGroupId"] = o.AdGroupId
	toSerialize["creative"] = o.Creative
	return toSerialize, nil
}

type NullableCreateVideoAd struct {
	value *CreateVideoAd
	isSet bool
}

func (v NullableCreateVideoAd) Get() *CreateVideoAd {
	return v.value
}

func (v *NullableCreateVideoAd) Set(val *CreateVideoAd) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVideoAd) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVideoAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVideoAd(val *CreateVideoAd) *NullableCreateVideoAd {
	return &NullableCreateVideoAd{value: val, isSet: true}
}

func (v NullableCreateVideoAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateVideoAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
