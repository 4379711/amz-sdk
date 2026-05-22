package merchant_fulfillment_v0

import (
	"github.com/bytedance/sonic"
)

// checks if the TemporarilyUnavailableCarrier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemporarilyUnavailableCarrier{}

// TemporarilyUnavailableCarrier A carrier who is temporarily unavailable, most likely due to a service outage experienced by the carrier.
type TemporarilyUnavailableCarrier struct {
	// The name of the carrier.
	CarrierName string `json:"CarrierName"`
}

type _TemporarilyUnavailableCarrier TemporarilyUnavailableCarrier

// NewTemporarilyUnavailableCarrier instantiates a new TemporarilyUnavailableCarrier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemporarilyUnavailableCarrier(carrierName string) *TemporarilyUnavailableCarrier {
	this := TemporarilyUnavailableCarrier{}
	this.CarrierName = carrierName
	return &this
}

// NewTemporarilyUnavailableCarrierWithDefaults instantiates a new TemporarilyUnavailableCarrier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemporarilyUnavailableCarrierWithDefaults() *TemporarilyUnavailableCarrier {
	this := TemporarilyUnavailableCarrier{}
	return &this
}

// GetCarrierName returns the CarrierName field value
func (o *TemporarilyUnavailableCarrier) GetCarrierName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierName
}

// GetCarrierNameOk returns a tuple with the CarrierName field value
// and a boolean to check if the value has been set.
func (o *TemporarilyUnavailableCarrier) GetCarrierNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierName, true
}

// SetCarrierName sets field value
func (o *TemporarilyUnavailableCarrier) SetCarrierName(v string) {
	o.CarrierName = v
}

func (o TemporarilyUnavailableCarrier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["CarrierName"] = o.CarrierName
	return toSerialize, nil
}

type NullableTemporarilyUnavailableCarrier struct {
	value *TemporarilyUnavailableCarrier
	isSet bool
}

func (v NullableTemporarilyUnavailableCarrier) Get() *TemporarilyUnavailableCarrier {
	return v.value
}

func (v *NullableTemporarilyUnavailableCarrier) Set(val *TemporarilyUnavailableCarrier) {
	v.value = val
	v.isSet = true
}

func (v NullableTemporarilyUnavailableCarrier) IsSet() bool {
	return v.isSet
}

func (v *NullableTemporarilyUnavailableCarrier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemporarilyUnavailableCarrier(val *TemporarilyUnavailableCarrier) *NullableTemporarilyUnavailableCarrier {
	return &NullableTemporarilyUnavailableCarrier{value: val, isSet: true}
}

func (v NullableTemporarilyUnavailableCarrier) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableTemporarilyUnavailableCarrier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
