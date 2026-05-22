package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the DeliveryPreferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryPreferences{}

// DeliveryPreferences Contains all of the delivery instructions provided by the customer for the shipping address.
type DeliveryPreferences struct {
	// Drop-off location selected by the customer.
	DropOffLocation       *string                `json:"DropOffLocation,omitempty"`
	PreferredDeliveryTime *PreferredDeliveryTime `json:"PreferredDeliveryTime,omitempty"`
	// Enumerated list of miscellaneous delivery attributes associated with the shipping address.
	OtherAttributes []OtherDeliveryAttributes `json:"OtherAttributes,omitempty"`
	// Building instructions, nearby landmark or navigation instructions.
	AddressInstructions *string `json:"AddressInstructions,omitempty"`
}

// NewDeliveryPreferences instantiates a new DeliveryPreferences object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeliveryPreferences() *DeliveryPreferences {
	this := DeliveryPreferences{}
	return &this
}

// NewDeliveryPreferencesWithDefaults instantiates a new DeliveryPreferences object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeliveryPreferencesWithDefaults() *DeliveryPreferences {
	this := DeliveryPreferences{}
	return &this
}

// GetDropOffLocation returns the DropOffLocation field value if set, zero value otherwise.
func (o *DeliveryPreferences) GetDropOffLocation() string {
	if o == nil || IsNil(o.DropOffLocation) {
		var ret string
		return ret
	}
	return *o.DropOffLocation
}

// GetDropOffLocationOk returns a tuple with the DropOffLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryPreferences) GetDropOffLocationOk() (*string, bool) {
	if o == nil || IsNil(o.DropOffLocation) {
		return nil, false
	}
	return o.DropOffLocation, true
}

// HasDropOffLocation returns a boolean if a field has been set.
func (o *DeliveryPreferences) HasDropOffLocation() bool {
	if o != nil && !IsNil(o.DropOffLocation) {
		return true
	}

	return false
}

// SetDropOffLocation gets a reference to the given string and assigns it to the DropOffLocation field.
func (o *DeliveryPreferences) SetDropOffLocation(v string) {
	o.DropOffLocation = &v
}

// GetPreferredDeliveryTime returns the PreferredDeliveryTime field value if set, zero value otherwise.
func (o *DeliveryPreferences) GetPreferredDeliveryTime() PreferredDeliveryTime {
	if o == nil || IsNil(o.PreferredDeliveryTime) {
		var ret PreferredDeliveryTime
		return ret
	}
	return *o.PreferredDeliveryTime
}

// GetPreferredDeliveryTimeOk returns a tuple with the PreferredDeliveryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryPreferences) GetPreferredDeliveryTimeOk() (*PreferredDeliveryTime, bool) {
	if o == nil || IsNil(o.PreferredDeliveryTime) {
		return nil, false
	}
	return o.PreferredDeliveryTime, true
}

// HasPreferredDeliveryTime returns a boolean if a field has been set.
func (o *DeliveryPreferences) HasPreferredDeliveryTime() bool {
	if o != nil && !IsNil(o.PreferredDeliveryTime) {
		return true
	}

	return false
}

// SetPreferredDeliveryTime gets a reference to the given PreferredDeliveryTime and assigns it to the PreferredDeliveryTime field.
func (o *DeliveryPreferences) SetPreferredDeliveryTime(v PreferredDeliveryTime) {
	o.PreferredDeliveryTime = &v
}

// GetOtherAttributes returns the OtherAttributes field value if set, zero value otherwise.
func (o *DeliveryPreferences) GetOtherAttributes() []OtherDeliveryAttributes {
	if o == nil || IsNil(o.OtherAttributes) {
		var ret []OtherDeliveryAttributes
		return ret
	}
	return o.OtherAttributes
}

// GetOtherAttributesOk returns a tuple with the OtherAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryPreferences) GetOtherAttributesOk() ([]OtherDeliveryAttributes, bool) {
	if o == nil || IsNil(o.OtherAttributes) {
		return nil, false
	}
	return o.OtherAttributes, true
}

// HasOtherAttributes returns a boolean if a field has been set.
func (o *DeliveryPreferences) HasOtherAttributes() bool {
	if o != nil && !IsNil(o.OtherAttributes) {
		return true
	}

	return false
}

// SetOtherAttributes gets a reference to the given []OtherDeliveryAttributes and assigns it to the OtherAttributes field.
func (o *DeliveryPreferences) SetOtherAttributes(v []OtherDeliveryAttributes) {
	o.OtherAttributes = v
}

// GetAddressInstructions returns the AddressInstructions field value if set, zero value otherwise.
func (o *DeliveryPreferences) GetAddressInstructions() string {
	if o == nil || IsNil(o.AddressInstructions) {
		var ret string
		return ret
	}
	return *o.AddressInstructions
}

// GetAddressInstructionsOk returns a tuple with the AddressInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryPreferences) GetAddressInstructionsOk() (*string, bool) {
	if o == nil || IsNil(o.AddressInstructions) {
		return nil, false
	}
	return o.AddressInstructions, true
}

// HasAddressInstructions returns a boolean if a field has been set.
func (o *DeliveryPreferences) HasAddressInstructions() bool {
	if o != nil && !IsNil(o.AddressInstructions) {
		return true
	}

	return false
}

// SetAddressInstructions gets a reference to the given string and assigns it to the AddressInstructions field.
func (o *DeliveryPreferences) SetAddressInstructions(v string) {
	o.AddressInstructions = &v
}

func (o DeliveryPreferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DropOffLocation) {
		toSerialize["DropOffLocation"] = o.DropOffLocation
	}
	if !IsNil(o.PreferredDeliveryTime) {
		toSerialize["PreferredDeliveryTime"] = o.PreferredDeliveryTime
	}
	if !IsNil(o.OtherAttributes) {
		toSerialize["OtherAttributes"] = o.OtherAttributes
	}
	if !IsNil(o.AddressInstructions) {
		toSerialize["AddressInstructions"] = o.AddressInstructions
	}
	return toSerialize, nil
}

type NullableDeliveryPreferences struct {
	value *DeliveryPreferences
	isSet bool
}

func (v NullableDeliveryPreferences) Get() *DeliveryPreferences {
	return v.value
}

func (v *NullableDeliveryPreferences) Set(val *DeliveryPreferences) {
	v.value = val
	v.isSet = true
}

func (v NullableDeliveryPreferences) IsSet() bool {
	return v.isSet
}

func (v *NullableDeliveryPreferences) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeliveryPreferences(val *DeliveryPreferences) *NullableDeliveryPreferences {
	return &NullableDeliveryPreferences{value: val, isSet: true}
}

func (v NullableDeliveryPreferences) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableDeliveryPreferences) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
