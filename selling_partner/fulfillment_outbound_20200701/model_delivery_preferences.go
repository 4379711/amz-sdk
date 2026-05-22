package fulfillment_outbound_20200701

import (
	"github.com/bytedance/sonic"
)

// checks if the DeliveryPreferences type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeliveryPreferences{}

// DeliveryPreferences The delivery preferences applied to the destination address. These preferences are applied when possible and are best effort. This feature is currently supported only in the JP marketplace and not applicable for other marketplaces. For eligible orders, the default delivery preference will be to deliver the package unattended at the front door, unless you specify otherwise.
type DeliveryPreferences struct {
	// Additional delivery instructions. For example, this could be instructions on how to enter a building, nearby landmark or navigation instructions, 'Beware of dogs', etc.
	DeliveryInstructions *string          `json:"deliveryInstructions,omitempty"`
	DropOffLocation      *DropOffLocation `json:"dropOffLocation,omitempty"`
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

// GetDeliveryInstructions returns the DeliveryInstructions field value if set, zero value otherwise.
func (o *DeliveryPreferences) GetDeliveryInstructions() string {
	if o == nil || IsNil(o.DeliveryInstructions) {
		var ret string
		return ret
	}
	return *o.DeliveryInstructions
}

// GetDeliveryInstructionsOk returns a tuple with the DeliveryInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryPreferences) GetDeliveryInstructionsOk() (*string, bool) {
	if o == nil || IsNil(o.DeliveryInstructions) {
		return nil, false
	}
	return o.DeliveryInstructions, true
}

// HasDeliveryInstructions returns a boolean if a field has been set.
func (o *DeliveryPreferences) HasDeliveryInstructions() bool {
	if o != nil && !IsNil(o.DeliveryInstructions) {
		return true
	}

	return false
}

// SetDeliveryInstructions gets a reference to the given string and assigns it to the DeliveryInstructions field.
func (o *DeliveryPreferences) SetDeliveryInstructions(v string) {
	o.DeliveryInstructions = &v
}

// GetDropOffLocation returns the DropOffLocation field value if set, zero value otherwise.
func (o *DeliveryPreferences) GetDropOffLocation() DropOffLocation {
	if o == nil || IsNil(o.DropOffLocation) {
		var ret DropOffLocation
		return ret
	}
	return *o.DropOffLocation
}

// GetDropOffLocationOk returns a tuple with the DropOffLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeliveryPreferences) GetDropOffLocationOk() (*DropOffLocation, bool) {
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

// SetDropOffLocation gets a reference to the given DropOffLocation and assigns it to the DropOffLocation field.
func (o *DeliveryPreferences) SetDropOffLocation(v DropOffLocation) {
	o.DropOffLocation = &v
}

func (o DeliveryPreferences) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeliveryInstructions) {
		toSerialize["deliveryInstructions"] = o.DeliveryInstructions
	}
	if !IsNil(o.DropOffLocation) {
		toSerialize["dropOffLocation"] = o.DropOffLocation
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
