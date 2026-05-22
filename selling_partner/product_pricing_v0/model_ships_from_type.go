package product_pricing_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the ShipsFromType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipsFromType{}

// ShipsFromType The state and country from where the item is shipped.
type ShipsFromType struct {
	// The state from where the item is shipped.
	State *string `json:"State,omitempty"`
	// The country from where the item is shipped.
	Country *string `json:"Country,omitempty"`
}

// NewShipsFromType instantiates a new ShipsFromType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipsFromType() *ShipsFromType {
	this := ShipsFromType{}
	return &this
}

// NewShipsFromTypeWithDefaults instantiates a new ShipsFromType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipsFromTypeWithDefaults() *ShipsFromType {
	this := ShipsFromType{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *ShipsFromType) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipsFromType) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *ShipsFromType) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *ShipsFromType) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *ShipsFromType) GetCountry() string {
	if o == nil || IsNil(o.Country) {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipsFromType) GetCountryOk() (*string, bool) {
	if o == nil || IsNil(o.Country) {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *ShipsFromType) HasCountry() bool {
	if o != nil && !IsNil(o.Country) {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *ShipsFromType) SetCountry(v string) {
	o.Country = &v
}

func (o ShipsFromType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.State) {
		toSerialize["State"] = o.State
	}
	if !IsNil(o.Country) {
		toSerialize["Country"] = o.Country
	}
	return toSerialize, nil
}

type NullableShipsFromType struct {
	value *ShipsFromType
	isSet bool
}

func (v NullableShipsFromType) Get() *ShipsFromType {
	return v.value
}

func (v *NullableShipsFromType) Set(val *ShipsFromType) {
	v.value = val
	v.isSet = true
}

func (v NullableShipsFromType) IsSet() bool {
	return v.isSet
}

func (v *NullableShipsFromType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipsFromType(val *ShipsFromType) *NullableShipsFromType {
	return &NullableShipsFromType{value: val, isSet: true}
}

func (v NullableShipsFromType) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableShipsFromType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
