package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the CollectOnDelivery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectOnDelivery{}

// CollectOnDelivery The amount to collect on delivery.
type CollectOnDelivery struct {
	Amount Currency `json:"amount"`
}

type _CollectOnDelivery CollectOnDelivery

// NewCollectOnDelivery instantiates a new CollectOnDelivery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectOnDelivery(amount Currency) *CollectOnDelivery {
	this := CollectOnDelivery{}
	this.Amount = amount
	return &this
}

// NewCollectOnDeliveryWithDefaults instantiates a new CollectOnDelivery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectOnDeliveryWithDefaults() *CollectOnDelivery {
	this := CollectOnDelivery{}
	return &this
}

// GetAmount returns the Amount field value
func (o *CollectOnDelivery) GetAmount() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *CollectOnDelivery) GetAmountOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *CollectOnDelivery) SetAmount(v Currency) {
	o.Amount = v
}

func (o CollectOnDelivery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	return toSerialize, nil
}

type NullableCollectOnDelivery struct {
	value *CollectOnDelivery
	isSet bool
}

func (v NullableCollectOnDelivery) Get() *CollectOnDelivery {
	return v.value
}

func (v *NullableCollectOnDelivery) Set(val *CollectOnDelivery) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectOnDelivery) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectOnDelivery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectOnDelivery(val *CollectOnDelivery) *NullableCollectOnDelivery {
	return &NullableCollectOnDelivery{value: val, isSet: true}
}

func (v NullableCollectOnDelivery) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCollectOnDelivery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
