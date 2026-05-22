package vendor_direct_fulfillment_sandbox_data_20211028

import (
	"github.com/bytedance/sonic"
)

// checks if the Scenario type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Scenario{}

// Scenario A scenario test case response returned when the request is successful.
type Scenario struct {
	// An identifier that identifies the type of scenario that user can use for testing.
	ScenarioId string `json:"scenarioId"`
	// A list of orders that can be used by the caller to test each life cycle or scenario.
	Orders []TestOrder `json:"orders"`
}

type _Scenario Scenario

// NewScenario instantiates a new Scenario object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScenario(scenarioId string, orders []TestOrder) *Scenario {
	this := Scenario{}
	this.ScenarioId = scenarioId
	this.Orders = orders
	return &this
}

// NewScenarioWithDefaults instantiates a new Scenario object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScenarioWithDefaults() *Scenario {
	this := Scenario{}
	return &this
}

// GetScenarioId returns the ScenarioId field value
func (o *Scenario) GetScenarioId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScenarioId
}

// GetScenarioIdOk returns a tuple with the ScenarioId field value
// and a boolean to check if the value has been set.
func (o *Scenario) GetScenarioIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScenarioId, true
}

// SetScenarioId sets field value
func (o *Scenario) SetScenarioId(v string) {
	o.ScenarioId = v
}

// GetOrders returns the Orders field value
func (o *Scenario) GetOrders() []TestOrder {
	if o == nil {
		var ret []TestOrder
		return ret
	}

	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value
// and a boolean to check if the value has been set.
func (o *Scenario) GetOrdersOk() ([]TestOrder, bool) {
	if o == nil {
		return nil, false
	}
	return o.Orders, true
}

// SetOrders sets field value
func (o *Scenario) SetOrders(v []TestOrder) {
	o.Orders = v
}

func (o Scenario) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["scenarioId"] = o.ScenarioId
	toSerialize["orders"] = o.Orders
	return toSerialize, nil
}

type NullableScenario struct {
	value *Scenario
	isSet bool
}

func (v NullableScenario) Get() *Scenario {
	return v.value
}

func (v *NullableScenario) Set(val *Scenario) {
	v.value = val
	v.isSet = true
}

func (v NullableScenario) IsSet() bool {
	return v.isSet
}

func (v *NullableScenario) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScenario(val *Scenario) *NullableScenario {
	return &NullableScenario{value: val, isSet: true}
}

func (v NullableScenario) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableScenario) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
