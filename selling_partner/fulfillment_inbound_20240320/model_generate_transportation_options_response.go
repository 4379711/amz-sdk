package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the GenerateTransportationOptionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateTransportationOptionsResponse{}

// GenerateTransportationOptionsResponse The `generateTransportationOptions` response.
type GenerateTransportationOptionsResponse struct {
	// UUID for the given operation.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _GenerateTransportationOptionsResponse GenerateTransportationOptionsResponse

// NewGenerateTransportationOptionsResponse instantiates a new GenerateTransportationOptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateTransportationOptionsResponse(operationId string) *GenerateTransportationOptionsResponse {
	this := GenerateTransportationOptionsResponse{}
	this.OperationId = operationId
	return &this
}

// NewGenerateTransportationOptionsResponseWithDefaults instantiates a new GenerateTransportationOptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateTransportationOptionsResponseWithDefaults() *GenerateTransportationOptionsResponse {
	this := GenerateTransportationOptionsResponse{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *GenerateTransportationOptionsResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *GenerateTransportationOptionsResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *GenerateTransportationOptionsResponse) SetOperationId(v string) {
	o.OperationId = v
}

func (o GenerateTransportationOptionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationId"] = o.OperationId
	return toSerialize, nil
}

type NullableGenerateTransportationOptionsResponse struct {
	value *GenerateTransportationOptionsResponse
	isSet bool
}

func (v NullableGenerateTransportationOptionsResponse) Get() *GenerateTransportationOptionsResponse {
	return v.value
}

func (v *NullableGenerateTransportationOptionsResponse) Set(val *GenerateTransportationOptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateTransportationOptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateTransportationOptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateTransportationOptionsResponse(val *GenerateTransportationOptionsResponse) *NullableGenerateTransportationOptionsResponse {
	return &NullableGenerateTransportationOptionsResponse{value: val, isSet: true}
}

func (v NullableGenerateTransportationOptionsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGenerateTransportationOptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
