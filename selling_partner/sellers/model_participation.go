package sellers

import (
	"github.com/bytedance/sonic"
)

// checks if the Participation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Participation{}

// Participation Information that is specific to a seller in a marketplace.
type Participation struct {
	// If `true`, the seller participates in the marketplace. Otherwise `false`.
	IsParticipating bool `json:"isParticipating"`
	// Specifies if the seller has suspended listings. `true` if the seller Listing Status is set to Inactive, otherwise `false`.
	HasSuspendedListings bool `json:"hasSuspendedListings"`
}

type _Participation Participation

// NewParticipation instantiates a new Participation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParticipation(isParticipating bool, hasSuspendedListings bool) *Participation {
	this := Participation{}
	this.IsParticipating = isParticipating
	this.HasSuspendedListings = hasSuspendedListings
	return &this
}

// NewParticipationWithDefaults instantiates a new Participation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParticipationWithDefaults() *Participation {
	this := Participation{}
	return &this
}

// GetIsParticipating returns the IsParticipating field value
func (o *Participation) GetIsParticipating() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsParticipating
}

// GetIsParticipatingOk returns a tuple with the IsParticipating field value
// and a boolean to check if the value has been set.
func (o *Participation) GetIsParticipatingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsParticipating, true
}

// SetIsParticipating sets field value
func (o *Participation) SetIsParticipating(v bool) {
	o.IsParticipating = v
}

// GetHasSuspendedListings returns the HasSuspendedListings field value
func (o *Participation) GetHasSuspendedListings() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasSuspendedListings
}

// GetHasSuspendedListingsOk returns a tuple with the HasSuspendedListings field value
// and a boolean to check if the value has been set.
func (o *Participation) GetHasSuspendedListingsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasSuspendedListings, true
}

// SetHasSuspendedListings sets field value
func (o *Participation) SetHasSuspendedListings(v bool) {
	o.HasSuspendedListings = v
}

func (o Participation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["isParticipating"] = o.IsParticipating
	toSerialize["hasSuspendedListings"] = o.HasSuspendedListings
	return toSerialize, nil
}

type NullableParticipation struct {
	value *Participation
	isSet bool
}

func (v NullableParticipation) Get() *Participation {
	return v.value
}

func (v *NullableParticipation) Set(val *Participation) {
	v.value = val
	v.isSet = true
}

func (v NullableParticipation) IsSet() bool {
	return v.isSet
}

func (v *NullableParticipation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParticipation(val *Participation) *NullableParticipation {
	return &NullableParticipation{value: val, isSet: true}
}

func (v NullableParticipation) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableParticipation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
