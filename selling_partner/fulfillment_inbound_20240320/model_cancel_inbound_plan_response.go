package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the CancelInboundPlanResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelInboundPlanResponse{}

// CancelInboundPlanResponse The `cancelInboundPlan` response.
type CancelInboundPlanResponse struct {
	// UUID for the given operation.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _CancelInboundPlanResponse CancelInboundPlanResponse

// NewCancelInboundPlanResponse instantiates a new CancelInboundPlanResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelInboundPlanResponse(operationId string) *CancelInboundPlanResponse {
	this := CancelInboundPlanResponse{}
	this.OperationId = operationId
	return &this
}

// NewCancelInboundPlanResponseWithDefaults instantiates a new CancelInboundPlanResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelInboundPlanResponseWithDefaults() *CancelInboundPlanResponse {
	this := CancelInboundPlanResponse{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *CancelInboundPlanResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *CancelInboundPlanResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *CancelInboundPlanResponse) SetOperationId(v string) {
	o.OperationId = v
}

func (o CancelInboundPlanResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationId"] = o.OperationId
	return toSerialize, nil
}

type NullableCancelInboundPlanResponse struct {
	value *CancelInboundPlanResponse
	isSet bool
}

func (v NullableCancelInboundPlanResponse) Get() *CancelInboundPlanResponse {
	return v.value
}

func (v *NullableCancelInboundPlanResponse) Set(val *CancelInboundPlanResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelInboundPlanResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelInboundPlanResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelInboundPlanResponse(val *CancelInboundPlanResponse) *NullableCancelInboundPlanResponse {
	return &NullableCancelInboundPlanResponse{value: val, isSet: true}
}

func (v NullableCancelInboundPlanResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableCancelInboundPlanResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
