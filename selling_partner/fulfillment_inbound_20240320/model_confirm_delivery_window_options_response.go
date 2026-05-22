package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the ConfirmDeliveryWindowOptionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConfirmDeliveryWindowOptionsResponse{}

// ConfirmDeliveryWindowOptionsResponse The `confirmDeliveryWindowOptions` response.
type ConfirmDeliveryWindowOptionsResponse struct {
	// UUID for the given operation.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _ConfirmDeliveryWindowOptionsResponse ConfirmDeliveryWindowOptionsResponse

// NewConfirmDeliveryWindowOptionsResponse instantiates a new ConfirmDeliveryWindowOptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConfirmDeliveryWindowOptionsResponse(operationId string) *ConfirmDeliveryWindowOptionsResponse {
	this := ConfirmDeliveryWindowOptionsResponse{}
	this.OperationId = operationId
	return &this
}

// NewConfirmDeliveryWindowOptionsResponseWithDefaults instantiates a new ConfirmDeliveryWindowOptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfirmDeliveryWindowOptionsResponseWithDefaults() *ConfirmDeliveryWindowOptionsResponse {
	this := ConfirmDeliveryWindowOptionsResponse{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *ConfirmDeliveryWindowOptionsResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *ConfirmDeliveryWindowOptionsResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *ConfirmDeliveryWindowOptionsResponse) SetOperationId(v string) {
	o.OperationId = v
}

func (o ConfirmDeliveryWindowOptionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationId"] = o.OperationId
	return toSerialize, nil
}

type NullableConfirmDeliveryWindowOptionsResponse struct {
	value *ConfirmDeliveryWindowOptionsResponse
	isSet bool
}

func (v NullableConfirmDeliveryWindowOptionsResponse) Get() *ConfirmDeliveryWindowOptionsResponse {
	return v.value
}

func (v *NullableConfirmDeliveryWindowOptionsResponse) Set(val *ConfirmDeliveryWindowOptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableConfirmDeliveryWindowOptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableConfirmDeliveryWindowOptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfirmDeliveryWindowOptionsResponse(val *ConfirmDeliveryWindowOptionsResponse) *NullableConfirmDeliveryWindowOptionsResponse {
	return &NullableConfirmDeliveryWindowOptionsResponse{value: val, isSet: true}
}

func (v NullableConfirmDeliveryWindowOptionsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableConfirmDeliveryWindowOptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
