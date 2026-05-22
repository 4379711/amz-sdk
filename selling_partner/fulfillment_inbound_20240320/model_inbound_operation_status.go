package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the InboundOperationStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InboundOperationStatus{}

// InboundOperationStatus GetInboundOperationStatus response.
type InboundOperationStatus struct {
	// The name of the operation in the asynchronous API call.
	Operation string `json:"operation"`
	// The operation ID returned by the asynchronous API call.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
	// The problems in the processing of the asynchronous operation.
	OperationProblems []OperationProblem `json:"operationProblems"`
	OperationStatus   OperationStatus    `json:"operationStatus"`
}

type _InboundOperationStatus InboundOperationStatus

// NewInboundOperationStatus instantiates a new InboundOperationStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInboundOperationStatus(operation string, operationId string, operationProblems []OperationProblem, operationStatus OperationStatus) *InboundOperationStatus {
	this := InboundOperationStatus{}
	this.Operation = operation
	this.OperationId = operationId
	this.OperationProblems = operationProblems
	this.OperationStatus = operationStatus
	return &this
}

// NewInboundOperationStatusWithDefaults instantiates a new InboundOperationStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInboundOperationStatusWithDefaults() *InboundOperationStatus {
	this := InboundOperationStatus{}
	return &this
}

// GetOperation returns the Operation field value
func (o *InboundOperationStatus) GetOperation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Operation
}

// GetOperationOk returns a tuple with the Operation field value
// and a boolean to check if the value has been set.
func (o *InboundOperationStatus) GetOperationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operation, true
}

// SetOperation sets field value
func (o *InboundOperationStatus) SetOperation(v string) {
	o.Operation = v
}

// GetOperationId returns the OperationId field value
func (o *InboundOperationStatus) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *InboundOperationStatus) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *InboundOperationStatus) SetOperationId(v string) {
	o.OperationId = v
}

// GetOperationProblems returns the OperationProblems field value
func (o *InboundOperationStatus) GetOperationProblems() []OperationProblem {
	if o == nil {
		var ret []OperationProblem
		return ret
	}

	return o.OperationProblems
}

// GetOperationProblemsOk returns a tuple with the OperationProblems field value
// and a boolean to check if the value has been set.
func (o *InboundOperationStatus) GetOperationProblemsOk() ([]OperationProblem, bool) {
	if o == nil {
		return nil, false
	}
	return o.OperationProblems, true
}

// SetOperationProblems sets field value
func (o *InboundOperationStatus) SetOperationProblems(v []OperationProblem) {
	o.OperationProblems = v
}

// GetOperationStatus returns the OperationStatus field value
func (o *InboundOperationStatus) GetOperationStatus() OperationStatus {
	if o == nil {
		var ret OperationStatus
		return ret
	}

	return o.OperationStatus
}

// GetOperationStatusOk returns a tuple with the OperationStatus field value
// and a boolean to check if the value has been set.
func (o *InboundOperationStatus) GetOperationStatusOk() (*OperationStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationStatus, true
}

// SetOperationStatus sets field value
func (o *InboundOperationStatus) SetOperationStatus(v OperationStatus) {
	o.OperationStatus = v
}

func (o InboundOperationStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operation"] = o.Operation
	toSerialize["operationId"] = o.OperationId
	toSerialize["operationProblems"] = o.OperationProblems
	toSerialize["operationStatus"] = o.OperationStatus
	return toSerialize, nil
}

type NullableInboundOperationStatus struct {
	value *InboundOperationStatus
	isSet bool
}

func (v NullableInboundOperationStatus) Get() *InboundOperationStatus {
	return v.value
}

func (v *NullableInboundOperationStatus) Set(val *InboundOperationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableInboundOperationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableInboundOperationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInboundOperationStatus(val *InboundOperationStatus) *NullableInboundOperationStatus {
	return &NullableInboundOperationStatus{value: val, isSet: true}
}

func (v NullableInboundOperationStatus) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableInboundOperationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
