package vendor_direct_fulfillment_orders_20211228

import (
	"github.com/bytedance/sonic"
)

// checks if the SubmitAcknowledgementRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubmitAcknowledgementRequest{}

// SubmitAcknowledgementRequest The request schema for the submitAcknowledgement operation.
type SubmitAcknowledgementRequest struct {
	// A list of one or more purchase orders.
	OrderAcknowledgements []OrderAcknowledgementItem `json:"orderAcknowledgements,omitempty"`
}

// NewSubmitAcknowledgementRequest instantiates a new SubmitAcknowledgementRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubmitAcknowledgementRequest() *SubmitAcknowledgementRequest {
	this := SubmitAcknowledgementRequest{}
	return &this
}

// NewSubmitAcknowledgementRequestWithDefaults instantiates a new SubmitAcknowledgementRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubmitAcknowledgementRequestWithDefaults() *SubmitAcknowledgementRequest {
	this := SubmitAcknowledgementRequest{}
	return &this
}

// GetOrderAcknowledgements returns the OrderAcknowledgements field value if set, zero value otherwise.
func (o *SubmitAcknowledgementRequest) GetOrderAcknowledgements() []OrderAcknowledgementItem {
	if o == nil || IsNil(o.OrderAcknowledgements) {
		var ret []OrderAcknowledgementItem
		return ret
	}
	return o.OrderAcknowledgements
}

// GetOrderAcknowledgementsOk returns a tuple with the OrderAcknowledgements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubmitAcknowledgementRequest) GetOrderAcknowledgementsOk() ([]OrderAcknowledgementItem, bool) {
	if o == nil || IsNil(o.OrderAcknowledgements) {
		return nil, false
	}
	return o.OrderAcknowledgements, true
}

// HasOrderAcknowledgements returns a boolean if a field has been set.
func (o *SubmitAcknowledgementRequest) HasOrderAcknowledgements() bool {
	if o != nil && !IsNil(o.OrderAcknowledgements) {
		return true
	}

	return false
}

// SetOrderAcknowledgements gets a reference to the given []OrderAcknowledgementItem and assigns it to the OrderAcknowledgements field.
func (o *SubmitAcknowledgementRequest) SetOrderAcknowledgements(v []OrderAcknowledgementItem) {
	o.OrderAcknowledgements = v
}

func (o SubmitAcknowledgementRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OrderAcknowledgements) {
		toSerialize["orderAcknowledgements"] = o.OrderAcknowledgements
	}
	return toSerialize, nil
}

type NullableSubmitAcknowledgementRequest struct {
	value *SubmitAcknowledgementRequest
	isSet bool
}

func (v NullableSubmitAcknowledgementRequest) Get() *SubmitAcknowledgementRequest {
	return v.value
}

func (v *NullableSubmitAcknowledgementRequest) Set(val *SubmitAcknowledgementRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubmitAcknowledgementRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubmitAcknowledgementRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubmitAcknowledgementRequest(val *SubmitAcknowledgementRequest) *NullableSubmitAcknowledgementRequest {
	return &NullableSubmitAcknowledgementRequest{value: val, isSet: true}
}

func (v NullableSubmitAcknowledgementRequest) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableSubmitAcknowledgementRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
