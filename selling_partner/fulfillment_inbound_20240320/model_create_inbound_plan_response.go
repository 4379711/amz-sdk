package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the CreateInboundPlanResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInboundPlanResponse{}

// CreateInboundPlanResponse The `createInboundPlan` response.
type CreateInboundPlanResponse struct {
	// Identifier of an inbound plan.
	InboundPlanId string `json:"inboundPlanId" validate:"regexp=^[a-zA-Z0-9-]*$"`
	// UUID for the given operation.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _CreateInboundPlanResponse CreateInboundPlanResponse

// NewCreateInboundPlanResponse instantiates a new CreateInboundPlanResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInboundPlanResponse(inboundPlanId string, operationId string) *CreateInboundPlanResponse {
	this := CreateInboundPlanResponse{}
	this.InboundPlanId = inboundPlanId
	this.OperationId = operationId
	return &this
}

// NewCreateInboundPlanResponseWithDefaults instantiates a new CreateInboundPlanResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInboundPlanResponseWithDefaults() *CreateInboundPlanResponse {
	this := CreateInboundPlanResponse{}
	return &this
}

// GetInboundPlanId returns the InboundPlanId field value
func (o *CreateInboundPlanResponse) GetInboundPlanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InboundPlanId
}

// GetInboundPlanIdOk returns a tuple with the InboundPlanId field value
// and a boolean to check if the value has been set.
func (o *CreateInboundPlanResponse) GetInboundPlanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InboundPlanId, true
}

// SetInboundPlanId sets field value
func (o *CreateInboundPlanResponse) SetInboundPlanId(v string) {
	o.InboundPlanId = v
}

// GetOperationId returns the OperationId field value
func (o *CreateInboundPlanResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *CreateInboundPlanResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *CreateInboundPlanResponse) SetOperationId(v string) {
	o.OperationId = v
}

func (o CreateInboundPlanResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inboundPlanId"] = o.InboundPlanId
	toSerialize["operationId"] = o.OperationId
	return toSerialize, nil
}

type NullableCreateInboundPlanResponse struct {
	value *CreateInboundPlanResponse
	isSet bool
}

func (v NullableCreateInboundPlanResponse) Get() *CreateInboundPlanResponse {
	return v.value
}

func (v *NullableCreateInboundPlanResponse) Set(val *CreateInboundPlanResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInboundPlanResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInboundPlanResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInboundPlanResponse(val *CreateInboundPlanResponse) *NullableCreateInboundPlanResponse {
	return &NullableCreateInboundPlanResponse{value: val, isSet: true}
}

func (v NullableCreateInboundPlanResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCreateInboundPlanResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
