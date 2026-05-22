package sp_v3

import "github.com/bytedance/sonic"

// checks if the SponsoredProductsErrorCause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SponsoredProductsErrorCause{}

// SponsoredProductsErrorCause Structure describing error cause - location in the payload and data causing error
type SponsoredProductsErrorCause struct {
	// Error location, JSON Path expression specifying element of API payload causing error
	Location string `json:"location"`
	// optional value causing error
	Trigger *string `json:"trigger,omitempty"`
}

type _SponsoredProductsErrorCause SponsoredProductsErrorCause

// NewSponsoredProductsErrorCause instantiates a new SponsoredProductsErrorCause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSponsoredProductsErrorCause(location string) *SponsoredProductsErrorCause {
	this := SponsoredProductsErrorCause{}
	this.Location = location
	return &this
}

// NewSponsoredProductsErrorCauseWithDefaults instantiates a new SponsoredProductsErrorCause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSponsoredProductsErrorCauseWithDefaults() *SponsoredProductsErrorCause {
	this := SponsoredProductsErrorCause{}
	return &this
}

// GetLocation returns the Location field value
func (o *SponsoredProductsErrorCause) GetLocation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *SponsoredProductsErrorCause) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *SponsoredProductsErrorCause) SetLocation(v string) {
	o.Location = v
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *SponsoredProductsErrorCause) GetTrigger() string {
	if o == nil || IsNil(o.Trigger) {
		var ret string
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SponsoredProductsErrorCause) GetTriggerOk() (*string, bool) {
	if o == nil || IsNil(o.Trigger) {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *SponsoredProductsErrorCause) HasTrigger() bool {
	if o != nil && !IsNil(o.Trigger) {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given string and assigns it to the Trigger field.
func (o *SponsoredProductsErrorCause) SetTrigger(v string) {
	o.Trigger = &v
}

func (o SponsoredProductsErrorCause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["location"] = o.Location
	if !IsNil(o.Trigger) {
		toSerialize["trigger"] = o.Trigger
	}
	return toSerialize, nil
}

type NullableSponsoredProductsErrorCause struct {
	value *SponsoredProductsErrorCause
	isSet bool
}

func (v NullableSponsoredProductsErrorCause) Get() *SponsoredProductsErrorCause {
	return v.value
}

func (v *NullableSponsoredProductsErrorCause) Set(val *SponsoredProductsErrorCause) {
	v.value = val
	v.isSet = true
}

func (v NullableSponsoredProductsErrorCause) IsSet() bool {
	return v.isSet
}

func (v *NullableSponsoredProductsErrorCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSponsoredProductsErrorCause(val *SponsoredProductsErrorCause) *NullableSponsoredProductsErrorCause {
	return &NullableSponsoredProductsErrorCause{value: val, isSet: true}
}

func (v NullableSponsoredProductsErrorCause) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSponsoredProductsErrorCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
