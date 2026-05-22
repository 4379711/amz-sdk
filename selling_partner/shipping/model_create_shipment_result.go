package shipping

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateShipmentResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateShipmentResult{}

// CreateShipmentResult The payload schema for the createShipment operation.
type CreateShipmentResult struct {
	// The unique shipment identifier.
	ShipmentId string `json:"shipmentId"`
	// A list of all the available rates that can be used to send the shipment.
	EligibleRates []Rate `json:"eligibleRates"`
}

type _CreateShipmentResult CreateShipmentResult

// NewCreateShipmentResult instantiates a new CreateShipmentResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShipmentResult(shipmentId string, eligibleRates []Rate) *CreateShipmentResult {
	this := CreateShipmentResult{}
	this.ShipmentId = shipmentId
	this.EligibleRates = eligibleRates
	return &this
}

// NewCreateShipmentResultWithDefaults instantiates a new CreateShipmentResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShipmentResultWithDefaults() *CreateShipmentResult {
	this := CreateShipmentResult{}
	return &this
}

// GetShipmentId returns the ShipmentId field value
func (o *CreateShipmentResult) GetShipmentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value
// and a boolean to check if the value has been set.
func (o *CreateShipmentResult) GetShipmentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentId, true
}

// SetShipmentId sets field value
func (o *CreateShipmentResult) SetShipmentId(v string) {
	o.ShipmentId = v
}

// GetEligibleRates returns the EligibleRates field value
func (o *CreateShipmentResult) GetEligibleRates() []Rate {
	if o == nil {
		var ret []Rate
		return ret
	}

	return o.EligibleRates
}

// GetEligibleRatesOk returns a tuple with the EligibleRates field value
// and a boolean to check if the value has been set.
func (o *CreateShipmentResult) GetEligibleRatesOk() ([]Rate, bool) {
	if o == nil {
		return nil, false
	}
	return o.EligibleRates, true
}

// SetEligibleRates sets field value
func (o *CreateShipmentResult) SetEligibleRates(v []Rate) {
	o.EligibleRates = v
}

func (o CreateShipmentResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipmentId"] = o.ShipmentId
	toSerialize["eligibleRates"] = o.EligibleRates
	return toSerialize, nil
}

type NullableCreateShipmentResult struct {
	value *CreateShipmentResult
	isSet bool
}

func (v NullableCreateShipmentResult) Get() *CreateShipmentResult {
	return v.value
}

func (v *NullableCreateShipmentResult) Set(val *CreateShipmentResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShipmentResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShipmentResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShipmentResult(val *CreateShipmentResult) *NullableCreateShipmentResult {
	return &NullableCreateShipmentResult{value: val, isSet: true}
}

func (v NullableCreateShipmentResult) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateShipmentResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
