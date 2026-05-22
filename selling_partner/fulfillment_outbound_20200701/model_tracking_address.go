package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the TrackingAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrackingAddress{}

// TrackingAddress Address information for tracking the package.
type TrackingAddress struct {
	// The city.
	City string `json:"city"`
	// The state.
	State string `json:"state"`
	// The country.
	Country string `json:"country"`
}

type _TrackingAddress TrackingAddress

// NewTrackingAddress instantiates a new TrackingAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrackingAddress(city string, state string, country string) *TrackingAddress {
	this := TrackingAddress{}
	this.City = city
	this.State = state
	this.Country = country
	return &this
}

// NewTrackingAddressWithDefaults instantiates a new TrackingAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrackingAddressWithDefaults() *TrackingAddress {
	this := TrackingAddress{}
	return &this
}

// GetCity returns the City field value
func (o *TrackingAddress) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *TrackingAddress) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *TrackingAddress) SetCity(v string) {
	o.City = v
}

// GetState returns the State field value
func (o *TrackingAddress) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *TrackingAddress) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *TrackingAddress) SetState(v string) {
	o.State = v
}

// GetCountry returns the Country field value
func (o *TrackingAddress) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *TrackingAddress) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *TrackingAddress) SetCountry(v string) {
	o.Country = v
}

func (o TrackingAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["city"] = o.City
	toSerialize["state"] = o.State
	toSerialize["country"] = o.Country
	return toSerialize, nil
}

type NullableTrackingAddress struct {
	value *TrackingAddress
	isSet bool
}

func (v NullableTrackingAddress) Get() *TrackingAddress {
	return v.value
}

func (v *NullableTrackingAddress) Set(val *TrackingAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableTrackingAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableTrackingAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrackingAddress(val *TrackingAddress) *NullableTrackingAddress {
	return &NullableTrackingAddress{value: val, isSet: true}
}

func (v NullableTrackingAddress) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTrackingAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
