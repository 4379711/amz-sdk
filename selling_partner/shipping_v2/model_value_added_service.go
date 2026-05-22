package shipping_v2

import (
	"github.com/bytedance/sonic"
)

// checks if the ValueAddedService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValueAddedService{}

// ValueAddedService A value-added service available for purchase with a shipment service offering.
type ValueAddedService struct {
	// The identifier for the value-added service.
	Id string `json:"id"`
	// The name of the value-added service.
	Name string   `json:"name"`
	Cost Currency `json:"cost"`
}

type _ValueAddedService ValueAddedService

// NewValueAddedService instantiates a new ValueAddedService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValueAddedService(id string, name string, cost Currency) *ValueAddedService {
	this := ValueAddedService{}
	this.Id = id
	this.Name = name
	this.Cost = cost
	return &this
}

// NewValueAddedServiceWithDefaults instantiates a new ValueAddedService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValueAddedServiceWithDefaults() *ValueAddedService {
	this := ValueAddedService{}
	return &this
}

// GetId returns the Id field value
func (o *ValueAddedService) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ValueAddedService) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ValueAddedService) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ValueAddedService) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ValueAddedService) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ValueAddedService) SetName(v string) {
	o.Name = v
}

// GetCost returns the Cost field value
func (o *ValueAddedService) GetCost() Currency {
	if o == nil {
		var ret Currency
		return ret
	}

	return o.Cost
}

// GetCostOk returns a tuple with the Cost field value
// and a boolean to check if the value has been set.
func (o *ValueAddedService) GetCostOk() (*Currency, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cost, true
}

// SetCost sets field value
func (o *ValueAddedService) SetCost(v Currency) {
	o.Cost = v
}

func (o ValueAddedService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["cost"] = o.Cost
	return toSerialize, nil
}

type NullableValueAddedService struct {
	value *ValueAddedService
	isSet bool
}

func (v NullableValueAddedService) Get() *ValueAddedService {
	return v.value
}

func (v *NullableValueAddedService) Set(val *ValueAddedService) {
	v.value = val
	v.isSet = true
}

func (v NullableValueAddedService) IsSet() bool {
	return v.isSet
}

func (v *NullableValueAddedService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValueAddedService(val *ValueAddedService) *NullableValueAddedService {
	return &NullableValueAddedService{value: val, isSet: true}
}

func (v NullableValueAddedService) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableValueAddedService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
