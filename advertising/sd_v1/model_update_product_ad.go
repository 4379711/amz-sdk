package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the UpdateProductAd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProductAd{}

// UpdateProductAd struct for UpdateProductAd
type UpdateProductAd struct {
	// The state of the campaign associated with the product ad.
	State *string `json:"state,omitempty"`
	// The identifier of the product ad.
	AdId int64 `json:"adId"`
}

type _UpdateProductAd UpdateProductAd

// NewUpdateProductAd instantiates a new UpdateProductAd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProductAd(adId int64) *UpdateProductAd {
	this := UpdateProductAd{}
	this.AdId = adId
	return &this
}

// NewUpdateProductAdWithDefaults instantiates a new UpdateProductAd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProductAdWithDefaults() *UpdateProductAd {
	this := UpdateProductAd{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *UpdateProductAd) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductAd) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *UpdateProductAd) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *UpdateProductAd) SetState(v string) {
	o.State = &v
}

// GetAdId returns the AdId field value
func (o *UpdateProductAd) GetAdId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.AdId
}

// GetAdIdOk returns a tuple with the AdId field value
// and a boolean to check if the value has been set.
func (o *UpdateProductAd) GetAdIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdId, true
}

// SetAdId sets field value
func (o *UpdateProductAd) SetAdId(v int64) {
	o.AdId = v
}

func (o UpdateProductAd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	toSerialize["adId"] = o.AdId
	return toSerialize, nil
}

type NullableUpdateProductAd struct {
	value *UpdateProductAd
	isSet bool
}

func (v NullableUpdateProductAd) Get() *UpdateProductAd {
	return v.value
}

func (v *NullableUpdateProductAd) Set(val *UpdateProductAd) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProductAd) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProductAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProductAd(val *UpdateProductAd) *NullableUpdateProductAd {
	return &NullableUpdateProductAd{value: val, isSet: true}
}

func (v NullableUpdateProductAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableUpdateProductAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
