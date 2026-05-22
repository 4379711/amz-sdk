package orders_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the FulfillmentInstruction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FulfillmentInstruction{}

// FulfillmentInstruction Contains the instructions about the fulfillment, such as the location from where you want the order filled.
type FulfillmentInstruction struct {
	// The `sourceId` of the location from where you want the order fulfilled.
	FulfillmentSupplySourceId *string `json:"FulfillmentSupplySourceId,omitempty"`
}

// NewFulfillmentInstruction instantiates a new FulfillmentInstruction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFulfillmentInstruction() *FulfillmentInstruction {
	this := FulfillmentInstruction{}
	return &this
}

// NewFulfillmentInstructionWithDefaults instantiates a new FulfillmentInstruction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFulfillmentInstructionWithDefaults() *FulfillmentInstruction {
	this := FulfillmentInstruction{}
	return &this
}

// GetFulfillmentSupplySourceId returns the FulfillmentSupplySourceId field value if set, zero value otherwise.
func (o *FulfillmentInstruction) GetFulfillmentSupplySourceId() string {
	if o == nil || IsNil(o.FulfillmentSupplySourceId) {
		var ret string
		return ret
	}
	return *o.FulfillmentSupplySourceId
}

// GetFulfillmentSupplySourceIdOk returns a tuple with the FulfillmentSupplySourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FulfillmentInstruction) GetFulfillmentSupplySourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.FulfillmentSupplySourceId) {
		return nil, false
	}
	return o.FulfillmentSupplySourceId, true
}

// HasFulfillmentSupplySourceId returns a boolean if a field has been set.
func (o *FulfillmentInstruction) HasFulfillmentSupplySourceId() bool {
	if o != nil && !IsNil(o.FulfillmentSupplySourceId) {
		return true
	}

	return false
}

// SetFulfillmentSupplySourceId gets a reference to the given string and assigns it to the FulfillmentSupplySourceId field.
func (o *FulfillmentInstruction) SetFulfillmentSupplySourceId(v string) {
	o.FulfillmentSupplySourceId = &v
}

func (o FulfillmentInstruction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FulfillmentSupplySourceId) {
		toSerialize["FulfillmentSupplySourceId"] = o.FulfillmentSupplySourceId
	}
	return toSerialize, nil
}

type NullableFulfillmentInstruction struct {
	value *FulfillmentInstruction
	isSet bool
}

func (v NullableFulfillmentInstruction) Get() *FulfillmentInstruction {
	return v.value
}

func (v *NullableFulfillmentInstruction) Set(val *FulfillmentInstruction) {
	v.value = val
	v.isSet = true
}

func (v NullableFulfillmentInstruction) IsSet() bool {
	return v.isSet
}

func (v *NullableFulfillmentInstruction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFulfillmentInstruction(val *FulfillmentInstruction) *NullableFulfillmentInstruction {
	return &NullableFulfillmentInstruction{value: val, isSet: true}
}

func (v NullableFulfillmentInstruction) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableFulfillmentInstruction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
