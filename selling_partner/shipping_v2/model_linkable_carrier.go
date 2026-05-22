package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the LinkableCarrier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LinkableCarrier{}

// LinkableCarrier Info About Linkable Carrier
type LinkableCarrier struct {
	// The carrier identifier for the offering, provided by the carrier.
	CarrierId *string `json:"carrierId,omitempty"`
	// A list of LinkableAccountType
	LinkableAccountTypes []LinkableAccountType `json:"linkableAccountTypes,omitempty"`
}

// NewLinkableCarrier instantiates a new LinkableCarrier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLinkableCarrier() *LinkableCarrier {
	this := LinkableCarrier{}
	return &this
}

// NewLinkableCarrierWithDefaults instantiates a new LinkableCarrier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLinkableCarrierWithDefaults() *LinkableCarrier {
	this := LinkableCarrier{}
	return &this
}

// GetCarrierId returns the CarrierId field value if set, zero value otherwise.
func (o *LinkableCarrier) GetCarrierId() string {
	if o == nil || IsNil(o.CarrierId) {
		var ret string
		return ret
	}
	return *o.CarrierId
}

// GetCarrierIdOk returns a tuple with the CarrierId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkableCarrier) GetCarrierIdOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierId) {
		return nil, false
	}
	return o.CarrierId, true
}

// HasCarrierId returns a boolean if a field has been set.
func (o *LinkableCarrier) HasCarrierId() bool {
	if o != nil && !IsNil(o.CarrierId) {
		return true
	}

	return false
}

// SetCarrierId gets a reference to the given string and assigns it to the CarrierId field.
func (o *LinkableCarrier) SetCarrierId(v string) {
	o.CarrierId = &v
}

// GetLinkableAccountTypes returns the LinkableAccountTypes field value if set, zero value otherwise.
func (o *LinkableCarrier) GetLinkableAccountTypes() []LinkableAccountType {
	if o == nil || IsNil(o.LinkableAccountTypes) {
		var ret []LinkableAccountType
		return ret
	}
	return o.LinkableAccountTypes
}

// GetLinkableAccountTypesOk returns a tuple with the LinkableAccountTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LinkableCarrier) GetLinkableAccountTypesOk() ([]LinkableAccountType, bool) {
	if o == nil || IsNil(o.LinkableAccountTypes) {
		return nil, false
	}
	return o.LinkableAccountTypes, true
}

// HasLinkableAccountTypes returns a boolean if a field has been set.
func (o *LinkableCarrier) HasLinkableAccountTypes() bool {
	if o != nil && !IsNil(o.LinkableAccountTypes) {
		return true
	}

	return false
}

// SetLinkableAccountTypes gets a reference to the given []LinkableAccountType and assigns it to the LinkableAccountTypes field.
func (o *LinkableCarrier) SetLinkableAccountTypes(v []LinkableAccountType) {
	o.LinkableAccountTypes = v
}

func (o LinkableCarrier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CarrierId) {
		toSerialize["carrierId"] = o.CarrierId
	}
	if !IsNil(o.LinkableAccountTypes) {
		toSerialize["linkableAccountTypes"] = o.LinkableAccountTypes
	}
	return toSerialize, nil
}

type NullableLinkableCarrier struct {
	value *LinkableCarrier
	isSet bool
}

func (v NullableLinkableCarrier) Get() *LinkableCarrier {
	return v.value
}

func (v *NullableLinkableCarrier) Set(val *LinkableCarrier) {
	v.value = val
	v.isSet = true
}

func (v NullableLinkableCarrier) IsSet() bool {
	return v.isSet
}

func (v *NullableLinkableCarrier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLinkableCarrier(val *LinkableCarrier) *NullableLinkableCarrier {
	return &NullableLinkableCarrier{value: val, isSet: true}
}

func (v NullableLinkableCarrier) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableLinkableCarrier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
