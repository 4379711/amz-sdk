package vendor_direct_fulfillment_sandbox_data_20211028

import (
	"github.com/bytedance/sonic"
)

// checks if the GenerateOrderScenarioRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateOrderScenarioRequest{}

// GenerateOrderScenarioRequest The request body for the generateOrderScenarios operation.
type GenerateOrderScenarioRequest struct {
	// The list of test orders requested as indicated by party identifiers.
	Orders []OrderScenarioRequest `json:"orders,omitempty"`
}

// NewGenerateOrderScenarioRequest instantiates a new GenerateOrderScenarioRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateOrderScenarioRequest() *GenerateOrderScenarioRequest {
	this := GenerateOrderScenarioRequest{}
	return &this
}

// NewGenerateOrderScenarioRequestWithDefaults instantiates a new GenerateOrderScenarioRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateOrderScenarioRequestWithDefaults() *GenerateOrderScenarioRequest {
	this := GenerateOrderScenarioRequest{}
	return &this
}

// GetOrders returns the Orders field value if set, zero value otherwise.
func (o *GenerateOrderScenarioRequest) GetOrders() []OrderScenarioRequest {
	if o == nil || IsNil(o.Orders) {
		var ret []OrderScenarioRequest
		return ret
	}
	return o.Orders
}

// GetOrdersOk returns a tuple with the Orders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenerateOrderScenarioRequest) GetOrdersOk() ([]OrderScenarioRequest, bool) {
	if o == nil || IsNil(o.Orders) {
		return nil, false
	}
	return o.Orders, true
}

// HasOrders returns a boolean if a field has been set.
func (o *GenerateOrderScenarioRequest) HasOrders() bool {
	if o != nil && !IsNil(o.Orders) {
		return true
	}

	return false
}

// SetOrders gets a reference to the given []OrderScenarioRequest and assigns it to the Orders field.
func (o *GenerateOrderScenarioRequest) SetOrders(v []OrderScenarioRequest) {
	o.Orders = v
}

func (o GenerateOrderScenarioRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Orders) {
		toSerialize["orders"] = o.Orders
	}
	return toSerialize, nil
}

type NullableGenerateOrderScenarioRequest struct {
	value *GenerateOrderScenarioRequest
	isSet bool
}

func (v NullableGenerateOrderScenarioRequest) Get() *GenerateOrderScenarioRequest {
	return v.value
}

func (v *NullableGenerateOrderScenarioRequest) Set(val *GenerateOrderScenarioRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateOrderScenarioRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateOrderScenarioRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateOrderScenarioRequest(val *GenerateOrderScenarioRequest) *NullableGenerateOrderScenarioRequest {
	return &NullableGenerateOrderScenarioRequest{value: val, isSet: true}
}

func (v NullableGenerateOrderScenarioRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGenerateOrderScenarioRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
