package sd_v1

import (
	"github.com/bytedance/sonic"
)

// checks if the BaseProductAd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaseProductAd{}

// BaseProductAd struct for BaseProductAd
type BaseProductAd struct {
	// The state of the campaign associated with the product ad.
	State *string `json:"state,omitempty"`
}

// NewBaseProductAd instantiates a new BaseProductAd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseProductAd() *BaseProductAd {
	this := BaseProductAd{}
	return &this
}

// NewBaseProductAdWithDefaults instantiates a new BaseProductAd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseProductAdWithDefaults() *BaseProductAd {
	this := BaseProductAd{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *BaseProductAd) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseProductAd) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *BaseProductAd) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *BaseProductAd) SetState(v string) {
	o.State = &v
}

func (o BaseProductAd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableBaseProductAd struct {
	value *BaseProductAd
	isSet bool
}

func (v NullableBaseProductAd) Get() *BaseProductAd {
	return v.value
}

func (v *NullableBaseProductAd) Set(val *BaseProductAd) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseProductAd) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseProductAd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseProductAd(val *BaseProductAd) *NullableBaseProductAd {
	return &NullableBaseProductAd{value: val, isSet: true}
}

func (v NullableBaseProductAd) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableBaseProductAd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
