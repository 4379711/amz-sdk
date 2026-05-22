package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the OneClickShipmentValueAddedService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OneClickShipmentValueAddedService{}

// OneClickShipmentValueAddedService A value-added service to be applied to a shipping service purchase.
type OneClickShipmentValueAddedService struct {
	// The identifier of the selected value-added service.
	Id     string    `json:"id"`
	Amount *Currency `json:"amount,omitempty"`
}

type _OneClickShipmentValueAddedService OneClickShipmentValueAddedService

// NewOneClickShipmentValueAddedService instantiates a new OneClickShipmentValueAddedService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOneClickShipmentValueAddedService(id string) *OneClickShipmentValueAddedService {
	this := OneClickShipmentValueAddedService{}
	this.Id = id
	return &this
}

// NewOneClickShipmentValueAddedServiceWithDefaults instantiates a new OneClickShipmentValueAddedService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOneClickShipmentValueAddedServiceWithDefaults() *OneClickShipmentValueAddedService {
	this := OneClickShipmentValueAddedService{}
	return &this
}

// GetId returns the Id field value
func (o *OneClickShipmentValueAddedService) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OneClickShipmentValueAddedService) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OneClickShipmentValueAddedService) SetId(v string) {
	o.Id = v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *OneClickShipmentValueAddedService) GetAmount() Currency {
	if o == nil || IsNil(o.Amount) {
		var ret Currency
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OneClickShipmentValueAddedService) GetAmountOk() (*Currency, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *OneClickShipmentValueAddedService) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given Currency and assigns it to the Amount field.
func (o *OneClickShipmentValueAddedService) SetAmount(v Currency) {
	o.Amount = &v
}

func (o OneClickShipmentValueAddedService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	return toSerialize, nil
}

type NullableOneClickShipmentValueAddedService struct {
	value *OneClickShipmentValueAddedService
	isSet bool
}

func (v NullableOneClickShipmentValueAddedService) Get() *OneClickShipmentValueAddedService {
	return v.value
}

func (v *NullableOneClickShipmentValueAddedService) Set(val *OneClickShipmentValueAddedService) {
	v.value = val
	v.isSet = true
}

func (v NullableOneClickShipmentValueAddedService) IsSet() bool {
	return v.isSet
}

func (v *NullableOneClickShipmentValueAddedService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOneClickShipmentValueAddedService(val *OneClickShipmentValueAddedService) *NullableOneClickShipmentValueAddedService {
	return &NullableOneClickShipmentValueAddedService{value: val, isSet: true}
}

func (v NullableOneClickShipmentValueAddedService) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableOneClickShipmentValueAddedService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
