package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the GenerateDeliveryWindowOptionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateDeliveryWindowOptionsResponse{}

// GenerateDeliveryWindowOptionsResponse The `generateDeliveryWindowOptions` response.
type GenerateDeliveryWindowOptionsResponse struct {
	// UUID for the given operation.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _GenerateDeliveryWindowOptionsResponse GenerateDeliveryWindowOptionsResponse

// NewGenerateDeliveryWindowOptionsResponse instantiates a new GenerateDeliveryWindowOptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateDeliveryWindowOptionsResponse(operationId string) *GenerateDeliveryWindowOptionsResponse {
	this := GenerateDeliveryWindowOptionsResponse{}
	this.OperationId = operationId
	return &this
}

// NewGenerateDeliveryWindowOptionsResponseWithDefaults instantiates a new GenerateDeliveryWindowOptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateDeliveryWindowOptionsResponseWithDefaults() *GenerateDeliveryWindowOptionsResponse {
	this := GenerateDeliveryWindowOptionsResponse{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *GenerateDeliveryWindowOptionsResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *GenerateDeliveryWindowOptionsResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *GenerateDeliveryWindowOptionsResponse) SetOperationId(v string) {
	o.OperationId = v
}

func (o GenerateDeliveryWindowOptionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationId"] = o.OperationId
	return toSerialize, nil
}

type NullableGenerateDeliveryWindowOptionsResponse struct {
	value *GenerateDeliveryWindowOptionsResponse
	isSet bool
}

func (v NullableGenerateDeliveryWindowOptionsResponse) Get() *GenerateDeliveryWindowOptionsResponse {
	return v.value
}

func (v *NullableGenerateDeliveryWindowOptionsResponse) Set(val *GenerateDeliveryWindowOptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateDeliveryWindowOptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateDeliveryWindowOptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateDeliveryWindowOptionsResponse(val *GenerateDeliveryWindowOptionsResponse) *NullableGenerateDeliveryWindowOptionsResponse {
	return &NullableGenerateDeliveryWindowOptionsResponse{value: val, isSet: true}
}

func (v NullableGenerateDeliveryWindowOptionsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGenerateDeliveryWindowOptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
