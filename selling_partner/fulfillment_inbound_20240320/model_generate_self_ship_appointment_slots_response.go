package fulfillment_inbound_20240320

import (
	"github.com/bytedance/sonic"
)

// checks if the GenerateSelfShipAppointmentSlotsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenerateSelfShipAppointmentSlotsResponse{}

// GenerateSelfShipAppointmentSlotsResponse The `generateSelfShipAppointmentSlots` response.
type GenerateSelfShipAppointmentSlotsResponse struct {
	// UUID for the given operation.
	OperationId string `json:"operationId" validate:"regexp=^[a-zA-Z0-9-]*$"`
}

type _GenerateSelfShipAppointmentSlotsResponse GenerateSelfShipAppointmentSlotsResponse

// NewGenerateSelfShipAppointmentSlotsResponse instantiates a new GenerateSelfShipAppointmentSlotsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenerateSelfShipAppointmentSlotsResponse(operationId string) *GenerateSelfShipAppointmentSlotsResponse {
	this := GenerateSelfShipAppointmentSlotsResponse{}
	this.OperationId = operationId
	return &this
}

// NewGenerateSelfShipAppointmentSlotsResponseWithDefaults instantiates a new GenerateSelfShipAppointmentSlotsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenerateSelfShipAppointmentSlotsResponseWithDefaults() *GenerateSelfShipAppointmentSlotsResponse {
	this := GenerateSelfShipAppointmentSlotsResponse{}
	return &this
}

// GetOperationId returns the OperationId field value
func (o *GenerateSelfShipAppointmentSlotsResponse) GetOperationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperationId
}

// GetOperationIdOk returns a tuple with the OperationId field value
// and a boolean to check if the value has been set.
func (o *GenerateSelfShipAppointmentSlotsResponse) GetOperationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperationId, true
}

// SetOperationId sets field value
func (o *GenerateSelfShipAppointmentSlotsResponse) SetOperationId(v string) {
	o.OperationId = v
}

func (o GenerateSelfShipAppointmentSlotsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operationId"] = o.OperationId
	return toSerialize, nil
}

type NullableGenerateSelfShipAppointmentSlotsResponse struct {
	value *GenerateSelfShipAppointmentSlotsResponse
	isSet bool
}

func (v NullableGenerateSelfShipAppointmentSlotsResponse) Get() *GenerateSelfShipAppointmentSlotsResponse {
	return v.value
}

func (v *NullableGenerateSelfShipAppointmentSlotsResponse) Set(val *GenerateSelfShipAppointmentSlotsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenerateSelfShipAppointmentSlotsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenerateSelfShipAppointmentSlotsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenerateSelfShipAppointmentSlotsResponse(val *GenerateSelfShipAppointmentSlotsResponse) *NullableGenerateSelfShipAppointmentSlotsResponse {
	return &NullableGenerateSelfShipAppointmentSlotsResponse{value: val, isSet: true}
}

func (v NullableGenerateSelfShipAppointmentSlotsResponse) MarshalJSON() ([]byte, error) {
	return sonic.Marshal(v.value)
}

func (v *NullableGenerateSelfShipAppointmentSlotsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return sonic.Unmarshal(src, &v.value)
}
