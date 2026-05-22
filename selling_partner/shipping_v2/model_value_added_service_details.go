package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the ValueAddedServiceDetails type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValueAddedServiceDetails{}

// ValueAddedServiceDetails A collection of supported value-added services.
type ValueAddedServiceDetails struct {
	CollectOnDelivery *CollectOnDelivery `json:"collectOnDelivery,omitempty"`
}

// NewValueAddedServiceDetails instantiates a new ValueAddedServiceDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueAddedServiceDetails() *ValueAddedServiceDetails {
	this := ValueAddedServiceDetails{}
	return &this
}

// NewValueAddedServiceDetailsWithDefaults instantiates a new ValueAddedServiceDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueAddedServiceDetailsWithDefaults() *ValueAddedServiceDetails {
	this := ValueAddedServiceDetails{}
	return &this
}

// GetCollectOnDelivery returns the CollectOnDelivery field value if set, zero value otherwise.
func (o *ValueAddedServiceDetails) GetCollectOnDelivery() CollectOnDelivery {
	if o == nil || IsNil(o.CollectOnDelivery) {
		var ret CollectOnDelivery
		return ret
	}
	return *o.CollectOnDelivery
}

// GetCollectOnDeliveryOk returns a tuple with the CollectOnDelivery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValueAddedServiceDetails) GetCollectOnDeliveryOk() (*CollectOnDelivery, bool) {
	if o == nil || IsNil(o.CollectOnDelivery) {
		return nil, false
	}
	return o.CollectOnDelivery, true
}

// HasCollectOnDelivery returns a boolean if a field has been set.
func (o *ValueAddedServiceDetails) HasCollectOnDelivery() bool {
	if o != nil && !IsNil(o.CollectOnDelivery) {
		return true
	}

	return false
}

// SetCollectOnDelivery gets a reference to the given CollectOnDelivery and assigns it to the CollectOnDelivery field.
func (o *ValueAddedServiceDetails) SetCollectOnDelivery(v CollectOnDelivery) {
	o.CollectOnDelivery = &v
}

func (o ValueAddedServiceDetails) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CollectOnDelivery) {
		toSerialize["collectOnDelivery"] = o.CollectOnDelivery
	}
	return toSerialize, nil
}

type NullableValueAddedServiceDetails struct {
	value *ValueAddedServiceDetails
	isSet bool
}

func (v NullableValueAddedServiceDetails) Get() *ValueAddedServiceDetails {
	return v.value
}

func (v *NullableValueAddedServiceDetails) Set(val *ValueAddedServiceDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableValueAddedServiceDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableValueAddedServiceDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueAddedServiceDetails(val *ValueAddedServiceDetails) *NullableValueAddedServiceDetails {
	return &NullableValueAddedServiceDetails{value: val, isSet: true}
}

func (v NullableValueAddedServiceDetails) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableValueAddedServiceDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
