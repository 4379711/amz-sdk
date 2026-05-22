package sb_v4

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateAd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAd{}

// UpdateAd struct for UpdateAd
type UpdateAd struct {
	// The product ad identifier.
	AdId string `json:"adId"`
	// The name of the ad.
	Name  *string                    `json:"name,omitempty"`
	State *CreateOrUpdateEntityState `json:"state,omitempty"`
}

type _UpdateAd UpdateAd

// NewUpdateAd instantiates a new UpdateAd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAd(adId string) *UpdateAd {
	this := UpdateAd{}
	this.AdId = adId
	return &this
}

// NewUpdateAdWithDefaults instantiates a new UpdateAd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAdWithDefaults() *UpdateAd {
	this := UpdateAd{}
	return &this
}

// GetAdId returns the AdId field value
func (o *UpdateAd) GetAdId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value
// and a boolean to check if the value has been set.
func (o *UpdateAd) GetAdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdId, true
}

// SetAdId sets field value
func (o *UpdateAd) SetAdId(v string) {
	o.AdId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateAd) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAd) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateAd) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateAd) SetName(v string) {
	o.Name = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UpdateAd) GetState() CreateOrUpdateEntityState {
	if o == nil || IsNil(o.State) {
		var ret CreateOrUpdateEntityState
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAd) GetStateOk() (*CreateOrUpdateEntityState, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UpdateAd) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given CreateOrUpdateEntityState and assigns it to the State field.
func (o *UpdateAd) SetState(v CreateOrUpdateEntityState) {
	o.State = &v
}

func (o UpdateAd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["adId"] = o.AdId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableUpdateAd struct {
	value *UpdateAd
	isSet bool
}

func (v NullableUpdateAd) Get() *UpdateAd {
	return v.value
}

func (v *NullableUpdateAd) Set(val *UpdateAd) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAd) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAd(val *UpdateAd) *NullableUpdateAd {
	return &NullableUpdateAd{value: val, isSet: true}
}

func (v NullableUpdateAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
